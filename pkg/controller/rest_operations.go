package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	userOperations "web-push/internal/database"
	userCredentials "web-push/internal/database/model"
	"web-push/internal/utils"
	createUserRequest "web-push/pkg/controller/model"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func SaveUser(ctx context.Context, userOperationsSvc userOperations.UserOperations) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// TODO: txn id is still null
		ctx = context.WithValue(ctx, "TXN_ID", uuid.New().String())
		writer.Header().Set("Content-Type", "application/json")

		// TODO: set error object to response body

		var newUserRequest createUserRequest.CreateUserRequest
		err := json.NewDecoder(request.Body).Decode(&newUserRequest)
		if err != nil {
			utils.LogError(ctx, utils.GetLogDetails("[REST_OPERATIONS]", "[SAVE_USER]", fmt.Errorf("there was an error decoding the request body into the struct: %w", err).Error()))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		validate := validator.New()
		err = validate.Struct(newUserRequest)
		if err != nil {
			utils.LogError(ctx, utils.GetLogDetails("[REST_OPERATIONS]", "[SAVE_USER]", fmt.Errorf("request validation failed: %w", err).Error()))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		var newUser userCredentials.UserCredentials
		newUser.UserId = newUserRequest.UserId
		newUser.DeviceId = newUserRequest.DeviceId // TODO: can get from header
		newUser.Credentials = newUserRequest.Credentials

		savedUserId, err := userOperationsSvc.SaveUser(ctx, newUser)
		if err != nil {
			utils.LogError(ctx, utils.GetLogDetails("[REST_OPERATIONS]", "[SAVE_USER]", fmt.Errorf("there was an error saving user record for userId %s: %w", newUserRequest.UserId, err).Error()))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(writer).Encode(savedUserId)
		if err != nil {
			utils.LogError(ctx, utils.GetLogDetails("[REST_OPERATIONS]", "[SAVE_USER]", fmt.Errorf("there was an error setting user record id in response body for userId %s: %w", newUserRequest.UserId, err).Error()))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		writer.WriteHeader(http.StatusCreated)
	}
}
