package router

import (
	"sistem-manajemen-gudang/controller"
	"sistem-manajemen-gudang/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, c *controller.AuthController) {
	public := r.Group("/api")
	{
		public.POST("/login", c.Login)
	}

	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.DELETE("/logout", c.Logout)
	}
}
