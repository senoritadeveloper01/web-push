// Code generated by MockGen. DO NOT EDIT.
// Source: user_operations.go

// Package database is a generated GoMock package.
package database

import (
	reflect "reflect"
	model "web-push/internal/database/model"

	gomock "github.com/golang/mock/gomock"
)

// MockUserOperations is a mock of UserOperations interface.
type MockUserOperations struct {
	ctrl     *gomock.Controller
	recorder *MockUserOperationsMockRecorder
}

// MockUserOperationsMockRecorder is the mock recorder for MockUserOperations.
type MockUserOperationsMockRecorder struct {
	mock *MockUserOperations
}

// NewMockUserOperations creates a new mock instance.
func NewMockUserOperations(ctrl *gomock.Controller) *MockUserOperations {
	mock := &MockUserOperations{ctrl: ctrl}
	mock.recorder = &MockUserOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserOperations) EXPECT() *MockUserOperationsMockRecorder {
	return m.recorder
}

// FindCredentialsByUserId mocks base method.
func (m *MockUserOperations) FindCredentialsByUserId(userId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCredentialsByUserId", userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCredentialsByUserId indicates an expected call of FindCredentialsByUserId.
func (mr *MockUserOperationsMockRecorder) FindCredentialsByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCredentialsByUserId", reflect.TypeOf((*MockUserOperations)(nil).FindCredentialsByUserId), userId)
}

// SaveUser mocks base method.
func (m *MockUserOperations) SaveUser(userCredentials model.UserCredentials) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUser", userCredentials)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveUser indicates an expected call of SaveUser.
func (mr *MockUserOperationsMockRecorder) SaveUser(userCredentials interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUser", reflect.TypeOf((*MockUserOperations)(nil).SaveUser), userCredentials)
}
