package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type City struct {
	CityID     uuid.UUID `json:"city_id" gorm:"column:city_id;primaryKey"`
	Name       string    `json:"name" gorm:"column:name;not null"`
	PostalCode string    `json:"postal_code" gorm:"column:postal_code;not null"`
	CountryID  uuid.UUID `json:"-" gorm:"not null"`
	Country    Country   `json:"country" gorm:"foreignKey:CountryID"`
}

func (city *City) BeforeCreate(scope *gorm.DB) error {
	city.CityID = uuid.New()
	return nil
}
