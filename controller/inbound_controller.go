package controller

import (
	"net/http"
	"sistem-manajemen-gudang/middleware"
	"sistem-manajemen-gudang/model"
	"sistem-manajemen-gudang/service"
	"time"
	"github.com/gin-gonic/gin"
)

type InboundController struct {
	service service.InboundService
}

func NewInboundController(h service.InboundService) *InboundController {
	return &InboundController{service: h}
}

func (c *InboundController) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/inbounds/add", middleware.StaffOrAdmin(), c.AddInbound)
	rg.GET("/inbounds", middleware.StaffOrAdmin(), c.ListInbound)
	rg.GET("/inbounds/:id", middleware.StaffOrAdmin(), c.GetByID)
	rg.PUT("/inbounds/:id", middleware.StaffOrAdmin(), c.UpdateInbound)
	rg.DELETE("/inbounds/:id", middleware.StaffOrAdmin(), c.DeleteInbound)
	rg.GET("/report-inbounds", middleware.AdminOnly(), c.LaporanInbound)
}

func (c *InboundController) ListInbound(ctx *gin.Context) {
	result, err := c.service.GetAll()
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil data"})
		return
	}
	ctx.JSON(http.StatusOK, result)
}
func (c *InboundController) AddInbound(ctx *gin.Context) {
	var p model.Inbound
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

func (c *InboundController) GetByID(ctx *gin.Context) {
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

func (c *InboundController) UpdateInbound(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	var input struct {
		ProductID  uint      `json:"product_id" binding:"required"`
		Quantity   int       `json:"quantity" binding:"required"`
		ReceivedAt time.Time `json:"received_at" binding:"required"`
		Supplier   string    `json:"supplier" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid"})
		return
	}

	inbound := model.Inbound{
		ID:         uri.ID,
		ProductID:  input.ProductID,
		Quantity:   input.Quantity,
		ReceivedAt: input.ReceivedAt,
		Supplier:   input.Supplier,
	}

	updated, err := c.service.Update(inbound)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update"})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}
func (c *InboundController) DeleteInbound(ctx *gin.Context) {
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

func (c *InboundController) LaporanInbound(ctx *gin.Context) {
	startDate := ctx.Query("start")
	endDate := ctx.Query("end")

	inbounds, err := c.service.GetLaporan(startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil laporan"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"laporan": inbounds,
	})
}