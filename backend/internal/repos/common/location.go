package common

import (
	"backend/internal/models"
	modelsCommon "backend/internal/models/common"
	"log"

	"gorm.io/gorm"
)

type LocationRepo interface {
	// countries
	GetCountryByISO3(iso3 string) (*modelsCommon.Country, error)
	// cities
	CreateCity(city *modelsCommon.City) error
	GetCityByCountryPostalCode(countryId uint, postalCode string) (*modelsCommon.City, error)
	UpdateCity(city *modelsCommon.City) error
	// addresses
	CreateAddress(address *modelsCommon.Address) error
	GetAddressByID(id uint) (*modelsCommon.Address, error)
	UpdateAddress(address *modelsCommon.Address) error
	// other
	FillCityAndCountryModels(city *modelsCommon.City, country *modelsCommon.Country) error
}

type locationRepo struct {
	db *gorm.DB
}

func NewLocationRepo(db *gorm.DB) LocationRepo {
	return &locationRepo{db: db}
}

func (r *locationRepo) GetCountryByISO3(iso3 string) (*modelsCommon.Country, error) {
	var country modelsCommon.Country
	err := r.db.Where("iso3 = ?", iso3).First(&country).Error
	return &country, err
}

func (r *locationRepo) CreateCity(city *modelsCommon.City) error {
	return r.db.Create(city).Error
}

func (repo *locationRepo) GetCityByCountryPostalCode(countryID uint, postalCode string) (*modelsCommon.City, error) {
	var city modelsCommon.City
	err := repo.db.Where("country_id = ? AND postal_code = ?", countryID, postalCode).First(&city).Error
	if err != nil {
		return nil, err
	}
	return &city, nil
}

func (r *locationRepo) UpdateCity(city *modelsCommon.City) error {
	return r.db.Save(city).Error
}

func (r *locationRepo) CreateAddress(address *modelsCommon.Address) error {
	if err := r.db.Omit("Country").FirstOrCreate(&address.City, modelsCommon.City{
		Name:      address.City.Name,
		CountryID: address.City.CountryID,
	}).Error; err != nil {
		return err
	}

	return r.db.Omit("City.Country").Create(address).Error
}

func (r *locationRepo) GetAddressByID(id uint) (*modelsCommon.Address, error) {
	var address modelsCommon.Address
	if err := r.db.Preload("City.Country").First(&address, id).Error; err != nil {
		return nil, err
	}

	return &address, nil
}

func (r *locationRepo) UpdateAddress(address *modelsCommon.Address) error {
	return r.db.Save(address).Error
}

func (s *locationRepo) FillCityAndCountryModels(city *modelsCommon.City, country *modelsCommon.Country) error {

	// check if country exists
	dbCountry, err := s.GetCountryByISO3(country.ISO3)
	if err != nil || dbCountry == nil {
		log.Println("country'", country.ISO3, "'does not exist")
		return models.ErrCountryNotFound
	}

	country.ID = dbCountry.ID

	// check if city exists in that country
	dbCity, err := s.GetCityByCountryPostalCode(country.ID, city.PostalCode)
	if err != nil || dbCity == nil {

		// if it does not, create one
		city.CountryID = dbCountry.ID
		city.Country = *dbCountry

		err = s.CreateCity(city)
		if err != nil {
			log.Println("error creating city:", err)
			return err
		}
	} else {
		city.ID = dbCity.ID
		city.CountryID = dbCountry.ID
		city.Country = *dbCountry
	}

	// now we have the city and the country filled out
	return nil
}
