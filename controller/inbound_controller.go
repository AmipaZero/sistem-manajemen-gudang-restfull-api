package controller
import (
	"net/http"
	"sistem-manajemen-gudang/model/domain"
	"sistem-manajemen-gudang/service"
	"sistem-manajemen-gudang/helper"
	"time"
	"github.com/gin-gonic/gin"
)

type InboundController struct {
	service service.InboundService
}

func NewInboundController(s service.InboundService) *InboundController {
	return &InboundController{service: s}
}

func (c *InboundController) ListInbound(ctx *gin.Context) {
	result, err := c.service.GetAll()
	if err != nil {
		helper.InternalServerError(ctx, "Gagal mengambil data inbounds")
		return
	}
	helper.Success(ctx, http.StatusOK, result)
}

func (c *InboundController) AddInbound(ctx *gin.Context) {
	var inbound domain.Inbound
	if err := ctx.ShouldBindJSON(&inbound); err != nil {
		helper.BadRequest(ctx, "Input tidak valid")
		return
	}

	result, err := c.service.Create(inbound)
	if err != nil {
		helper.InternalServerError(ctx, "Gagal menyimpan data inbound")
		return
	}

	helper.Success(ctx, http.StatusCreated, result)
}

func (c *InboundController) GetByID(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		helper.BadRequest(ctx, "ID tidak valid")
		return
	}

	inbound, err := c.service.GetByID(uri.ID)
	if err != nil {
		helper.NotFound(ctx, "Inbound tidak ditemukan")
		return
	}

	helper.Success(ctx, http.StatusOK, inbound)
}

func (c *InboundController) UpdateInbound(ctx *gin.Context) {
    var uri struct {
        ID uint `uri:"id" binding:"required"`
    }

    if err := ctx.ShouldBindUri(&uri); err != nil {
        helper.BadRequest(ctx, "ID tidak valid")
        return
    }

    var input struct {
        Supplier   string    `json:"supplier" binding:"required"`
        ReceivedAt time.Time `json:"received_at" binding:"required"`
    }

    if err := ctx.ShouldBindJSON(&input); err != nil {
        helper.BadRequest(ctx, "Data input tidak valid")
        return
    }

    updated, err := c.service.UpdateData(uri.ID, input.Supplier, input.ReceivedAt)
    if err != nil {
        helper.InternalServerError(ctx, "Gagal memperbarui data inbound")
        return
    }

    helper.Success(ctx, http.StatusOK, updated)
}



func (c *InboundController) DeleteInbound(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		helper.BadRequest(ctx, "ID tidak valid")
		return
	}

	err := c.service.Delete(uri.ID)
	if err != nil {
		helper.InternalServerError(ctx, "Gagal menghapus data inbound")
		return
	}

	helper.Success(ctx, http.StatusOK, gin.H{"message": "Data inbound berhasil dihapus"})
}


