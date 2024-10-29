package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Country struct {
	CountryID uuid.UUID `json:"country_id" gorm:"column:country_id;primaryKey"`
	Name      string    `json:"name" gorm:"column:name;not null;unique"`
	ISOCode3  string    `json:"iso_code_3" gorm:"column:iso_code_3;not null;unique"`
}

func (country *Country) BeforeCreate(scope *gorm.DB) error {
	country.CountryID = uuid.New()
	return nil
}
