package controller

import (
	"net/http"
	"time"

	"sistem-manajemen-gudang/helper"
	"sistem-manajemen-gudang/model/domain"
	"sistem-manajemen-gudang/service"

	"github.com/gin-gonic/gin"
)

type OutboundController struct {
	service service.OutboundService
}

func NewOutboundController(s service.OutboundService) *OutboundController {
	return &OutboundController{service: s}
}

//  GET /api/outbounds
func (c *OutboundController) ListOutbound(ctx *gin.Context) {
	result, err := c.service.GetAll()
	if err != nil {
		helper.InternalServerError(ctx, "gagal mengambil data")
		return
	}
	helper.Success(ctx, http.StatusOK, result)

}

//  POST /api/outbounds
func (c *OutboundController) AddOutbound(ctx *gin.Context) {
	var req domain.Outbound
	if err := ctx.ShouldBindJSON(&req); err != nil || req.ProductID == 0 {
		helper.BadRequest(ctx, "Input tidak valid")
		return
	}

	result, err := c.service.Create(req)
	if err != nil {
		helper.InternalServerError(ctx, "Gagal menyimpan data")
		return
	}

	helper.Success(ctx, http.StatusCreated, result)
}

//  GET /api/outbounds/:id
func (c *OutboundController) GetByID(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		helper.BadRequest(ctx, "ID tidak valid")
		return
	}

	outbound, err := c.service.GetByID(uri.ID)
	if err != nil {
		helper.NotFound(ctx, "Outbound tidak ditemukan")
		return
	}

	helper.Success(ctx, http.StatusOK, outbound)
}

//  PUT /api/outbounds/:id
func (c *OutboundController) UpdateOutbound(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	var req struct {
		ProductID    uint      `json:"product_id" binding:"required"`
		Quantity     int       `json:"quantity" binding:"required"`
		SentAt       time.Time `json:"sent_at" binding:"required"`
		Destination  string    `json:"destination" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		helper.BadRequest(ctx, "ID tidak valid")
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.BadRequest(ctx, "Data input tidak valid")
		return
	}

	outbound := domain.Outbound{
		ID:          uri.ID,
		ProductID:   req.ProductID,
		Quantity:    req.Quantity,
		SentAt:      req.SentAt,
		Destination: req.Destination,
	}

	updated, err := c.service.Update(outbound)
	if err != nil {
		helper.InternalServerError(ctx, "Gagal memperbarui data")
		return
	}

	helper.Success(ctx, http.StatusOK, updated)
}

//  DELETE /api/outbounds/:id
func (c *OutboundController) DeleteOutbound(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		helper.BadRequest(ctx, "ID tidak valid")
		return
	}

	if err := c.service.Delete(uri.ID); err != nil {
		helper.InternalServerError(ctx, "Gagal menghapus data outbound")
		return
	}

	helper.Success(ctx, http.StatusOK, gin.H{"message": "Data outbound berhasil dihapus"})
}


