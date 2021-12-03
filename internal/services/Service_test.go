package services

import "testing"

var token = "token"

func TestApplyTokenFromRegistryOK(t *testing.T) {
	RegisterToken(token)
	err := ApplyTokenFromRegistry(token)
	if err != nil {
		t.Error("Expected no error")
	}
}

func TestApplyTokenFromRegistryKO(t *testing.T) {
	err := ApplyTokenFromRegistry("notExistingToken")
	if err == nil {
		t.Error("Expected error")
	}
}

func TestCacheService(t *testing.T)  {
	AddUserToCache("user1")
	AddUserToCache("user2")
	AddUserToCache("user3")
	res := GetTotalActiveUsers()
	if res != 3 {
		t.Error("Expected 3 user")
	}
}

func TestCacheServiceWithRemove(t *testing.T)  {
	AddUserToCache("user1")
	AddUserToCache("user2")
	AddUserToCache("user3")
	RemoveUserFromCache("user1")
	res := GetTotalActiveUsers()
	if res != 2 {
		t.Error("Expected 2 user")
	}
}