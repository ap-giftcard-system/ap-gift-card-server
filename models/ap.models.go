package models

import "time"

// @notice struct for AP Gift Holder
type ApGiftHolder struct {
	GiftHolderID string		`json:"giftHolderId" bson:"_id" validate:"omitempty"`
	BarCode string 			`json:"barCode" bson:"bar_code" validate:"required"`
	HolderName string 		`json:"holderName" bson:"holder_name" validate:"required"`
	HolderPhone string		`json:"holderPhone" bson:"holder_phone" validate:"omitempty,e164"` // e164Pattern = `^\+\d{1,15}$`
	HolderEmail string		`json:"holderEmail" bson:"holder_email" validate:"omitempty,email"`
	GiftAmount float64		`json:"giftAmount" bson:"gift_amount" validate:"omitempty,number"`
	CreatedAt time.Time 		`json:"createdAt" bson:"created_at,omitempty" validate:"omitempty"`
	UpdatedAt time.Time 		`json:"updateAt" bson:"updated_at,omitempty" validate:"omitempty"`
}