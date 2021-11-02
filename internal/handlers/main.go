package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"lets-go-chat/pkg/hasher"
	Api "lets-go-chat/pkg/openapi3"
	"math/rand"
	"net/http"
)

var userStorage = make(map[int32]Api.User)

type Server struct {

}

func (s Server) FindUserByID(c echo.Context, id int32) error  {
	if user, ok := userStorage[id]; ok {
		j, _ := json.Marshal(user)
		c.JSON(http.StatusOK, string(j))
	}
	return nil
}

func (s Server) AddUser(c echo.Context) error {
	var req Api.AddUserJSONRequestBody
	c.Bind(&req)
	passwordHashed, err := hasher.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest,"{}")
		return nil
	}
	user := Api.User{
		Name: req.Name,
		UserName: req.UserName,
		Password:  passwordHashed,
	}
	userID := rand.Int31()
	user.Id = &userID
	userStorage[userID] = user
	j, _ := json.Marshal(user)
	c.JSON(http.StatusOK, string(j))
	return nil
}
