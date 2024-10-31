package repos

import (
	"location-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddressRepo interface {
	Create(address *models.Address) error
	Get(addressID uuid.UUID) (*models.Address, error)
	GetByID(addressID uuid.UUID) (*models.Address, error)
	GetAll() ([]models.Address, error)
	Update(address *models.Address) error
}

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepo(db *gorm.DB) AddressRepo {
	return &addressRepo{db}
}

func (r *addressRepo) Create(address *models.Address) error {
	return r.db.Create(address).Error
}

func (r *addressRepo) Get(addressID uuid.UUID) (*models.Address, error) {
	var address models.Address
	err := r.db.Preload("City").Where("id = ?", addressID).First(&address).Error
	return &address, err
}

func (r *addressRepo) GetByID(addressID uuid.UUID) (*models.Address, error) {
	var address models.Address
	err := r.db.Preload("City").Where("id = ?", addressID).First(&address).Error
	return &address, err
}

func (r *addressRepo) GetAll() ([]models.Address, error) {
	var addresses []models.Address
	err := r.db.Find(&addresses).Error
	return addresses, err
}

func (r *addressRepo) Update(address *models.Address) error {
	return r.db.Save(address).Error
}
