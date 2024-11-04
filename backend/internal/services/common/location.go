package services

import (
	modelsError "backend/internal/models"
	modelsCommon "backend/internal/models/common"
	repositoriesCommon "backend/internal/repositories/common"
	"log"
)

type LocationService interface {
	// Sta nam sve treba:
	// Kreiranje nove
	CreateAddress(address *modelsCommon.Address) error
	// Update neke adrese
	// Brisanje adrese
	// Dohvatanje svih drzava
	// Dohvatanje gradova u drzavi
	// Dohvatanje svih gradova
	GetCities() ([]modelsCommon.City, error)
	// Dohvatanje adresa u gradu
	// Dohvatanje adrese po id-u

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

	// Check if country exists
	dbCountry, err := s.countryRepo.GetByISO3(address.City.Country.ISO3)
	if err != nil || dbCountry == nil {
		log.Println("country'", address.City.Country.ISO3, "'does not exist")
		return modelsError.ErrCountryNotFound
	}

	// Check if city exists
	// We need to try to find a city in the found country with the given name or postal code,
	// If it's not found, then we create it
	dbCity, err := s.cityRepo.GetByCountryAndPostalCode(dbCountry.CountryID, address.City.PostalCode)
	if err != nil || dbCity == nil {
		log.Println("city does not exist, creating new city")
		city := modelsCommon.City{
			Name:       address.City.Name,
			CountryID:  dbCountry.CountryID,
			PostalCode: address.City.PostalCode,
		}
		if err := s.cityRepo.Create(&city); err != nil {
			log.Println("error creating city:", err)
			return err
		}
		dbCity = &city
	}

	address.CityID = dbCity.CityID
	address.City = modelsCommon.City{CityID: dbCity.CityID, Country: modelsCommon.Country{CountryID: dbCountry.CountryID}}

	err = s.addressRepo.Create(address)
	if err != nil {
		log.Println("error creating address:", err)
		return err
	}

	log.Println("address created successfully:", address.AddressID)
	return nil
}

func (s *locationService) GetCities() ([]modelsCommon.City, error) {
	return s.cityRepo.GetAll()
}
