package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	RoleStudent    UserRole = "student"
	RoleInstructor UserRole = "instructor"
	RoleAdmin      UserRole = "admin"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}
