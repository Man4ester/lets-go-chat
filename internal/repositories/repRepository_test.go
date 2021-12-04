package repositories

import (
	"testing"
)

func TestUserRepository(t *testing.T) {
	userRepository := usersDataRepository {
		dbCon: nil,
	}
	RegisterUserRepository(&userRepository)
	res := GetUserRepository()
	if res == nil {
		t.Error("Expected not nil")
	}
}