package services

import (
	"location-service/internal/models"
	"location-service/internal/repos"
)

type LocationService interface {
	CreateAddress(address *models.Address, city *models.City, country *models.Country) error
	GetCities() ([]models.City, error)
}

type locationService struct {
	addressRepo repos.AddressRepo
	cityRepo    repos.CityRepo
	countryRepo repos.CountryRepo
}

func NewLocationService(addressRepo repos.AddressRepo, cityRepo repos.CityRepo, countryRepo repos.CountryRepo) LocationService {
	return &locationService{
		addressRepo: addressRepo,
		cityRepo:    cityRepo,
		countryRepo: countryRepo,
	}
}

func (s *locationService) CreateAddress(address *models.Address, city *models.City, country *models.Country) error {
	return s.addressRepo.Create(address)
}

func (s *locationService) GetCities() ([]models.City, error) {
	return s.cityRepo.GetAll()
}
