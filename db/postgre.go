// db/postgres.go
package db

import (
	"fmt"
	"log"
	"os"

	"github.com/AbdullahAlzariqi/Pearls/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	log.Println("Connected to PostgreSQL")
}

func AutoMigrateModels() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Team{},
		&models.Role{},
		&models.Permission{},
		&models.UserRole{},
		&models.RolePermission{},
		// Add other models here
	)
	if err != nil {
		log.Fatal("Failed to auto-migrate models:", err)
	}

	// Optionally, initialize UUID generation if necessary
	// GORM doesn't handle UUID generation by default, you may need hooks or default values
}
