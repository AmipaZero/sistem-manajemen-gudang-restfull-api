package controller

import (
	"sistem-manajemen-gudang/service"
	"net/http"
	"github.com/gin-gonic/gin"
)
// "sistem-manajemen-gudang/middleware"
type UserService struct {
	service service.UserService
}

func NewUserController(s service.UserService) *UserService {
	return &UserService{s}
}
// func AuthMiddleware() gin.HandlerFunc {
// 	return middleware.JWTAuthMiddleware()
// }


func (c *UserService) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", c.Register)
	// // rg.POST("/login", c.Login)

	// auth := rg.Group("/")
	// auth.Use(AuthMiddleware()) // middleware JWT
	// {
	// 	auth.DELETE("/logout", c.Logout)
	// }
}

func (c *UserService) Register(ctx *gin.Context) {
		var req service.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Register(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "register success"})
}


