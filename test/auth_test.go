package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"sistem-manajemen-gudang/config"
	"sistem-manajemen-gudang/controller"
	"sistem-manajemen-gudang/middleware"
	"sistem-manajemen-gudang/model/domain"
	"sistem-manajemen-gudang/repository"
	"sistem-manajemen-gudang/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var authTestToken string

func setupAuthTestDB() *gorm.DB {
	dsn := "root:mipa@tcp(localhost:3306)/tests?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&domain.User{})
	config.DB = db
	return db
}

func truncateAuthTables(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE users")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func setAuthRouter(db *gorm.DB) *gin.Engine {
	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	r := gin.New()
	r.Use(gin.Recovery())

	public := r.Group("/api/auth")
	{
		public.POST("/login", authController.Login)
	}

	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.DELETE("/logout", authController.Logout)
	}

	return r
}

func seedAuthUser(db *gorm.DB) domain.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := domain.User{
		ID:       1,
		Username: "testuser",
		Password: string(password),
		Role:     domain.Admin,
	}
	db.Create(&user)
	return user
}


func TestLoginSuccess(t *testing.T) {
	db := setupAuthTestDB()
	truncateAuthTables(db)
	user := seedAuthUser(db)
	router := setAuthRouter(db)

	payload := map[string]string{
		"username": user.Username,
		"password": "password123",
	}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	data := resp["data"].(map[string]interface{})
	token := data["token"].(string)
	assert.NotEmpty(t, token)

	authTestToken = token
}

func TestLoginFailed_WrongPassword(t *testing.T) {
	db := setupAuthTestDB()
	truncateAuthTables(db)
	user := seedAuthUser(db)
	router := setAuthRouter(db)

	payload := map[string]string{
		"username": user.Username,
		"password": "wrongpassword",
	}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestLogoutSuccess(t *testing.T) {
	db := setupAuthTestDB()
	truncateAuthTables(db)
	user := seedAuthUser(db)
	router := setAuthRouter(db)

	authTestToken, _ := service.NewAuthService(repository.NewAuthRepository(db)).Login(user.Username, "password123")

	req := httptest.NewRequest(http.MethodDelete, "/api/logout", nil)
	req.Header.Set("Authorization", "Bearer "+authTestToken)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var updated domain.User
	db.First(&updated, user.ID)
	assert.Nil(t, updated.Token)
}
