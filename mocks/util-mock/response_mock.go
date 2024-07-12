package utilmock

import "github.com/stretchr/testify/mock"

type ResponseMock struct {
	mock.Mock
}

func (r *ResponseMock) SendEmail(to, subject, body string) error {
	args := r.Called(to, subject, body)
	return args.Error(0)
}

func (r *ResponseMock) NotifyOwner(message string) error {
	args := r.Called(message)
	return args.Error(0)
}
