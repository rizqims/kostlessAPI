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
		util.SendErrRes(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := t.service.CreateTrans(payload)
	if err != nil {
		util.SendErrRes(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleRes(c, http.StatusCreated, "success create", response)
}

func (t *TransController) Route() {
	group := t.rg.Group("trans")
	group.POST("/create", t.CreateTransHandler)
}

func NewTransController(rg *gin.RouterGroup, service service.TransService) *TransController {
	return &TransController{
		rg:      rg,
		service: service,
	}
}
