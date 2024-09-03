package main

import (
	"backend/db" // Use the correct import path
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	// Database connection URL
	databaseUrl := "postgres://username:password@localhost:5432/pawgdb"

	// Initialize database connection
	db.InitDB(ctx, databaseUrl)
	defer db.CloseDB()

	// Define the SQL schema
	schema := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        password VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS get_in_touch (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) NOT NULL,
        message TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS appointments (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) NOT NULL,
        phone_number VARCHAR(20),
        appointment_date DATE NOT NULL,
        appointment_time TIME NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS adoption_applications (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) NOT NULL,
        phone_number VARCHAR(20),
        address TEXT,
        interest_in_adopting VARCHAR(50) CHECK (interest_in_adopting IN ('Dog', 'Cat', 'Both')),
        type_of_animal VARCHAR(50) CHECK (type_of_animal IN ('Puppy', 'Adult Dog', 'Kitten', 'Adult Cat')),
        special_needs_animal VARCHAR(10) CHECK (special_needs_animal IN ('Yes', 'No', 'Maybe')),
        own_pet_before VARCHAR(10) CHECK (own_pet_before IN ('Yes', 'No', 'Maybe')),
        working_time TEXT,
        living_situation TEXT,
        other_animals TEXT,
        animal_access TEXT,
        travel TEXT,
        leave_cambodia BOOLEAN DEFAULT FALSE,
        feed TEXT,
        anything_else TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    `

	// Execute the schema creation
	_, err := db.Pool.Exec(ctx, schema)
	if err != nil {
		log.Fatalf("Failed to execute schema: %v", err)
	}

	log.Println("Database schema created successfully!")
}
