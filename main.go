// main.go (continued)
package main

import (
	"context"
	"log"
	"time"

	"github.com/AbdullahAlzariqi/Pearls/db"
	"github.com/AbdullahAlzariqi/Pearls/models"
	"github.com/AbdullahAlzariqi/Pearls/services"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	// ... (previous connection code)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Connect to PostgreSQL
	db.ConnectPostgres()

	// Connect to MongoDB
	db.ConnectMongoDB()

	// Auto-migrate PostgreSQL models
	db.AutoMigrateModels()
	// Example: Create a new user
	newUser := &models.User{
		UserID:       uuid.New(),
		Username:     "john_doe",
		Email:        "john@example.com",
		PasswordHash: "hashed_password",
		Status:       "active",
		TeamID:       uuid.New(), // Ensure the team exists
	}

	err = services.CreateUser(newUser)
	if err != nil {
		log.Fatal("Error creating user:", err)
	}

	// Example: Create new content in MongoDB
	content := &models.Content{
		JobID:       "some-job-id",
		URL:         "https://example.com",
		HTMLContent: "<html>...</html>",
		TextContent: "Example text content.",
		Media: []models.Media{
			{
				MediaType: "image",
				MediaURL:  "https://example.com/image.png",
			},
		},
		Metadata: models.Metadata{
			Title:       "Example Title",
			Description: "Example Description",
			Keywords:    []string{"example", "content"},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	insertResult, err := services.CreateContent(ctx, content)
	if err != nil {
		log.Fatal("Error creating content:", err)
	}
	log.Println("Inserted content with ID:", insertResult.InsertedID)
}
