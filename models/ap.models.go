package models

import "time"

// @notice struct for AP Gift Holder
type ApGiftHolder struct {
	BarCode string 			`json:"barCode" bson:"bar_code" validate:"required"`
	HolderName string 		`json:"holderName" bson:"holder_name" validate:"required"`
	HolderPhone string		`json:"holderPhone" bson:"holder_phone" validate:"e164"`
	GiftAmount uint64		`json:"giftAmount" bson:"gift_amount" validate:"required,number"`
	CreatedAt time.Time 	`json:"createdAt" bson:"created_at,omitempty" validate:"omitempty"`
	UpdatedAt time.Time 	`json:"updateAt" bson:"updated_at,omitempty" validate:"omitempty"`
	DeletedAt time.Time 	`json:"deleteAt" bson:"deleted_at,omitempty" validate:"omitempty"`
}