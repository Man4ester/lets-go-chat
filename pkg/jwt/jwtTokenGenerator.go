package jwt

import (
	"time"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var secret string

func ApplySecret(secretValue string) {
	secret = secretValue
}

func GenerateJWT(userName string) (string, error) {
	var mySigningKey = []byte(secret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userName"] = userName
	claims["role"] = "user"
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func DecodeJWT(tokenString string)  (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	var userName string
	for key, val := range claims {
		if key == "userName" {
			userName = fmt.Sprintf("%v", val)
		}
	}
	return userName, nil
}
