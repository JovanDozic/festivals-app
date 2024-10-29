package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID    uuid.UUID `json:"user_id" gorm:"column:user_id;type:uuid;primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username" gorm:"column:username;unique;not null"`
	Password  string    `json:"password" gorm:"column:password;not null"`
	Email     string    `json:"email" gorm:"column:email;unique;not null"`
	Role      string    `json:"role" gorm:"column:role;not null"`
}

func (u *User) BeforeCreate(scope *gorm.DB) error {
	u.UserID = uuid.New()
	return nil
}

func (u *User) Validate() error {
	// todo: implement validation
	return nil
}
