package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	userOperationsImpl "web-push/internal/database"
	userCredentials "web-push/internal/database/model"
	createUserRequest "web-push/pkg/controller/model"
)

func TestSaveUserSuccess(t *testing.T) {
	controller := NewController(t)
	defer controller.Finish()

	userOperationsSrv := userOperationsImpl.NewMockUserOperations(controller)
	userOperationsSrv.EXPECT().SaveUser(Any()).Return(int64(999), nil).Times(1)

	router := mux.NewRouter()
	router.HandleFunc("/users/create", SaveUser(userOperationsSrv)).Methods("POST")

	newUserRequest := createUserRequest.CreateUserRequest{UserId: "testUserId", DeviceId: "testDeviceId", Credentials: "testCredentials"}
	newUserRequestStr, err := json.Marshal(newUserRequest)
	if err != nil {
		fmt.Printf("Error object to str: %s", err)
		return
	}
	req, err := http.NewRequest("POST", "/users/create", strings.NewReader(string(newUserRequestStr)))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusCreated, responseRecorder.Code)

	var createdUserCredentials userCredentials.UserCredentials
	err = json.NewDecoder(responseRecorder.Body).Decode(&createdUserCredentials)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, newUserRequest.UserId, createdUserCredentials.UserId, "Unexpected userId value in TestRegister")
	assert.Equal(t, newUserRequest.DeviceId, createdUserCredentials.DeviceId, "Unexpected deviceId value in TestRegister")
	assert.Equal(t, newUserRequest.Credentials, createdUserCredentials.Credentials, "Unexpected credentials value in TestRegister")
}

func TestSaveUserInvalidRequest(t *testing.T) {
	controller := NewController(t)
	defer controller.Finish()

	userOperationsSrv := userOperationsImpl.NewMockUserOperations(controller)
	userOperationsSrv.EXPECT().SaveUser(Any()).Return(int64(999), nil).Times(1)

	router := mux.NewRouter()
	router.HandleFunc("/users/create", SaveUser(userOperationsSrv)).Methods("POST")

	req, err := http.NewRequest("POST", "/users/create", strings.NewReader(string("invalidRequest")))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}
