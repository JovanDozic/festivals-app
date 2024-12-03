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

func (repo *cityRepo) GetByCountryPostalCode(countryID uint, postalCode string) (*modelsCommon.City, error) {
	var city modelsCommon.City
	err := repo.db.Where("country_id = ? AND postal_code = ?", countryID, postalCode).First(&city).Error
	if err != nil {
		return nil, err
	}
	return &city, nil
}

func (r *cityRepo) Update(city *modelsCommon.City) error {
	return r.db.Save(city).Error
}
