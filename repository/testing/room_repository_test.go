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

type RoomRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    repository.RoomRepository
}

func (suite *RoomRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = repository.NewRoomRepository(suite.mockDb)
}

func TestRoomRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RoomRepositoryTestSuite))
}

var mockingRoom = model.Room{
	ID:          "1",
	KosID:       "1",
	Name:        "Room 1",
	Type:        "Private Room",
	Description: "KM Dalam",
	Avail:       "true",
	Price:       100000,
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
}

func (suite *RoomRepositoryTestSuite) TestCreateRoom_Success() {
	suite.mockSql.ExpectQuery("INSERT INTO rooms").WithArgs(mockingRoom.KosID, mockingRoom.Name, mockingRoom.Type, mockingRoom.Description, mockingRoom.Avail, mockingRoom.Price, time.Now(), time.Now()).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(mockingRoom.ID, mockingRoom.CreatedAt, mockingRoom.UpdatedAt))

	actual, err := suite.repo.CreateRoom(mockingRoom)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockingRoom, actual)
}

func (suite *RoomRepositoryTestSuite) TestCreateRoom_Failed() {
	suite.mockSql.ExpectQuery("INSERT INTO rooms").WithArgs(mockingRoom.KosID, mockingRoom.Name, mockingRoom.Type, mockingRoom.Description, mockingRoom.Avail, mockingRoom.Price, time.Now(), time.Now()).WillReturnError(errors.New("Insert New Room Failed"))

	_, err := suite.repo.CreateRoom(mockingRoom)
	assert.Error(suite.T(), err)
}

func (suite *RoomRepositoryTestSuite) TestGetAllRoom_Success() {
	suite.mockSql.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "kos_id", "name", "type", "description", "avail", "price", "created_at", "updated_at"}).AddRow(mockingRoom.ID, mockingRoom.KosID, mockingRoom.Name, mockingRoom.Type, mockingRoom.Description, mockingRoom.Avail, mockingRoom.Price, mockingRoom.CreatedAt, mockingRoom.UpdatedAt))

	actual, err := suite.repo.GetAllRooms()
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), actual, 1)
	assert.Equal(suite.T(), mockingRoom, actual[0])
}

func (suite *RoomRepositoryTestSuite) TestGetAllRoom_Failed() {
	suite.mockSql.ExpectQuery("SELECT").WillReturnError(errors.New("Get All Room Failed"))

	_, err := suite.repo.GetAllRooms()
	assert.Error(suite.T(), err)
}

func (suite *RoomRepositoryTestSuite) TestGetAllRoom_FailedScan() {
	suite.mockSql.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "kos_id", "name", "type", "description", "avail", "price", "created_at", "updated_at"}).AddRow(mockingRoom.ID, mockingRoom.KosID, mockingRoom.Name, mockingRoom.Type, mockingRoom.Description, mockingRoom.Avail, "Invalid Price", mockingRoom.CreatedAt, mockingRoom.UpdatedAt))

	_, err := suite.repo.GetAllRooms()
	assert.Error(suite.T(), err)
}

func (suite *RoomRepositoryTestSuite) TestGetRoomByID_Success() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingKos.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "kos_id", "name", "type", "description", "avail", "price", "created_at", "updated_at"}).AddRow(mockingRoom.ID, mockingRoom.KosID, mockingRoom.Name, mockingRoom.Type, mockingRoom.Description, mockingRoom.Avail, mockingRoom.Price, mockingRoom.CreatedAt, mockingRoom.UpdatedAt))

	actual, err := suite.repo.GetRoomByID(mockingRoom.ID)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockingRoom, actual)
}

func (suite *RoomRepositoryTestSuite) TestGetRoomByID_Failed() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingKos.ID).WillReturnError(errors.New("Get Room By ID Failed"))

	_, err := suite.repo.GetRoomByID(mockingRoom.ID)
	assert.Error(suite.T(), err)
}

func (suite *RoomRepositoryTestSuite) TestGetRoomByAvailability_Success() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingRoom.Avail).WillReturnRows(sqlmock.NewRows([]string{"id", "kos_id", "name", "type", "description", "avail", "price", "created_at", "updated_at"}).AddRow(mockingRoom.ID, mockingRoom.KosID, mockingRoom.Name, mockingRoom.Type, mockingRoom.Description, mockingRoom.Avail, mockingRoom.Price, mockingRoom.CreatedAt, mockingRoom.UpdatedAt))

	actual, err := suite.repo.GetRoomByAvailability(mockingRoom.Avail)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), actual, 1)
	assert.Equal(suite.T(), mockingRoom, actual[0])
}

