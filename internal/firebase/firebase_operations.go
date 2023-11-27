package firebase

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"google.golang.org/api/option"
	"web-push/internal/utils"
)

var fcmClient *messaging.Client

func InitFirebase(ctx context.Context, filePath string) {
	// Use the path to your service account credential json file
	opt := option.WithCredentialsFile(filePath)
	// Create a new firebase app
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err.Error())
	}
	// Get the FCM object
	fcmClient, err = app.Messaging(ctx)
	if err != nil {
		utils.LogFatalAndStop(ctx, utils.GetLogDetails("[FIREBASE_OPERATIONS]", "[INIT_FIREBASE]", fmt.Errorf("error getting messaging client: %w", err).Error()))
	}
}

func SendNotification(ctx context.Context, userToken string, title string, body string, imageUrl string) {
	_, err := fcmClient.Send(ctx, &messaging.Message{
		Notification: &messaging.Notification{
			Title:    title,
			Body:     body,
			ImageURL: imageUrl,
		},
		Token: userToken, // a token that you received from a client
	})

	/*
		log.Debug("Response success count : ", response.SuccessCount)
		log.Debug("Response failure count : ", response.FailureCount)
	*/

	if err != nil {
		utils.LogError(ctx, utils.GetLogDetails("[FIREBASE_OPERATIONS]", "[SEND_NOTIFICATION]", fmt.Errorf("error sending notification: %w", err).Error()))
	} else {
		utils.LogInfo(ctx, utils.GetLogDetails("[FIREBASE_OPERATIONS]", "[SEND_NOTIFICATION]", "notification successfully sent"))
	}
}
