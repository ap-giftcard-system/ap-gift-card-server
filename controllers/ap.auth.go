package controllers

import (
	"ap-gift-card-server/utils"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// @route `POST/generate-access-token`
func GenerateAccessToken(gc *gin.Context) {
	// prepare param struct
	type paramStruct struct {
		AdminUsername string `json:"admin-username" validate:"required"`
		AdminPassword string `json:"admin-password" validate:"required"`
	}

	// prepare param placeholder
	param := &paramStruct{}

	// bind json post data to `param`
	if err := gc.ShouldBindJSON(param); err != nil {
		gc.AbortWithStatusJSON(400, gin.H{
			"access-token": nil,
			"error": gin.H{
				"key": "!BAD_REQUEST",
				"msg": err.Error(),
		}}); return;
	}
	log.Println(param)

	// validate struct
	if err := validate.Struct(param); err != nil {
		gc.AbortWithStatusJSON(400, gin.H{
			"access-token": nil,
			"error": gin.H{
				"key": "!BAD_REQUEST",
				"msg": err.Error(),
		}}); return;
	}

	// security check on param
	ok, hash := utils.ValidateAdminLogin(param.AdminUsername, param.AdminPassword) 
	if !ok {
		gc.AbortWithStatusJSON(401, gin.H{"error": gin.H{
			"key": "!UNAUTHORIZED",
			"msg": "Invalid admin credentials.",
		}}); return;
	}

	// Create a new jwt token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"adminCredentialsHash": hash,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		gc.AbortWithStatusJSON(500, gin.H{"access-token": nil, "error": gin.H{
			"key": "!JWT_ISSUE",
			"msg": "Unable to sign new jwt access token.",
		}}); return;
	}

	// return 200 OK to client
	gc.JSON(200, gin.H{"access-token": accessToken, "error": nil})
}