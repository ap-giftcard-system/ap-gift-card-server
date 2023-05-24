package dao

import (
	"ap-gift-card-server/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// @notice Root struct for other methods in dao-impl
type ApGiftDaoImpl struct {
	ctx		context.Context
	mongoCollection 	*mongo.Collection
}

// @dev GiftDaoConstructor
func ApGiftDaoConstructor(ctx context.Context, mongoClient *mongo.Collection) ApGiftDao {
	return &ApGiftDaoImpl{
		ctx: ctx,
		mongoCollection: mongoClient,
	}
}

// @dev Create and add a new ApGiftHolder to internal database
// 
// @param giftHolder *models.ApGiftHolder
// 
// @return error
func (gdi *ApGiftDaoImpl) RegisterNewApGiftHoder(giftHolder *models.ApGiftHolder) (error) {
	// prepare filter
	filter := bson.D{{Key: "bar_code", Value: giftHolder.BarCode}}

	// check if there's already a document matching `filter`
	dbRes := gdi.mongoCollection.FindOne(gdi.ctx, filter)
	
	// @logic if dbRes.Err() == nil => a document already exists with given `giftHolder.BarCode` => 409 - CONFLICT
	// @logic if dbRes.Err() == mongo.ErrNoDocuments => document with given `giftHolder.BarCode` is a valid new document => add document to database
	// @logic else => uknown error occurs in internal database
	if dbRes.Err() == nil {
		return errors.New("ErrDocumentConflict")
	} else if dbRes.Err() == mongo.ErrNoDocuments {
		// prepare uuid for giftHolder.GiftHolderID
		giftHolder.GiftHolderID = primitive.NewObjectID().Hex()

		// prepare timing for the `giftHolder` document
		giftHolder.CreatedAt = time.Now() // UTC
		giftHolder.UpdatedAt = time.Now() // UTC

		// add new `giftHolder` to database
		if _, err := gdi.mongoCollection.InsertOne(gdi.ctx, giftHolder); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return dbRes.Err()
	}
}

// @dev Update an existed ApGiftHolder in internal database
// 
// @param giftHolder *models.ApGiftHolder
// 
// @return error
func (gdi *ApGiftDaoImpl) UpdateApGiftHolder(giftHolder *models.ApGiftHolder) (error) {
	// prepare filter
	filter := bson.D{{Key: "_id", Value: giftHolder.GiftHolderID}}

	// find the document matching `filter`
	dbRes := gdi.mongoCollection.FindOne(gdi.ctx, filter)

	// @logic if dbRes.Err() == nil => document found => update document
	// @logic if dbRes.Err() == mongo.ErrNoDocuments => document not found => 404 !OK
	// @logic else => unknown error occurs => 500
	if dbRes.Err() == nil {
		// prepare update query
		update := bson.D{
			{Key: "$set", Value: bson.D{{Key: "bar_code", Value: giftHolder.BarCode}}},
			{Key: "$set", Value: bson.D{{Key: "holder_name", Value: giftHolder.HolderName}}},
			{Key: "$set", Value: bson.D{{Key: "holder_phone", Value: giftHolder.HolderPhone}}},
			{Key: "$set", Value: bson.D{{Key: "holder_email", Value: giftHolder.HolderEmail}}},
			{Key: "$set", Value: bson.D{{Key: "gift_amount", Value: giftHolder.GiftAmount}}},
			{Key: "$set", Value: bson.D{{Key: "updated_at", Value: time.Now()}}},
		}

		// update document
		if _, err := gdi.mongoCollection.UpdateOne(gdi.ctx, filter, update); err != nil {
			return errors.New("cannot update document")
		} else {
			return nil
		}
	} else if dbRes.Err() == mongo.ErrNoDocuments {
		return errors.New("ErrNoDocuments")
	} else {
		return dbRes.Err()
	}
}


// @dev Get ApGiftHolders by params
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
// @return *model.ApGiftHolder
// 
// @return *[]model.ApGiftHolder
// 
// @return error
func (gdi *ApGiftDaoImpl) GetApGiftHolder(barCode, holderName, holderPhone, holderEmail string) (*models.ApGiftHolder, error) {
	return nil, nil
}

// @dev Remove a specific ApGiftHolder
// 
// @param holderID string
// 
// @return error
func (gdi *ApGiftDaoImpl) DeleteApHolder(holderID string) (error) {
	return nil
}