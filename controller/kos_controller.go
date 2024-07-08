package controller

import (
	"kostless/model/dto"
	"kostless/service"
	"kostless/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KosController struct {
	service service.KosService
	rg      *gin.RouterGroup
}

func NewKosController(service service.KosService, rg *gin.RouterGroup) *KosController {
	return &KosController{service, rg}
}

func (k *KosController) Route() {
	group := k.rg.Group("/kos")
	group.POST("/", k.createKos)
}

func (k *KosController) createKos(ctx *gin.Context) {
	var request dto.KosRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		util.SendErrRes(ctx, http.StatusBadRequest, err.Error())
		return
	}

	kos, err := k.service.CreateKos(request)
	if err != nil {
		util.SendErrRes(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleRes(ctx, http.StatusCreated, "Success", kos)
}
