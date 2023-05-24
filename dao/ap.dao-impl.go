package dao

import (
	"ap-gift-card-server/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// @notice Root struct for other methods in dao-impl
type ApGiftDaoImpl struct {
	ctx		context.Context
	mongoCollection 	*mongo.Client
}

// @dev GiftDaoConstructor
func ApGiftDaoConstructor(ctx context.Context, mongoClient *mongo.Client) ApGiftDao {
	return &ApGiftDaoImpl{
		ctx: ctx,
		mongoCollection: mongoClient,
	}
}

// @dev Create and add a new ApGiftHolder to internal database
// 
// @param giftHolder *models.ApGiftHolder
// 
// @return bool
// 
// @return error
func (gdi *ApGiftDaoImpl) CreateApGiftHolder(giftHolder *models.ApGiftHolder) (bool, error) {
	return false, nil
}

// @dev Update an existed ApGiftHolder in internal database
// 
// @param giftHolder *models.ApGiftHolder
// 
// @return bool
// 
// @return error
func (gdi *ApGiftDaoImpl) UpdateApGiftHolder(giftHolder *models.ApGiftHolder) (bool, error) {
	return false, nil
}

// @dev Get a list of all ApGiftHolders in internal database
// 
// @return *[]models.ApGiftHolder
// 
// @return error
func (gdi *ApGiftDaoImpl) GetAllApGiftHolders() (*[]models.ApGiftHolder, error) {
	return nil, nil
}

// @dev Get a specific ApGiftHolder by `BarCode`
// 
// @param barCode string
// 	
// @param holderName string
// 
// @param holderPhone string
// 
// @param holderEmail string
// 
// @return *model.ApGiftHolder
// 
// @return *model.ApGiftHolder
// 
// @return error
func (gdi *ApGiftDaoImpl) GetApGiftHolder(barCode, holderName, holderPhone, holderEmail string) (*models.ApGiftHolder, error) {
	return nil, nil
}

// @dev Remove a specific ApGiftHolder by `BarCode`
// 
// @param barCode string
// 
// @return bool
// 
// @return error
func (gdi *ApGiftDaoImpl) DeleteApHolderByBarCode(barCode string) (bool, error) {
	return false, nil
}