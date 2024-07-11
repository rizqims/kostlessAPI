package testing

import (
	"database/sql"
	"errors"
	"kostless/model"
	"kostless/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SeekerRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    repository.SeekerRepo
}

func (suite *SeekerRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = repository.NewUserSeeker(suite.mockDb)
}

func TestSeekerRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(SeekerRepositoryTestSuite))
}

var mockingSeeker = model.Seekers{
	Id:           "1",
	Fullname:     "Seeker 1",
	Username:     "seeker1",
	Password:     "seeker1",
	Email:        "seeker@gmail.com",
	PhoneNumber:  "08123456789",
	PhotoProfile: "seeker1.jpg",
	CreatedAt:    time.Now(),
	UpdatedAt:    time.Now(),
}

func (suite *SeekerRepositoryTestSuite) TestDeleteSeeker_Success() {
	suite.mockSql.ExpectExec("DELETE FROM seekers").WithArgs(mockingSeeker.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	err := suite.repo.DeleteSeeker(mockingSeeker.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *SeekerRepositoryTestSuite) TestDeleteSeeker_Failed() {
	suite.mockSql.ExpectExec("DELETE FROM seekers").WithArgs(mockingSeeker.Id).WillReturnError(errors.New("Delete Room Failed"))

	err := suite.repo.DeleteSeeker(mockingSeeker.Id)
	assert.Error(suite.T(), err)
}

func (suite *SeekerRepositoryTestSuite) TestGetAllSeekers_Success() {
	mockingSeekerAll := &model.Seekers{
		Id:           "1",
		Fullname:     "Seeker 1",
		Username:     "seeker1",
		Email:        "seeker@gmail.com",
		PhoneNumber:  "08123456789",
		PhotoProfile: "seeker1.jpg",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	suite.mockSql.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "fullname", "username", "email", "phone_number", "attiude_points", "status", "photo_profile", "room_id", "created_at", "updated_at"}).AddRow(mockingSeekerAll.Id, mockingSeekerAll.Fullname, mockingSeekerAll.Username, mockingSeekerAll.Email, mockingSeekerAll.PhoneNumber, mockingSeekerAll.AtitudePoits, mockingSeekerAll.Status, mockingSeekerAll.PhotoProfile, mockingSeekerAll.RoomId, mockingSeekerAll.CreatedAt, mockingSeekerAll.UpdatedAt))

	actual, err := suite.repo.GetAllSeekers()
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), actual, 1)
	assert.Equal(suite.T(), mockingSeekerAll, actual[0])
}

func (suite *SeekerRepositoryTestSuite) TestGetAllSeekers_Failed() {
	suite.mockSql.ExpectQuery("SELECT").WillReturnError(errors.New("Get All Seekers Failed"))

	_, err := suite.repo.GetAllSeekers()
	assert.Error(suite.T(), err)
}

func (suite *SeekerRepositoryTestSuite) TestGetAllSeekers_FailedScan() {
	suite.mockSql.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "fullname", "username", "password", "email", "phone_number", "attiude_points", "status", "photo_profile", "room_id", "created_at", "updated_at"}).AddRow(mockingSeeker.Id, mockingSeeker.Fullname, mockingSeeker.Username, mockingSeeker.Password, mockingSeeker.Email, mockingSeeker.PhoneNumber, "Invalid Attitude Points", mockingSeeker.Status, mockingSeeker.PhotoProfile, mockingSeeker.RoomId, mockingSeeker.CreatedAt, mockingSeeker.UpdatedAt))

	_, err := suite.repo.GetAllSeekers()
	assert.Error(suite.T(), err)
}

func (suite *SeekerRepositoryTestSuite) TestGetSeekerByID_Success() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingSeeker.Id).WillReturnRows(sqlmock.NewRows([]string{"id", "fullname", "username", "password", "email", "phone_number", "attiude_points", "status", "photo_profile", "room_id", "created_at", "updated_at"}).AddRow(mockingSeeker.Id, mockingSeeker.Fullname, mockingSeeker.Username, mockingSeeker.Password, mockingSeeker.Email, mockingSeeker.PhoneNumber, mockingSeeker.AtitudePoits, mockingSeeker.Status, mockingSeeker.PhotoProfile, mockingSeeker.RoomId, mockingSeeker.CreatedAt, mockingSeeker.UpdatedAt))

	actual, err := suite.repo.GetSeekerByID(mockingSeeker.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockingSeeker, actual)
}

func (suite *SeekerRepositoryTestSuite) TestGetSeekerByID_Failed() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingSeeker.Id).WillReturnError(errors.New("Get Seeker By ID Failed"))

	_, err := suite.repo.GetSeekerByID(mockingSeeker.Id)
	assert.Error(suite.T(), err)
}

func (suite *SeekerRepositoryTestSuite) TestUpdateAttitudePoints_Success() {
	suite.mockSql.ExpectExec("UPDATE seekers SET").WithArgs(mockingSeeker.AtitudePoits, time.Now(), mockingSeeker.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	err := suite.repo.UpdateAttitudePoints(mockingSeeker.Id, mockingSeeker.AtitudePoits)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *SeekerRepositoryTestSuite) TestUpdateSeeker_Success() {
	suite.mockSql.ExpectExec("UPDATE seekers SET").WithArgs(mockingSeeker.Fullname, mockingSeeker.Username, mockingSeeker.Password, mockingSeeker.Email, mockingSeeker.PhoneNumber, mockingSeeker.Status, mockingSeeker.PhotoProfile, time.Now(), mockingSeeker.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	err := suite.repo.UpdateSeeker(mockingSeeker.Id, mockingSeeker)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *SeekerRepositoryTestSuite) TestGetBySeeker_Success() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingSeeker.Username).WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "fullname", "phone_number", "status", "photo_profile", "created_at", "updated_at"}).AddRow(mockingSeeker.Id, mockingSeeker.Username, mockingSeeker.Password, mockingSeeker.Fullname, mockingSeeker.PhoneNumber, mockingSeeker.Status, mockingSeeker.PhotoProfile, mockingSeeker.CreatedAt, mockingSeeker.UpdatedAt))

	actual, err := suite.repo.GetBySeeker(mockingSeeker.Username)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockingSeeker, actual)
}
