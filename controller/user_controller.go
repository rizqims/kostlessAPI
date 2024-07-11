package controller

import (
	"fmt"
	"kostless/middleware"
	"kostless/model"
	"kostless/model/dto"
	"kostless/service"
	"kostless/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct
type UserContr struct {
	ser        service.UserServ
	serS service.SeekerServ
	rg         *gin.RouterGroup
	middleware middleware.AuthMiddleware
}

// register func
func (u *UserContr) regisHandler(ctx *gin.Context) {
	var User model.User
	if err := ctx.ShouldBindBodyWithJSON(&User); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"failed to register": err.Error()})
		return
	}

	data, err := u.ser.CreatedNewUser(User)
	if err != nil {
		fmt.Print("err ====", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	util.SendSingleResponse(ctx, http.StatusOK, "Succesfully Created data", data)
}

// login handler
func (u *UserContr) login(ctx *gin.Context) {
	var payload dto.LoginDto
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"failed to parsing": err.Error()})
		return
	}
	resp, errors := u.ser.Login(payload)
	if errors != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"failed to parsing": errors.Error()})
		return
	}
	util.SendSingleResponse(ctx, http.StatusOK, "Success Login", resp)
}

// get id
func (u *UserContr) getUserId(ctx *gin.Context) {
	id := ctx.Param("id")
	users, err := u.ser.GetUser(id)
	if err != nil {
		fmt.Print("err===", err)
		util.SendErrRes(ctx, http.StatusInternalServerError, "id not found")
		return
	}
	util.SendSingleResponse(ctx, http.StatusOK, "Id found", users)
}

func (u *UserContr) updateUser(ctx *gin.Context) {
	var updatedUser model.User
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		util.SendErrRes(ctx, http.StatusInternalServerError, "error updated")
		return
	}
	id := ctx.Param("id")
	err := u.ser.UpdateProfile(id, updatedUser)
	if err != nil {
		util.SendErrRes(ctx, http.StatusInternalServerError, "failed to updated")
		return
	}
	user, err := u.ser.GetUser(id)
	if err != nil {
		util.SendErrRes(ctx, http.StatusInternalServerError, "updated error")
		return
	}
	util.SendSingleResponse(ctx, http.StatusOK, "Success Updated", user)
}

func (u *UserContr) UpdateAttitudePoints(ctx *gin.Context) {
	seekerID := ctx.Param("id")
	var request struct {
		AttitudePoints int    `json:"attitudePoints"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		util.SendErrRes(ctx, http.StatusBadRequest, "failed not found")
		return
	}
	if err := u.serS.UpdateAttitudePoints(seekerID , request.AttitudePoints); err != nil {
		util.SendErrRes(ctx, http.StatusInternalServerError, "attitude failed updated")
		return
	}
	util.SendSingleResponse(ctx, http.StatusOK, "seekers update attititude success", request)

}

// router
func (u *UserContr) Route() {
	router := u.rg.Group("/users")
	router.POST("/register", u.regisHandler)
	router.POST("/login", u.login)
	router.GET("/profile/:id", u.middleware.CheckToken(), u.getUserId)
	router.PUT("/profile/:id", u.middleware.CheckToken(), u.updateUser)
	router.PUT("/seekers/attitude/:id", u.middleware.CheckToken(), u.UpdateAttitudePoints)
}

func NewUserContr(uS service.UserServ, sS service.SeekerServ, rg *gin.RouterGroup, aM middleware.AuthMiddleware) *UserContr {
	return &UserContr{ser: uS, serS: sS, rg: rg, middleware: aM}
}
