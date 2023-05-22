package routers

import "github.com/gin-gonic/gin"

// @dev Declares list of endpoints
func ApRouter (rg *gin.RouterGroup) {
	rg.GET("/ping", func(gc *gin.Context) {
		gc.JSON(200, "pong")
	})
}