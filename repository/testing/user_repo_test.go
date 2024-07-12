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

type UserRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    repository.UserRepo
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = repository.NewUserRepo(suite.mockDb)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

var mockingUser = model.User{
	Id:           "1",
	Fullname:     "User 1",
	Username:     "user1",
	Password:     "user1",
	Email:        "user@gmail.com",
	PhoneNumber:  "08123456789",
	PhotoProfile: "user1.jpg",
	CreatedAt:    time.Now(),
	UpdatedAt:    time.Now(),
}

func (suite *UserRepositoryTestSuite) TestGetUserById_Success() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingUser.Id).WillReturnRows(sqlmock.NewRows([]string{"id", "fullname", "username", "password", "email", "phone_number", "photo_profile", "created_at", "updated_at"}).AddRow(mockingUser.Id, mockingUser.Fullname, mockingUser.Username, mockingUser.Password, mockingUser.Email, mockingUser.PhoneNumber, mockingUser.PhotoProfile, mockingUser.CreatedAt, mockingUser.UpdatedAt))

	actual, err := suite.repo.GetUserById(mockingUser.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockingUser, actual)
}

func (suite *UserRepositoryTestSuite) TestGetUserById_Failed() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingUser.Id).WillReturnError(sql.ErrNoRows)

	_, err := suite.repo.GetUserById(mockingUser.Id)
	assert.Error(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestPutUpdateUserProf_Sucsess() {
	suite.mockSql.ExpectExec("UPDATE users SET").WithArgs(mockingUser.Fullname, mockingUser.Username, mockingUser.Password, mockingUser.Email, mockingUser.PhoneNumber, mockingUser.PhotoProfile, time.Now(), mockingUser.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	err := suite.repo.PutUpdateUserProf(mockingUser.Id, mockingUser)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestGetByUsername_Success() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingUser.Username).WillReturnRows(sqlmock.NewRows([]string{"id", "fullname", "username", "password", "email", "phone_number", "photo_profile", "created_at", "updated_at"}).AddRow(mockingUser.Id, mockingUser.Fullname, mockingUser.Username, mockingUser.Password, mockingUser.Email, mockingUser.PhoneNumber, mockingUser.PhotoProfile, mockingUser.CreatedAt, mockingUser.UpdatedAt))

	actual, err := suite.repo.GetByUsername(mockingUser.Username)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockingUser, actual)
}

func (suite *UserRepositoryTestSuite) TestGetByUsername_Failed() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingUser.Username).WillReturnError(errors.New("Get By Username Failed"))

	_, err := suite.repo.GetByUsername(mockingUser.Username)
	assert.Error(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestCreatedNewUser_Success() {
	mockingUserCreate := model.User{
		Id:           "1",
		Fullname:     "User 1",
		Username:     "user1",
		Email:        "user@gmail.com",
		PhoneNumber:  "08123456789",
		PhotoProfile: "user1.jpg",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	suite.mockSql.ExpectQuery("INSERT INTO users").WithArgs(mockingUser.Fullname, mockingUser.Username, mockingUser.Password, mockingUser.Email, mockingUser.PhoneNumber, mockingUser.PhotoProfile, time.Now(), time.Now()).WillReturnRows(sqlmock.NewRows([]string{"id", "fullname", "username", "email", "phone_number", "photo_profile", "created_at", "updated_at"}).AddRow(mockingUserCreate.Id, mockingUserCreate.Fullname, mockingUserCreate.Username, mockingUserCreate.Email, mockingUserCreate.PhoneNumber, mockingUserCreate.PhotoProfile, time.Now(), time.Now()))

	actual, err := suite.repo.CreatedNewUser(mockingUser)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockingUserCreate, actual)
}

func (suite *UserRepositoryTestSuite) TestCreatedNewUser_Failed() {
	suite.mockSql.ExpectQuery("INSERT INTO users").WithArgs(mockingUser.Fullname, mockingUser.Username, mockingUser.Password, mockingUser.Email, mockingUser.PhoneNumber, mockingUser.PhotoProfile, time.Now(), time.Now()).WillReturnError(errors.New("Created New User Failed"))

	_, err := suite.repo.CreatedNewUser(mockingUser)
	assert.Error(suite.T(), err)
}
