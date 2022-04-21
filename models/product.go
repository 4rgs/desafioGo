package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//tad
type Product struct {
	_id         primitive.ObjectID `bson:"_id,omitempty"`
	ID          int32              `json:"id,omitempty"`
	Brand       string             `json:"brand"`
	Description string             `json:"description"`
	Image       string             `json:"image"`
	Price       float32            `json:"price"`
}

// lista de productos
type Products []*Product
