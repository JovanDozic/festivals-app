package services

import (
	modelsError "backend/internal/models"
	modelsCommon "backend/internal/models/common"
	repositoriesCommon "backend/internal/repositories/common"
	"log"
	"strings"
)

type LocationService interface {
	CreateAddress(address *modelsCommon.Address) error
	GetAddressByID(id uint) (*modelsCommon.Address, error)
	UpdateAddress(addressId uint, updatedAddress *modelsCommon.Address) error
	GetCityAndCountry(city *modelsCommon.City, country *modelsCommon.Country) error
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

func (s *locationService) UpdateAddress(addressId uint, updatedAddress *modelsCommon.Address) error {
	address, err := s.addressRepo.Get(addressId)
	if err != nil {
		log.Println("error getting address from the database:", err)
		return err
	}

	// Update basic address fields
	address.Street = updatedAddress.Street
	address.Number = updatedAddress.Number
	address.ApartmentSuite = updatedAddress.ApartmentSuite

	// Ensure the city is correct
	city, err := s.cityRepo.GetByPostalCode(updatedAddress.City.PostalCode)
	if err != nil && !strings.Contains(err.Error(), "record not found") {
		log.Println("error getting city from the database:", err)
		return err
	}

	if city == nil || city.ID == 0 || city.Name != updatedAddress.City.Name || city.Country.ID != updatedAddress.City.Country.ID {
		// Validate the country
		country, err := s.countryRepo.GetByISO3(updatedAddress.City.Country.ISO3)
		if err != nil {
			log.Println("error getting country from the database:", err)
			return err
		}

		if country == nil {
			log.Println("country not found:", updatedAddress.City.Country.ISO3)
			return modelsError.ErrCountryNotFound
		}

		// Create or update the city
		if city == nil || city.ID == 0 {
			city = &modelsCommon.City{
				Name:       updatedAddress.City.Name,
				PostalCode: updatedAddress.City.PostalCode,
				CountryID:  country.ID,
			}

			err = s.cityRepo.Create(city)
			if err != nil {
				log.Println("error creating city:", err)
				return err
			}
		} else {
			city.Name = updatedAddress.City.Name
			city.CountryID = country.ID
			city.Country = *country
			err = s.cityRepo.Update(city)
			if err != nil {
				log.Println("error updating city:", err)
				return err
			}
		}
	}

	// Update address with new city
	address.CityID = city.ID
	address.City = *city

	// Persist updated address
	err = s.addressRepo.Update(address)
	if err != nil {
		log.Println("error updating address:", err)
		return err
	}

	log.Println("address updated successfully:", address.ID)
	return nil
}

func (s *locationService) GetCityAndCountry(city *modelsCommon.City, country *modelsCommon.Country) error {

	// check if country exists

	dbCountry, err := s.countryRepo.GetByISO3(country.ISO3)
	if err != nil || dbCountry == nil {
		log.Println("country'", country.ISO3, "'does not exist")
		return modelsError.ErrCountryNotFound
	}

	country.ID = dbCountry.ID

	// check if city exists in that country
	dbCity, err := s.cityRepo.GetByCountryPostalCode(country.ID, city.PostalCode)
	if err != nil || dbCity == nil {
		// if it does not, create one

		city.CountryID = dbCountry.ID
		city.Country = *dbCountry

		err = s.cityRepo.Create(city)
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
