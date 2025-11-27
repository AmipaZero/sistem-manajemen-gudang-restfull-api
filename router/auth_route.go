package router

import (
	"sistem-manajemen-gudang/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, c *controller.AuthController) {

		r.POST("/login", c.Login)
		r.DELETE("/logout", c.Logout)

}
