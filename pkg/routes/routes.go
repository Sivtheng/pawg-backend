package routes

import (
	"backend/pkg"
	"net/http"
)

// SetupRoutes initializes the routes for the application
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Pawg Adoption Center!"))
	})

	mux.HandleFunc("/pets", pkg.GetPets) // Add route for getting pets

	// Wrap mux with middleware
	return mux
}
