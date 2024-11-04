package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Country struct {
	CountryID uuid.UUID `json:"countryId" gorm:"column:country_id;primaryKey;type:uuid"`
	Name      string    `json:"name" gorm:"column:name;not null;unique"`
	NiceName  string    `json:"niceName" gorm:"column:nice_name;not null;unique"`
	ISO       string    `json:"iso" gorm:"column:iso;unique"`
	ISO3      string    `json:"iso3" gorm:"column:iso3;unique"`
	NumCode   int       `json:"numCode" gorm:"column:num_code"`
	PhoneCode int       `json:"phoneCode" gorm:"column:phone_code"`
}

func (c *Country) BeforeCreate(scope *gorm.DB) error {
	c.CountryID = uuid.New()
	return nil
}

type City struct {
	CityID     uuid.UUID `json:"cityId" gorm:"column:city_id;primaryKey;type:uuid"`
	Name       string    `json:"name" gorm:"column:name;size:100;not null"`
	PostalCode string    `json:"postalCode" gorm:"column:postal_code;size:20;not null"`
	CountryID  uuid.UUID `json:"countryId" gorm:"column:country_id;type:uuid;not null"`
	Country    Country   `json:"country" gorm:"foreignKey:CountryID;references:CountryID"`
}

func (c *City) BeforeCreate(scope *gorm.DB) error {
	c.CityID = uuid.New()
	return nil
}

type Address struct {
	AddressID      uuid.UUID `json:"addressId" gorm:"column:address_id;primaryKey;type:uuid"`
	Street         string    `json:"street" gorm:"column:street;size:255;not null"`
	Number         string    `json:"number" gorm:"column:number;size:20;not null"`
	ApartmentSuite *string   `json:"apartmentSuite" gorm:"column:apartment_suite;size:100"`
	CityID         uuid.UUID `json:"cityId" gorm:"column:city_id;type:uuid;not null"`
	City           City      `json:"city" gorm:"foreignKey:CityID;references:CityID"`
}

func (a *Address) BeforeCreate(scope *gorm.DB) error {
	a.AddressID = uuid.New()
	return nil
}
