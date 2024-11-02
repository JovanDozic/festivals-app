package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (u *User) BeforeCreate(scope *gorm.DB) error {
	u.UserID = uuid.New()
	return nil
}

func (u *User) Validate() error {
	// todo: implement validation
	return nil
}
