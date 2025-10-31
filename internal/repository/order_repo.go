package repository

import (
	"context"
	"fmt"
	"time"

	"presentation-demo/internal/database"
	"presentation-demo/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		collection: database.MongoDB.Collection("orders"),
	}
}

// Create creates a new order
func (r *OrderRepository) Create(req models.OrderCreateRequest) (*models.Order, error) {
	order := models.Order{
		AccountID:    req.AccountID,
		FoodID:       req.FoodID,
		RestaurantID: req.RestaurantID,
		TotalPrice:   req.TotalPrice,
		CreatedAt:    time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("error creating order: %w", err)
	}

	order.ID = result.InsertedID.(primitive.ObjectID)
	return &order, nil
}

// GetByID retrieves an order by ID
func (r *OrderRepository) GetByID(id string) (*models.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid order ID: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var order models.Order
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&order)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("order not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error getting order: %w", err)
	}

	return &order, nil
}

// GetByAccountID retrieves all orders for an account
func (r *OrderRepository) GetByAccountID(accountID int) ([]models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"account_id": accountID})
	if err != nil {
		return nil, fmt.Errorf("error finding orders: %w", err)
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	if err := cursor.All(ctx, &orders); err != nil {
		return nil, fmt.Errorf("error decoding orders: %w", err)
	}

	return orders, nil
}

// GetAll retrieves all orders
func (r *OrderRepository) GetAll() ([]models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error finding orders: %w", err)
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	if err := cursor.All(ctx, &orders); err != nil {
		return nil, fmt.Errorf("error decoding orders: %w", err)
	}

	return orders, nil
}
