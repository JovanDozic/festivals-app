package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	AddressID      uuid.UUID `json:"address_id" gorm:"column:address_id;primaryKey"`
	Street         string    `json:"street" gorm:"column:street"`
	Number         string    `json:"number" gorm:"column:number"`
	ApartmentSuite string    `json:"apartment_suite" gorm:"column:apartment_suite"`
	CityID         uuid.UUID `json:"-" gorm:"not null"`
	City           City      `json:"city" gorm:"foreignKey:CityID"`
}

func (address *Address) BeforeCreate(scope *gorm.DB) error {
	address.AddressID = uuid.New()
	return nil
}
