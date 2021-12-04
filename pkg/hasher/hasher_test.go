package hasher

import (
	"testing"
)

var password = "password"
var expectedHash = "5f4dcc3b5aa765d61d8327deb882cf99"

func TestHashPasswordOK(t *testing.T) {
	res, err := HashPassword(password)
	if err != nil {
		t.Error("Expected no error")
	}

	if res != expectedHash {
		t.Error("Expected result is not:" + expectedHash)
	}
}

func TestHashPasswordKO(t *testing.T) {
	_, err := HashPassword("sh")
	if err == nil {
		t.Error("Expected error")
	}
}

func TestCheckPasswordHashOK(t *testing.T) {
	res := CheckPasswordHash(password, expectedHash)
	if !res {
		t.Error("Expected hash wasn't successful")
	}
}