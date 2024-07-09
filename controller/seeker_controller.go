package controller

import (
	"kostless-api/model"
	"kostless-api/model/dto"
	"kostless-api/service"
	"kostless-api/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct.
type SeekerContr struct {
	ser service.SeekerServ
	rg  *gin.RouterGroup
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	util.SendSingleResponse(ctx, "Succesfully Created data", data, http.StatusOK)
}

//login handler seeker
func (s *SeekerContr) loginSeeker(ctx *gin.Context) {
	var payload dto.LoginDto
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"failed to parsing": err.Error()})
		return
	}
	resp , errors := s.ser.Login(payload)
	if errors != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"failed to parsing": errors.Error()})
		return
	}
	util.SendSingleResponse(ctx, "Success Login", resp, http.StatusOK)
}

// router
func (s *SeekerContr) Route() {
	router := s.rg.Group("/seekers")
	router.POST("/register", s.regisHandlerSeeker)
	router.POST("/login", s.loginSeeker)
}

func NewSeekerContr(sS service.SeekerServ, rg *gin.RouterGroup) *SeekerContr {
	return &SeekerContr{ser: sS, rg: rg}
}
