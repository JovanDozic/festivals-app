package common

import "github.com/google/uuid"

// Country struct
type Country struct {
	CountryID uuid.UUID `json:"countryId" gorm:"column:country_id;primaryKey;type:uuid"`
	Name      string    `json:"name" gorm:"column:name;size:100;not null"`
	ISOCode2  string    `json:"isoCode2" gorm:"column:iso_code_2;unique;size:2;not null"`
}

// City struct
type City struct {
	CityID     uuid.UUID `json:"cityId" gorm:"column:city_id;primaryKey;type:uuid"`
	Name       string    `json:"name" gorm:"column:name;size:100;not null"`
	PostalCode string    `json:"postalCode" gorm:"column:postal_code;size:20;not null"`
	CountryID  uuid.UUID `json:"countryId" gorm:"column:country_id;type:uuid;not null"`
}

// Address struct
type Address struct {
	AddressID      uuid.UUID `json:"addressId" gorm:"column:address_id;primaryKey;type:uuid"`
	Street         string    `json:"street" gorm:"column:street;size:255;not null"`
	Number         string    `json:"number" gorm:"column:number;size:20;not null"`
	ApartmentSuite *string   `json:"apartmentSuite" gorm:"column:apartment_suite;size:100"`
	CityID         uuid.UUID `json:"cityId" gorm:"column:city_id;type:uuid;not null"`
}
