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

type SeekerServiceTestSuite struct {
	suite.Suite
	repomock     *repomock.SeekerRepoMock
	sS           service.SeekerServ
	jwtmock      *utilmock.JwtTokenMock
	responsemock *utilmock.ResponseMock
	hashmock     *utilmock.HashPasswordMock
}

func (suite *SeekerServiceTestSuite) SetupTest() {
	suite.repomock = new(repomock.SeekerRepoMock)
	suite.jwtmock = new(utilmock.JwtTokenMock)
	suite.responsemock = new(utilmock.ResponseMock)
	suite.hashmock = new(utilmock.HashPasswordMock)
	suite.sS = service.NewSeekerServ(suite.repomock, suite.jwtmock)
}

func TestSeekerServiceTestSuite(t *testing.T) {
	suite.Run(t, new(SeekerServiceTestSuite))
}

var mockSeeker = model.Seekers{
	Id:           "1",
	Fullname:     "Seeker 1",
	Username:     "seeker1username",
	Password:     "seeker1password",
	Email:        "seeker@gmail.com",
	PhoneNumber:  "08123456789",
	PhotoProfile: "seeker1.jpg",
}

var updateSeekermock = model.Seekers{
	Id:           "1",
	Fullname:     "Seeker 1",
	Username:     "seeker1",
	Password:     "seeker1",
	Email:        "seeker@gmail.com",
	PhoneNumber:  "08123456789",
	AtitudePoits: 10,
	PhotoProfile: "seeker1.jpg",
}

func (suite *SeekerServiceTestSuite) TestDeleteSeeker_Success() {
	suite.repomock.On("DeleteSeeker", mockSeeker.Id).Return(nil)

	err := suite.sS.DeleteSeeker(mockSeeker.Id)
	assert.NoError(suite.T(), err)
}

func (suite *SeekerServiceTestSuite) TestGetAllSeekers() {
	suite.repomock.On("GetAllSeekers").Return([]*model.Seekers{&mockSeeker}, nil)

	_, err := suite.sS.GetAllSeekers()

	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}

func (suite *SeekerServiceTestSuite) TestGetSeekeryID() {
	suite.repomock.On("GetSeekerByID", mockSeeker.Id).Return(mockSeeker, nil)

	_, err := suite.sS.GetSeekerByID(mockSeeker.Id)

	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}

func (suite *SeekerServiceTestSuite) TestUpdateAttitudePoints_Success() {
	suite.repomock.On("GetSeekerByID", mockSeeker.Id).Return(mockSeeker, nil)
	suite.repomock.On("UpdateAttitudePoints", updateSeekermock.Id, updateSeekermock.AtitudePoits).Return(nil)

	err := suite.sS.UpdateAttitudePoints(updateSeekermock.Id, updateSeekermock.AtitudePoits)

	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}

func (suite *SeekerServiceTestSuite) TestUpdateAttitudePoints_FailedGetSeeker() {
	suite.repomock.On("GetSeekerByID", mockSeeker.Id).Return(model.Seekers{}, errors.New("failed get seeker"))
	suite.repomock.On("UpdateAttitudePoints", updateSeekermock.Id, updateSeekermock.AtitudePoits).Return(nil)

	err := suite.sS.UpdateAttitudePoints(updateSeekermock.Id, updateSeekermock.AtitudePoits)

	assert.Error(suite.T(), err)
}

func (suite *SeekerServiceTestSuite) TestUpdateAttitudePoints_LowerThan5() {
	suite.repomock.On("GetSeekerByID", mockSeeker.Id).Return(mockSeeker, nil)
	suite.responsemock.On("SendEmail", mockSeeker.Email, "", "").Return(nil)
	suite.repomock.On("UpdateAttitudePoints", updateSeekermock.Id, 4).Return(nil)

	err := suite.sS.UpdateAttitudePoints(updateSeekermock.Id, 4)

	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}

func (suite *SeekerServiceTestSuite) TestUpdateAttitudePoints_HigherThan10() {
	suite.repomock.On("GetSeekerByID", mockSeeker.Id).Return(mockSeeker, nil)
	suite.responsemock.On("NotifyOwner", "").Return(nil)
	suite.repomock.On("UpdateAttitudePoints", updateSeekermock.Id, 11).Return(nil)

	err := suite.sS.UpdateAttitudePoints(updateSeekermock.Id, 11)

	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}

func (suite *SeekerServiceTestSuite) TestUpdateProfile() {
	var mockSeekerUpdate = model.Seekers{
		Id:           "1",
		Fullname:     "Seeker 1",
		Username:     "seeker1",
		Password:     "seeker1",
		Email:        "seeker@gmail.com",
		PhoneNumber:  "08123456789",
		PhotoProfile: "seeker1.jpg",
		UpdatedAt:    time.Now(),
	}

	suite.repomock.On("UpdateSeeker", mockSeeker.Id, mockSeekerUpdate).Return(nil)

	err := suite.sS.UpdateProfile(mockSeeker.Id, mockSeekerUpdate)

	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}

func (suite *SeekerServiceTestSuite) TestLogin_Success() {
	mockPayload := dto.LoginDto{
		Username: mockSeeker.Username,
		Password: mockSeeker.Password,
	}

	mockLoginResponse := dto.LoginResponse{
		Token: "token",
	}

	mockSeeker.Password = "$2a$10$mATz2DeGgNWp5OkmMeXh.uZk6uVte1G8BOZdLXoNV7ZouYytlI3Me"
	suite.repomock.On("GetBySeeker", mockSeeker.Username).Return(mockSeeker, nil)

	suite.hashmock.On("CheckPasswordHash", mockSeeker.Password, mockSeeker.Password).Return(nil)
	suite.jwtmock.On("GenerateToken", mockSeeker.Id, mockSeeker.Username).Return(mockLoginResponse, nil)
	_, err := suite.sS.Login(mockPayload)

	assert.NoError(suite.T(), err)
}