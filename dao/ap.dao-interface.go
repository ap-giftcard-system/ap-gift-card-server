package dao

import "ap-gift-card-server/models"

// @notice GiftDao interface
type ApGiftDao interface {

	// @notice Create and add a new ApGiftHolder to internal database
	// 
	// @param giftHolder *models.ApGiftHolder
	// 
	// @return error
	RegisterNewApGiftHoder(giftHolder *models.ApGiftHolder) (error)

	// @notice Update an existed ApGiftHolder in internal database
	// 
	// @param giftHolder *models.ApGiftHolder
	// 
	// @return error
	UpdateApGiftHolder(giftHolder *models.ApGiftHolder) (error)

	// @notice Get a list of ApGiftHolders by params
	// 
	// @notice if all params are empty (i.e. no params provided), retrives all Gift Holders
	// 
	// @param barCode string
	// 
	// @param holderName string
	// 
	// @param holderPhone string
	// 
	// @param holderEmail string
	// 
	// @return *[]model.ApGiftHolder
	// 
	// @return error
	GetApGiftHolders(barCode, holderName, holderPhone, holderEmail string) (*[]models.ApGiftHolder, error)

	// @notice Remove a specific ApGiftHolder
	// 
	// @param holderID string
	// 
	// @return error
	DeleteApHolder(holderID string) (error)

}