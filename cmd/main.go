package main

import (
	"github.com/breezjirasak/triptales/config"
	"github.com/breezjirasak/triptales/internal/model"
	"github.com/breezjirasak/triptales/internal/route"
)

func main() {
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

	r := route.SetupRouter()
	r.Run(":8080")
}
