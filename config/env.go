// config/env.go
package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	if os.Getenv("PRODUCTION") != "true" {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Println("Warning: No .env file found, relying on system environment variables")
		}
	}
}
