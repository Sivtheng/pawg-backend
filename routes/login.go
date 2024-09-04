package routes

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := models.AuthenticateUser(db.DB, creds.Name, creds.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func SetupLoginRoutes(r *mux.Router) {
	r.HandleFunc("/login", LoginHandler).Methods("POST")
}
