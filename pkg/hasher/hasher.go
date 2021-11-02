// Package hasher is a module for hashing txt valuse and validation
package hasher

import (
	"crypto/md5"
	"encoding/hex"

	"errors"
)


// HashPassword function for hashing txt
func HashPassword(password string) (string, error) {
	if len(password) < 3 {
		return "", errors.New("password to short")
	}
	data := []byte(password)
	b := md5.Sum(data)
	pass := hex.EncodeToString(b[:])
	return pass, nil
}

// CheckPasswordHash function for checking hashed txt via hash
func CheckPasswordHash(password, hash string) bool {
	receivedHash, err := HashPassword(password)
	if err !=nil {
		return false
	}
	return receivedHash == hash
}

