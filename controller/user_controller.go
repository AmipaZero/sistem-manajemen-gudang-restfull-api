package controller

import (
	"sistem-manajemen-gudang/helper"
	"sistem-manajemen-gudang/middleware"
	"sistem-manajemen-gudang/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{service: s}
}

// Middleware JWT untuk user
func UserMiddleware() gin.HandlerFunc {
	return middleware.JWTAuthMiddleware()
}

// POST /api/register
func (c *UserController) Register(ctx *gin.Context) {
	var req service.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.BadRequest(ctx, "Input tidak valid: "+err.Error())
		return
	}

	if err := c.service.Register(&req); err != nil {
		helper.InternalServerError(ctx, "Gagal register: "+err.Error())
		return
	}

	helper.Success(ctx, 200, gin.H{"message": "Register berhasil"})
}

// GET /api/users/me
func (c *UserController) Current(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("userID")
	if !exists {
		helper.Unauthorized(ctx, "Unauthorized")
		return
	}

	userID := userIDVal.(uint)
	res, err := c.service.CurrentUser(userID)
	if err != nil {
		helper.BadRequest(ctx, "User tidak ditemukan: "+err.Error())
		return
	}

	helper.Success(ctx, 200, res)
}
