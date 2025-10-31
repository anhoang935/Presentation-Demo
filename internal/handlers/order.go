package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"presentation-demo/internal/models"
	"presentation-demo/internal/repository"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	repo *repository.OrderRepository
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{
		repo: repository.NewOrderRepository(),
	}
}

// CreateOrder handles POST /api/orders
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req models.OrderCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.AccountID == 0 || req.FoodID == 0 || req.RestaurantID == 0 || req.TotalPrice <= 0 {
		respondWithError(w, http.StatusBadRequest, "All fields are required and total_price must be positive")
		return
	}

	// Validate that the food exists
	food := models.GetFoodByID(req.FoodID)
	if food == nil {
		respondWithError(w, http.StatusBadRequest, "Invalid food ID")
		return
	}

	// Validate that the restaurant exists
	restaurant := models.GetRestaurantByID(req.RestaurantID)
	if restaurant == nil {
		respondWithError(w, http.StatusBadRequest, "Invalid restaurant ID")
		return
	}

	// Validate that the food belongs to the restaurant
	if food.RestaurantID != req.RestaurantID {
		respondWithError(w, http.StatusBadRequest, "Food does not belong to the specified restaurant")
		return
	}

	order, err := h.repo.Create(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, order)
}

// GetOrder handles GET /api/orders/{id}
func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	order, err := h.repo.GetByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, order)
}

// GetOrdersByAccountID handles GET /api/orders/account/{account_id}
func (h *OrderHandler) GetOrdersByAccountID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID, err := strconv.Atoi(vars["account_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid account ID")
		return
	}

	orders, err := h.repo.GetByAccountID(accountID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, orders)
}

// GetAllOrders handles GET /api/orders
func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.repo.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, orders)
}
