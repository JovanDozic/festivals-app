package repositories

import (
	modelsCommon "backend/internal/models/common"

	"gorm.io/gorm"
)

type AddressRepo interface {
	Create(address *modelsCommon.Address) error
	Get(id uint) (*modelsCommon.Address, error)
	Update(address *modelsCommon.Address) error
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

func (r *addressRepo) Get(id uint) (*modelsCommon.Address, error) {
	var address modelsCommon.Address
	if err := r.db.Preload("City.Country").First(&address, id).Error; err != nil {
		return nil, err
	}

	return &address, nil
}

func (r *addressRepo) Update(address *modelsCommon.Address) error {
	return r.db.Save(address).Error
}
