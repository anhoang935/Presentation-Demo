package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order represents an order in MongoDB
type Order struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AccountID    int                `bson:"account_id" json:"account_id"`
	FoodID       int                `bson:"food_id" json:"food_id"`
	RestaurantID int                `bson:"restaurant_id" json:"restaurant_id"`
	TotalPrice   float64            `bson:"total_price" json:"total_price"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
}

// OrderCreateRequest is the request body for creating an order
type OrderCreateRequest struct {
	AccountID    int     `json:"account_id"`
	FoodID       int     `json:"food_id"`
	RestaurantID int     `json:"restaurant_id"`
	TotalPrice   float64 `json:"total_price"`
}
