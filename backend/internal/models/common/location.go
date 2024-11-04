package models

import (
	"gorm.io/gorm"
)

type Country struct {
	gorm.Model
	Name      string
	NiceName  string
	ISO       string
	ISO3      string
	NumCode   int
	PhoneCode int
}

type City struct {
	gorm.Model
	Name       string
	PostalCode string
	CountryID  uint
	Country    Country
}

type Address struct {
	gorm.Model
	Street         string
	Number         string
	ApartmentSuite *string
	CityID         uint
	City           City
}
