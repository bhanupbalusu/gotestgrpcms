package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductDetails struct {
	ProductName string `json:"product_name,omitempty" bson:"product_name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	ImageUrl    string `json:"ImageUrl,omitempty" bson:"ImageUrl,omitempty"`
}

type BulkQuantity struct {
	Volume string `json:"volume,omitempty" bson:"volume,omitempty"`
	Units  string `json:"units,omitempty" bson:"units,omitempty"`
}

type Price struct {
	Amount   string `json:"amount,omitempty" bson:"amount,omitempty"`
	Currency string `json:"currency,omitempty" bson:"currency,omitempty"`
	PerUnit  string `json:"per_unit,omitempty" bson:"per_unit,omitempty"`
	Units    string `json:"units,omitempty" bson:"units,omitempty"`
}

type QuantityDetails struct {
	BulkQuantity BulkQuantity `json:"bulk_quantity,omitempty" bson:"bulk_quantity,omitempty"`
	Price        Price        `json:"price,omitempty" bson:"price,omitempty"`
}

type Schedular struct {
	StartDate string `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty" bson:"end_date,omitempty"`
}

type ProductModel struct {
	ProductID         primitive.ObjectID `json:"product_id,omitempty" bson:"_id,omitempty"`
	PreOrderRequestId string             `json:"pre_order_request_id,omitempty" bson:"pre_order_request_id,omitempty"`
	CustomerId        string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
	ProductDetails    ProductDetails     `json:"product_details,omitempty" bson:"product_details,omitempty"`
	QuantityDetails   QuantityDetails    `json:"quantity_details,omitempty" bson:"quantity_details,omitempty"`
	Schedular         Schedular          `json:"schedular,omitempty" bson:"schedular,omitempty"`
	CreatedAt         int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt         int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
