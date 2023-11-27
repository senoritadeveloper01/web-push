package database

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"

	userCredentials "web-push/internal/database/model"
)

func Test1SaveUserSuccess(t *testing.T) {
	// t.parallel() for parallel run
	var mock sqlmock.Sqlmock
	var err error

	ctx := context.Background()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	connection := Connection{DB: db}

	user := userCredentials.UserCredentials{
		UserId:      "testClientId",
		DeviceId:    "testDeviceId",
		Credentials: "testCredentials",
	}

	mock.ExpectExec("INSERT INTO user_credentials").
		WithArgs(user.UserId, user.DeviceId, user.Credentials).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = connection.SaveUser(ctx, user)
	if err != nil {
		t.Errorf("error '%s' was not expected, while inserting a row", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test2FindCredentialsByUserSuccess(t *testing.T) {
	var mock sqlmock.Sqlmock
	var err error

	ctx := context.Background()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	connection := Connection{DB: db}

	rows := sqlmock.NewRows([]string{"credentials"}).AddRow("testCredentials")
	mock.ExpectQuery("SELECT (.+) FROM user_credentials WHERE user_id = ?").
		WithArgs("testClientId").
		WillReturnRows(rows)

	foundCredentials, err := connection.FindCredentialsByUserId(ctx, "testClientId")
	if err != nil {
		t.Errorf("error '%s' was not expected, while selecting a row", err)
	}

	assert.Equal(t, "testCredentials", foundCredentials, "expected and found user credentials do not match")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
