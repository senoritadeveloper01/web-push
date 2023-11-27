package database

import (
	"context"
	userCredentials "web-push/internal/database/model"
)

type UserOperations interface {
	SaveUser(ctx context.Context, userCredentials userCredentials.UserCredentials) (int64, error)
	FindCredentialsByUserId(ctx context.Context, userId string) (string, error)
}
