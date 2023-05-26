package routers

import (
	"ap-gift-card-server/controllers"

	"github.com/gin-gonic/gin"
)

// @notice Root struct for other methods in router
type ApGiftRouter struct {
	ApGiftController *controllers.ApGiftController
}

// @dev Constructor
func ApGiftRouterConstructor(apGiftController *controllers.ApGiftController) *ApGiftRouter {
	return &ApGiftRouter{
		ApGiftController: apGiftController,
	}
}

// @dev Declares list of gift holder endpoints
func (agr *ApGiftRouter) ApRouter (rg *gin.RouterGroup) {
	rg.POST("/register", agr.ApGiftController.RegisterNewApGiftHoder)
	rg.PATCH("/update", agr.ApGiftController.UpdateApGiftHolder)
	rg.GET("/find-gift-holders", agr.ApGiftController.GetApGiftHolder)
	rg.DELETE("/delete", agr.ApGiftController.DeleteApGiftHolder)
}

// @dev Declares auth endpoint
func AuthRouter (rg *gin.RouterGroup) {
	rg.POST("/generate-access-token", controllers.GenerateAccessToken)
}