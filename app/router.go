package app

import (
	"sistem-manajemen-gudang/controller"
	"sistem-manajemen-gudang/repository"
	"sistem-manajemen-gudang/router"
	"sistem-manajemen-gudang/service"
	"sistem-manajemen-gudang/config"
	"sistem-manajemen-gudang/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// inbounds
	inboundRepo := repository.NewInboundRepository(db)
	inboundService := service.NewInboundService(inboundRepo)
	inboundController := controller.NewInboundController(inboundService)
	// outbounds
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
	// outbound
	outboundRepo := repository.NewOutboundRepository(db)
	outboundService := service.NewOutboundService(outboundRepo)
	outboundController := controller.NewOutboundController(outboundService)

// --- Group tanpa JWT (public) ---
	public := r.Group("/api")
	{
		public.POST("/login", authController.Login)
		public.POST("/register", userController.Register)
	}

	// --- Group dengan JWT (protected) ---
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		router.InboundRoutes(protected, inboundController)
		router.OutboundRoutes(protected, outboundController)
		router.ProductRoutes(protected, productController)
		protected.GET("/current", userController.Current)
		protected.DELETE("/logout", authController.Logout)
	}


	// protected := r.Group("/api")
	
	// protected.Use(middleware.JWTAuthMiddleware())
	// {
	// 	router.InboundRoutes(protected, inboundController)
	// router.OutboundRoutes(protected, outboundController)
	// router.ProductRoutes(protected, productController)
	// router.UserRoutes(protected, userController)
	// router.AuthRoutes(protected, authController)
	// }
	// router.InboundRoutes(protected, inboundController)
	// router.OutboundRoutes(protected, outboundController)
	// router.ProductRoutes(protected, productController)
	// router.UserRoutes(protected, userController)
	// router.AuthRoutes(protected, authController)


// 	r := gin.Default()
// 	api := r.Group("/")
// 	// Public route
// 	userController.RegisterPublicRoutes(api)
// 	authController.RegisterRoutes(api)

// 		// Protected route
// 		protected := api.Group("/api")
// 		protected.Use(middleware.JWTAuthMiddleware())
// 		{
// 			userController.RegisterProtectedRoutes(protected)
// 			productController.RegisterRoutes(protected)
// 			inboundController.RegisterRoutes(protected)
// 			outboundController.RegisterRoutes(protected)
// 		}




	return r
}
