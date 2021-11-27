package services

import "errors"

var tokenRegistry = make(map[string] bool)

func RegisterToken(token string) {
	tokenRegistry[token] = true
}

func ApplyTokenFromRegistry(token string) error{
	if _, ok := tokenRegistry[token]; ok {
		delete(tokenRegistry, token)
		return nil
	}

	return errors.New("token was used before")

}

