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

//  GET /api/inbounds
func (c *InboundController) ListInbound(ctx *gin.Context) {
	result, err := c.service.GetAll()
	if err != nil {
		helper.InternalServerError(ctx, "Gagal mengambil data inbounds")
		return
	}
	helper.Success(ctx, http.StatusOK, result)
}

//  POST /api/inbounds
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
// GET /api/inbounds/:id
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

//  PUT /api/inbounds/:id
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
		helper.BadRequest(ctx, "ID tidak valid")
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		helper.BadRequest(ctx, "Data input tidak valid")
		return
	}

	inbound := domain.Inbound{
		ID:         uri.ID,
		ProductID:  input.ProductID,
		Quantity:   input.Quantity,
		ReceivedAt: input.ReceivedAt,
		Supplier:   input.Supplier,
	}

	updated, err := c.service.Update(inbound)
	if err != nil {
		helper.InternalServerError(ctx, "Gagal memperbarui data inbound")
		return
	}

	helper.Success(ctx, http.StatusOK, updated)
}

//  DELETE /api/inbounds/:id
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


