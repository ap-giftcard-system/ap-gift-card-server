package common

import (
	"ap-gift-card-server/models"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/sha3"
)

// @dev sanitize struct received from request's body
//
// @param gc *gin.Context
//
// @param validate *validator.Validate
//
// @param giftHolder *models.ApGiftHolder
//
// @return error
func SanitizeStruct( gc *gin.Context, validate *validator.Validate, giftHolder *models.ApGiftHolder) error {
	// validate struct
	if err := validate.Struct(giftHolder); err != nil {
		gc.AbortWithStatusJSON(400, gin.H{"error": gin.H{
			"key": "!BAD_REQUEST",
			"msg": err.Error(),
		}})
		return errors.New("!BAD_REQUEST")
	}

	// @logic giftHolder must have either HolderPhone or HolderEmail
	if strings.EqualFold(giftHolder.HolderEmail, "") && strings.EqualFold(giftHolder.HolderPhone, "") {
		gc.AbortWithStatusJSON(400, gin.H{"error": gin.H{
			"key": "!MISSING_CONTACT", 
			"msg": "Gift Holder must provide at least either email or phone information for contacting purposes",
		}})
		return errors.New("!MISSING_CONTACT")
	}

	return nil
}

// @dev calculate a keccak256 hashing solution for admin credentials
// 
// @param input []byte
// 
// @return string
func CalculateHash(input []byte) string {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(input)
	result := hash.Sum(nil)

	return hex.EncodeToString(result)
}