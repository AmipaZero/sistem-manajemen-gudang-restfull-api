package app

import (
	"sistem-manajemen-gudang/config"
	"sistem-manajemen-gudang/controller"

	"sistem-manajemen-gudang/middleware"
	"sistem-manajemen-gudang/repository"
	"sistem-manajemen-gudang/router"

	"sistem-manajemen-gudang/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		
		AllowOrigins:     []string{"http://localhost:5173","http://localhost:10917",},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	//auth
		authRepo := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	// inbounds
	inboundRepo := repository.NewInboundRepository(db)
	inboundService := service.NewInboundService(inboundRepo)
	inboundController := controller.NewInboundController(inboundService)
	
	// product
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)
	// outbound
	outboundRepo := repository.NewOutboundRepository(db)
	outboundService := service.NewOutboundService(outboundRepo)
	outboundController := controller.NewOutboundController(outboundService)


	

	public := r.Group("/api/auth")
	{
		public.POST("/login", authController.Login)
		public.POST("/register", userController.Register)
	

	}

	protected := r.Group("/api")
	
	protected.Use(middleware.JWTAuthMiddleware())
	{
		router.OutboundRoutes(protected, outboundController)
		router.ProductRoutes(protected, productController)
		router.InboundRoutes(protected, inboundController)
		protected.GET("/current", userController.Current)
		protected.DELETE("/logout", authController.Logout)
	}

	return r
}
