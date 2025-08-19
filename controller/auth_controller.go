package controller

import (
	"sistem-manajemen-gudang/service"
	"sistem-manajemen-gudang/middleware"
	"net/http"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service service.AuthService
}

func NewAuthController(s service.AuthService) *AuthController {
	return &AuthController{s}
}
func AuthMiddleware() gin.HandlerFunc {
	return middleware.JWTAuthMiddleware()
}


func (c *AuthController) RegisterRoutes(rg *gin.RouterGroup) {
	// rg.POST("/register", c.Register)
	rg.POST("/login", c.Login)

	auth := rg.Group("/")
	auth.Use(AuthMiddleware()) // middleware JWT
	{
		auth.DELETE("/logout", c.Logout)
	}
}

// func (c *AuthController) Register(ctx *gin.Context) {
// 		var req service.RegisterRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := c.service.Register(&req); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"message": "register success"})
// }

func (c *AuthController) Login(ctx *gin.Context) {
		var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.service.Login(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *AuthController) Logout(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := c.service.Logout(userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal logout"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Logout berhasil"})
}
