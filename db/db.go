package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB // Exported variable

// InitDB initializes the database connection
func InitDB() {
	connStr := "user=pawg password=pawg dbname=pawg sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
}

func SeedAdminUser() {
	var userCount int
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE name = 'admin'").Scan(&userCount)
	if err != nil {
		log.Fatalf("Failed to check if admin user exists: %v", err)
	}

	if userCount == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash admin password: %v", err)
		}

		_, err = DB.Exec("INSERT INTO users (name, password) VALUES ($1, $2)", "admin", string(hashedPassword))
		if err != nil {
			log.Fatalf("Failed to seed admin user: %v", err)
		}

		log.Println("Admin user seeded successfully.")
	} else {
		log.Println("Admin user already exists, skipping seeding.")
	}
}
