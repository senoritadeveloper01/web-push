package database

import (
	"context"
	"database/sql"
	"fmt"
	"web-push/internal/utils"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnect(ctx context.Context, URI string) *sql.DB {
	var err error
	db, err := sql.Open("mysql", URI)
	if err != nil {
		utils.LogFatalAndStop(ctx, utils.GetLogDetails("[DB_OPERATIONS]", "[DB_CONNECT]", fmt.Errorf("error opening db: %w", err).Error()))
	}

	err = db.Ping()
	if err != nil {
		utils.LogFatalAndStop(ctx, utils.GetLogDetails("[DB_OPERATIONS]", "[DB_CONNECT]", fmt.Errorf("error pinging db: %w", err).Error()))
	}

	return db
}

func InitDB(ctx context.Context, db *sql.DB) {
	stmt, err := db.Prepare(
		"CREATE TABLE IF NOT EXISTS user_credentials(id int primary key NOT NULL AUTO_INCREMENT, " +
			"user_id varchar(255) NOT NULL UNIQUE, device_id varchar(1000) NOT NULL, credentials varchar(2000) NOT NULL, " +
			"created_date datetime default CURRENT_TIMESTAMP, modified_date datetime default CURRENT_TIMESTAMP)")
	if err != nil {
		utils.LogFatalAndStop(ctx, utils.GetLogDetails("[DB_OPERATIONS]", "[INIT_DB]", fmt.Errorf("error preparing db table stmt: %w", err).Error()))
	}

	// TODO: a timeout can be added
	_, err = stmt.Exec()
	if err != nil {
		utils.LogFatalAndStop(ctx, utils.GetLogDetails("[DB_OPERATIONS]", "[INIT_DB]", fmt.Errorf("error creating db table: %w", err).Error()))
	} else {
		utils.LogInfo(ctx, utils.GetLogDetails("[DB_OPERATIONS]", "[INIT_DB]", "table created successfully."))
	}
}

func DbDisconnect(ctx context.Context, db *sql.DB) {
	err := db.Close()
	if err != nil {
		utils.LogPanic(ctx, utils.GetLogDetails("[DB_OPERATIONS]", "[DB_DISCONNECT]", fmt.Errorf("error closing db: %w", err).Error()))
	}
}
