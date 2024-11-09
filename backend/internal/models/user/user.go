package models

import (
	"time"

	"gorm.io/gorm"

	modelsCommon "backend/internal/models/common"
)

type User struct {
	gorm.Model
	Username string
	Email    string `gorm:"unique"`
	Password string
	Role     string
}

type UserProfile struct {
	ID          uint
	FirstName   string
	LastName    string
	DateOfBirth time.Time
	PhoneNumber string
	UserID      uint `gorm:"unique"`
	User        User
	AddressID   *uint
	Address     *modelsCommon.Address
	ImageID     *uint
	Image       *modelsCommon.Image
}

type Attendee struct {
	UserID uint
	User   User
}

type Employee struct {
	UserID uint
	User   User
}

type Organizer struct {
	UserID uint
	User   User
}