func (suite *RoomRepositoryTestSuite) TestGetRoomByAvailability_Failed() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingRoom.Avail).WillReturnError(errors.New("Get Room By Availability Failed"))

	_, err := suite.repo.GetRoomByAvailability(mockingRoom.Avail)
	assert.Error(suite.T(), err)
}

func (suite *RoomRepositoryTestSuite) TestGetRoomByAvailability_FailedScan() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingRoom.Avail).WillReturnRows(sqlmock.NewRows([]string{"id", "kos_id", "name", "type", "description", "avail", "price", "created_at", "updated_at"}).AddRow(mockingRoom.ID, mockingRoom.KosID, mockingRoom.Name, mockingRoom.Type, mockingRoom.Description, mockingRoom.Avail, "Invalid Price", mockingRoom.CreatedAt, mockingRoom.UpdatedAt))

	_, err := suite.repo.GetRoomByAvailability(mockingRoom.Avail)
	assert.Error(suite.T(), err)
}

func (suite *RoomRepositoryTestSuite) TestGetRoomByPriceLowerThanOrEqual_Success() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingRoom.Price).WillReturnRows(sqlmock.NewRows([]string{"id", "kos_id", "name", "type", "description", "avail", "price", "created_at", "updated_at"}).AddRow(mockingRoom.ID, mockingRoom.KosID, mockingRoom.Name, mockingRoom.Type, mockingRoom.Description, mockingRoom.Avail, mockingRoom.Price, mockingRoom.CreatedAt, mockingRoom.UpdatedAt))

	actual, err := suite.repo.GetRoomByPriceLowerThanOrEqual(mockingRoom.Price)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), actual, 1)
	assert.Equal(suite.T(), mockingRoom, actual[0])
}

func (suite *RoomRepositoryTestSuite) TestGetRoomByPriceLowerThanOrEqual_Failed() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingRoom.Price).WillReturnError(errors.New("Get Room By Price Failed"))

	_, err := suite.repo.GetRoomByPriceLowerThanOrEqual(mockingRoom.Price)
	assert.Error(suite.T(), err)
}

func (suite *RoomRepositoryTestSuite) TestGetRoomByPriceLowerThanOrEqual_FailedScan() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingRoom.Price).WillReturnRows(sqlmock.NewRows([]string{"id", "kos_id", "name", "type", "description", "avail", "price", "created_at", "updated_at"}).AddRow(mockingRoom.ID, mockingRoom.KosID, mockingRoom.Name, mockingRoom.Type, mockingRoom.Description, mockingRoom.Avail, "Invalid Price", mockingRoom.CreatedAt, mockingRoom.UpdatedAt))

	_, err := suite.repo.GetRoomByPriceLowerThanOrEqual(mockingRoom.Price)
	assert.Error(suite.T(), err)
}

func (suite *RoomRepositoryTestSuite) TestUpdateRoom_Success() {
	suite.mockSql.ExpectQuery("UPDATE rooms SET").WithArgs(mockingRoom.KosID, mockingRoom.Name, mockingRoom.Type, mockingRoom.Description, mockingRoom.Avail, mockingRoom.Price, time.Now(), mockingRoom.ID).WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).AddRow(mockingRoom.CreatedAt, mockingRoom.UpdatedAt))

	actual, err := suite.repo.UpdateRoom(mockingRoom)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockingRoom, actual)
}

func (suite *RoomRepositoryTestSuite) TestUpdateRoom_Failed() {
	suite.mockSql.ExpectQuery("UPDATE rooms SET").WillReturnError(errors.New("Update Room Failed"))

	_, err := suite.repo.UpdateRoom(mockingRoom)
	assert.Error(suite.T(), err)
}

func (suite *RoomRepositoryTestSuite) TestDeleteRoom_Success() {
	suite.mockSql.ExpectExec("DELETE FROM rooms").WithArgs(mockingRoom.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := suite.repo.DeleteRoom(mockingRoom.ID)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomRepositoryTestSuite) TestDeleteRoom_Failed() {
	suite.mockSql.ExpectExec("DELETE FROM rooms").WithArgs(mockingRoom.ID).WillReturnError(errors.New("Delete Room Failed"))

	err := suite.repo.DeleteRoom(mockingRoom.ID)
	assert.Error(suite.T(), err)
}
