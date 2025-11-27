package router

import (
	"sistem-manajemen-gudang/controller"
	"sistem-manajemen-gudang/middleware"

	"github.com/gin-gonic/gin"
)

func OutboundRoutes(rg *gin.RouterGroup, c *controller.OutboundController) {
	outboundsGroup := rg.Group("/outbound")
	outboundsGroup.Use(middleware.StaffOrAdmin())
	outboundsGroup.GET("", c.ListOutbound)
	outboundsGroup.POST("/add", c.AddOutbound)
	outboundsGroup.GET("/:id", c.GetByID)
	outboundsGroup.PATCH("/:id", c.UpdateOutbound)
	outboundsGroup.DELETE("/:id", c.DeleteOutbound)

	
}
