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

// Insert new user into database with a hashed password, returns the created user with ID and creation timestamp
func CreateUser(db *sql.DB, name, password string) (*User, error) {
	var user User
	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	err = db.QueryRow(`
        INSERT INTO users (name, password)
        VALUES ($1, $2) RETURNING id, created_at
    `, name, hashedPassword).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	user.Name = name
	user.Password = string(hashedPassword)
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

var jwtKey = []byte("your_secret_key")

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

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
