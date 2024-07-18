package app

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CheckJwt(config *Configuration, tokenStr string) bool {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Jwt.ServerSecret), nil
		},
	)
	if err != nil {
		return false
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return false
	}

	if claims.Issuer != "digitao-official-website" {
		return false
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		return false
	}

	return true
}

func SignJwt(config *Configuration) string {
	currentTime := time.Now()
	expireIn, err := time.ParseDuration(config.Jwt.ExpireIn)
	if err != nil {
		return ""
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    "digitao-official-website",
			IssuedAt:  jwt.NewNumericDate(currentTime),
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(expireIn)),
		},
	)

	jwt, err := token.SignedString([]byte(config.Jwt.ServerSecret))
	if err != nil {
		return ""
	}

	return jwt
}
