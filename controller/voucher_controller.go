package controller

import (
	"kostless/model/dto"
	"kostless/service"
	"kostless/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VoucherController struct {
	rg      *gin.RouterGroup
	service service.VoucherService
}

func (v *VoucherController) CreateVoucherHandler(c *gin.Context) {
	var payload dto.CreateVoucherReq
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		util.SendErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := v.service.CreateVoucher(payload)
	if err != nil {
		util.SendErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	util.SendSingleResponse(c, http.StatusCreated, "success creating voucher", response)
}

func (v *VoucherController) DeleteExpiredVoucherHandler(c *gin.Context) {
	err := v.service.DeleteExpiredVoucher()
	if err != nil {
		util.SendErrResponse(c, http.StatusInternalServerError, "delete voucher failed")
		return
	}

	util.SendSingleResponse(c, http.StatusOK, "success delete expired voucher", 0)
}

func (v *VoucherController) GetAllVoucherHandler(c *gin.Context) {
	response, err := v.service.GetAllVoucher()
	if err != nil {
		util.SendErrResponse(c, http.StatusInternalServerError, "failed when retrieving")
		return
	}
	util.SendSingleResponse(c, http.StatusOK, "success get all voucher", response)
}

func (v *VoucherController) GetVoucherBySeekerIDHandler(c *gin.Context){
	id := c.Param("id")
	response, err := v.service.GetVoucherBySeekerID(id)
	if err != nil {
		util.SendErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	util.SendSingleResponse(c, http.StatusOK, "success get all voucher", response)
}
func (t *VoucherController) Route() {
	group := t.rg.Group("voucher")
	group.POST("/create", t.CreateVoucherHandler) //
	group.DELETE("/", t.DeleteExpiredVoucherHandler) //
	group.GET("/", t.GetAllVoucherHandler) //
	group.GET("/:id", t.GetVoucherBySeekerIDHandler) //
}

func NewVoucherController(service service.VoucherService, rg *gin.RouterGroup) *VoucherController {
	return &VoucherController{
		rg:      rg,
		service: service,
	}
}
