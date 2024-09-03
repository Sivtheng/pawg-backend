package main

import (
	"backend/db"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	databaseUrl := "postgres://username:password@localhost:5432/pawgdb"

	// Initialize database connection
	db.InitDB(ctx, databaseUrl)
	defer db.CloseDB()

	// Other setup code, like starting the server, goes here

	log.Println("Server started successfully!")
}
