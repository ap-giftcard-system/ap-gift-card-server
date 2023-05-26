package middleware

import (
	"ap-gift-card-server/common"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// @dev Examine access token in headers
//
// @return gin.HandleFunc
func Authenticate() gin.HandlerFunc {
	return func(gc *gin.Context) {
		// prepare custom claim for decoding JWT
		type ApJWTClaims struct {
			AdminCredentialsHash string `json:"adminCredentialsHash"`
			jwt.StandardClaims
		}

		// Get bearer token from Authorization headers
		bearerToken := gc.GetHeader("Authorization")
		if bearerToken == "" {
			gc.AbortWithStatusJSON(401, gin.H{"error": gin.H{
				"key": "!ACCESS_TOKEN",
				"msg": "No authorization header found.",
			}}); return;
		}

		// extra the jwt from bearer token
		accessToken := strings.Split(bearerToken, " ")[1]

		// check if accessToken is not empty
		if accessToken == "" {
			gc.AbortWithStatusJSON(401, gin.H{"error": gin.H{
				"key": "!ACCESS_TOKEN",
				"msg": "Authorization token is empty.",
			}}); return;
		}

		// Decode/validate accessToken
		token, err := jwt.ParseWithClaims(accessToken, &ApJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil {
			gc.AbortWithStatusJSON(401, gin.H{"error": gin.H{
				"key": "!ACCESS_TOKEN",
				"msg": "Cannot parse auth JWT.",
			}}); return;
		}

		// Implement authenticating logic
		if claims, ok := token.Claims.(*ApJWTClaims); ok && token.Valid {
			// prepare real username & password
			iUsername := os.Getenv("AP_ADMIN_USERNAME")
			iPassword := os.Getenv("AP_ADMIN_PASSWORD")
			
			// prepare hash
			credentialHash := common.CalculateHash([]byte(iUsername + iPassword))

			// compare credentialHash and hash from jwt
			if strings.Compare(credentialHash, claims.AdminCredentialsHash) != 0 {
				gc.AbortWithStatusJSON(401, gin.H{"error": gin.H{
					"key": "!ACCESS_TOKEN",
					"msg": "Invalid access token.",
				}}); return;
			}

			// move on
			gc.Next()
		} else {
			gc.AbortWithStatusJSON(401, gin.H{"error": gin.H{
				"key": "!ACCESS_TOKEN",
				"msg": "Invalid access token.",
			}}); return;
		}
	}
}