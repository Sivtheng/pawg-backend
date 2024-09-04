package routes

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// CreateAppointmentHandler handles creating a new appointment
func CreateAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	var appointment models.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := time.Parse("2006-01-02", appointment.AppointmentDate)
	if err != nil {
		http.Error(w, "Invalid appointment date format", http.StatusBadRequest)
		return
	}

	_, err = time.Parse("15:04:05", appointment.AppointmentTime)
	if err != nil {
		http.Error(w, "Invalid appointment time format", http.StatusBadRequest)
		return
	}

	newAppointment, err := models.CreateAppointment(db.DB, appointment.Name, appointment.Email, appointment.PhoneNumber, appointment.AppointmentDate, appointment.AppointmentTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAppointment)
}

// GetAppointmentHandler retrieves an appointment by its ID
func GetAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}

	appointment, err := models.GetAppointmentByID(db.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if appointment == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(appointment)
}

// UpdateAppointmentHandler updates an existing appointment by its ID
func UpdateAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}

	var appointment models.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = time.Parse("2006-01-02", appointment.AppointmentDate)
	if err != nil {
		http.Error(w, "Invalid appointment date format", http.StatusBadRequest)
		return
	}

	_, err = time.Parse("15:04:05", appointment.AppointmentTime)
	if err != nil {
		http.Error(w, "Invalid appointment time format", http.StatusBadRequest)
		return
	}

	updatedAppointment, err := models.UpdateAppointment(db.DB, id, appointment.Name, appointment.Email, appointment.PhoneNumber, appointment.AppointmentDate, appointment.AppointmentTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if updatedAppointment == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(updatedAppointment)
}

// DeleteAppointmentHandler deletes an appointment by its ID
func DeleteAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}

	err = models.DeleteAppointment(db.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// SetupRoutes initializes the routes for appointments
func SetupAppointmentRoutes(router *mux.Router) {
	router.HandleFunc("/appointments", CreateAppointmentHandler).Methods("POST")
	router.HandleFunc("/appointments/{id:[0-9]+}", GetAppointmentHandler).Methods("GET")
	router.HandleFunc("/appointments/{id:[0-9]+}", UpdateAppointmentHandler).Methods("PUT")
	router.HandleFunc("/appointments/{id:[0-9]+}", DeleteAppointmentHandler).Methods("DELETE")
}
