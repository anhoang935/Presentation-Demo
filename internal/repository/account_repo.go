package repository

import (
	"database/sql"
	"fmt"

	"presentation-demo/internal/database"
	"presentation-demo/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type AccountRepository struct{}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

// Create creates a new account
func (r *AccountRepository) Create(req models.AccountCreateRequest) (*models.Account, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	// Insert into database
	result, err := database.MySQLDB.Exec(
		"INSERT INTO Account (email, password) VALUES (?, ?)",
		req.Email, string(hashedPassword),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating account: %w", err)
	}

	// Get the inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert ID: %w", err)
	}

	// Retrieve the created account
	return r.GetByID(int(id))
}

// GetByID retrieves an account by ID
func (r *AccountRepository) GetByID(id int) (*models.Account, error) {
	account := &models.Account{}
	err := database.MySQLDB.QueryRow(
		"SELECT id, email, created_at, updated_at FROM Account WHERE id = ?",
		id,
	).Scan(&account.ID, &account.Email, &account.CreatedAt, &account.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("account not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error getting account: %w", err)
	}

	return account, nil
}

// GetByEmail retrieves an account by email
func (r *AccountRepository) GetByEmail(email string) (*models.Account, error) {
	account := &models.Account{}
	err := database.MySQLDB.QueryRow(
		"SELECT id, email, password, created_at, updated_at FROM Account WHERE email = ?",
		email,
	).Scan(&account.ID, &account.Email, &account.Password, &account.CreatedAt, &account.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("account not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error getting account: %w", err)
	}

	return account, nil
}

// ValidatePassword validates the password for an account
func (r *AccountRepository) ValidatePassword(account *models.Account, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
}
