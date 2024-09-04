package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateUser(db *sql.DB, name, password string) (*User, error) {
	var user User
	err := db.QueryRow(`
        INSERT INTO users (name, password)
        VALUES ($1, $2) RETURNING id, created_at
    `, name, password).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	user.Name = name
	user.Password = password
	return &user, nil
}

func GetUserByID(db *sql.DB, id int) (*User, error) {
	var user User
	err := db.QueryRow(`
        SELECT id, name, password, created_at
        FROM users WHERE id = $1
    `, id).Scan(&user.ID, &user.Name, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(db *sql.DB, id int, name, password string) (*User, error) {
	var user User
	err := db.QueryRow(`
        UPDATE users
        SET name = $2, password = $3
        WHERE id = $1
        RETURNING id, created_at
    `, id, name, password).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	user.Name = name
	user.Password = password
	return &user, nil
}

func DeleteUser(db *sql.DB, id int) error {
	_, err := db.Exec(`
        DELETE FROM users WHERE id = $1
    `, id)
	return err
}
