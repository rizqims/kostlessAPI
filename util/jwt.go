package util

import (
	"kostless-api/model/dto"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = os.Getenv("JWT_KEY")

type JwtClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (dto.LoginResponse, error) {
	claims := &JwtClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: os.Getenv("JWT_ISSUER_NAME"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	ss , err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return dto.LoginResponse{}, err
	}
	return dto.LoginResponse{Token: ss}, nil
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