package services

import (
	modelsError "backend/internal/models"
	modelsCommon "backend/internal/models/common"
	repositoriesCommon "backend/internal/repositories/common"
	"log"
)

type LocationService interface {
	CreateAddress(address *modelsCommon.Address) error
	GetAddressByID(id uint) (*modelsCommon.Address, error)
}

type locationService struct {
	addressRepo repositoriesCommon.AddressRepo
	cityRepo    repositoriesCommon.CityRepo
	countryRepo repositoriesCommon.CountryRepo
}

func NewLocationService(addressRepo repositoriesCommon.AddressRepo, cityRepo repositoriesCommon.CityRepo, countryRepo repositoriesCommon.CountryRepo) LocationService {
	return &locationService{
		addressRepo: addressRepo,
		cityRepo:    cityRepo,
		countryRepo: countryRepo,
	}
}

func (s *locationService) CreateAddress(address *modelsCommon.Address) error {

	dbCountry, err := s.countryRepo.GetByISO3(address.City.Country.ISO3)
	if err != nil || dbCountry == nil {
		log.Println("country'", address.City.Country.ISO3, "'does not exist")
		return modelsError.ErrCountryNotFound
	}

	address.City.CountryID = dbCountry.ID

	err = s.addressRepo.Create(address)
	if err != nil {
		log.Println("error creating address:", err)
		return err
	}

	log.Println("address created successfully:", address.ID)
	return nil
}

func (s *locationService) GetAddressByID(id uint) (*modelsCommon.Address, error) {
	address, err := s.addressRepo.Get(id)
	if err != nil {
		log.Println("error getting address by id:", err)
		return nil, err
	}

	return address, nil
}
