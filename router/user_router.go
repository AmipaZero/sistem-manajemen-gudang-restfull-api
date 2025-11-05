package router

import (
	"sistem-manajemen-gudang/controller"
	"sistem-manajemen-gudang/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, c *controller.UserController) {
	// --- PUBLIC ROUTES 
	public := r.Group("/api")
	{
		public.POST("/register", c.Register)
	}

	// --- PROTECTED ROUTES 
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/current", c.Current)
	}
}
