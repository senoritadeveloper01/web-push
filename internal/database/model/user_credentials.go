package model

import "time"

type UserCredentials struct {
	Id          int64
	UserId      string
	DeviceId    string
	Credentials string
	CreatedDate time.Time
	UpdatedDate time.Time
}
