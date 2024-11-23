package models

import (
	"gorm.io/gorm"
)

type Country struct {
	gorm.Model
	Name      string `json:"name"`
	NiceName  string `json:"niceName"`
	ISO       string `json:"iso"`
	ISO3      string `json:"iso3"`
	NumCode   int    `json:"numCode"`
	PhoneCode int    `json:"phoneCode"`
}

type City struct {
	gorm.Model
	Name       string  `json:"name"`
	PostalCode string  `json:"postalCode"`
	CountryID  uint    `json:"countryId"`
	Country    Country `json:"country"`
}

type Address struct {
	gorm.Model
	Street         string  `json:"street"`
	Number         string  `json:"number"`
	ApartmentSuite *string `json:"apartmentSuite"`
	CityID         uint    `json:"cityId"`
	City           City    `json:"city"`
}
