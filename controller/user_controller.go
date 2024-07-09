package controller

import (
	"fmt"
	"kostless-api/model"
	"kostless-api/model/dto"
	"kostless-api/service"
	"kostless-api/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct
type UserContr struct {
	ser service.UserServ
	rg  *gin.RouterGroup
}

// register func
func (u *UserContr) regisHandler(ctx *gin.Context) {
	var owner model.User
	if err := ctx.ShouldBindBodyWithJSON(&owner); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"failed to register": err.Error()})
		return
	}

	data, err := u.ser.CreatedNewUser(owner)
	if err != nil {
		fmt.Print("err ====", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	util.SendSingleResponse(ctx, "Succesfully Created data", data, http.StatusOK)
}

//login handler
func (u *UserContr) login(ctx *gin.Context) {
	var payload dto.LoginDto
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"failed to parsing": err.Error()})
		return
	}
	resp , errors := u.ser.Login(payload)
	if errors != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"failed to parsing": errors.Error()})
		return
	}
	util.SendSingleResponse(ctx, "Success Login", resp, http.StatusOK)
}

// router
func (u *UserContr) Route() {
	router := u.rg.Group("/users")
	router.POST("/register", u.regisHandler)
	router.POST("/login", u.login)
}

func NewUserContr(uS service.UserServ, rg *gin.RouterGroup) *UserContr{
	return &UserContr{ser: uS, rg: rg}
}
