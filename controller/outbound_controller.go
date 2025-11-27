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

func (c *OutboundController) ListOutbound(ctx *gin.Context) {
	result, err := c.service.GetAll()
	if err != nil {
		helper.InternalServerError(ctx, "gagal mengambil data")
		return
	}
	helper.Success(ctx, http.StatusOK, result)

}

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

func (c *OutboundController) UpdateOutbound(ctx *gin.Context) {
    var uri struct {
        ID uint `uri:"id" binding:"required"`
    }

    if err := ctx.ShouldBindUri(&uri); err != nil {
        helper.BadRequest(ctx, "ID tidak valid")
        return
    }

    var input struct {
        Destination   string    `json:"destination" binding:"required"`
        SentAt time.Time `json:"sent_at" binding:"required"`
    }

    if err := ctx.ShouldBindJSON(&input); err != nil {
        helper.BadRequest(ctx, "Data input tidak valid")
        return
    }

    updated, err := c.service.Update(uri.ID, input.Destination, input.SentAt)
    if err != nil {
        helper.InternalServerError(ctx, "Gagal memperbarui data outbound")
        return
    }

    helper.Success(ctx, http.StatusOK, updated)
}

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


