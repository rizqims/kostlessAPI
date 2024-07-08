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
		ctx.JSON(http.StatusBadRequest, util.SingleRes{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	kos, err := k.service.CreateKos(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.SingleRes{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, util.SingleRes{
		Code:    http.StatusCreated,
		Message: "Success",
		Data:    kos,
	})
}
