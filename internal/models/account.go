package models

import "time"

// Account represents a user account in MySQL
type Account struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AccountCreateRequest is the request body for creating an account
type AccountCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AccountLoginRequest is the request body for login
type AccountLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
