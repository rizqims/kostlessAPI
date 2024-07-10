package dto

import "github.com/golang-jwt/jwt/v5"

type JwtTokenClaims struct {
	jwt.RegisteredClaims
	Id string `json:"id"`
	Username string `json:"username"`
}