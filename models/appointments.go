package models

import (
	"database/sql"
	"time"
)

// Appointment represents the appointment model
type Appointment struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	PhoneNumber     string    `json:"phone_number"`
	AppointmentDate string    `json:"appointment_date"` // Use string type
	AppointmentTime string    `json:"appointment_time"` // Use string type
	CreatedAt       time.Time `json:"created_at"`
}

// CreateAppointment inserts a new appointment into the database
func CreateAppointment(db *sql.DB, name, email, phoneNumber, appointmentDate, appointmentTime string) (*Appointment, error) {
	var appointment Appointment
	query := `INSERT INTO appointments (name, email, phone_number, appointment_date, appointment_time) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`
	err := db.QueryRow(query, name, email, phoneNumber, appointmentDate, appointmentTime).Scan(&appointment.ID, &appointment.CreatedAt)
	if err != nil {
		return nil, err
	}
	appointment.Name = name
	appointment.Email = email
	appointment.PhoneNumber = phoneNumber
	return &appointment, nil
}

// GetAppointmentByID retrieves an appointment by its ID
func GetAppointmentByID(db *sql.DB, id int) (*Appointment, error) {
	var appointment Appointment
	query := `SELECT id, name, email, phone_number, appointment_date, appointment_time, created_at 
              FROM appointments WHERE id = $1`
	row := db.QueryRow(query, id)
	err := row.Scan(&appointment.ID, &appointment.Name, &appointment.Email, &appointment.PhoneNumber,
		&appointment.AppointmentDate, &appointment.AppointmentTime, &appointment.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Appointment not found
		}
		return nil, err
	}
	return &appointment, nil
}

// UpdateAppointment updates an existing appointment
func UpdateAppointment(db *sql.DB, id int, name, email, phoneNumber, appointmentDate, appointmentTime string) (*Appointment, error) {
	var appointment Appointment
	query := `UPDATE appointments SET name = $1, email = $2, phone_number = $3, appointment_date = $4, appointment_time = $5
              WHERE id = $6 RETURNING id, created_at`
	err := db.QueryRow(query, name, email, phoneNumber, appointmentDate, appointmentTime, id).Scan(&appointment.ID, &appointment.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Appointment not found
		}
		return nil, err
	}
	appointment.Name = name
	appointment.Email = email
	appointment.PhoneNumber = phoneNumber
	appointment.AppointmentDate = appointmentDate
	appointment.AppointmentTime = appointmentTime
	return &appointment, nil
}

// DeleteAppointment removes an appointment by ID
func DeleteAppointment(db *sql.DB, id int) error {
	query := `DELETE FROM appointments WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}
