package auth

import (
	"time"

	"github.com/b1g-Dr4goN/hussatapi/configs/env"
	"github.com/golang-jwt/jwt"
)

func CreateJWT(secret []byte, userId string, role string) (string, error) {
	expiration := time.Second * time.Duration(env.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"role":      role,
		"issuedAt":  time.Now().Unix(),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
