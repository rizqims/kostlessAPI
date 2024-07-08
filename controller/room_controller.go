package controller

import (
	"kostless/model/dto"
	"kostless/service"
	"kostless/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	service service.RoomService
	rg      *gin.RouterGroup
}

func NewRoomController(service service.RoomService, rg *gin.RouterGroup) *RoomController {
	return &RoomController{service, rg}
}

func (r *RoomController) Route() {
	group := r.rg.Group("/room")
	group.POST("/", r.createRoom)
	group.GET("/:id", r.getRoomByID)
	group.GET("/availability/:avail", r.getRoomByAvailability)
	group.GET("/", r.getRoomByPriceLowerThan)
}

func (r *RoomController) createRoom(ctx *gin.Context) {
	var request dto.RoomRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		util.SendErrRes(ctx, http.StatusBadRequest, err.Error())
		return
	}

	room, err := r.service.CreateRoom(request)
	if err != nil {
		util.SendErrRes(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleRes(ctx, http.StatusCreated, "Success", room)
}

func (r *RoomController) getRoomByID(ctx *gin.Context) {
	id := ctx.Param("id")
	room, err := r.service.GetRoomByID(id)
	if err != nil {
		util.SendErrRes(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleRes(ctx, http.StatusOK, "Success", room)
}

func (r *RoomController) getRoomByAvailability(ctx *gin.Context) {
	availability := ctx.Param("avail")
	rooms, err := r.service.GetRoomByAvailability(availability)
	if err != nil {
		util.SendErrRes(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleRes(ctx, http.StatusOK, "Success", rooms)
}

func (r *RoomController) getRoomByPriceLowerThan(ctx *gin.Context) {
	budget := ctx.Query("budget")
	rooms, err := r.service.GetRoomByPriceLowerThan(budget)
	if err != nil {
		util.SendErrRes(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleRes(ctx, http.StatusOK, "Success", rooms)
}
