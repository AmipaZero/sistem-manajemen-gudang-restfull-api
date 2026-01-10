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

var userTestToken string

// ================= SETUP =================
func setupUserTestDB() *gorm.DB {
	dsn := "root:mipa@tcp(localhost:3306)/tests?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&domain.User{},
	)
	config.DB = db
	return db
}

func truncateUserTables(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE users")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func setUserRouter(db *gorm.DB) *gin.Engine {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := gin.New()
	r.Use(gin.Recovery())

	api := r.Group("/api/users")
	{
		userController.RegisterPublicRoutes(api)    // register
		protected := api.Group("/")
		protected.Use(middleware.JWTAuthMiddleware())
		userController.RegisterProtectedRoutes(protected) // current
	}

	return r
}

func seedUsers(db *gorm.DB) domain.User {
	hash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := domain.User{
		ID:       1,
		Username: "testuser",
		Password: string(hash),
		Role:     domain.Admin,
	}
	db.Create(&user)
	return user
}

func TestRegisterUserSuccess(t *testing.T) {
	db := setupUserTestDB()
	truncateUserTables(db)
	router := setUserRouter(db)

	payload := map[string]interface{}{
		"username": "newuser",
		"password": "mypassword",
		"role":     "admin",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/users/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var user domain.User
	db.First(&user, "username = ?", "newuser")
	assert.Equal(t, "newuser", user.Username)
	assert.Equal(t, domain.Admin, user.Role)
	assert.NotEmpty(t, user.Password) 
}

func TestCurrentUserSuccess(t *testing.T) {
	db := setupUserTestDB()
	truncateUserTables(db)

	user := seedUsers(db)

	router := setUserRouter(db)

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)

	token, err := authService.Login(user.Username, "password123")
	if err != nil {
		t.Fatalf("failed to login user: %v", err)
	}
	userTestToken = token

	req := httptest.NewRequest(http.MethodGet, "/api/users/current", nil)
	req.Header.Set("Authorization", "Bearer "+userTestToken)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)

	data, ok := resp["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected data in response, got: %v", resp)
	}

	assert.Equal(t, user.Username, data["username"])
	assert.Equal(t, string(user.Role), data["role"])
}

