package models

// Food represents a food item (constant data, not from database)
type Food struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	RestaurantID int     `json:"restaurant_id"`
	Category     string  `json:"category"`
}

// GetFoods returns all available food items
func GetFoods() []Food {
	return []Food{
		{ID: 1, Name: "Margherita Pizza", Price: 12.99, RestaurantID: 1, Category: "Pizza"},
		{ID: 2, Name: "Pepperoni Pizza", Price: 14.99, RestaurantID: 1, Category: "Pizza"},
		{ID: 3, Name: "California Roll", Price: 8.99, RestaurantID: 2, Category: "Sushi"},
		{ID: 4, Name: "Salmon Nigiri", Price: 10.99, RestaurantID: 2, Category: "Sushi"},
		{ID: 5, Name: "Classic Burger", Price: 9.99, RestaurantID: 3, Category: "Burger"},
		{ID: 6, Name: "Cheese Burger", Price: 10.99, RestaurantID: 3, Category: "Burger"},
		{ID: 7, Name: "Spaghetti Carbonara", Price: 13.99, RestaurantID: 4, Category: "Pasta"},
		{ID: 8, Name: "Fettuccine Alfredo", Price: 12.99, RestaurantID: 4, Category: "Pasta"},
		{ID: 9, Name: "Beef Tacos", Price: 7.99, RestaurantID: 5, Category: "Tacos"},
		{ID: 10, Name: "Chicken Quesadilla", Price: 9.99, RestaurantID: 5, Category: "Mexican"},
	}
}

// GetFoodByID returns a food item by ID
func GetFoodByID(id int) *Food {
	foods := GetFoods()
	for _, f := range foods {
		if f.ID == id {
			return &f
		}
	}
	return nil
}

// GetFoodsByRestaurantID returns all food items for a specific restaurant
func GetFoodsByRestaurantID(restaurantID int) []Food {
	foods := GetFoods()
	var result []Food
	for _, f := range foods {
		if f.RestaurantID == restaurantID {
			result = append(result, f)
		}
	}
	return result
}
