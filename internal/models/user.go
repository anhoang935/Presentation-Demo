package models

import "time"

// User represents a user profile in MySQL
type User struct {
	ID        int       `json:"id"`
	AccountID int       `json:"account_id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserCreateRequest is the request body for creating a user
type UserCreateRequest struct {
	AccountID int    `json:"account_id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
}

// UserUpdateRequest is the request body for updating a user
type UserUpdateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
