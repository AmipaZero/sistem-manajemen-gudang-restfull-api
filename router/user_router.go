package router

import (
	"sistem-manajemen-gudang/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, c *controller.UserController) {
		// --- PUBLIC ROUTES 

		r.POST("/register", c.Register)
	// --- PROTECTED ROUTES 
		r.GET("/current", c.Current)

}
