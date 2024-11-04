package repositories

import (
	modelsCommon "backend/internal/models/common"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CityRepo interface {
	Create(city *modelsCommon.City) error
	Get(name string) (*modelsCommon.City, error)
	GetByID(cityID uuid.UUID) (*modelsCommon.City, error)
	GetByCountryAndPostalCode(countryID uuid.UUID, postalCode string) (*modelsCommon.City, error)
	GetAll() ([]modelsCommon.City, error)
	Update(city *modelsCommon.City) error
}

type cityRepo struct {
	db *gorm.DB
}

func NewCityRepo(db *gorm.DB) CityRepo {
	return &cityRepo{db}
}

func (r *cityRepo) Create(city *modelsCommon.City) error {
	return r.db.Create(city).Error
}

func (r *cityRepo) Get(name string) (*modelsCommon.City, error) {
	var city modelsCommon.City
	err := r.db.Preload("Country").Where("name = ?", name).First(&city).Error
	return &city, err
}

func (r *cityRepo) GetByID(cityID uuid.UUID) (*modelsCommon.City, error) {
	var city modelsCommon.City
	err := r.db.Preload("Country").Where("city_id = ?", cityID).First(&city).Error
	return &city, err
}

func (r *cityRepo) GetByCountryAndPostalCode(countryID uuid.UUID, postalCode string) (*modelsCommon.City, error) {
	var city modelsCommon.City
	err := r.db.Preload("Country").Where("country_id = ? AND postal_code = ?", countryID, postalCode).First(&city).Error
	return &city, err
}

func (r *cityRepo) GetAll() ([]modelsCommon.City, error) {
	var cities []modelsCommon.City
	err := r.db.Find(&cities).Error
	return cities, err
}

func (r *cityRepo) Update(city *modelsCommon.City) error {
	return r.db.Save(city).Error
}
