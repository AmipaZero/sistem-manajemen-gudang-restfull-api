package router

import (
	"sistem-manajemen-gudang/controller"
	"sistem-manajemen-gudang/middleware"
	"github.com/gin-gonic/gin"
)

func InboundRoutes(rg *gin.RouterGroup, c *controller.InboundController) {
	inboundGroup := rg.Group("/inbound")
	inboundGroup.Use(middleware.StaffOrAdmin())
	inboundGroup.GET("", c.ListInbound)
	inboundGroup.POST("/add", c.AddInbound)
	inboundGroup.GET("/:id", c.GetByID)
	inboundGroup.PATCH("/:id", c.UpdateInbound)
	inboundGroup.DELETE("/:id", c.DeleteInbound)
}
