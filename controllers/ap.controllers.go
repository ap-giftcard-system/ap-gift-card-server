package controllers

import (
	"ap-gift-card-server/dao"

	"github.com/gin-gonic/gin"
)

// @notice Root struct for other methods in controller
type ApGiftController struct {
	ApGiftDao dao.ApGiftDao
}

// @dev ApGiftControllerContructor
func ApGiftControllerContructor(apGiftDao *dao.ApGiftDao) *ApGiftController {
	return &ApGiftController{
		ApGiftDao: *apGiftDao,
	}
}

// @route `POST/register`
// 
// @dev Create and add a new ApGiftHolder to internal database
// 
// @param gc *gin.Context
func (agc *ApGiftController) RegisterNewApGiftHoder(gc *gin.Context) {
	gc.JSON(200, "true")
}

// @route `PUT/update`
// 
// @dev Update an existed ApGiftHolder in internal database
// 
// @param gc *gin.Context
func (agc *ApGiftController) UpdateApGiftHolder(gc *gin.Context) {
	gc.JSON(200, "true")
}

// @route `GET/all`
// 
// @dev Get a list of all ApGiftHolders in internal database
// 
// @param gc *gin.Context
func (agc *ApGiftController) GetAllApGiftHolders(gc *gin.Context) {
	gc.JSON(200, "true")
}

// @route `GET/single?bar-code=?holder-name=?holder-phone=?holder-email=`
// 
// @dev Get a specific ApGiftHolder
// 
// @param gc *gin.Context
func (agc *ApGiftController) GetApGiftHolder(gc *gin.Context) {
	gc.JSON(200, "true")
}

// @route `DELETE/delete?holder-id=`
// 
// @dev Remove a specific ApGiftHolder
// 
// @param gc *gin.Context
func (agc *ApGiftController) DeleteApGiftHolder(gc *gin.Context) {
	gc.JSON(200, "true")
}