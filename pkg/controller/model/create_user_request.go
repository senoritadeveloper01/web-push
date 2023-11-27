package model

type CreateUserRequest struct {
	UserId      string `json:"userId" validate:"email,required"`
	DeviceId    string `json:"clientId" validate:"required"`
	Credentials string `json:"credentials" validate:"required"`
}
