package handlers

import (
	"net/http"
	"strconv"

	"presentation-demo/internal/models"

	"github.com/gorilla/mux"
)

type StaticHandler struct{}

func NewStaticHandler() *StaticHandler {
	return &StaticHandler{}
}

// GetRestaurants handles GET /api/restaurants
func (h *StaticHandler) GetRestaurants(w http.ResponseWriter, r *http.Request) {
	restaurants := models.GetRestaurants()
	respondWithJSON(w, http.StatusOK, restaurants)
}

// GetRestaurant handles GET /api/restaurants/{id}
func (h *StaticHandler) GetRestaurant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid restaurant ID")
		return
	}

	restaurant := models.GetRestaurantByID(id)
	if restaurant == nil {
		respondWithError(w, http.StatusNotFound, "Restaurant not found")
		return
	}

	respondWithJSON(w, http.StatusOK, restaurant)
}

// GetFoods handles GET /api/foods
func (h *StaticHandler) GetFoods(w http.ResponseWriter, r *http.Request) {
	foods := models.GetFoods()
	respondWithJSON(w, http.StatusOK, foods)
}

// GetFood handles GET /api/foods/{id}
func (h *StaticHandler) GetFood(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid food ID")
		return
	}

	food := models.GetFoodByID(id)
	if food == nil {
		respondWithError(w, http.StatusNotFound, "Food not found")
		return
	}

	respondWithJSON(w, http.StatusOK, food)
}

// GetFoodsByRestaurant handles GET /api/restaurants/{id}/foods
func (h *StaticHandler) GetFoodsByRestaurant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid restaurant ID")
		return
	}

	foods := models.GetFoodsByRestaurantID(id)
	respondWithJSON(w, http.StatusOK, foods)
}
