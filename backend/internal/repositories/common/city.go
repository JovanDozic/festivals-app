package repositories

import (
	modelsCommon "backend/internal/models/common"

	"gorm.io/gorm"
)

type CityRepo interface {
	Get(cityID uint) (*modelsCommon.City, error)
	GetAll() ([]modelsCommon.City, error)
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
