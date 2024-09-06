package routes

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ListAdoptionApplicationsHandler handles GET requests to list all adoption applications
func ListAdoptionApplicationsHandler(w http.ResponseWriter, r *http.Request) {
	applications, err := models.ListAdoptionApplications(db.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(applications)
}

// handles POST requests to create a new adoption application
func CreateAdoptionApplicationHandler(w http.ResponseWriter, r *http.Request) {
	// check if request method is post
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// decode the request body into an AdoptionApplication struct
	var application models.AdoptionApplication
	if err := json.NewDecoder(r.Body).Decode(&application); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call the model function to create the adoption application in the database
	newApplication, err := models.CreateAdoptionApplication(db.DB, application.Name, application.Email, application.PhoneNumber, application.Address, application.InterestInAdopting, application.TypeOfAnimal, application.SpecialNeedsAnimal, application.OwnPetBefore, application.WorkingTime, application.LivingSituation, application.OtherAnimals, application.AnimalAccess, application.Travel, application.LeaveCambodia, application.Feed, application.AnythingElse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set the response status to Created (201) and encode the new application as JSON
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newApplication)
}

// handles GET requests to retrieve an application by ID
func GetAdoptionApplicationHandler(w http.ResponseWriter, r *http.Request) {
	// extract id from URL parameters
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Call the model function to get application from database
	application, err := models.GetAdoptionApplicationByID(db.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// encode tehe application as JSON and send it in the response
	json.NewEncoder(w).Encode(application)
}

// handles put request to update an existing adoption application
func UpdateAdoptionApplicationHandler(w http.ResponseWriter, r *http.Request) {
	// extract id from URL parameters
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// decode teh request body into AdoptionApplication struct
	var application models.AdoptionApplication
	if err := json.NewDecoder(r.Body).Decode(&application); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call teh model function to update the application in database
	updatedApplication, err := models.UpdateAdoptionApplication(db.DB, id, application.Name, application.Email, application.PhoneNumber, application.Address, application.InterestInAdopting, application.TypeOfAnimal, application.SpecialNeedsAnimal, application.OwnPetBefore, application.WorkingTime, application.LivingSituation, application.OtherAnimals, application.AnimalAccess, application.Travel, application.LeaveCambodia, application.Feed, application.AnythingElse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// encode the updated application as JSON and send in response
	json.NewEncoder(w).Encode(updatedApplication)
}

// Handles delete request to delete application by ID
func DeleteAdoptionApplicationHandler(w http.ResponseWriter, r *http.Request) {
	// get id from URL paramters
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// call the model function to delete the adoption application from databse
	if err := models.DeleteAdoptionApplication(db.DB, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set the response status to No Content (204) indicating the request was successfully but there is no content to send
	w.WriteHeader(http.StatusNoContent)
}

// set up the routes for handling adoption applications
func SetupAdoptionApplicationRoutes(r *mux.Router) {
	r.HandleFunc("/adoption_applications", CreateAdoptionApplicationHandler).Methods("POST")
	r.HandleFunc("/adoption_applications", ListAdoptionApplicationsHandler).Methods("GET")
	r.HandleFunc("/adoption_applications/{id}", GetAdoptionApplicationHandler).Methods("GET")
	r.HandleFunc("/adoption_applications/{id}", UpdateAdoptionApplicationHandler).Methods("PUT")
	r.HandleFunc("/adoption_applications/{id}", DeleteAdoptionApplicationHandler).Methods("DELETE")
}
