package dao

import (
	"ap-gift-card-server/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func (gdi *ApGiftDaoImpl) UpdateApGiftHolder(giftHolder *models.ApGiftHolder) (*models.ApGiftHolder, error) {
	// prepare apUpdatedGiftHolder
	apUpdatedGiftHolder := &models.ApGiftHolder{}

	// prepare filter
	filter := bson.D{{Key: "_id", Value: giftHolder.GiftHolderID}}

	// find the document matching `filter`
	dbRes := gdi.mongoCollection.FindOne(gdi.ctx, filter)

	// Set the options to return the updated document
	mongoOptions := options.FindOneAndUpdate().SetReturnDocument(options.After)

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
		if err := gdi.mongoCollection.FindOneAndUpdate(gdi.ctx, filter, update, mongoOptions).Decode(apUpdatedGiftHolder); err != nil {
			return nil, errors.New("cannot update document")
		} else {
			return apUpdatedGiftHolder, nil
		}
	} else if dbRes.Err() == mongo.ErrNoDocuments {
		return nil, errors.New("ErrNoDocuments")
	} else {
		return nil, dbRes.Err()
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
func (gdi *ApGiftDaoImpl) GetApGiftHolders(barCode, holderName, holderPhone, holderEmail string) (*[]models.ApGiftHolder, error) {
	// prepare apGiftHolder placeholder
	apGiftHolders := &[]models.ApGiftHolder{}

	// prepare filter 
	filter := bson.M{}
	if barCode != "" {
		filter["bar_code"] = barCode
	}
	if holderName != "" {
		filter["holder_name"] = primitive.Regex{Pattern: "^" + holderName + "$", Options: "i"}
	}
	if holderPhone != "" {
		filter["holder_phone"] = holderPhone
	}
	if holderEmail != "" {
		filter["holder_email"] = primitive.Regex{Pattern: "^" + holderEmail + "$", Options: "i"}
	}

	// sort documents in ascending order of "created_at" field
	mongoOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	// find documents matches filter
	cursor, err := gdi.mongoCollection.Find(gdi.ctx, filter, mongoOptions)
	if err != nil {
		return nil, err
	}

	// decode cursor to `apGiftHolders`
	err = cursor.All(gdi.ctx, apGiftHolders)
	
	return apGiftHolders, err
}

// @dev Remove a specific ApGiftHolder
// 
// @param holderId string
// 
// @return int64 - delete document counts
// 
// @return error
func (gdi *ApGiftDaoImpl) DeleteApGiftHolder(holderId string) (int64, error) {
	// prepare filter
	filter := bson.D{{Key: "_id", Value: holderId}}

	// remove document matching filter
	dbRes, err := gdi.mongoCollection.DeleteOne(gdi.ctx, filter)

	return dbRes.DeletedCount, err
}