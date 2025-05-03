package main

import (
	"log"
	"os"

	"github.com/breezjirasak/triptales/config"
	"github.com/breezjirasak/triptales/internal/service"
	"github.com/breezjirasak/triptales/internal/repository"
	"github.com/breezjirasak/triptales/internal/model"
	"github.com/breezjirasak/triptales/internal/route"
)

func main() {
	// Load environment variables
	config.LoadEnv()
	
	// Initialize database connection
	config.InitDB()

	// Auto migrate tables
	config.DB.AutoMigrate(
		&model.Country{},
		&model.Trip{},
		&model.User{},
		&model.Friend{},
		&model.TripCompanion{},
		&model.Bookmark{},
	)

	// Ensure JWT secret is set
	if os.Getenv("JWT_SECRET_KEY") == "" {
		// For development only
		os.Setenv("JWT_SECRET_KEY", "your-default-secret-key-change-in-production")
		log.Println("WARNING: Using default JWT secret key. Set JWT_SECRET_KEY environment variable in production.")
	}

	// Create a repository instance
	userRepo := repository.NewUserRepository(config.DB)
	countryRepo := repository.NewCountryRepository(config.DB)
	tripRepo := repository.NewTripRepository(config.DB)
	tripCompanionRepo := repository.NewTripCompanionRepository(config.DB)

	// Create a service instance using the repository
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)
	countryService := service.NewCountryService(countryRepo)
	tripService := service.NewTripService(tripRepo)
	tripCompanionService := service.NewTripCompanionService(tripCompanionRepo)

	// Set up router with auth service
	r := route.SetupRouter(authService, userService, countryService, tripService, tripCompanionService)
	
	// Start server
	log.Println("Server starting on port 8080...")
	r.Run(":8080")
}