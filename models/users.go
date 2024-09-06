package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	Token     string    `json:"token,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// ListUsers retrieves all users from the database
func ListUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`SELECT id, name, password, created_at FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userList []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.CreatedAt); err != nil {
			return nil, err
		}
		userList = append(userList, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userList, nil
}

// Insert new user into database with a hashed password, returns the created user with ID and creation timestamp
func CreateUser(db *sql.DB, name, password string) (*User, error) {
	var user User

	// Create a new User instance
	user.Name = name

	// Set the hashed password using the SetPassword method
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}

	// Insert the user into the database
	err := db.QueryRow(`
        INSERT INTO users (name, password)
        VALUES ($1, $2) RETURNING id, created_at
    `, user.Name, user.Password).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Retrieves a user from database by ID and return details
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

// Update an existing user's details in database
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

// Remove a user from database by ID
func DeleteUser(db *sql.DB, id int) error {
	_, err := db.Exec(`
        DELETE FROM users WHERE id = $1
    `, id)
	return err
}

// jwtKey is used to sign and verify JWT tokens
var jwtKey = []byte("your_secret_key")

// Hashes and set the user's password
// This is used when creating or updating an user's password
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Compares the given password with the stored hashed password.
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// Creates a jwt token for the user with a 72 hours expiration time.
func (u *User) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Verifies the user's credentials and generate a JWT token
func AuthenticateUser(db *sql.DB, name, password string) (*User, error) {
	user := &User{}
	err := db.QueryRow("SELECT id, name, password FROM users WHERE name=$1", name).Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		return nil, err
	}

	if !user.CheckPassword(password) {
		return nil, errors.New("invalid password")
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, err
	}

	user.Token = token
	return user, nil
}
