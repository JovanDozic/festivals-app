package repositories

import (
	modelsCommon "backend/internal/models/common"

	"gorm.io/gorm"
)

type AddressRepo interface {
	Create(address *modelsCommon.Address) error
}

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepo(db *gorm.DB) AddressRepo {
	return &addressRepo{db}
}

func (r *addressRepo) Create(address *modelsCommon.Address) error {
	if err := r.db.Omit("Country").FirstOrCreate(&address.City, modelsCommon.City{
		Name:      address.City.Name,
		CountryID: address.City.CountryID,
	}).Error; err != nil {
		return err
	}

	return r.db.Omit("City.Country").Create(address).Error
}
