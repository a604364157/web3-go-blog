package utils

import (
	"fmt"
	"time"
	"web3-go-blog/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint, username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"usr": username,
		"exp": time.Now().Add(config.TokenTTl).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTSecret)
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.JWTSecret, nil
	})
}
