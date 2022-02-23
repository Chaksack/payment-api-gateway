package main

import (
	"log"
)

func main() {

	// Database Configuration
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Routes Configuration
	setupRouter()

	// Server Configuration

}
