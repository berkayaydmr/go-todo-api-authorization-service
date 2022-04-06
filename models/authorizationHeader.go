package models

import (
	"authorization-service/common"

	"github.com/dgrijalva/jwt-go"
)

type AuthHeader struct {
	Token string `header:"Authorization"`
}

func (auth *AuthHeader) Validate() bool {
	envSecretKey := common.GetEnvironment().SecretKey
	err := jwt.SigningMethodHS256.Verify(auth.Token, envSecretKey, envSecretKey)
	if err != nil {
		return true
	}
	return false
}
