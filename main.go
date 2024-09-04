package main

import (
	"backend/db"
	"backend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Create a new router
	r := mux.NewRouter()

	// Set up routes
	routes.SetupRoutes(r)

	// Start the server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
