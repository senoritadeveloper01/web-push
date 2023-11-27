package database

import (
	"context"
	"database/sql"
	"fmt"
	userCredentials "web-push/internal/database/model"
	"web-push/internal/utils"
)

type Connection struct {
	DB *sql.DB
}

func (c Connection) SaveUser(ctx context.Context, userCredentials userCredentials.UserCredentials) (int64, error) {
	result, err := c.DB.Exec("INSERT INTO user_credentials (user_id, device_id, credentials) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE device_id = ?, credentials = ?",
		userCredentials.UserId, userCredentials.DeviceId, userCredentials.Credentials, userCredentials.DeviceId, userCredentials.Credentials)
	if err != nil {
		utils.LogError(ctx, utils.GetLogDetails("[DB_OPERATIONS]", "[SAVE_USER]", fmt.Errorf("error saving user: %w", err).Error()))
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		utils.LogError(ctx, utils.GetLogDetails("[DB_OPERATIONS]", "[SAVE_USER]", fmt.Errorf("error getting last saved user id: %w", err).Error()))
		return 0, fmt.Errorf("saveUser: %w", err)
	}
	return id, nil
}

func (c Connection) FindCredentialsByUserId(ctx context.Context, userId string) (string, error) {
	row := c.DB.QueryRow("SELECT credentials FROM user_credentials WHERE user_id = ?", userId)
	var credentials string
	err := row.Scan(&credentials)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.LogInfo(ctx, utils.GetLogDetails("[DB_OPERATIONS]", "[FIND_CREDENTIALS_BY_USER_ID]", fmt.Errorf("cannot find credentials by userId %s: no such user", userId).Error()))
			return "", err
		}

		utils.LogError(ctx, utils.GetLogDetails("[DB_OPERATIONS]", "[FIND_CREDENTIALS_BY_USER_ID]", fmt.Errorf("error trying to find credentials by userId %s: %w", userId, err).Error()))
		return "", err
	}

	return credentials, nil
}
