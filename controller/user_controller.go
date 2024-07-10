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
	util.SendSingleResponse(ctx, "Succesfully Created data", data, http.StatusOK)
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
	util.SendSingleResponse(ctx, "Success Login", resp, http.StatusOK)
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
	util.SendSingleResponse(ctx, "Id found", users, http.StatusOK)
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
	util.SendSingleResponse(ctx, "Success Updated", user, http.StatusOK)
}

// router
func (u *UserContr) Route() {
	router := u.rg.Group("/users")
	router.POST("/register", u.regisHandler)
	router.POST("/login", u.login)
	router.GET("/profile/:id", u.middleware.CheckToken(), u.getUserId)
	router.PUT("/profile/id", u.middleware.CheckToken(), u.updateUser)
}

func NewUserContr(uS service.UserServ, rg *gin.RouterGroup, aM middleware.AuthMiddleware) *UserContr {
	return &UserContr{ser: uS, rg: rg, middleware: aM}
}
