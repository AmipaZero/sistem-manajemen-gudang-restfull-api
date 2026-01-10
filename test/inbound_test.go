package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

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

var inboundTestToken string



func setupInboundTestDB() *gorm.DB {
	dsn := "root:mipa@tcp(localhost:3306)/tests?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&domain.User{},
		&domain.Product{},
		&domain.Inbound{},
	)

	config.DB = db
	return db
}

func truncateInboundTables(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE inbounds")
	db.Exec("TRUNCATE TABLE products")
	db.Exec("TRUNCATE TABLE users")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

}

func setInboundRouter(db *gorm.DB) *gin.Engine {
	inboundRepo := repository.NewInboundRepository(db)
	inboundService := service.NewInboundService(inboundRepo)
	inboundController := controller.NewInboundController(inboundService)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.JWTAuthMiddleware())

	api := r.Group("/api")
	{
		api.POST("/inbounds", inboundController.AddInbound)
		api.GET("/inbounds", inboundController.ListInbound)
		api.GET("/inbounds/:id", inboundController.GetByID)
		api.PUT("/inbounds/:id", inboundController.UpdateInbound)
		api.DELETE("/inbounds/:id", inboundController.DeleteInbound)
	}

	return r
}



func seedInboundUser(db *gorm.DB) {
	token, _ := util.GenerateToken(1, "admin")

	user := domain.User{
		ID:       1,
		Username: "Admin Inbound",
		Role:     "admin",
		Token:    &token,
	}

	db.Create(&user)
	inboundTestToken = token
}

func seedInboundProduct(db *gorm.DB) domain.Product {
	product := domain.Product{
		Name:  "Produk Inbound",
		Stock: 10,
	}
	db.Create(&product)
	return product
}



func TestCreateInboundSuccess(t *testing.T) {
	db := setupInboundTestDB()
	truncateInboundTables(db)
	seedInboundUser(db)
	router := setInboundRouter(db)

	product := seedInboundProduct(db)

	payload := map[string]interface{}{
		"product_id": product.ID,
		"quantity":   5,
		"supplier":   "PT Supplier",
		"received_at": time.Now(),
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/inbounds", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+inboundTestToken)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	var updated domain.Product
	db.First(&updated, product.ID)

	
	assert.Equal(t, 15, updated.Stock)
}

func TestGetInboundByIDSuccess(t *testing.T) {
	db := setupInboundTestDB()
	truncateInboundTables(db)
	seedInboundUser(db)
	router := setInboundRouter(db)

	product := seedInboundProduct(db)

	inbound := domain.Inbound{
		ProductID:  product.ID,
		Quantity:   2,
		Supplier:   "PT A",
		ReceivedAt: time.Now(),
	}
	db.Create(&inbound)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/inbounds/"+strconv.Itoa(int(inbound.ID)),
		nil,
	)
	req.Header.Set("Authorization", "Bearer "+inboundTestToken)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateInboundSuccess(t *testing.T) {
	db := setupInboundTestDB()
	truncateInboundTables(db)
	seedInboundUser(db)
	router := setInboundRouter(db)

	product := seedInboundProduct(db)

	inbound := domain.Inbound{
		ProductID:  product.ID,
		Quantity:   2,
		Supplier:   "Supplier Lama",
		ReceivedAt: time.Now(),
	}
	db.Create(&inbound)

	payload := map[string]interface{}{
		"supplier":    "Supplier Baru",
		"received_at": time.Now(),
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(
		http.MethodPut,
		"/api/inbounds/"+strconv.Itoa(int(inbound.ID)),
		bytes.NewBuffer(body),
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+inboundTestToken)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteInboundSuccess(t *testing.T) {
	db := setupInboundTestDB()
	truncateInboundTables(db)
	seedInboundUser(db)
	router := setInboundRouter(db)

	product := seedInboundProduct(db)

	inbound := domain.Inbound{
		ProductID: product.ID,
		Quantity:  1,
	}
	db.Create(&inbound)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/inbounds/"+strconv.Itoa(int(inbound.ID)),
		nil,
	)
	req.Header.Set("Authorization", "Bearer "+inboundTestToken)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var count int64
	db.Model(&domain.Inbound{}).Count(&count)
	assert.Equal(t, int64(0), count)
}