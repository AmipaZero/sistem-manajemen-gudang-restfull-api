// package router
// import (
// 	"sistem-manajemen-gudang/controller"
// 	"sistem-manajemen-gudang/middleware"
// 	"github.com/gin-gonic/gin"
// )

// func InboundRoutes(rg *gin.RouterGroup, c *controller.InboundController) {
// 	group := rg.Group("/inbounds")
// 	group.Use(middleware.StaffOrAdmin())
// 	group.POST("/inbounds/add", c.AddInbound)
// 	group.GET("/inbounds", c.ListInbound)
// 	group.GET("/inbounds/:id", c.GetByID)
// 	group.PUT("/inbounds/:id", c.UpdateInbound)
// 	group.DELETE("/inbounds/:id", c.DeleteInbound)
// 	group.GET("/report-inbounds", c.LaporanInbound)

// }
package router

import (
	"sistem-manajemen-gudang/controller"
	"sistem-manajemen-gudang/middleware"

	"github.com/gin-gonic/gin"
)

func InboundRoutes(rg *gin.RouterGroup, c *controller.InboundController) {
	inboundGroup := rg.Group("/inbounds")
	inboundGroup.Use(middleware.StaffOrAdmin())

	inboundGroup.GET("/", c.ListInbound)
	inboundGroup.POST("/add", c.AddInbound)
	inboundGroup.GET("/:id", c.GetByID)
	inboundGroup.PUT("/:id", c.UpdateInbound)
	inboundGroup.DELETE("/:id", c.DeleteInbound)
	
	
}
