package utilmock

import (
	"kostless/model/dto"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
)

type JwtTokenMock struct {
	mock.Mock
}

func (j *JwtTokenMock) GenerateToken(id, username string) (dto.LoginResponse, error) {
	args := j.Called(id, username)
	return args.Get(0).(dto.LoginResponse), args.Error(1)
}

func (j *JwtTokenMock) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	args := j.Called(tokenString)
	return args.Get(0).(jwt.MapClaims), args.Error(1)
}
