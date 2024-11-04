package repositories

import (
	modelsCommon "backend/internal/models/common"

	"gorm.io/gorm"
)

type AddressRepo interface {
	Create(address *modelsCommon.Address) error
	Get(addressID uint) (*modelsCommon.Address, error)
	GetAll() ([]modelsCommon.Address, error)
	Update(address *modelsCommon.Address) error
}

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepo(db *gorm.DB) AddressRepo {
	return &addressRepo{db}
}

func (r *addressRepo) Create(address *modelsCommon.Address) error {
	if err := r.db.FirstOrCreate(&address.City, modelsCommon.City{Name: address.City.Name, CountryID: address.City.CountryID}).Error; err != nil {
		return err
	}

	return r.db.Create(address).Error
}

func (r *addressRepo) Get(addressID uint) (*modelsCommon.Address, error) {
	var address modelsCommon.Address
	err := r.db.Preload("City").Where("id = ?", addressID).First(&address).Error
	return &address, err
}

func (r *addressRepo) GetAll() ([]modelsCommon.Address, error) {
	var addresses []modelsCommon.Address
	err := r.db.Find(&addresses).Error
	return addresses, err
}

func (r *addressRepo) Update(address *modelsCommon.Address) error {
	return r.db.Save(address).Error
}
