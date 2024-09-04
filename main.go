package main

import (
	"net/http"
)

func main() {
	initDB() // Initialize the database connection

	http.ListenAndServe(":8080", nil)
}
