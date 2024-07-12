package testing

import (
	"errors"
	repomock "kostless/mocks/repo-mock"
	utilmock "kostless/mocks/util-mock"
	"kostless/model"
	"kostless/model/dto"
	"kostless/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	repomock     *repomock.UserRepoMock
	uS           service.UserServ
	jwtmock      *utilmock.JwtTokenMock
	responsemock *utilmock.ResponseMock
	hashmock     *utilmock.HashPasswordMock
}

func (suite *UserServiceTestSuite) SetupTest() {
	suite.repomock = new(repomock.UserRepoMock)
	suite.jwtmock = new(utilmock.JwtTokenMock)
	suite.responsemock = new(utilmock.ResponseMock)
	suite.hashmock = new(utilmock.HashPasswordMock)
	suite.uS = service.NewUserServ(suite.repomock, suite.jwtmock)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

var mockUser = model.User{
	Id:           "1",
	Fullname:     "User 1",
	Username:     "user1username",
	Password:     "user1password",
	Email:        "user1@gmail.com",
	PhoneNumber:  "08123456789",
	PhotoProfile: "user1.jpg",
	UpdatedAt:    time.Now(),
}

func (suite *UserServiceTestSuite) TestCreatedNewUser_Success() {
	hashedPassword := "$2a$10$mATz2DeGgNWp5OkmMeXh.uZk6uVte1G8BOZdLXoNV7ZouYytlI3Me"
	suite.hashmock.On("HashPassword", mockUser.Password).Return(hashedPassword, nil)
	suite.repomock.On("CreatedNewUser", mockUser).Return(mockUser, nil)

	newUser, err := suite.uS.CreatedNewUser(mockUser)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockUser, newUser)
}

func (suite *UserServiceTestSuite) TestLogin_Success() {
	mockPayload := dto.LoginDto{
		Username: mockUser.Username,
		Password: mockUser.Password,
	}

	mockLoginResponse := dto.LoginResponse{
		Token: "token",
	}

	mockUser.Password = "$2a$10$mATz2DeGgNWp5OkmMeXh.uZk6uVte1G8BOZdLXoNV7ZouYytlI3Me"
	suite.repomock.On("GetByUsername", mockUser.Username).Return(mockUser, nil)
	suite.hashmock.On("CheckPasswordHash", mockUser.Password, mockPayload.Password).Return(nil)
	suite.jwtmock.On("GenerateToken", mockUser.Id, mockUser.Username).Return(mockLoginResponse, nil)

	response, err := suite.uS.Login(mockPayload)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockLoginResponse, response)
}

func (suite *UserServiceTestSuite) TestUpdateProfile_Success() {
	var updatedUser = model.User{
		Id:           "1",
		Fullname:     "User 1 Updated",
		Username:     "user1updated",
		Password:     "user1updatedpassword",
		Email:        "user1updated@gmail.com",
		PhoneNumber:  "08123456789",
		PhotoProfile: "user1updated.jpg",
		UpdatedAt:    time.Now(),
	}

	suite.repomock.On("PutUpdateUserProf", updatedUser.Id, updatedUser).Return(nil)

	err := suite.uS.UpdateProfile(updatedUser.Id, updatedUser)
	assert.NoError(suite.T(), err)
}

func (suite *UserServiceTestSuite) TestGetUser_Success() {
	suite.repomock.On("GetUserById", mockUser.Id).Return(mockUser, nil)

	user, err := suite.uS.GetUser(mockUser.Id)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockUser, user)
}

func (suite *UserServiceTestSuite) TestCreatedNewUser_Failed() {
	suite.hashmock.On("HashPassword", mockUser.Password).Return("", errors.New("hashing error"))

	_, err := suite.uS.CreatedNewUser(mockUser)
	assert.Error(suite.T(), err)
}

func (suite *UserServiceTestSuite) TestLogin_FailedInvalidUsername() {
	mockPayload := dto.LoginDto{
		Username: mockUser.Username,
		Password: mockUser.Password,
	}

	suite.repomock.On("GetByUsername", mockUser.Username).Return(model.User{}, errors.New("username invalid"))

	_, err := suite.uS.Login(mockPayload)
	assert.Error(suite.T(), err)
}

func (suite *UserServiceTestSuite) TestLogin_FailedInvalidPassword() {
	mockPayload := dto.LoginDto{
		Username: mockUser.Username,
		Password: mockUser.Password,
	}

	mockUser.Password = "$2a$10$mATz2DeGgNWp5OkmMeXh.uZk6uVte1G8BOZdLXoNV7ZouYytlI3Me"
	suite.repomock.On("GetByUsername", mockUser.Username).Return(mockUser, nil)
	suite.hashmock.On("CheckPasswordHash", mockUser.Password, mockPayload.Password).Return(errors.New("password incorrect"))

	_, err := suite.uS.Login(mockPayload)
	assert.Error(suite.T(), err)
}

func (suite *UserServiceTestSuite) TestUpdateProfile_Failed() {
	var updatedUser = model.User{
		Id:           "1",
		Fullname:     "User 1 Updated",
		Username:     "user1updated",
		Password:     "user1updatedpassword",
		Email:        "user1updated@gmail.com",
		PhoneNumber:  "08123456789",
		PhotoProfile: "user1updated.jpg",
		UpdatedAt:    time.Now(),
	}

	suite.repomock.On("PutUpdateUserProf", updatedUser.Id, updatedUser).Return(errors.New("update failed"))

	err := suite.uS.UpdateProfile(updatedUser.Id, updatedUser)
	assert.Error(suite.T(), err)
}

func (suite *UserServiceTestSuite) TestGetUser_Failed() {
	suite.repomock.On("GetUserById", mockUser.Id).Return(model.User{}, errors.New("user not found"))

	_, err := suite.uS.GetUser(mockUser.Id)
	assert.Error(suite.T(), err)
}
