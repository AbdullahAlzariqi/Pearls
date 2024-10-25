// models/postgres.go
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents the Users table
type User struct {
	UserID       uuid.UUID `gorm:"type:uuid;primaryKey;" json:"user_id"`
	Username     string    `gorm:"unique;not null" json:"username"`
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Status       string    `gorm:"type:enum_status;default:'active'" json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	TeamID       uuid.UUID `gorm:"type:uuid" json:"team_id"`

	Team      Team       `gorm:"foreignKey:TeamID" json:"team"`
	UserRoles []UserRole `gorm:"foreignKey:UserID" json:"user_roles"`
}

// Team represents the Teams table
type Team struct {
	TeamID    uuid.UUID `gorm:"type:uuid;primaryKey;" json:"team_id"`
	TeamName  string    `gorm:"not null" json:"team_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Users []User `gorm:"foreignKey:TeamID" json:"users"`
}

// Role represents the Roles table
type Role struct {
	RoleID      uuid.UUID `gorm:"type:uuid;primaryKey;" json:"role_id"`
	RoleName    string    `gorm:"not null" json:"role_name"`
	Description string    `json:"description"`

	RolePermissions []RolePermission `gorm:"foreignKey:RoleID" json:"role_permissions"`
	UserRoles       []UserRole       `gorm:"foreignKey:RoleID" json:"user_roles"`
}

// Permission represents the Permissions table
type Permission struct {
	PermissionID   uuid.UUID `gorm:"type:uuid;primaryKey;" json:"permission_id"`
	PermissionName string    `gorm:"not null" json:"permission_name"`
	Description    string    `json:"description"`

	RolePermissions []RolePermission `gorm:"foreignKey:PermissionID" json:"role_permissions"`
}

// UserRole represents the UserRoles table
type UserRole struct {
	UserID uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`
	RoleID uuid.UUID `gorm:"type:uuid;primaryKey" json:"role_id"`

	User User `gorm:"foreignKey:UserID" json:"user"`
	Role Role `gorm:"foreignKey:RoleID" json:"role"`
}

// RolePermission represents the RolePermissions table
type RolePermission struct {
	RoleID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"role_id"`
	PermissionID uuid.UUID `gorm:"type:uuid;primaryKey" json:"permission_id"`

	Role       Role       `gorm:"foreignKey:RoleID" json:"role"`
	Permission Permission `gorm:"foreignKey:PermissionID" json:"permission"`
}

// models/postgres.go (continued)
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UserID == uuid.Nil {
		u.UserID = uuid.New()
	}
	return
}

func (t *Team) BeforeCreate(tx *gorm.DB) (err error) {
	if t.TeamID == uuid.Nil {
		t.TeamID = uuid.New()
	}
	return
}

// Similarly add BeforeCreate hooks for other models
