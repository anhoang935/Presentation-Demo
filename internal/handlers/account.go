package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"presentation-demo/internal/models"
	"presentation-demo/internal/repository"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	repo *repository.AccountRepository
}

func NewAccountHandler() *AccountHandler {
	return &AccountHandler{
		repo: repository.NewAccountRepository(),
	}
}

// CreateAccount handles POST /api/accounts
func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req models.AccountCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Email == "" || req.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Email and password are required")
		return
	}

	account, err := h.repo.Create(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, account)
}

// GetAccount handles GET /api/accounts/{id}
func (h *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid account ID")
		return
	}

	account, err := h.repo.GetByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, account)
}

// Login handles POST /api/accounts/login
func (h *AccountHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.AccountLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Email == "" || req.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Email and password are required")
		return
	}

	account, err := h.repo.GetByEmail(req.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	if err := h.repo.ValidatePassword(account, req.Password); err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Don't send password in response
	account.Password = ""
	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"account": account,
	})
}
