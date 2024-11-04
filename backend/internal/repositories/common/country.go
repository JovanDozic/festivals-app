package repositories

import (
	modelsCommon "backend/internal/models/common"

	"gorm.io/gorm"
)

type CountryRepo interface {
	Create(country *modelsCommon.Country) error
	Get(name string) (*modelsCommon.Country, error)
	GetByID(countryID uint) (*modelsCommon.Country, error)
	GetByISO3(iso3 string) (*modelsCommon.Country, error)
	GetAll() ([]modelsCommon.Country, error)
	Update(country *modelsCommon.Country) error
}

type countryRepo struct {
	db *gorm.DB
}

func NewCountryRepo(db *gorm.DB) CountryRepo {
	return &countryRepo{db}
}

func (r *countryRepo) Create(country *modelsCommon.Country) error {
	return r.db.Create(country).Error
}

func (r *countryRepo) Get(name string) (*modelsCommon.Country, error) {
	var country modelsCommon.Country
	err := r.db.Where("name = ?", name).First(&country).Error
	return &country, err
}

func (r *countryRepo) GetByID(countryID uint) (*modelsCommon.Country, error) {
	var country modelsCommon.Country
	err := r.db.Where("country_id = ?", countryID).First(&country).Error
	return &country, err
}

func (r *countryRepo) GetByISO3(iso3 string) (*modelsCommon.Country, error) {
	var country modelsCommon.Country
	err := r.db.Where("iso3 = ?", iso3).First(&country).Error
	return &country, err
}

func (r *countryRepo) GetAll() ([]modelsCommon.Country, error) {
	var countries []modelsCommon.Country
	err := r.db.Find(&countries).Error
	return countries, err
}

func (r *countryRepo) Update(country *modelsCommon.Country) error {
	return r.db.Save(country).Error
}
