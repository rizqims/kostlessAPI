package controller

import (
	"kostless-api/model"
	"kostless-api/service"
	"kostless-api/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct
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

// router
func (s *SeekerContr) Route() {
	router := s.rg.Group("/seekers")
	router.POST("/register", s.regisHandlerSeeker)
}

func NewSeekerContr(sS service.SeekerServ, rg *gin.RouterGroup) *SeekerContr {
	return &SeekerContr{ser: sS, rg: rg}
}
