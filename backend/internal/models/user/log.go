package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Type        string
	Description string
	Data        datatypes.JSON
	UserID      uint
	User        User
}
