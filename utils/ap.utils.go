package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// @dev Loads environment variables
func LoadEnvVars() {
	err := godotenv.Load();
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
