package controller

import (
	"kostless-api/middleware"
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
	aM middleware.AuthMiddleware
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

func (s *SeekerContr) GetSeekerByID(ctx *gin.Context) {
    id := ctx.Param("id")
    seeker, err := s.ser.GetSeekerByID(id)
    if err != nil {
       util.SendErrRes(ctx, http.StatusInternalServerError, "id not found")
	   return
    }
    util.SendSingleResponse(ctx, "seeker found", seeker, http.StatusOK)
}

func (s *SeekerContr) GetAllSeekers(ctx *gin.Context) {
    seekers, err := s.ser.GetAllSeekers()
    if err != nil {
        util.SendErrRes(ctx, http.StatusInternalServerError, "seeker not found")
		return
    }
	util.SendSingleResponse(ctx, "seekers found", seekers, http.StatusOK)
}

func (s *SeekerContr) UpdateProfile(ctx *gin.Context) {
    id := ctx.Param("id")
    var seeker model.Seekers
    if err := ctx.ShouldBindJSON(&seeker); err != nil {
       util.SendErrRes(ctx, http.StatusBadRequest, "failed not found")
	   return
    }
    if err := s.ser.UpdateProfile(id, seeker); err != nil {
        util.SendErrRes(ctx, http.StatusInternalServerError, "seeker error")
		return
    }
    util.SendSingleResponse(ctx, "seekers updated", seeker, http.StatusOK)
}

func (s *SeekerContr) DeleteSeeker(ctx *gin.Context) {
    id := ctx.Param("id")
    if err := s.ser.DeleteSeeker(id); err != nil {
        util.SendErrRes(ctx, http.StatusInternalServerError, "seeker failed deleted")
        return
    }
	util.SendSingleResponse(ctx, "seekers deleted", id, http.StatusOK)
}

func (s *SeekerContr) UpdateAttitudePoints(ctx *gin.Context) {
    var request struct {
        ID            string `json:"id"`
        AttitudePoints int    `json:"attitudePoints"`
    }
    if err := ctx.ShouldBindJSON(&request); err != nil {
        util.SendErrRes(ctx, http.StatusBadRequest, "failed not found")
        return
    }
    if err := s.ser.UpdateAttitudePoints(request.ID, request.AttitudePoints); err != nil {
        util.SendErrRes(ctx, http.StatusInternalServerError, "attitude failed updated")
        return
    }
    util.SendSingleResponse(ctx, "seekers update attititude success", request, http.StatusOK)
}

// router
func (s *SeekerContr) Route() {
	router := s.rg.Group("/seekers")
	router.POST("/register", s.regisHandlerSeeker)
	router.POST("/login", s.loginSeeker)
	router.GET("/profile/:id", s.aM.CheckToken(), s.GetSeekerByID)
	router.PUT("/profile/:id", s.aM.CheckToken(), s.UpdateProfile)
	router.DELETE("/profile/:id", s.aM.CheckToken(), s.DeleteSeeker)
	router.GET("/profile/getall", s.GetAllSeekers)
	router.POST("/profile/update", s.UpdateAttitudePoints)
}

func NewSeekerContr(sS service.SeekerServ, rg *gin.RouterGroup, aM middleware.AuthMiddleware) *SeekerContr {
	return &SeekerContr{ser: sS, rg: rg, aM: aM}
}
