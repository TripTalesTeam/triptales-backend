// config/env.go
package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Warning: No .env file found, relying on system environment variables")
	}
}
