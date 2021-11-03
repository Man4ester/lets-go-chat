package models

type LoginUserRequest struct {
	UserName     string  `json:"userName" description:" The user name for login"`
	Password     string  `json:"password" description:"The password for login in clear text"`
}