package controller

import (
	"sistem-manajemen-gudang/service"
	"net/http"
	"github.com/gin-gonic/gin"
)
type UserService struct {
	service service.UserService
}

func NewUserController(s service.UserService) *UserService {
	return &UserService{s}
}


func (c *UserService) RegisterPublicRoutes(rg *gin.RouterGroup) {
    rg.POST("/register", c.Register)
}

func (c *UserService) RegisterProtectedRoutes(rg *gin.RouterGroup) {
    rg.GET("/current", c.Current)
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

func (c *UserService) Current(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID := userIDVal.(uint)

	res, err := c.service.CurrentUser(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

