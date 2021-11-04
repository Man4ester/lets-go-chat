package models

type User struct {
	Id       string `json:"id"`
	UserName string `json:"userName" validate:"string,min=4"`
	Password string `json:"password" validate:"string,min=8"`
}
