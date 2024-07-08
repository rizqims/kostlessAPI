package controller

import (
	"kostless-api/model"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	util.SendSingleResponse(ctx, "Succesfully Created data", data, http.StatusOK)
}

// router
func (u *UserContr) Route() {
	router := u.rg.Group("/users")
	router.POST("/register", u.regisHandler)
}

func NewUserContr(uS service.UserServ, rg *gin.RouterGroup) *UserContr{
	return &UserContr{ser: uS, rg: rg}
}
