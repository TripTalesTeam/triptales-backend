package route

import (
	"github.com/breezjirasak/triptales/internal/middleware"
	"github.com/breezjirasak/triptales/internal/handler"
	"github.com/breezjirasak/triptales/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authService *service.AuthService) *gin.Engine {
	r := gin.Default()

	// Public routes
	r.GET("/users", handler.GetUsers)
	
	// Auth routes
	authHandler := handler.NewAuthHandler(authService)
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
		
		// Protected auth routes
		protected := authGroup.Group("/")
		protected.Use(middleware.JWTMiddleware())
		{
			protected.GET("/me", authHandler.GetMe)
			protected.POST("/profile-image", authHandler.UploadProfileImage)
		}
	}
	
	// API routes that require authentication
	api := r.Group("/api")
	api.Use(middleware.JWTMiddleware())
	{
		// Protected routes go here
		// api.GET("/protected-resource", handler.ProtectedResource)
		
		// You can add your other protected routes here
		// e.g., api.GET("/trips", tripHandler.GetTrips)
	}

	return r
}