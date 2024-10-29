package repos

import (
	"location-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CountryRepo interface {
	Create(user *models.Country) error
	Get(username string) (*models.Country, error)
	GetByID(id uuid.UUID) (*models.Country, error)
	GetAll() ([]models.Country, error)
	Update(user *models.Country) error
}

type countryRepo struct {
	db *gorm.DB
}

func NewCountryRepo(db *gorm.DB) CountryRepo {
	return &countryRepo{db}
}

func (r *countryRepo) Create(country *models.Country) error {
	return r.db.Create(country).Error
}

func (r *countryRepo) Get(name string) (*models.Country, error) {
	var country models.Country
	err := r.db.Preload("Role").Where("username = ?", name).First(&country).Error
	return &country, err
}

func (r *countryRepo) GetByID(countryID uuid.UUID) (*models.Country, error) {
	var country models.Country
	err := r.db.Preload("Role").Where("id = ?", countryID).First(&country).Error
	return &country, err
}

func (r *countryRepo) GetAll() ([]models.Country, error) {
	var countries []models.Country
	err := r.db.Find(&countries).Error
	return countries, err
}

func (r *countryRepo) Update(country *models.Country) error {
	return r.db.Save(country).Error
}
