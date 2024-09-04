package routes

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateGetInTouchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var getInTouch models.GetInTouch
	if err := json.NewDecoder(r.Body).Decode(&getInTouch); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newGetInTouch, err := models.CreateGetInTouch(db.DB, getInTouch.Name, getInTouch.Email, getInTouch.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newGetInTouch)
}

func GetGetInTouchHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	getInTouch, err := models.GetGetInTouchByID(db.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(getInTouch)
}

func UpdateGetInTouchHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var getInTouch models.GetInTouch
	if err := json.NewDecoder(r.Body).Decode(&getInTouch); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedGetInTouch, err := models.UpdateGetInTouch(db.DB, id, getInTouch.Name, getInTouch.Email, getInTouch.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedGetInTouch)
}

func DeleteGetInTouchHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := models.DeleteGetInTouch(db.DB, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func SetupGetInTouchRoutes(r *mux.Router) {
	r.HandleFunc("/get_in_touch", CreateGetInTouchHandler).Methods("POST")
	r.HandleFunc("/get_in_touch/{id:[0-9]+}", GetGetInTouchHandler).Methods("GET")
	r.HandleFunc("/get_in_touch/{id:[0-9]+}", UpdateGetInTouchHandler).Methods("PUT")
	r.HandleFunc("/get_in_touch/{id:[0-9]+}", DeleteGetInTouchHandler).Methods("DELETE")
}
