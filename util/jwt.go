package util

import (
	"errors"
	"fmt"
	"kostless/config"
	"kostless/model/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken interface {
	GenerateToken(username string) (dto.LoginResponse, error)
	ValidateToken(tokenString string) (jwt.MapClaims, error)
}

type jwtClaims struct {
	config config.JwtConfig
}

func (j *jwtClaims) GenerateToken(username string) (dto.LoginResponse, error) {
	claims := dto.JwtTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.config.Issues,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.Durasi * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Username: username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(j.config.Key))
	if err != nil {
		return dto.LoginResponse{}, err
	}
	return dto.LoginResponse{Token: ss}, nil
}

func (j *jwtClaims) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.config.Key), nil
	})
	if err != nil {
		return nil, errors.New("failed verify token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok || claims["iss"] != j.config.Issues {
		fmt.Print("error===", err)
		return nil, errors.New("invalid issuer or claims")
	}
	return claims, nil
}

func NewJwtUtil(cg config.JwtConfig) JwtToken {
	return &jwtClaims{config: cg}
}
