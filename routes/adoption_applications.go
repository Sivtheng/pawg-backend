package routes

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateAdoptionApplicationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var application models.AdoptionApplication
	if err := json.NewDecoder(r.Body).Decode(&application); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newApplication, err := models.CreateAdoptionApplication(db.DB, application.Name, application.Email, application.PhoneNumber, application.Address, application.InterestInAdopting, application.TypeOfAnimal, application.SpecialNeedsAnimal, application.OwnPetBefore, application.WorkingTime, application.LivingSituation, application.OtherAnimals, application.AnimalAccess, application.Travel, application.LeaveCambodia, application.Feed, application.AnythingElse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newApplication)
}

func GetAdoptionApplicationHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	application, err := models.GetAdoptionApplicationByID(db.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(application)
}

func UpdateAdoptionApplicationHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var application models.AdoptionApplication
	if err := json.NewDecoder(r.Body).Decode(&application); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedApplication, err := models.UpdateAdoptionApplication(db.DB, id, application.Name, application.Email, application.PhoneNumber, application.Address, application.InterestInAdopting, application.TypeOfAnimal, application.SpecialNeedsAnimal, application.OwnPetBefore, application.WorkingTime, application.LivingSituation, application.OtherAnimals, application.AnimalAccess, application.Travel, application.LeaveCambodia, application.Feed, application.AnythingElse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedApplication)
}

func DeleteAdoptionApplicationHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := models.DeleteAdoptionApplication(db.DB, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func SetupAdoptionApplicationRoutes(r *mux.Router) {
	r.HandleFunc("/adoption_applications", CreateAdoptionApplicationHandler).Methods("POST")
	r.HandleFunc("/adoption_applications/{id}", GetAdoptionApplicationHandler).Methods("GET")
	r.HandleFunc("/adoption_applications/{id}", UpdateAdoptionApplicationHandler).Methods("PUT")
	r.HandleFunc("/adoption_applications/{id}", DeleteAdoptionApplicationHandler).Methods("DELETE")
}
