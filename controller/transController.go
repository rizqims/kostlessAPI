package controller

import (
	"kostless/model/req"
	"kostless/service"
	"kostless/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransController struct {
	rg      *gin.RouterGroup
	service service.TransService
}

func (t *TransController) CreateTransHandler(c *gin.Context) {
	var payload req.TransCreateReq
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

func (t *TransController) Route() {
	group := t.rg.Group("trans")
	group.POST("/create", t.CreateTransHandler)
	group.GET("/:id", t.GetTransByIDHandler)
	group.GET("/", t.GetTransHistoryHandler)
	group.GET("/paylaterlist", t.GetPaylaterListHandler)
}

func NewTransController(rg *gin.RouterGroup, service service.TransService) *TransController {
	return &TransController{
		rg:      rg,
		service: service,
	}
}
