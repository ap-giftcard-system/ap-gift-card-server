package routers

import (
	"ap-gift-card-server/controllers"
	"ap-gift-card-server/middleware"

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
	rg.POST("/register", middleware.Authenticate(), agr.ApGiftController.RegisterNewApGiftHoder)
	rg.PATCH("/update", middleware.Authenticate(), agr.ApGiftController.UpdateApGiftHolder)
	rg.GET("/find-gift-holders", middleware.Authenticate(), agr.ApGiftController.GetApGiftHolder)
	rg.DELETE("/delete", middleware.Authenticate(), agr.ApGiftController.DeleteApGiftHolder)
}

// @dev Declares auth endpoint
func AuthRouter (rg *gin.RouterGroup) {
	rg.POST("/generate-access-token", controllers.GenerateAccessToken)
}