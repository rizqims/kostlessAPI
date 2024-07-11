package testing

import (
	repomock "kostless/mocks/repo-mock"
	"kostless/model"
	"kostless/model/dto"
	"kostless/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type KostServiceTestSuite struct {
	suite.Suite
	repomock *repomock.KosRepoMock
	kS       service.KosService
}

func (suite *KostServiceTestSuite) SetupTest() {
	suite.repomock = new(repomock.KosRepoMock)
	suite.kS = service.NewKosService(suite.repomock)
}

func TestKostServiceTestSuite(t *testing.T) {
	suite.Run(t, new(KostServiceTestSuite))
}

var createKos = model.Kos{
	Name:        "Kost 1",
	Address:     "Jl. Kost 1",
	RoomCount:   1,
	Coordinate:  "Makassar",
	Description: "Kost 1 Indah",
	Rules:       "No Female in the room",
	Rooms:       nil,
}

var mockKos = model.Kos{
	ID:          "1",
	Name:        "Kost 1",
	Address:     "Jl. Kost 1",
	RoomCount:   1,
	Coordinate:  "Makassar",
	Description: "Kost 1 Indah",
	Rules:       "No Female in the room",
	Rooms:       nil,
}

func (suite *KostServiceTestSuite) TestCreateKos() {

	var mockingKosPayload = dto.KosRequest{
		Name:        "Kost 1",
		Address:     "Jl. Kost 1",
		RoomCount:   1,
		Coordinate:  "Makassar",
		Description: "Kost 1 Indah",
		Rules:       "No Female in the room",
	}

	suite.repomock.On("CreateKos", createKos).Return(createKos, nil)
	_, err := suite.kS.CreateKos(mockingKosPayload)

	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}

func (suite *KostServiceTestSuite) TestUpdateKos() {
	var mockingKosPayload = dto.KosRequest{
		Name:        "Kost 1",
		Address:     "Jl. Kost 1",
		RoomCount:   1,
		Coordinate:  "Makassar",
		Description: "Kost 1 Indah",
		Rules:       "No Female in the room",
	}

	suite.repomock.On("UpdateKos", mockKos).Return(mockKos, nil)
	_, err := suite.kS.UpdateKos(mockKos.ID, mockingKosPayload)

	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}

func (suite *KostServiceTestSuite) TestDeleteKos() {
	suite.repomock.On("DeleteKos", mockKos.ID).Return(nil)

	err := suite.kS.DeleteKos(mockKos.ID)
	assert.NoError(suite.T(), err)
}

func (suite *KostServiceTestSuite) TestGetKosByID() {
	suite.repomock.On("GetKosByID", mockKos.ID).Return(mockKos, nil)

	_, err := suite.kS.GetKosByID(mockKos.ID)

	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}
