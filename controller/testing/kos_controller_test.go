package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"kostless/controller"
	middlemock "kostless/mocks/middleware-mock"
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

type KosControllerTestSuite struct {
	suite.Suite
	serviceMock    *servicemock.KostServiceMock
	middlewareMock *middlemock.AuthMiddlewareMock
	router         *gin.Engine
	kosController  *controller.KosController
}

func (suite *KosControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.serviceMock = new(servicemock.KostServiceMock)
	suite.middlewareMock = new(middlemock.AuthMiddlewareMock)
	suite.router = gin.New()
	suite.kosController = controller.NewKosController(suite.serviceMock, suite.router.Group("/"), suite.middlewareMock)
	suite.kosController.Route()
}

func TestKosControllerTestSuite(t *testing.T) {
	suite.Run(t, new(KosControllerTestSuite))
}

func (suite *KosControllerTestSuite) TestUpdateKos_Success() {
	mockID := "1"
	mockRequest := dto.KosRequest{
		Name:        "Updated Kos",
		Address:     "Updated Address",
		RoomCount:   20,
		Coordinate:  "Updated Coordinate",
		Description: "Updated Description",
		Rules:       "Updated Rules",
	}
	mockResponse := model.Kos{
		ID:          "1",
		Name:        "Updated Kos",
		Address:     "Updated Address",
		RoomCount:   20,
		Coordinate:  "Updated Coordinate",
		Description: "Updated Description",
		Rules:       "Updated Rules",
	}

	suite.serviceMock.On("UpdateKos", mockID, mockRequest).Return(mockResponse, nil)

	body, _ := json.Marshal(mockRequest)
	req, _ := http.NewRequest(http.MethodPut, "/kos/"+mockID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.serviceMock.AssertExpectations(suite.T())
}

func (suite *KosControllerTestSuite) TestUpdateKos_InvalidRequestBody() {
	req, _ := http.NewRequest(http.MethodPut, "/kos/1", bytes.NewBuffer([]byte(`invalid json`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *KosControllerTestSuite) TestUpdateKos_ServiceError() {
	mockID := "1"
	mockRequest := dto.KosRequest{
		Name:        "Updated Kos",
		Address:     "Updated Address",
		RoomCount:   20,
		Coordinate:  "Updated Coordinate",
		Description: "Updated Description",
		Rules:       "Updated Rules",
	}
	suite.serviceMock.On("UpdateKos", mockID, mockRequest).Return(model.Kos{}, errors.New("service error"))

	body, _ := json.Marshal(mockRequest)
	req, _ := http.NewRequest(http.MethodPut, "/kos/"+mockID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	suite.serviceMock.AssertExpectations(suite.T())
}

func (suite *KosControllerTestSuite) TestDeleteKos_Success() {
	mockID := "1"
	suite.serviceMock.On("DeleteKos", mockID).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/kos/"+mockID, nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.serviceMock.AssertExpectations(suite.T())
}

func (suite *KosControllerTestSuite) TestDeleteKos_ServiceError() {
	mockID := "1"
	suite.serviceMock.On("DeleteKos", mockID).Return(errors.New("service error"))

	req, _ := http.NewRequest(http.MethodDelete, "/kos/"+mockID, nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	suite.serviceMock.AssertExpectations(suite.T())
}

func (suite *KosControllerTestSuite) TestGetKosByID_Success() {
	mockID := "1"
	mockResponse := model.Kos{
		ID:          "1",
		Name:        "Test Kos",
		Address:     "Test Address",
		RoomCount:   10,
		Coordinate:  "Test Coordinate",
		Description: "Test Description",
		Rules:       "Test Rules",
	}

	suite.serviceMock.On("GetKosByID", mockID).Return(mockResponse, nil)

	req, _ := http.NewRequest(http.MethodGet, "/kos/"+mockID, nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.serviceMock.AssertExpectations(suite.T())
}

func (suite *KosControllerTestSuite) TestGetKosByID_ServiceError() {
	mockID := "1"
	suite.serviceMock.On("GetKosByID", mockID).Return(model.Kos{}, errors.New("service error"))

	req, _ := http.NewRequest(http.MethodGet, "/kos/"+mockID, nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	suite.serviceMock.AssertExpectations(suite.T())
}
