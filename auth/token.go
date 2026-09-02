package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	UserID int    `json:"userId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

const secretKey = "..."

func GenerateAccessToken(userID int, name string, email string) (string, error) {

	claims := AccessTokenClaims{
		UserID:           userID,
		Name:             name,
		Email:            email,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour))},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
}

func VerifyAccessToken(token string) (AccessTokenClaims, error) {

	claims := AccessTokenClaims{}

	_, err := jwt.ParseWithClaims(token, &claims, func(*jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})

	return claims, err
}
