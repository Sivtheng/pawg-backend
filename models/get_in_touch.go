package models

import (
	"database/sql"
	"time"
)

type GetInTouch struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

// ListGetInTouch retrieves all "Get In Touch" entries from the database
func ListGetInTouch(db *sql.DB) ([]GetInTouch, error) {
	rows, err := db.Query(`SELECT id, name, email, message, created_at FROM get_in_touch`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var getInTouchList []GetInTouch
	for rows.Next() {
		var getInTouch GetInTouch
		if err := rows.Scan(&getInTouch.ID, &getInTouch.Name, &getInTouch.Email, &getInTouch.Message, &getInTouch.CreatedAt); err != nil {
			return nil, err
		}
		getInTouchList = append(getInTouchList, getInTouch)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return getInTouchList, nil
}

// Insert new record into database and return with ID and creation timestamp
func CreateGetInTouch(db *sql.DB, name, email, message string) (*GetInTouch, error) {
	var getInTouch GetInTouch
	err := db.QueryRow(
		`INSERT INTO get_in_touch (name, email, message) VALUES ($1, $2, $3) RETURNING id, created_at`,
		name, email, message,
	).Scan(&getInTouch.ID, &getInTouch.CreatedAt)
	if err != nil {
		return nil, err
	}
	getInTouch.Name = name
	getInTouch.Email = email
	getInTouch.Message = message
	return &getInTouch, nil
}

// Retrieves record from database by ID and return details
func GetGetInTouchByID(db *sql.DB, id int) (*GetInTouch, error) {
	var getInTouch GetInTouch
	err := db.QueryRow(
		`SELECT id, name, email, message, created_at FROM get_in_touch WHERE id = $1`,
		id,
	).Scan(&getInTouch.ID, &getInTouch.Name, &getInTouch.Email, &getInTouch.Message, &getInTouch.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &getInTouch, nil
}

// Update existing record in database with new info and return updated details
func UpdateGetInTouch(db *sql.DB, id int, name, email, message string) (*GetInTouch, error) {
	var getInTouch GetInTouch
	err := db.QueryRow(
		`UPDATE get_in_touch SET name = $1, email = $2, message = $3 WHERE id = $4 RETURNING id, created_at`,
		name, email, message, id,
	).Scan(&getInTouch.ID, &getInTouch.CreatedAt)
	if err != nil {
		return nil, err
	}
	getInTouch.Name = name
	getInTouch.Email = email
	getInTouch.Message = message
	return &getInTouch, nil
}

// Remove record from database by ID
func DeleteGetInTouch(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM get_in_touch WHERE id = $1`, id)
	return err
}
