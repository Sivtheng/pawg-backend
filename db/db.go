package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB // Exported variable for the database connection

// InitDB initializes the database connection
func InitDB() {
	// Define the connection string with the database credentials and options
	connStr := "user=pawg password=pawg dbname=pawg sslmode=disable"
	var err error

	// Open a connection to the database
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to ensure it is reachable
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
}

// SeedAdminUser seeds the database with an admin user if it does not already exist
func SeedAdminUser() {
	var userCount int

	// Query the database to check if an admin user already exists
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE name = 'admin'").Scan(&userCount)
	if err != nil {
		log.Fatalf("Failed to check if admin user exists: %v", err)
	}

	// If no admin user is found, proceed to create one
	if userCount == 0 {
		// Hash the password using bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash admin password: %v", err)
		}

		// Insert the new admin user into the database
		_, err = DB.Exec("INSERT INTO users (name, password) VALUES ($1, $2)", "admin", string(hashedPassword))
		if err != nil {
			log.Fatalf("Failed to seed admin user: %v", err)
		}

		log.Println("Admin user seeded successfully.")
	} else {
		log.Println("Admin user already exists, skipping seeding.")
	}
}
