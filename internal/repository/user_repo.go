package repository

import (
	"database/sql"
	"fmt"

	"presentation-demo/internal/database"
	"presentation-demo/internal/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Create creates a new user
func (r *UserRepository) Create(req models.UserCreateRequest) (*models.User, error) {
	result, err := database.MySQLDB.Exec(
		"INSERT INTO User (account_id, name, address) VALUES (?, ?, ?)",
		req.AccountID, req.Name, req.Address,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert ID: %w", err)
	}

	return r.GetByID(int(id))
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(id int) (*models.User, error) {
	user := &models.User{}
	err := database.MySQLDB.QueryRow(
		"SELECT id, account_id, name, address, created_at, updated_at FROM User WHERE id = ?",
		id,
	).Scan(&user.ID, &user.AccountID, &user.Name, &user.Address, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return user, nil
}

// GetByAccountID retrieves a user by account ID
func (r *UserRepository) GetByAccountID(accountID int) (*models.User, error) {
	user := &models.User{}
	err := database.MySQLDB.QueryRow(
		"SELECT id, account_id, name, address, created_at, updated_at FROM User WHERE account_id = ?",
		accountID,
	).Scan(&user.ID, &user.AccountID, &user.Name, &user.Address, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return user, nil
}

// Update updates a user
func (r *UserRepository) Update(id int, req models.UserUpdateRequest) (*models.User, error) {
	_, err := database.MySQLDB.Exec(
		"UPDATE User SET name = ?, address = ? WHERE id = ?",
		req.Name, req.Address, id,
	)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}

	return r.GetByID(id)
}
