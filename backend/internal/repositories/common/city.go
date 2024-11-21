package repositories

import (
	modelsCommon "backend/internal/models/common"

	"gorm.io/gorm"
)

type CityRepo interface {
	Get(cityID uint) (*modelsCommon.City, error)
	GetByPostalCode(postalCode string) (*modelsCommon.City, error)
	GetAll() ([]modelsCommon.City, error)
	Create(city *modelsCommon.City) error
	GetByCountryPostalCode(countryId uint, postalCode string) (*modelsCommon.City, error)
	Update(city *modelsCommon.City) error
}

type cityRepo struct {
	db *gorm.DB
}

func NewCityRepo(db *gorm.DB) CityRepo {
	return &cityRepo{db}
}

func (r *cityRepo) Get(cityID uint) (*modelsCommon.City, error) {
	var city modelsCommon.City
	err := r.db.Preload("Country").Where("ID = ?", cityID).First(&city).Error // todo: preload might not work??
	return &city, err
}

func (r *cityRepo) GetAll() ([]modelsCommon.City, error) {
	var cities []modelsCommon.City
	err := r.db.Find(&cities).Error
	return cities, err
}

func (r *cityRepo) GetByPostalCode(postalCode string) (*modelsCommon.City, error) {
	var city modelsCommon.City
	err := r.db.Preload("Country").Where("postal_code = ?", postalCode).First(&city).Error
	return &city, err
}

func (r *cityRepo) Create(city *modelsCommon.City) error {
	return r.db.Create(city).Error
}

func (r *cityRepo) GetByCountryPostalCode(countryId uint, postalCode string) (*modelsCommon.City, error) {
	var city modelsCommon.City
	err := r.db.Preload("Country").Joins("JOIN countries ON cities.country_id = countries.id").Where("countries.id = ? AND cities.postal_code = ?", countryId, postalCode).First(&city).Error
	return &city, err
}

func (r *cityRepo) Update(city *modelsCommon.City) error {
	return r.db.Save(city).Error
}
