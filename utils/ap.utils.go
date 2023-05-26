package utils

import (
	"ap-gift-card-server/common"
	"log"
	"os"
	"strings"
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

// @dev Security check on admin login
// 
// @param username string
// 
// @param password string
// 
// @return ok bool
// 
// @return hash string - keccak256 hex string of the iUsername and iPassword
func ValidateAdminLogin(username, password string) (bool, string) {
	// prepare real username & password
	iUsername := os.Getenv("AP_ADMIN_USERNAME")
	iPassword := os.Getenv("AP_ADMIN_PASSWORD")
	
	// prepare hash
	credentialHash := common.CalculateHash([]byte(iUsername + iPassword))

	// security checks
	if strings.Compare(username, iUsername) != 0 || strings.Compare(password, iPassword) != 0 {
		return false, ""
	} else {
		return true, credentialHash
	}
}