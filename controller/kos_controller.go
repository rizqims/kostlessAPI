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
	group.PUT("/:id", k.updateKos)
	group.DELETE("/:id", k.deleteKos)
	group.GET("/:id", k.getKosByID)
}

func (k *KosController) createKos(ctx *gin.Context) {
	var request dto.KosRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		util.SendErrResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	kos, err := k.service.CreateKos(request)
	if err != nil {
		util.SendErrResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleResponse(ctx, http.StatusCreated, "Success", kos)
}

func (k *KosController) updateKos(ctx *gin.Context) {
	id := ctx.Param("id")
	var request dto.KosRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		util.SendErrResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	kos, err := k.service.UpdateKos(id, request)
	if err != nil {
		util.SendErrResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleResponse(ctx, http.StatusOK, "Success", kos)
}

func (k *KosController) deleteKos(ctx *gin.Context) {
	id := ctx.Param("id")

	err := k.service.DeleteKos(id)
	if err != nil {
		util.SendErrResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleResponse(ctx, http.StatusOK, "Success delete", nil)
}

func (k *KosController) getKosByID(ctx *gin.Context) {
	id := ctx.Param("id")

	kos, err := k.service.GetKosByID(id)
	if err != nil {
		util.SendErrResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleResponse(ctx, http.StatusOK, "Success", kos)
}
