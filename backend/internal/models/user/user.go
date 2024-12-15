package models

import (
	"time"

	"gorm.io/gorm"

	modelsCommon "backend/internal/models/common"
)

type UserRole string

const (
	RoleAttendee  UserRole = "ATTENDEE"
	RoleEmployee  UserRole = "EMPLOYEE"
	RoleOrganizer UserRole = "ORGANIZER"
	RoleAdmin     UserRole = "ADMINISTRATOR"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
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
	UserID uint `gorm:"primaryKey"`
	User   User
}

type Employee struct {
	UserID uint `gorm:"primaryKey"`
	User   User
}

type Organizer struct {
	UserID uint `gorm:"primaryKey"`
	User   User
}

type Administrator struct {
	UserID uint `gorm:"primaryKey"`
	User   User
}
