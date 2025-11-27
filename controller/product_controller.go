package controller

import (
	"net/http"
	"sistem-manajemen-gudang/helper"
	"sistem-manajemen-gudang/model/domain"
	"sistem-manajemen-gudang/service"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service service.ProductService
}

func NewProductController(s service.ProductService) *ProductController {
	return &ProductController{service: s}
}

func (c *ProductController) AddProduct(ctx *gin.Context) {
	var req domain.Product
	if err := ctx.ShouldBindJSON(&req); err != nil || req.Name == "" {
		helper.BadRequest(ctx, "Input tidak valid")
		return
	}

	result, err := c.service.Create(req)
	if err != nil {
		helper.InternalServerError(ctx, "Gagal menyimpan data produk")
		return
	}

	helper.Success(ctx, http.StatusOK, result)
}

func (c *ProductController) ListProduct(ctx *gin.Context) {
	result, err := c.service.GetAll()
	if err != nil {
		helper.InternalServerError(ctx, "Gagal mengambil data produk")
		return
	}
	helper.Success(ctx, http.StatusOK, result)
}

func (c *ProductController) GetByID(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		helper.BadRequest(ctx, "ID tidak valid")
		return
	}

	product, err := c.service.GetByID(uri.ID)
	if err != nil {
		helper.NotFound(ctx, "Data produk tidak ditemukan")
		return
	}

	helper.Success(ctx, http.StatusOK, product)
}
func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		helper.BadRequest(ctx, "ID tidak valid")
		return
	}

	var req struct {
		Name     string `json:"name"`
		SKU      string `json:"sku"`
		Category string `json:"category"`
		Unit     string `json:"unit"`
		// Stock tidak diterima di sini
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.BadRequest(ctx, "Data input tidak valid")
		return
	}

	product := domain.Product{
		ID:       uri.ID,
		Name:     req.Name,
		SKU:      req.SKU,
		Category: req.Category,
		Unit:     req.Unit,
	}

	updated, err := c.service.Update(product)
	if err != nil {
		helper.InternalServerError(ctx, err.Error())
		return
	}

	helper.Success(ctx, http.StatusOK, updated)
}



//  DELETE
func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		helper.BadRequest(ctx, "ID tidak valid")
		return
	}

	if err := c.service.Delete(uri.ID); err != nil {
		helper.InternalServerError(ctx, "Gagal menghapus data produk")
		return
	}

	helper.Success(ctx, http.StatusOK, gin.H{"message": "Data produk berhasil dihapus"})
}


