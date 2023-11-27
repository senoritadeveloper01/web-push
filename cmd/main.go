package main

import (
	"context"
	viperOperations "web-push/internal/config"
	userOperations "web-push/internal/database"
	firebaseOperations "web-push/internal/firebase"
	kafkaOperations "web-push/internal/kafka"
	kafkaConfig "web-push/internal/kafka/model"
	"web-push/internal/utils"
	appRouting "web-push/pkg/config"
)

func main() {
	ctx := context.Background()

	// init logging
	utils.InitLog()

	// TODO: you can get from command line
	// init viper config
	properties := viperOperations.InitConfiguration(ctx)

	// TODO: check if properties are ok

	// init db
	db := userOperations.DbConnect(ctx, properties.Mysql.ConnectionUrl)
	defer userOperations.DbDisconnect(ctx, db)
	userOperations.InitDB(ctx, db)

	userOperationsSrv := userOperations.Connection{DB: db}

	// init firebase
	firebaseOperations.InitFirebase(ctx, properties.Firebase.ConfigFilePath)

	go func() {
		// init kafka consumer
		kafkaProperties := kafkaConfig.KafkaConfig{
			Brokers:               properties.Kafka.Servers,
			Topic:                 properties.Kafka.Topic,
			Id:                    properties.Kafka.Id,
			GroupId:               properties.Kafka.GroupId,
		}
		kafkaOperations.InitAndConsumeKafka(ctx, kafkaProperties, properties.Firebase.LogoUrl, firebaseOperations.SendNotification, userOperationsSrv.FindCredentialsByUserId)
		defer kafkaOperations.ShutDownKafka(ctx)
	}()

	// init rest api
	routingSrv := appRouting.Routing{UserOperationsSrv: userOperationsSrv}
	routingSrv.InitRouting(ctx)
}
