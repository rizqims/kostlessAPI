package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"kostless/controller"
	servicemock "kostless/mocks/service-mock"
	"kostless/model"
	"kostless/model/dto"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RoomControllerTestSuite struct {
	suite.Suite
	serviceMock    *servicemock.RoomServiceMock
	router         *gin.Engine
	roomController *controller.RoomController
}

func (suite *RoomControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.serviceMock = new(servicemock.RoomServiceMock)
	suite.router = gin.New()
	suite.roomController = controller.NewRoomController(suite.serviceMock, suite.router.Group("/"))
	suite.roomController.Route()
}

func TestRoomControllerTestSuite(t *testing.T) {
	suite.Run(t, new(RoomControllerTestSuite))
}

func (suite *RoomControllerTestSuite) TestCreateRoom_Success() {
	mockRequest := dto.RoomRequest{
		KosID:       "1",
		Name:        "Room A",
		Type:        "Single",
		Description: "Nice room",
		Avail:       "open",
		Price:       1000,
	}
	mockResponse := model.Room{
		ID:          "1",
		KosID:       "1",
		Name:        "Room A",
		Type:        "Single",
		Description: "Nice room",
		Avail:       "open",
		Price:       1000,
	}

	suite.serviceMock.On("CreateRoom", mockRequest).Return(mockResponse, nil)

	body, _ := json.Marshal(mockRequest)
	req, _ := http.NewRequest(http.MethodPost, "/room/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	suite.serviceMock.AssertExpectations(suite.T())
}

func (suite *RoomControllerTestSuite) TestCreateRoom_InvalidRequestBody() {
	req, _ := http.NewRequest(http.MethodPost, "/room/", bytes.NewBuffer([]byte(`invalid json`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *RoomControllerTestSuite) TestCreateRoom_ServiceError() {
	mockRequest := dto.RoomRequest{
		KosID:       "1",
		Name:        "Room A",
		Type:        "Single",
		Description: "Nice room",
		Avail:       "open",
		Price:       1000,
	}

	suite.serviceMock.On("CreateRoom", mockRequest).Return(model.Room{}, errors.New("service error"))

	body, _ := json.Marshal(mockRequest)
	req, _ := http.NewRequest(http.MethodPost, "/room/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	suite.serviceMock.AssertExpectations(suite.T())
}

func (suite *RoomControllerTestSuite) TestUpdateRoom_Success() {
	mockID := "1"
	mockRequest := dto.RoomRequest{
		KosID:       "1",
		Name:        "Updated Room A",
		Type:        "Double",
		Description: "Updated nice room",
		Avail:       "occupied",
		Price:       1200,
	}
	mockRoom := model.Room{
		ID:          "1",
		KosID:       "1",
		Name:        "Room A",
		Type:        "Single",
		Description: "Nice room",
		Avail:       "open",
		Price:       1000,
	}
	mockUpdatedRoom := model.Room{
		ID:          "1",
		KosID:       "1",
		Name:        "Updated Room A",
		Type:        "Double",
		Description: "Updated nice room",
		Avail:       "occupied",
		Price:       1200,
	}

	suite.serviceMock.On("GetRoomByID", mockID).Return(mockRoom, nil)
	suite.serviceMock.On("UpdateRoom", mockUpdatedRoom).Return(mockUpdatedRoom, nil)

	body, _ := json.Marshal(mockRequest)
	req, _ := http.NewRequest(http.MethodPut, "/room/"+mockID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.serviceMock.AssertExpectations(suite.T())
}

func (suite *RoomControllerTestSuite) TestUpdateRoom_InvalidRequestBody() {
	req, _ := http.NewRequest(http.MethodPut, "/room/1", bytes.NewBuffer([]byte(`invalid json`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *RoomControllerTestSuite) TestUpdateRoom_ServiceError() {
	mockID := "1"
	mockRequest := dto.RoomRequest{
		KosID:       "1",
		Name:        "Updated Room A",
		Type:        "Double",
		Description: "Updated nice room",
		Avail:       "occupied",
		Price:       1200,
	}
	mockRoom := model.Room{
		ID:          "1",
		KosID:       "1",
		Name:        "Room A",
		Type:        "Single",
		Description: "Nice room",
		Avail:       "open",
		Price:       1000,
	}
	mockUpdatedRoom := model.Room{
		ID:          "1",
		KosID:       "1",
		Name:        "Updated Room A",
		Type:        "Double",
		Description: "Updated nice room",
		Avail:       "occupied",
		Price:       1200,
	}

	suite.serviceMock.On("GetRoomByID", mockID).Return(mockRoom, nil)
	suite.serviceMock.On("UpdateRoom", mockUpdatedRoom).Return(model.Room{}, errors.New("service error"))

	body, _ := json.Marshal(mockRequest)
	req, _ := http.NewRequest(http.MethodPut, "/room/"+mockID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	suite.serviceMock.AssertExpectations(suite.T())
}

func (suite *RoomControllerTestSuite) TestGetAllRooms_Success() {
	mockRooms := []model.Room{
		{
			ID:          "1",
			KosID:       "1",
			Name:        "Room A",
			Type:        "Single",
			Description: "Nice room",
			Avail:       "open",
			Price:       1000,
		},
		{
			ID:          "2",
			KosID:       "1",
			Name:        "Room B",
			Type:        "Double",
			Description: "Another nice room",
			Avail:       "occupied",
			Price:       1500,
		},
	}

	suite.serviceMock.On("GetAllRooms").Return(mockRooms, nil)

	req, _ := http.NewRequest(http.MethodGet, "/room/rooms", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.serviceMock.AssertExpectations(suite.T())
}

func (suite *RoomControllerTestSuite) TestGetAllRooms_ServiceError() {
	suite.serviceMock.On("GetAllRooms").Return(nil, errors.New("service error"))

	req, _ := http.NewRequest(http.MethodGet, "/room/rooms", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	suite.serviceMock.AssertExpectations(suite.T())
}

func (suite *RoomControllerTestSuite) TestGetRoomByID_Success() {
	mockID := "1"
	mockRoom := model.Room{
		ID:          "1",
		KosID:       "1",
		Name:        "Room A",
		Type:        "Single",
		Description: "Nice room",
		Avail:       "open",
		Price:       1000,
	}

	suite.serviceMock.On("GetRoomByID", mockID).Return(mockRoom, nil)

	req, _ := http.NewRequest(http.MethodGet, "/room/"+mockID, nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.serviceMock.AssertExpectations(suite.T())
}

func (suite *RoomControllerTestSuite) TestGetRoomByID_ServiceError() {
	mockID := "1"
	suite.serviceMock.On("GetRoomByID", mockID).Return(model.Room{}, errors.New("service error"))

	req, _ := http.NewRequest(http.MethodGet, "/room/"+mockID, nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	suite.serviceMock.AssertExpectations(suite.T())
}

func (suite *RoomControllerTestSuite) TestDeleteRoom_Success() {
	mockID := "1"
	suite.serviceMock.On("DeleteRoom", mockID).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/room/"+mockID, nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.serviceMock.AssertExpectations(suite.T())
}

func (suite *RoomControllerTestSuite) TestDeleteRoom_ServiceError() {
	mockID := "1"
	suite.serviceMock.On("DeleteRoom", mockID).Return(errors.New("service error"))

	req, _ := http.NewRequest(http.MethodDelete, "/room/"+mockID, nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	suite.serviceMock.AssertExpectations(suite.T())
}
