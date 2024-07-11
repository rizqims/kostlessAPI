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
	group.GET("/budget/", r.getRoomByPriceLowerThanOrEqual)
	group.GET("/rooms", r.getAllRooms)
	group.PUT("/:id", r.updateRoom)
	group.DELETE("/:id", r.deleteRoom)
}

func (r *RoomController) createRoom(ctx *gin.Context) {
	var request dto.RoomRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		util.SendErrResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	room, err := r.service.CreateRoom(request)
	if err != nil {
		util.SendErrResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleResponse(ctx, http.StatusCreated, "Success", room)
}

func (r *RoomController) updateRoom(ctx *gin.Context) {
	var request dto.RoomRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		util.SendErrResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("id")
	room, err := r.service.GetRoomByID(id)
	if err != nil {
		util.SendErrResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	room.KosID = request.KosID
	room.Name = request.Name
	room.Type = request.Type
	room.Description = request.Description
	room.Avail = request.Avail
	room.Price = request.Price

	updatedRoom, err := r.service.UpdateRoom(room)
	if err != nil {
		util.SendErrResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleResponse(ctx, http.StatusOK, "Success", updatedRoom)
}

func (r *RoomController) getAllRooms(ctx *gin.Context) {
	rooms, err := r.service.GetAllRooms()
	if err != nil {
		util.SendErrResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleResponse(ctx, http.StatusOK, "Success", rooms)
}

func (r *RoomController) getRoomByID(ctx *gin.Context) {
	id := ctx.Param("id")
	room, err := r.service.GetRoomByID(id)
	if err != nil {
		util.SendErrResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleResponse(ctx, http.StatusOK, "Success", room)
}

func (r *RoomController) getRoomByAvailability(ctx *gin.Context) {
	availability := ctx.Param("avail")
	if availability != "open" && availability != "occupied" {
		util.SendErrResponse(ctx, http.StatusBadRequest, "Availability must be 'open' or 'occupied'")
		return
	}

	rooms, err := r.service.GetRoomByAvailability(availability)
	if err != nil {
		util.SendErrResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleResponse(ctx, http.StatusOK, "Success", rooms)
}

func (r *RoomController) getRoomByPriceLowerThanOrEqual(ctx *gin.Context) {
	budget := ctx.Query("budget")
	rooms, err := r.service.GetRoomByPriceLowerThanOrEqual(budget)
	if err != nil {
		util.SendErrResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleResponse(ctx, http.StatusOK, "Success", rooms)
}

func (r *RoomController) deleteRoom(ctx *gin.Context) {
	id := ctx.Param("id")
	err := r.service.DeleteRoom(id)
	if err != nil {
		util.SendErrResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendSingleResponse(ctx, http.StatusOK, "Success delete", nil)
}
