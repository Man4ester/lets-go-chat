package repositories

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"lets-go-chat/internal/models"
)

func TestUserRepository(t *testing.T) {
	userRepository := usersDataRepository{
		dbCon: nil,
	}
	RegisterUserRepository(&userRepository)
	res := GetUserRepository()
	if res == nil {
		t.Error("Expected not nil")
	}
}

func TestUserSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := models.User{
		UserName: "Alex",
		Password: "password",
		Id:       "id",
	}
	repository := NewUsersDataRepository(db)
	mock.ExpectExec("insert into \"users\"").WithArgs(user.Id, user.UserName, user.Password).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repository.SaveUser(user)
	if err != nil {
		t.Error("Expected no errors")
	}
}

func TestGetUserByUserName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := models.User{
		UserName: "Alex",
		Password: "password",
		Id:       "id",
	}
	columns := []string{"id", "username", "password"}
	repository := NewUsersDataRepository(db)
	mock.ExpectQuery("SELECT id, username, password FROM users").WithArgs(user.UserName).WillReturnRows(
		sqlmock.NewRows(columns).AddRow(user.Id, user.UserName, user.Password))

	response, err := repository.GetUserByUserName(user.UserName)
	if err != nil {
		t.Error("Expected no errors")
	}

	if response.UserName != user.UserName {
		t.Errorf("Expected '%s'", user.UserName)
	}
}