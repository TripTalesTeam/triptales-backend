package route

import (
	"github.com/breezjirasak/triptales/internal/middleware"
	"github.com/breezjirasak/triptales/internal/handler"
	"github.com/breezjirasak/triptales/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authService *service.AuthService, userService *service.UserService, countryService *service.CountryService) *gin.Engine {
	r := gin.Default()
	
	// Auth routes
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	countryHandler := handler.NewCountryHandler(countryService)
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
		user := api.Group("/users")
		{
			user.GET("/", userHandler.GetUsers)
		}


		country := api.Group("/countries")
		{
			country.POST("/", countryHandler.CreateCountry)
			country.GET("/", countryHandler.GetAllCountries)
			country.GET("/by-name", countryHandler.GetCountryByName)
			country.GET("/:id", countryHandler.GetCountryByID)
			country.PUT("/", countryHandler.UpdateCountry)
			country.DELETE("/:id", countryHandler.DeleteCountry)
		}
	}

	return r
}