package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"presentation-demo/internal/models"
	"presentation-demo/internal/repository"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		repo: repository.NewUserRepository(),
	}
}

// CreateUser handles POST /api/users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.AccountID == 0 || req.Name == "" {
		respondWithError(w, http.StatusBadRequest, "Account ID and name are required")
		return
	}

	user, err := h.repo.Create(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

// GetUser handles GET /api/users/{id}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.repo.GetByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

// GetUserByAccountID handles GET /api/users/account/{account_id}
func (h *UserHandler) GetUserByAccountID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID, err := strconv.Atoi(vars["account_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid account ID")
		return
	}

	user, err := h.repo.GetByAccountID(accountID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

// UpdateUser handles PUT /api/users/{id}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var req models.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Name == "" {
		respondWithError(w, http.StatusBadRequest, "Name is required")
		return
	}

	user, err := h.repo.Update(id, req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
