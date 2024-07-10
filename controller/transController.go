package controller

import (
	"kostless/model/dto"
	"kostless/service"
	"kostless/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TransController struct {
	rg      *gin.RouterGroup
	service service.TransService
}

func (t *TransController) CreateTransHandler(c *gin.Context) {
	var payload dto.TransCreateReq
	err := c.ShouldBindBodyWithJSON(&payload)
	if err != nil {
		util.SendErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := t.service.CreateTrans(payload)
	if err != nil {
		util.SendErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleResponse(c, http.StatusCreated, "success create", response)
}

func (t *TransController) GetTransByIDHandler(c *gin.Context) {
	id := c.Param("id")

	response, err := t.service.GetTransByID(id)
	if err != nil {
		util.SendErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	util.SendSingleResponse(c, http.StatusOK, "success retrieve", response)
}

func (t *TransController) GetTransHistoryHandler(c *gin.Context) {
	response, err := t.service.GetTransHistory()
	if err != nil {
		util.SendErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleResponse(c, http.StatusOK, "success retrieve", response)
}

func (t *TransController) GetPaylaterListHandler(c *gin.Context) {
	response, err := t.service.GetPaylaterList()
	if err != nil {
		util.SendErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleResponse(c, http.StatusOK, "success retrieve paylater list", response)
}

func (t *TransController) GetTransByMonth(c *gin.Context){
	var month, year string
	month = c.DefaultQuery("month", time.Now().Format(`January`))
	year = c.DefaultQuery("year", time.Now().Format(`2006`))

	response, err := t.service.GetTransByMonth(month, year)
	if err != nil {
		util.SendErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleResponse(c, http.StatusOK, "success retrieve getbymonth", response)
}

func (t *TransController) UpdatePaylaterHandler(c *gin.Context){
	var payload dto.UpdatePaylaterReq
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		util.SendErrResponse(c, http.StatusBadRequest, err.Error())
	}

	response, err := t.service.UpdatePaylater(payload)
	if err != nil {
		util.SendErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	util.SendSingleResponse(c, http.StatusOK, "success update paylater", response)
}

func (t *TransController) Route() {
	group := t.rg.Group("trans")
	group.POST("/create", t.CreateTransHandler)
	group.GET("/:id", t.GetTransByIDHandler)
	group.GET("/", t.GetTransHistoryHandler)
	group.GET("/paylaterlist", t.GetPaylaterListHandler)
	group.GET("/getbymonth", t.GetTransByMonth)
	group.PUT("/updatepaylater", t.UpdatePaylaterHandler)
}

func NewTransController(rg *gin.RouterGroup, service service.TransService) *TransController {
	return &TransController{
		rg:      rg,
		service: service,
	}
}
