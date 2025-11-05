package controller

import (
	"net/http"
	"sistem-manajemen-gudang/model/domain"
	"sistem-manajemen-gudang/service"
	"sistem-manajemen-gudang/helper"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service service.ProductService
}

func NewProductController(s service.ProductService) *ProductController {
	return &ProductController{service: s}
}

//  POST /api/products
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
// GET /api/products
func (c *ProductController) ListProduct(ctx *gin.Context) {
	result, err := c.service.GetAll()
	if err != nil {
		helper.InternalServerError(ctx, "Gagal mengambil data produk")
		return
	}
	helper.Success(ctx, http.StatusOK, result)
}

//  GET /api/products/:id
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

//  PUT /api/products/:id
func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	var req struct {
		Name     string `json:"name" binding:"required"`
		SKU      string `json:"sku" binding:"required"`
		Category string `json:"category" binding:"required"`
		Unit     string `json:"unit" binding:"required"`
		Stock    int    `json:"stock" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		helper.BadRequest(ctx, "ID tidak valid")
		return
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
		Stock:    req.Stock,
	}

	updated, err := c.service.Update(product)
	if err != nil {
		helper.InternalServerError(ctx, "Gagal memperbarui data produk")
		return
	}

	helper.Success(ctx, http.StatusOK, updated)
}

//  DELETE /api/products/:id
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


