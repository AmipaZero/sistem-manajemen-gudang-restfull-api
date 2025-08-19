package main

import (
	"sistem-manajemen-gudang/config"
	"sistem-manajemen-gudang/controller"
	"sistem-manajemen-gudang/repository"
	"sistem-manajemen-gudang/service"
	"sistem-manajemen-gudang/middleware"

	"github.com/gin-gonic/gin"
)

func main(){
	config.ConnectDB()
	db := config.DB
	authRepo := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	// product
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)
	// inbound
	inboundRepo := repository.NewInboundRepository(db)
	inboundService := service.NewInboundService(inboundRepo)
	inboundController := controller.NewInboundController(inboundService)
	// outbound
	outboundRepo := repository.NewOutboundRepository(db)
	outboundService := service.NewOutboundService(outboundRepo)
	outboundController := controller.NewOutboundController(outboundService)

	r := gin.Default()
	api := r.Group("/")
	// Public route
	userController.RegisterRoutes(api)
	authController.RegisterRoutes(api)

		// Protected route (wajib token)
		protected := api.Group("/api")
		protected.Use(middleware.JWTAuthMiddleware())
		{
			productController.RegisterRoutes(protected)
			inboundController.RegisterRoutes(protected)
			outboundController.RegisterRoutes(protected)
		}

	r.Run(":8080")
	
}