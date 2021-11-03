package models

type CreateUserRequest struct {
	UserName     string  `json:"userName" validate:"string,min=4"`
	Password     string  `json:"password" validate:"string,min=8"`
}