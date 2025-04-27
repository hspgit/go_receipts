package database

import (
	"log"

	"github.com/hspgit/go_receipts/internal/models"
	"gorm.io/gorm"
)

// MigrateDB handles database schema migrations
func MigrateDB(db *gorm.DB) error {
	log.Println("Running migrations...")

	if err := db.Migrator().DropTable(&models.User{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	users := []models.User{
		{Name: "Hrishikesh Pradhan", Email: "hspsat@gmail.com"},
		{Name: "Jane Smith", Email: "jane@example.com"},
		{Name: "Bob Johnson", Email: "bob@example.com"},
	}

	result := db.Create(&users)
	if result.Error != nil {
		return result.Error
	}

	log.Printf("Migrations complete. Created %d users.", len(users))
	return nil
}
