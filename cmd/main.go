package main

import (
	"log"
	"os"

	"github.com/breezjirasak/triptales/config"
	"github.com/breezjirasak/triptales/internal/service"
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

	// Create uploads directory if it doesn't exist
	// if err := os.MkdirAll("uploads/profiles", 0755); err != nil {
	// 	log.Fatalf("Failed to create uploads directory: %v", err)
	// }

	// Initialize auth service
	authService := service.NewAuthService(config.DB)

	// Set up router with auth service
	r := route.SetupRouter(authService)
	
	// Start server
	log.Println("Server starting on port 8080...")
	r.Run(":8080")
}