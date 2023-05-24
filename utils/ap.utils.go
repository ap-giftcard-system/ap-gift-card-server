package utils

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @dev Loads environment variables
func LoadEnvVars() {
	err := godotenv.Load();
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// @dev Sets up config for cors
// 
// @return gin.HandlerFunc
func SetupCorsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: 		[]string{os.Getenv("CORS_ALLOW_LOCAL_ORIGIN"), os.Getenv("CORS_ALLOW_PRODUCTION_CLIENT_ORIGIN"), os.Getenv("CORS_ALLOW_STAGING_CLIENT_ORIGIN")},
		AllowMethods:		[]string{"POST", "PATCH", "PUT", "DELETE", "GET"},
		AllowHeaders: 		[]string{"Origin", "Authorization", "Access-Control-Allow-Origin"},	
		AllowCredentials: 	true,
		MaxAge: 			12*time.Hour,
	})
}