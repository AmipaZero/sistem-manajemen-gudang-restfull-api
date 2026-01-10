package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"sistem-manajemen-gudang/config"
	"sistem-manajemen-gudang/controller"
	"sistem-manajemen-gudang/middleware"
	"sistem-manajemen-gudang/model/domain"
	"sistem-manajemen-gudang/repository"
	"sistem-manajemen-gudang/service"
	"sistem-manajemen-gudang/util"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var productTestToken string


func setupProductTestDB() *gorm.DB {
	dsn := "root:mipa@tcp(localhost:3306)/tests?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&domain.User{},
		&domain.Product{},
	)

	config.DB = db
	return db
}

func truncateProductTables(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE inbounds")
	db.Exec("TRUNCATE TABLE products")
	db.Exec("TRUNCATE TABLE users")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

}

func setProductRouter(db *gorm.DB) *gin.Engine {
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.JWTAuthMiddleware())

	api := r.Group("/api")
	{
		api.POST("/products", productController.AddProduct)
		api.GET("/products", productController.ListProduct)
		api.GET("/products/:id", productController.GetByID)
		api.PUT("/products/:id", productController.UpdateProduct)
		api.DELETE("/products/:id", productController.DeleteProduct)
	}

	return r
}


func seedProductUser(db *gorm.DB) {
	token, _ := util.GenerateToken(1, "admin")

	user := domain.User{
		ID:       1,
		Username: "Admin Product",
		Role:     "admin",
		Token:    &token,
	}

	db.Create(&user)
	productTestToken = token
}


func TestCreateProductSuccess(t *testing.T) {
	db := setupProductTestDB()
	truncateProductTables(db)
	seedProductUser(db)
	router := setProductRouter(db)

	payload := map[string]interface{}{
		"name":     "Produk Test",
		"sku":      "SKU-001",
		"category": "Elektronik",
		"unit":     "pcs",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/products", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+productTestToken)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var product domain.Product
	db.First(&product)

	assert.Equal(t, 0, product.Stock)
}


func TestGetProductByIDSuccess(t *testing.T) {
	db := setupProductTestDB()
	truncateProductTables(db)
	seedProductUser(db)
	router := setProductRouter(db)

	product := domain.Product{
		Name:     "Produk A",
		SKU:      "SKU-A",
		Category: "Kategori A",
		Unit:     "pcs",
		Stock:    10,
	}
	db.Create(&product)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/products/"+strconv.Itoa(int(product.ID)),
		nil,
	)
	req.Header.Set("Authorization", "Bearer "+productTestToken)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateProductSuccess_StockNotChanged(t *testing.T) {
	db := setupProductTestDB()
	truncateProductTables(db)
	seedProductUser(db)
	router := setProductRouter(db)

	product := domain.Product{
		Name:     "Produk Lama",
		SKU:      "SKU-OLD",
		Category: "Kategori Lama",
		Unit:     "pcs",
		Stock:    50,
	}
	db.Create(&product)

	payload := map[string]interface{}{
		"name":     "Produk Baru",
		"sku":      "SKU-NEW",
		"category": "Kategori Baru",
		"unit":     "box",
		"stock":    999, 
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(
		http.MethodPut,
		"/api/products/"+strconv.Itoa(int(product.ID)),
		bytes.NewBuffer(body),
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+productTestToken)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var updated domain.Product
	db.First(&updated, product.ID)

	assert.Equal(t, 50, updated.Stock)
	assert.Equal(t, "SKU-NEW", updated.SKU)
}
func TestUpdateProductFailed_EmptySKU(t *testing.T) {
	db := setupProductTestDB()
	truncateProductTables(db)
	seedProductUser(db)
	router := setProductRouter(db)

	product := domain.Product{
		Name:     "Produk Error",
		SKU:      "SKU-X",
		Category: "Kategori",
		Unit:     "pcs",
		Stock:    5,
	}
	db.Create(&product)

	payload := map[string]interface{}{
		"name": "Produk Error Update",
		"sku":  "",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(
		http.MethodPut,
		"/api/products/"+strconv.Itoa(int(product.ID)),
		bytes.NewBuffer(body),
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+productTestToken)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}


func TestDeleteProductSuccess(t *testing.T) {
	db := setupProductTestDB()
	truncateProductTables(db)
	seedProductUser(db)
	router := setProductRouter(db)

	product := domain.Product{
		Name:     "Produk Delete",
		SKU:      "SKU-DEL",
		Category: "Kategori",
		Unit:     "pcs",
	}
	db.Create(&product)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/products/"+strconv.Itoa(int(product.ID)),
		nil,
	)
	req.Header.Set("Authorization", "Bearer "+productTestToken)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var count int64
	db.Model(&domain.Product{}).Count(&count)
	assert.Equal(t, int64(0), count)
}
