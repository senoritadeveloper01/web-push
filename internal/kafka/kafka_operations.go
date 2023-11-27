package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	kafkaConfig "web-push/internal/kafka/model"
	"web-push/internal/utils"

	"github.com/Shopify/sarama"
)

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ctx   context.Context
	ready chan bool
}

var endSignal chan struct{}
var consumedMessageFunction func(context.Context, string, string, string, string)
var findFirebaseUserFunction func(context.Context, string) (string, error)
var firebaseLogoUrl string

func InitAndConsumeKafka(ctx context.Context, kafkaConfig kafkaConfig.KafkaConfig, logoUrl string, sendNotificationFunction func(context.Context, string, string, string, string), FindCredentialsByUser func(context.Context, string) (string, error)) {
	consumedMessageFunction = sendNotificationFunction
	findFirebaseUserFunction = FindCredentialsByUser
	firebaseLogoUrl = logoUrl

	keepRunning := true
	utils.LogInfo(ctx, utils.GetLogDetails("[KAFKA_OPERATIONS]", "[INIT_KAFKA]", "Starting a new Sarama consumer"))

	/**
	 * Setup a new Sarama consumer group
	 */
	consumer := Consumer{
		ctx:   ctx,
		ready: make(chan bool),
	}

	saramaConfig := sarama.NewConfig()
	saramaConfig.Consumer.Return.Errors = true
	// TODO: additional timeout config may be added

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(kafkaConfig.Brokers, ","), kafkaConfig.GroupId, saramaConfig)
	if err != nil {
		utils.LogFatalAndStop(ctx, utils.GetLogDetails("[KAFKA_OPERATIONS]", "[INIT_KAFKA]", fmt.Errorf("error creating consumer group client: %w", err).Error()))
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, strings.Split(kafkaConfig.Topic, ","), &consumer); err != nil {
				utils.LogFatalAndStop(ctx, utils.GetLogDetails("[KAFKA_OPERATIONS]", "[INIT_KAFKA]", fmt.Errorf("error from consumer: %w", err).Error()))
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	// receive from channel
	<-consumer.ready // Await till the consumer has been set up
	utils.LogInfo(ctx, utils.GetLogDetails("[KAFKA_OPERATIONS]", "[INIT_KAFKA]", "kafka consumer up and running"))

	endSignal = make(chan struct{})

	for keepRunning {
		for range endSignal {
		}
		utils.LogInfo(ctx, utils.GetLogDetails("[KAFKA_OPERATIONS]", "[SIGNAL_KAFKA]", "terminating: kafka shutDown has been called"))

		keepRunning = false
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		utils.LogPanic(ctx, utils.GetLogDetails("[KAFKA_OPERATIONS]", "[SHUT_DOWN_KAFKA]", fmt.Errorf("error closing kafka client: %w", err).Error()))
	}
}

func ShutDownKafka(ctx context.Context) {
	utils.LogInfo(ctx, utils.GetLogDetails("[KAFKA_OPERATIONS]", "[SHUT_DOWN_KAFKA]", "shutting down kafka"))
	close(endSignal)
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	utils.LogInfo(consumer.ctx, utils.GetLogDetails("[KAFKA_OPERATIONS]", "[SET_UP_KAFKA]", "setting up kafka"))
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	utils.LogInfo(consumer.ctx, utils.GetLogDetails("[KAFKA_OPERATIONS]", "[CLEAN_UP_KAFKA]", "cleaning up kafka"))
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message := <-claim.Messages():
			utils.LogInfo(session.Context(), utils.GetLogDetails("[KAFKA_OPERATIONS]", "[INIT_KAFKA]", fmt.Sprintf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)))
			var eventMessage kafkaConfig.EventMessage
			err := json.Unmarshal(message.Value, &eventMessage)
			if err != nil {
				utils.LogError(session.Context(), utils.GetLogDetails("[REST_OPERATIONS]", "[SAVE_USER]", fmt.Errorf("parse error ocurred in ConsumeClaim: %w", err).Error()))
			} else {
				// TODO: maybe saved to redis
				// TODO: should unregister user after logout
				userToken, err := findFirebaseUserFunction(session.Context(), eventMessage.AccountMail)
				if err == nil {
					consumedMessageFunction(session.Context(), userToken, eventMessage.Subject, eventMessage.Fragment, firebaseLogoUrl)
				}
			}
			session.MarkMessage(message, "")

		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/Shopify/sarama/issues/1192
		case <-session.Context().Done():
			utils.LogInfo(session.Context(), utils.GetLogDetails("[KAFKA_OPERATIONS]", "[SESSION_CONTEXT_DONE]", "kafka context done"))
			return nil
		}
	}
}
