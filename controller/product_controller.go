package controller

import (
	"net/http"
	"sistem-manajemen-gudang/middleware"
	"sistem-manajemen-gudang/model"
	"sistem-manajemen-gudang/service"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service service.ProductService
}

func NewProductController(h service.ProductService) *ProductController {
	return &ProductController{service: h}
}

func (c *ProductController) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/products/add", middleware.StaffOrAdmin(), c.AddProduct)
	rg.GET("/products",middleware.StaffOrAdmin(), c.ListProduct)
	rg.GET("/products/:id",middleware.StaffOrAdmin(), c.GetByID)
	rg.PUT("/products/:id",middleware.StaffOrAdmin(), c.UpdateProduct)
	rg.DELETE("/products/:id", middleware.StaffOrAdmin(),c.DeleteProduct)
	rg.GET("/report-products",  middleware.AdminOnly() ,c.LaporanProduct)
}

func (c *ProductController) AddProduct(ctx *gin.Context) {
	var p model.Product
	if err := ctx.ShouldBindJSON(&p); err != nil || p.Name == "" {
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

func (c *ProductController) ListProduct(ctx *gin.Context) {
	result, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil data"})
		return
	}
	ctx.JSON(http.StatusOK, result)
}
func (c *ProductController) GetByID(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	product, err := c.service.GetByID(uri.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	var input struct {
		Name     string      `json:"name" binding:"required"`
		SKU      string      `gorm:"uniqueIndex" json:"sku" binding:"required"`
		Category string      `json:"category" binding:"required"`
		Unit     string      `json:"unit" binding:"required"`
		Stock    int         `json:"stock" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid"})
		return
	}

	product := model.Product{
		ID:   uri.ID,
		Name: input.Name,
		SKU: input.SKU,
		Category: input.Category,
		Unit: input.Unit,
		Stock: input.Stock,
	}

	updated, err := c.service.Update(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update"})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
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
func (c *ProductController) LaporanProduct(ctx *gin.Context) {
	startDate := ctx.Query("start")
	endDate := ctx.Query("end")

	products, err := c.service.GetLaporan(startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil laporan"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"laporan": products,
	})
}