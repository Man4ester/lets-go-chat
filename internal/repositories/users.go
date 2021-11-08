package configs

import (
	"database/sql"
	"errors"
	"lets-go-chat/internal/models"
)

type DBConnection struct {
	dbCon *sql.DB
}

var UserNotFound = errors.New("user not found")

var UserWasNotSaved = errors.New("user was not saved in db")

var dbConnection DBConnection

func NewUsersRepository(db *sql.DB){
	dbConnection = DBConnection{
		dbCon : db,
	}
}
func SaveUser(user models.User) error {

	insertStmt := `insert into "users"("id", "username", "password") values($1, $2, $3)`
	_, err := dbConnection.dbCon.Exec(insertStmt, user.Id, user.UserName, user.Password)
	if err != nil {
		return UserWasNotSaved
	}
	return nil
}

func GetUserByUserName(userName string) (models.User, error) {

	var userDB models.User
	userSql := "SELECT id, username, password FROM users WHERE username = $1"

	err := dbConnection.dbCon.QueryRow(userSql, userName).Scan(&userDB.Id, &userDB.UserName, &userDB.Password)
	if err != nil {
		return models.User{}, UserNotFound
	}

	return userDB, nil
}