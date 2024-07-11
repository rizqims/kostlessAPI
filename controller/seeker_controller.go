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

// struct.
type SeekerContr struct {
	ser service.SeekerServ
	rg  *gin.RouterGroup
	aM  middleware.AuthMiddleware
}

// register func
func (s *SeekerContr) regisHandlerSeeker(ctx *gin.Context) {
	var seeker model.Seekers
	if err := ctx.ShouldBindBodyWithJSON(&seeker); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"failed to register": err.Error()})
		return
	}

	data, err := s.ser.CreatedNewSeeker(seeker)
	if err != nil {
		fmt.Print("err====", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	util.SendSingleResponse(ctx, http.StatusOK, "Succesfully Created data", data)
}

// login handler seeker
func (s *SeekerContr) loginSeeker(ctx *gin.Context) {
	var payload dto.LoginDto
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"failed to parsing": err.Error()})
		return
	}
	resp, errors := s.ser.Login(payload)
	if errors != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"failed to parsing": errors.Error()})
		return
	}
	util.SendSingleResponse(ctx, http.StatusOK, "Success Login", resp)
}

func (s *SeekerContr) GetSeekerByID(ctx *gin.Context) {
	id := ctx.Param("id")
	seeker, err := s.ser.GetSeekerByID(id)
	if err != nil {
		fmt.Print("err====", err)
		util.SendErrRes(ctx, http.StatusInternalServerError, "id not found")
		return
	}
	util.SendSingleResponse(ctx, http.StatusOK, "seeker found", seeker)
}

func (s *SeekerContr) GetAllSeekers(ctx *gin.Context) {
	seekers, err := s.ser.GetAllSeekers()
	if err != nil {
		util.SendErrRes(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	for _, seeker := range seekers {
		seeker.Password = ""
	}
	util.SendSingleResponse(ctx, http.StatusOK, "seekers found", seekers)
}

func (s *SeekerContr) UpdateProfile(ctx *gin.Context) {
	var updateSeeker model.Seekers
	if err := ctx.ShouldBindJSON(&updateSeeker); err != nil {
		util.SendErrRes(ctx, http.StatusBadRequest, "failed not found")
		return
	}
	id := ctx.Param("id")
	if err := s.ser.UpdateProfile(id, updateSeeker); err != nil {
		util.SendErrRes(ctx, http.StatusInternalServerError, "seeker error")
		return
	}
	seeker, err := s.ser.GetSeekerByID(id)
	if err != nil {
		util.SendErrRes(ctx, http.StatusInternalServerError, "updated error")
		return
	}
	util.SendSingleResponse(ctx, http.StatusOK, "seekers updated", seeker)
}

func (s *SeekerContr) DeleteSeeker(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := s.ser.DeleteSeeker(id); err != nil {
		util.SendErrRes(ctx, http.StatusInternalServerError, "seeker failed deleted")
		return
	}
	util.SendSingleResponse(ctx, http.StatusOK, "seekers deleted", id)
}



// router
func (s *SeekerContr) Route() {
	router := s.rg.Group("/seekers")
	{
	router.POST("/register", s.regisHandlerSeeker)
	router.POST("/login", s.loginSeeker)
	router.GET("/profile/:id", s.aM.CheckToken(), s.GetSeekerByID)
	router.PUT("/profile/:id", s.aM.CheckToken(), s.UpdateProfile)
	router.DELETE("/profile/:id", s.aM.CheckToken(), s.DeleteSeeker)
	router.GET("/profile/getall", s.GetAllSeekers)
	}

}

func NewSeekerContr(sS service.SeekerServ, rg *gin.RouterGroup, aM middleware.AuthMiddleware) *SeekerContr {
	return &SeekerContr{ser: sS, rg: rg, aM: aM}
}
