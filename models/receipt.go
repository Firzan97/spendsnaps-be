package models

import "github.com/kamva/mgm/v3"

type Receipt struct {
	mgm.DefaultModel `bson:",inline"`
	ShopName         string  `json:"shop_name,omitempty" bson:"shop_name,omitempty"`
	CompanyName      string  `json:"company_name,omitempty" bson:"company_name,omitempty"`
	Total            float64 `json:"total,omitempty" bson:"total,omitempty"`
	Text             string  `json:"text,omitempty" bson:"text,omitempty"`
	ImageUrl         string  `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Currency         string  `json:"currency,omitempty" bson:"currency,omitempty" `
}

type UpdateReceipt struct {
	mgm.DefaultModel `bson:",inline"`
	ShopName         *string  `json:"shop_name,omitempty" bson:"shop_name,omitempty"`
	CompanyName      *string  `json:"company_name,omitempty" bson:"company_name,omitempty"`
	Total            *float64 `json:"total,omitempty" bson:"total,omitempty"`
	Text             *string  `json:"text,omitempty" bson:"text,omitempty"`
	ImageUrl         *string  `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Currency         *string  `json:"currency,omitempty" bson:"currency,omitempty" `
}
