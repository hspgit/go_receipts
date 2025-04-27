package main

import (
	"log"

	"github.com/hspgit/go_receipts/internal/database"
	"github.com/hspgit/go_receipts/internal/server"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := database.MigrateDB(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	server.StartServer(":8080")
}
