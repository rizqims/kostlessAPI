package utilmock

import "github.com/stretchr/testify/mock"

type HashPasswordMock struct {
	mock.Mock
}

func (h *HashPasswordMock) HashPassword(password string) (string, error) {
	args := h.Called(password)
	return args.Get(0).(string), args.Error(1)
}

func (h *HashPasswordMock) CheckPasswordHash(password, hash string) error {
	args := h.Called(password, hash)
	return args.Error(0)
}
