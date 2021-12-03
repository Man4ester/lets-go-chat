package jwt

import (
	"testing"
)

var userName = "Admin@gmail.com"


func TestGenerateJWT(t *testing.T) {
	token, err := GenerateJWT(userName)
	if err != nil {
		t.Error("Expected no err")
	}

	if token == "" {
		t.Error("Expected not empty token")
	}
}

func TestDecodeJWT(t *testing.T) {
	token, _ := GenerateJWT(userName)
	result, err := DecodeJWT(token)
	if err != nil {
		t.Error("Expected no err")
	}

	if result != userName {
		t.Error("Expected result is not " + userName)
	}
}

func BenchmarkGenerateJWT(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateJWT(userName)
	}
}
