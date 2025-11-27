package controller

import (
	"sistem-manajemen-gudang/helper"
	"sistem-manajemen-gudang/middleware"
	"sistem-manajemen-gudang/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service service.AuthService
}
func NewAuthController(s service.AuthService) *AuthController {
	return &AuthController{service: s}
}

func AuthMiddleware() gin.HandlerFunc {
	return middleware.JWTAuthMiddleware()
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.BadRequest(ctx, "Input tidak valid: "+err.Error())
		return
	}

	token, err := c.service.Login(req.Username, req.Password)
	if err != nil {
		helper.Unauthorized(ctx, "Login gagal: "+err.Error())
		return
	}

	helper.Success(ctx, 200, gin.H{"token": token})
}

func (c *AuthController) Logout(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		helper.Unauthorized(ctx, "Token tidak valid atau tidak ditemukan")
		return
	}

	if err := c.service.Logout(userID.(uint)); err != nil {
		helper.InternalServerError(ctx, "Gagal logout: "+err.Error())
		return
	}

	helper.Success(ctx, 200, gin.H{"message": "Logout berhasil"})
}
