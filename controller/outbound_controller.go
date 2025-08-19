package controller

import (
	"net/http"
	"sistem-manajemen-gudang/model"
	"sistem-manajemen-gudang/service"
	"sistem-manajemen-gudang/middleware"
	"time"
	"github.com/gin-gonic/gin"
)
type OutboundController struct {
	service service.OutboundService
}

func NewOutboundController(s service.OutboundService) *OutboundController {
	return &OutboundController{service: s}
}

func (c *OutboundController) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/outbounds/add", middleware.StaffOrAdmin(), c.AddOutbound)
	rg.GET("/outbounds", middleware.StaffOrAdmin(),  c.ListOutbound)
	rg.GET("/outbounds/:id",middleware.StaffOrAdmin(),c.GetByID)
	rg.PUT("/outbounds/:id", middleware.StaffOrAdmin(),c.UpdateOutbound)
	rg.DELETE("/outbounds/:id",middleware.StaffOrAdmin(), c.DeleteOutbound)
	rg.GET("/report-outbounds", middleware.AdminOnly(), c.LaporanOutbound)

}

func (c *OutboundController) ListOutbound(ctx *gin.Context) {
	result, err := c.service.GetAll()
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil data"})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *OutboundController) AddOutbound(ctx *gin.Context) {
	var p model.Outbound
	if err := ctx.ShouldBindJSON(&p); err != nil || p.ProductID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	result, err := c.service.Create(p)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal simpan"})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *OutboundController) GetByID(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	inbound, err := c.service.GetByID(uri.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	ctx.JSON(http.StatusOK, inbound)
}
func (c *OutboundController) UpdateOutbound(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	var input struct {
		ProductID  uint      `json:"product_id" binding:"required"`
		Quantity   int       `json:"quantity" binding:"required"`
		SentAt time.Time `json:"sent_at" binding:"required"`
		Destination   string    `json:"destination" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid"})
		return
	}

	inbound := model.Outbound{
		ID:         uri.ID,
		ProductID:  input.ProductID,
		Quantity:   input.Quantity,
		SentAt:     input.SentAt,
		Destination: input.Destination,
	}

	updated, err := c.service.Update(inbound)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update"})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}
func (c *OutboundController) DeleteOutbound(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	err := c.service.Delete(uri.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal hapus"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
func (c *OutboundController) LaporanOutbound(ctx *gin.Context) {
	startDate := ctx.Query("start")
	endDate := ctx.Query("end")

	outbounds, err := c.service.GetLaporan(startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil laporan"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"laporan": outbounds,
	})
}