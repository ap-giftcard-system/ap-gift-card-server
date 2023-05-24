package controllers

import (
	"ap-gift-card-server/dao"
	"ap-gift-card-server/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @notice global var
var validate = validator.New()

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
	// prepare `param` placeholder
	param := &models.ApGiftHolder{}

	// bind json post data to `param`
	if err := gc.ShouldBindJSON(param); err != nil {
		gc.AbortWithStatusJSON(400, gin.H{"error": gin.H{
			"key": "!BAD_REQUEST",
			"msg": err.Error(),
		}}); return;
	}

	// validate struct
	if err := validate.Struct(param); err != nil {
		gc.AbortWithStatusJSON(400, gin.H{"error": gin.H{
			"key": "!BAD_REQUEST",
			"msg": err.Error(),
		}}); return;
	}

	// @logic post data must have either HolderPhone or HolderEmail
	if strings.EqualFold(param.HolderEmail, "") && strings.EqualFold(param.HolderPhone, "") {
		gc.AbortWithStatusJSON(400, gin.H{"error": gin.H{
			"key": "!MISSING_CONTACT", 
			"msg": "Gift Holder must provide at least either email or phone information for contacting purposes",
		}}); return;
	}

	// invoke dao.RegisterNewApGiftHoder
	if err := agc.ApGiftDao.RegisterNewApGiftHoder(param); err != nil {
		// @logic abort request with a conflict status if Gift Holder already exists
		if strings.EqualFold(err.Error(), "ErrDocumentConflict") {
			gc.AbortWithStatusJSON(409, gin.H{"error": gin.H{
				"key": "!DOCUMENT_CONFLICT",
				"msg": "Existed document found in result",
			}}); return;
		} else {
			// @logic abort request with a 500 if there's a unknown error from internal database
			gc.AbortWithStatusJSON(500, gin.H{"error": gin.H{
				"key": "!INTERNAL_SERVER",
				"msg": err.Error(),
			}}); return;
		}
	}
	
	// return 200 OK to client
	gc.JSON(200, gin.H{"error": nil})
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