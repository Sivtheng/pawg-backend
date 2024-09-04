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
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	// Insert the admin user
	query := `INSERT INTO users (name, password) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err = DB.Exec(query, "admin", string(hashedPassword))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Admin user seeded successfully!")
}
