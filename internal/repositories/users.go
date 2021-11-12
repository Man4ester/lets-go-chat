package repositories

import (
	"database/sql"
	"errors"
	"lets-go-chat/internal/models"
)

type UserRepository interface {
	SaveUser(user models.User) error
	GetUserByUserName(userName string) (models.User, error)
}

type usersDataRepository struct {
	dbCon *sql.DB
}

func NewUsersDataRepository(db *sql.DB) *usersDataRepository {
	return &usersDataRepository{
		dbCon: db,
	}
}

var UserNotFound = errors.New("user not found")

var UserWasNotSaved = errors.New("user was not saved in db")

func (usersDataRep usersDataRepository) SaveUser (user models.User) error {

	insertStmt := `insert into "users"("id", "username", "password") values($1, $2, $3)`
	_, err := usersDataRep.dbCon.Exec(insertStmt, user.Id, user.UserName, user.Password)
	if err != nil {
		return UserWasNotSaved
	}
	return nil
}

func (usersDataRep usersDataRepository)GetUserByUserName(userName string) (models.User, error) {

	var userDB models.User
	userSql := "SELECT id, username, password FROM users WHERE username = $1"

	err := usersDataRep.dbCon.QueryRow(userSql, userName).Scan(&userDB.Id, &userDB.UserName, &userDB.Password)
	if err != nil {
		return models.User{}, UserNotFound
	}

	return userDB, nil
}