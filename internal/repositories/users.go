package configs

import (
	"database/sql"
	"errors"
	"fmt"
	"lets-go-chat/configs"
	"lets-go-chat/internal/models"
	"log"
)

var DB *sql.DB

var UserNotFound = errors.New("USER NOT FOUND")

var UserWasNotSaved = errors.New("USER WAS NOT SAVED IN DB")

var psqlconn string

func init() {
	psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", configs.Config.DBConfig.DBHost, configs.Config.DBConfig.DBPort, configs.Config.DBConfig.DBUser, configs.Config.DBConfig.DBPassword, configs.Config.DBConfig.DBName)
}

func SaveUser(user models.User) error {
	openDBConnection()
	defer DB.Close()
	insertStmt := `insert into "users"("id", "username", "password") values($1, $2, $3)`
	_, err := DB.Exec(insertStmt, user.Id, user.UserName, user.Password)
	if err != nil {
		return UserWasNotSaved
	}
	return nil
}

func GetUserByUserName(userName string) (models.User, error) {
	openDBConnection()
	defer DB.Close()
	var userDB models.User
	userSql := "SELECT id, username, password FROM users WHERE username = $1"

	err := DB.QueryRow(userSql, userName).Scan(&userDB.Id, &userDB.UserName, &userDB.Password)
	if err != nil {
		return models.User{}, UserNotFound
	}

	return userDB, nil
}

func openDBConnection() {
	var err error
	DB, err = sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal("CANT CONNECT TO DB")
	}
}
