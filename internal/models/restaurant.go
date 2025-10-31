package models

// Restaurant represents a restaurant (constant data, not from database)
type Restaurant struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Cuisine string `json:"cuisine"`
}

// GetRestaurants returns all available restaurants
func GetRestaurants() []Restaurant {
	return []Restaurant{
		{ID: 1, Name: "Pizza Palace", Address: "123 Main St", Cuisine: "Italian"},
		{ID: 2, Name: "Sushi World", Address: "456 Oak Ave", Cuisine: "Japanese"},
		{ID: 3, Name: "Burger House", Address: "789 Elm St", Cuisine: "American"},
		{ID: 4, Name: "Pasta Paradise", Address: "321 Pine Rd", Cuisine: "Italian"},
		{ID: 5, Name: "Taco Town", Address: "654 Maple Dr", Cuisine: "Mexican"},
	}
}

// GetRestaurantByID returns a restaurant by ID
func GetRestaurantByID(id int) *Restaurant {
	restaurants := GetRestaurants()
	for _, r := range restaurants {
		if r.ID == id {
			return &r
		}
	}
	return nil
}
