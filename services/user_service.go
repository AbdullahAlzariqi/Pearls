// services/user_service.go
package services

import (
	"github.com/AbdullahAlzariqi/Pearls/db"
	"github.com/AbdullahAlzariqi/Pearls/models"
	"github.com/google/uuid"
)

// CreateUser creates a new user in PostgreSQL
func CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

// GetUserByID retrieves a user by ID
func GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	err := db.DB.Preload("Team").Preload("UserRoles.Role").First(&user, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates user details
func UpdateUser(user *models.User) error {
	return db.DB.Save(user).Error
}

// DeleteUser deletes a user by ID
func DeleteUser(userID uuid.UUID) error {
	return db.DB.Delete(&models.User{}, "user_id = ?", userID).Error
}
