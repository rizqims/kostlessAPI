package testing

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"kostless/model"
	"kostless/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type KostRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    repository.KosRepository
}

func (suite *KostRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = repository.NewKosRepository(suite.mockDb)
}

func TestKostRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(KostRepositoryTestSuite))
}

var mockingKos = model.Kos{
	ID:          "1",
	Name:        "Kost 1",
	Address:     "Jl. Kost 1",
	RoomCount:   3,
	Coordinate:  "Makassar",
	Description: "Kost 1 Indah",
	Rules:       "No Female in the room",
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
	Rooms: []model.Room{
		{
			ID:          "2",
			KosID:       "1",
			Name:        "Room 2",
			Type:        "Private Room",
			Description: "KM Dalam",
			Avail:       "true",
			Price:       100000,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	},
}

func (suite *KostRepositoryTestSuite) TestCreateKos_Success() {
	suite.mockSql.ExpectQuery("INSERT INTO kos").WithArgs(mockingKos.Name, mockingKos.Address, mockingKos.RoomCount, mockingKos.Coordinate, mockingKos.Description, mockingKos.Rules, time.Now(), time.Now()).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(mockingKos.ID, mockingKos.CreatedAt, mockingKos.UpdatedAt))

	actual, err := suite.repo.CreateKos(mockingKos)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockingKos.ID, actual.ID)
}

func (suite *KostRepositoryTestSuite) TestCreateKos_Failed() {
	suite.mockSql.ExpectQuery("INSERT INTO kos").WithArgs(mockingKos.Name, mockingKos.Address, mockingKos.RoomCount, mockingKos.Coordinate, mockingKos.Description, mockingKos.Rules, time.Now(), time.Now()).WillReturnError(errors.New("Insert New Kos Failed"))

	_, err := suite.repo.CreateKos(mockingKos)
	assert.Error(suite.T(), err)
}

func (suite *KostRepositoryTestSuite) TestUpdateKos_Success() {
	suite.mockSql.ExpectQuery("UPDATE kos SET").WithArgs(mockingKos.Name, mockingKos.Address, mockingKos.RoomCount, mockingKos.Coordinate, mockingKos.Description, mockingKos.Rules, time.Now(), mockingKos.ID).WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).AddRow(mockingKos.CreatedAt, mockingKos.UpdatedAt))

	actual, err := suite.repo.UpdateKos(mockingKos)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockingKos.ID, actual.ID)
}

func (suite *KostRepositoryTestSuite) TestUpdateKos_Failed() {
	suite.mockSql.ExpectQuery("UPDATE kos SET").WithArgs(mockingKos.Name, mockingKos.Address, mockingKos.RoomCount, mockingKos.Coordinate, mockingKos.Description, mockingKos.Rules, time.Now(), mockingKos.ID).WillReturnError(errors.New("Update Kos Failed"))

	_, err := suite.repo.UpdateKos(mockingKos)
	assert.Error(suite.T(), err)
}

func (suite *KostRepositoryTestSuite) TestDeleteKos_Success() {
	suite.mockSql.ExpectExec("DELETE FROM kos").WithArgs(mockingKos.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := suite.repo.DeleteKos(mockingKos.ID)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *KostRepositoryTestSuite) TestDeleteKos_Failed() {
	suite.mockSql.ExpectExec("DELETE FROM kos").WithArgs(mockingKos.ID).WillReturnError(errors.New("Delete Kos Failed"))

	err := suite.repo.DeleteKos(mockingKos.ID)
	assert.Error(suite.T(), err)
}

func (suite *KostRepositoryTestSuite) TestGetKosByID_Success() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingKos.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address", "room_count", "coordinate", "description", "rules", "created_at", "updated_at"}).AddRow(mockingKos.ID, mockingKos.Name, mockingKos.Address, mockingKos.RoomCount, mockingKos.Coordinate, mockingKos.Description, mockingKos.Rules, mockingKos.CreatedAt, mockingKos.UpdatedAt))

	for _, room := range mockingKos.Rooms {
		rows := sqlmock.NewRows([]string{"id", "kos_id", "name", "type", "description", "avail", "price", "created_at", "updated_at"}).AddRow(room.ID, room.KosID, room.Name, room.Type, room.Description, room.Avail, room.Price, room.CreatedAt, room.UpdatedAt)
		suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingKos.ID).WillReturnRows(rows)
	}

	actual, err := suite.repo.GetKosByID(mockingKos.ID)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockingKos.ID, actual.ID)
}

func (suite *KostRepositoryTestSuite) TestGetKosByID_Failed() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingKos.ID).WillReturnError(errors.New("Get Kos By ID Failed"))

	_, err := suite.repo.GetKosByID(mockingKos.ID)
	assert.Error(suite.T(), err)
}

func (suite *KostRepositoryTestSuite) TestGetKosByID_FailedRoom() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingKos.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address", "room_count", "coordinate", "description", "rules", "created_at", "updated_at"}).AddRow(mockingKos.ID, mockingKos.Name, mockingKos.Address, mockingKos.RoomCount, mockingKos.Coordinate, mockingKos.Description, mockingKos.Rules, mockingKos.CreatedAt, mockingKos.UpdatedAt))

	for _, room := range mockingKos.Rooms {
		suite.mockSql.ExpectQuery("SELECT").WithArgs(room.KosID).WillReturnError(errors.New("Get Room Failed"))
		_, err := suite.repo.GetKosByID(mockingKos.ID)
		assert.Error(suite.T(), err)
	}
}

func (suite *KostRepositoryTestSuite) TestGetKostByID_FailedScan() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingKos.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address", "room_count", "coordinate", "description", "rules", "created_at", "updated_at"}).AddRow(mockingKos.ID, mockingKos.Name, mockingKos.Address, mockingKos.RoomCount, mockingKos.Coordinate, mockingKos.Description, mockingKos.Rules, mockingKos.CreatedAt, mockingKos.UpdatedAt))

	for _, room := range mockingKos.Rooms {
		rows := sqlmock.NewRows([]string{"id", "kos_id", "name", "type", "description", "avail", "price", "created_at", "updated_at"}).AddRow(room.ID, room.KosID, room.Name, room.Type, room.Description, room.Avail, "Invalid Price", room.CreatedAt, room.UpdatedAt)
		suite.mockSql.ExpectQuery("SELECT").WithArgs(mockingKos.ID).WillReturnRows(rows)
	}

	_, err := suite.repo.GetKosByID(mockingKos.ID)
	assert.Error(suite.T(), err)
}
