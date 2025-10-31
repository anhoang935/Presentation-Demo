package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"presentation-demo/internal/database"
	"presentation-demo/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize databases
	if err := database.InitMySQL(); err != nil {
		log.Fatalf("Failed to initialize MySQL: %v", err)
	}
	defer database.CloseMySQL()

	if err := database.InitMongoDB(); err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
	}
	defer database.CloseMongoDB()

	// Initialize router
	router := mux.NewRouter()

	// Add middleware
	router.Use(loggingMiddleware)
	router.Use(corsMiddleware)

	// Initialize handlers
	accountHandler := handlers.NewAccountHandler()
	userHandler := handlers.NewUserHandler()
	orderHandler := handlers.NewOrderHandler()
	staticHandler := handlers.NewStaticHandler()

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// Account routes
	api.HandleFunc("/accounts", accountHandler.CreateAccount).Methods("POST")
	api.HandleFunc("/accounts/{id}", accountHandler.GetAccount).Methods("GET")
	api.HandleFunc("/accounts/login", accountHandler.Login).Methods("POST")

	// User routes
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/users/account/{account_id}", userHandler.GetUserByAccountID).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")

	// Order routes
	api.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	api.HandleFunc("/orders/{id}", orderHandler.GetOrder).Methods("GET")
	api.HandleFunc("/orders/account/{account_id}", orderHandler.GetOrdersByAccountID).Methods("GET")
	api.HandleFunc("/orders", orderHandler.GetAllOrders).Methods("GET")

	// Restaurant and Food routes (static data)
	api.HandleFunc("/restaurants", staticHandler.GetRestaurants).Methods("GET")
	api.HandleFunc("/restaurants/{id}", staticHandler.GetRestaurant).Methods("GET")
	api.HandleFunc("/restaurants/{id}/foods", staticHandler.GetFoodsByRestaurant).Methods("GET")
	api.HandleFunc("/foods", staticHandler.GetFoods).Methods("GET")
	api.HandleFunc("/foods/{id}", staticHandler.GetFood).Methods("GET")

	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")

	// Serve static files from web directory
	fs := http.FileServer(http.Dir("./web"))
	router.PathPrefix("/").Handler(fs)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("📝 API documentation available at http://localhost:%s/api", port)
	log.Printf("🌐 Web interface available at http://localhost:%s", port)

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down server...")
}

// loggingMiddleware logs incoming requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// corsMiddleware adds CORS headers
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
