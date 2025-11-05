package router

import (
	"sistem-manajemen-gudang/controller"
	"sistem-manajemen-gudang/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(rg *gin.RouterGroup, c *controller.ProductController) {
	productGroup := rg.Group("/products")
	productGroup.Use(middleware.StaffOrAdmin())
	productGroup.GET("/", c.ListProduct)
	productGroup.POST("/add", c.AddProduct)
	productGroup.GET("/:id", c.GetByID)
	productGroup.PUT("/:id", c.UpdateProduct)
	productGroup.DELETE("/:id", c.DeleteProduct)
}
