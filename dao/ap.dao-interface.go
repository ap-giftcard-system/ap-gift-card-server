package dao

import "ap-gift-card-server/models"

// @notice GiftDao interface
type ApGiftDao interface {

	// @notice Create and add a new ApGiftHolder to internal database
	// 
	// @param giftHolder *models.ApGiftHolder
	// 
	// @return bool
	// 
	// @return error
	CreateApGiftHolder(giftHolder *models.ApGiftHolder) (bool, error)

	// @notice Update an existed ApGiftHolder in internal database
	// 
	// @param giftHolder *models.ApGiftHolder
	// 
	// @return bool
	// 
	// @return error
	UpdateApGiftHolder(giftHolder *models.ApGiftHolder) (bool, error)

	// @notice Get a list of all ApGiftHolders in internal database
	// 
	// @return *[]models.ApGiftHolder
	// 
	// @return error
	GetAllApGiftHolders() (*[]models.ApGiftHolder, error)

	// @notice Get a specific ApGiftHolder by `BarCode`
	// 
	// @param barCode string
	// 
	// @return *model.ApGiftHolder
	// 
	// @return error
	GetApGiftHolderByBarCode(barCode string) (*models.ApGiftHolder, error)

	// @notice Get a specific ApGiftHolder by `HolderName`. The return object is an array as many records can have same name
	// 
	// @param holderName string
	// 
	// @return *[]model.ApGiftHolder
	// 
	// @return error
	GetApGiftHolderByHolderName(holderName string) (*[]models.ApGiftHolder, error)

	// @notice Get a specific ApGiftHolder by `HolderPhone`
	// 
	// @param holderPhone string
	// 
	// @return *model.ApGiftHolder
	// 
	// @return error
	GetApGiftHolderByHolderPhone(holderPhone string) (*models.ApGiftHolder, error)

	// @notice Get a specific ApGiftHolder by `HolderEmail`
	// 
	// @param holderEmail string
	// 
	// @return *model.ApGiftHolder
	// 
	// @return error
	GetApGiftHolderByHolderEmail(holderEmail string) (*models.ApGiftHolder, error)
	
	// @notice Remove a specific ApGiftHolder by `BarCode`
	// 
	// @param barCode string
	// 
	// @return bool
	// 
	// @return error
	DeleteApHolderByBarCode(barCode string) (bool, error)

}