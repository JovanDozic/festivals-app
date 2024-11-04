package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID    uuid.UUID `json:"userId" gorm:"column:user_id;primaryKey;type:uuid"`
	Username  string    `json:"username" gorm:"column:username;unique;size:255;not null"`
	Email     string    `json:"email" gorm:"column:email;unique;size:255;not null"`
	Password  string    `json:"password" gorm:"column:password;size:255;not null"`
	Role      string    `json:"role" gorm:"column:role;size:50;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;not null"`
}

func (u *UserProfile) BeforeCreate(scope *gorm.DB) error {
	u.UserProfileID = uuid.New()
	return nil
}

type UserProfile struct {
	UserProfileID uuid.UUID  `json:"userProfileId" gorm:"column:user_profile_id;primaryKey;type:uuid"`
	FirstName     string     `json:"firstName" gorm:"column:first_name;size:255;not null"`
	LastName      string     `json:"lastName" gorm:"column:last_name;size:255;not null"`
	DateOfBirth   time.Time  `json:"dateOfBirth" gorm:"column:date_of_birth;not null"`
	PhoneNumber   *string    `json:"phoneNumber" gorm:"column:phone_number;unique;size:20"`
	UserID        uuid.UUID  `json:"userId" gorm:"column:user_id;unique;type:uuid;not null"`
	AddressID     *uuid.UUID `json:"addressId" gorm:"column:address_id;type:uuid"`
	ImageID       *uuid.UUID `json:"imageId" gorm:"column:image_id;type:uuid"`
}

func (u *User) BeforeCreate(scope *gorm.DB) error {
	u.UserID = uuid.New()
	return nil
}

type Attendee struct {
	UserID uuid.UUID `json:"userId" gorm:"column:user_id;primaryKey;type:uuid"`
}

type Employee struct {
	UserID uuid.UUID `json:"userId" gorm:"column:user_id;primaryKey;type:uuid"`
}

type Organizer struct {
	UserID uuid.UUID `json:"userId" gorm:"column:user_id;primaryKey;type:uuid"`
}
