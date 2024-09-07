package routes

import (
	"backend/db"
	"backend/middleware"
	"backend/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Public endpoint for GetInTouch form submission
func CreatePublicGetInTouchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var getInTouch models.GetInTouch
	if err := json.NewDecoder(r.Body).Decode(&getInTouch); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newGetInTouch, err := models.CreateGetInTouch(db.DB, getInTouch.Name, getInTouch.Email, getInTouch.Message)
	if err != nil {
		log.Printf("Error inserting into database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newGetInTouch)
}

// CreatePublicAdoptionApplicationHandler handles POST requests to create a new adoption application
func CreatePublicAdoptionApplicationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var application models.AdoptionApplication
	if err := json.NewDecoder(r.Body).Decode(&application); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newApplication, err := models.CreateAdoptionApplication(
		db.DB,
		application.Name,
		application.Email,
		application.PhoneNumber,
		application.Address,
		application.InterestInAdopting,
		application.TypeOfAnimal,
		application.SpecialNeedsAnimal,
		application.OwnPetBefore,
		application.WorkingTime,
		application.LivingSituation,
		application.OtherAnimals,
		application.AnimalAccess,
		application.Travel,
		application.LeaveCambodia,
		application.Feed,
		application.AnythingElse,
	)
	if err != nil {
		log.Printf("Error inserting into database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newApplication)
}

func SetupRoutes(r *mux.Router) {
	// Public routes
	SetupLoginRoutes(r)
	r.HandleFunc("/public/get_in_touch", CreatePublicGetInTouchHandler).Methods("POST")
	r.HandleFunc("/public/submit-adoption-form", CreatePublicAdoptionApplicationHandler).Methods("POST")

	// Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTAuthMiddleware)

	SetupUserRoutes(api)
	SetupGetInTouchRoutes(api)
	SetupAppointmentRoutes(api)
	SetupAdoptionApplicationRoutes(api)
}
