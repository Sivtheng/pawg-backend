package main

import (
	"backend/db"
	"backend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CORS Middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions { // Handle preflight requests
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Initialize the database
	db.InitDB()

	// Seed the admin user
	db.SeedAdminUser()

	// Create a new router
	r := mux.NewRouter()

	// Set up routes
	routes.SetupRoutes(r)

	// Start the server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", enableCORS(r)); err != nil {
		log.Fatal(err)
	}
}
