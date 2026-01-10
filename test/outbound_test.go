package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"strconv"
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

var testToken string


func setupTestDB() *gorm.DB {
	dsn := "root:mipa@tcp(localhost:3306)/tests?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&domain.User{},
		&domain.Product{},
		&domain.Outbound{},
	)

	config.DB = db

	return db
}

func truncateTables(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("DELETE FROM outbounds")
	db.Exec("DELETE FROM products")
	db.Exec("DELETE FROM users")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func setRouter(db *gorm.DB) *gin.Engine {
	outboundRepo := repository.NewOutboundRepository(db)
	outboundService := service.NewOutboundService(outboundRepo)
	outboundController := controller.NewOutboundController(outboundService)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.JWTAuthMiddleware())

	api := r.Group("/api")
	{
		api.POST("/outbounds", outboundController.AddOutbound)
		api.GET("/outbounds", outboundController.ListOutbound)
		api.GET("/outbounds/:id", outboundController.GetByID)
		api.PUT("/outbounds/:id", outboundController.UpdateOutbound)
		api.DELETE("/outbounds/:id", outboundController.DeleteOutbound)
	}

	return r
}


func seedUser(db *gorm.DB) {
	token, _ := util.GenerateToken(1, "admin")

	user := domain.User{
		ID:    1,
		Username:  "Admin Test",
		Role:  "admin",
		Token: &token,
	}

	db.Create(&user)
	testToken = token
}

func seedProduct(db *gorm.DB) domain.Product {
	product := domain.Product{
		Name:  "Produk A",
		Stock: 10,
	}
	db.Create(&product)
	return product
}


func TestCreateOutboundSuccess(t *testing.T) {
	db := setupTestDB()
	truncateTables(db)
	seedUser(db)
	router := setRouter(db)

	product := seedProduct(db)

	payload := map[string]interface{}{
		"product_id": product.ID,
		"quantity":   3,
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/outbounds", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("unexpected status %d, body: %s", rec.Code, rec.Body.String())
	}

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, float64(3), response["data"].(map[string]interface{})["quantity"])

	var updated domain.Product
	db.First(&updated, product.ID)
	assert.Equal(t, 7, updated.Stock)
}

func TestCreateOutboundFailedStockNotEnough(t *testing.T) {
	db := setupTestDB()
	truncateTables(db)
	seedUser(db)
	router := setRouter(db)

	product := seedProduct(db)

	payload := map[string]interface{}{
		"product_id": product.ID,
		"quantity":   20,
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/outbounds", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
func TestGetOutboundByIDSuccess(t *testing.T) {
	db := setupTestDB()
	truncateTables(db)
	seedUser(db)
	router := setRouter(db)

	product := seedProduct(db)

	outbound := domain.Outbound{
		ProductID: product.ID,
		Quantity:  2,
	}
	db.Create(&outbound)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/outbounds/"+strconv.Itoa(int(outbound.ID)),
		nil,
	)
	req.Header.Set("Authorization", "Bearer "+testToken)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status %d, body: %s", rec.Code, rec.Body.String())
	}
}

func TestUpdateOutboundSuccess(t *testing.T) {
	db := setupTestDB()
	truncateTables(db)
	seedUser(db)
	router := setRouter(db)

	product := seedProduct(db)

	outbound := domain.Outbound{
		ProductID: product.ID,
		Quantity:  2,
	}
	db.Create(&outbound)

	payload := map[string]interface{}{
		"destination": "Jakarta",
		"sent_at":     time.Now(),
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(
		http.MethodPut,
		"/api/outbounds/"+strconv.Itoa(int(outbound.ID)),
		bytes.NewBuffer(body),
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status %d, body: %s", rec.Code, rec.Body.String())
	}
}

func TestDeleteOutboundSuccess(t *testing.T) {
	db := setupTestDB()
	truncateTables(db)
	seedUser(db)
	router := setRouter(db)

	product := seedProduct(db)

	outbound := domain.Outbound{
		ProductID: product.ID,
		Quantity:  1,
	}
	db.Create(&outbound)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/outbounds/"+strconv.Itoa(int(outbound.ID)),
		nil,
	)
	req.Header.Set("Authorization", "Bearer "+testToken)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status %d, body: %s", rec.Code, rec.Body.String())
	}

	var count int64
	db.Model(&domain.Outbound{}).Count(&count)
	assert.Equal(t, int64(0), count)
}
