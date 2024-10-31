package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Country struct {
	CountryID uuid.UUID `json:"countryId" gorm:"column:country_id;primaryKey"`
	Name      string    `json:"name" gorm:"column:name;not null;unique"`
	NiceName  string    `json:"niceName" gorm:"column:nice_name;not null;unique"`
	ISO       string    `json:"iso" gorm:"column:iso;unique"`
	ISO3      string    `json:"iso3" gorm:"column:iso3;unique"`
	NumCode   int       `json:"numCode" gorm:"column:num_code"`
	PhoneCode int       `json:"phoneCode" gorm:"column:phone_code"`
}

func (country *Country) BeforeCreate(scope *gorm.DB) error {
	country.CountryID = uuid.New()
	return nil
}
