package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("Secret_key")

type JwtClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	claims := &JwtClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	ss , err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func ValidateToken(tokenstring string) (*JwtClaims, error) {
	claims := &JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token)(interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}