package repos

import (
	"location-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CityRepo interface {
	Create(city *models.City) error
	Get(name string) (*models.City, error)
	GetByID(cityID uuid.UUID) (*models.City, error)
	GetAll() ([]models.City, error)
	Update(city *models.City) error
}

type cityRepo struct {
	db *gorm.DB
}

func NewCityRepo(db *gorm.DB) CityRepo {
	return &cityRepo{db}
}

func (r *cityRepo) Create(city *models.City) error {
	return r.db.Create(city).Error
}

func (r *cityRepo) Get(name string) (*models.City, error) {
	var city models.City
	err := r.db.Preload("Role").Where("name = ?", name).First(&city).Error
	return &city, err
}

func (r *cityRepo) GetByID(cityID uuid.UUID) (*models.City, error) {
	var city models.City
	err := r.db.Preload("Role").Where("id = ?", cityID).First(&city).Error
	return &city, err
}

func (r *cityRepo) GetAll() ([]models.City, error) {
	var cities []models.City
	err := r.db.Find(&cities).Error
	return cities, err
}

func (r *cityRepo) Update(city *models.City) error {
	return r.db.Save(city).Error
}
