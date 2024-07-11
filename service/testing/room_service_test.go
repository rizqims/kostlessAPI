package testing

import (
	"errors"
	repomock "kostless/mocks/repo-mock"
	"kostless/model"
	"kostless/model/dto"
	"kostless/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RoomServiceTestSuite struct {
	suite.Suite
	repomock *repomock.RoomRepoMock
	rS       service.RoomService
}

func (suite *RoomServiceTestSuite) SetupTest() {
	suite.repomock = new(repomock.RoomRepoMock)
	suite.rS = service.NewRoomService(suite.repomock)
}

func TestRoomServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RoomServiceTestSuite))
}

var createRoom = model.Room{
	KosID:       "1",
	Name:        "Room 1",
	Type:        "Private",
	Description: "Room 1 Indah",
	Avail:       "Available",
	Price:       100000,
}

func (suite *RoomServiceTestSuite) TestCreateRoom() {
	var mockingRoomPayload = dto.RoomRequest{
		KosID:       "1",
		Name:        "Room 1",
		Type:        "Private",
		Description: "Room 1 Indah",
		Avail:       "Available",
		Price:       100000,
	}

	suite.repomock.On("CreateRoom", createRoom).Return(createRoom, nil)
	_, err := suite.rS.CreateRoom(mockingRoomPayload)

	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}

func (suite *RoomServiceTestSuite) TestUpdateRoom() {
	suite.repomock.On("UpdateRoom", createRoom).Return(createRoom, nil)

	_, err := suite.rS.UpdateRoom(createRoom)
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}

func (suite *RoomServiceTestSuite) TestGetAllRooms() {
	suite.repomock.On("GetAllRooms").Return([]model.Room{createRoom}, nil)
	_, err := suite.rS.GetAllRooms()
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}

func (suite *RoomServiceTestSuite) TestGetRoomByID() {
	suite.repomock.On("GetRoomByID", "1").Return(createRoom, nil)
	_, err := suite.rS.GetRoomByID("1")
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}

func (suite *RoomServiceTestSuite) TestGetRoomByAvailability() {
	suite.repomock.On("GetRoomByAvailability", "Available").Return([]model.Room{createRoom}, nil)
	_, err := suite.rS.GetRoomByAvailability("Available")
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}

func (suite *RoomServiceTestSuite) TestGetRoomByPriceLowerThanOrEqual() {
	suite.repomock.On("GetRoomByPriceLowerThanOrEqual", 100000).Return([]model.Room{createRoom}, nil)
	_, err := suite.rS.GetRoomByPriceLowerThanOrEqual("100000")
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}

func (suite *RoomServiceTestSuite) TestGetRoomByPriceLowerThanOrEqual_Failed() {
	suite.repomock.On("GetRoomByPriceLowerThanOrEqual", 100000).Return([]model.Room{}, errors.New("error"))
	_, err := suite.rS.GetRoomByPriceLowerThanOrEqual("")
	assert.Error(suite.T(), err)
}

func (suite *RoomServiceTestSuite) TestDeleteRoom() {
	suite.repomock.On("DeleteRoom", "1").Return(nil)
	err := suite.rS.DeleteRoom("1")
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}
