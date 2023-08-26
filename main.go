package main

import (
	"log"
	. "test_api/database"
	. "test_api/server"
)

func main() {
	db, err := NewDatabase("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	server := NewServer(db)

	server.Router.Run("localhost:3000")
}
