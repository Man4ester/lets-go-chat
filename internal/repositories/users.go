package configs

import (
	"errors"
	"lets-go-chat/internal/models"
)

var userStorage = make(map[string]models.User)



func SaveUser(user models.User) {
	userStorage[user.UserName] = user
}

func GetUserByUserName(userName string)  (models.User, error){
	if user, ok := userStorage[userName]; ok {
		return user, nil
	}
	return models.User{}, errors.New("USER NOT FOUND")
}
