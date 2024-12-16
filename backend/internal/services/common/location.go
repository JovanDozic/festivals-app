package common

import (
	modelsError "backend/internal/models"
	modelsCommon "backend/internal/models/common"
	"backend/internal/repos/common"
	"log"
	"strings"
)

type LocationService interface {
	UpdateAddress(addressId uint, updatedAddress *modelsCommon.Address) error
}

type locationService struct {
	locationRepo common.LocationRepo
}

func NewLocationService(lr common.LocationRepo) LocationService {
	return &locationService{
		locationRepo: lr,
	}
}

func (s *locationService) UpdateAddress(addressId uint, updatedAddress *modelsCommon.Address) error {

	// Fetch the existing address
	address, err := s.locationRepo.GetAddressByID(addressId)
	if err != nil {
		log.Println("error getting address from the database:", err)
		return err
	}

	// Update basic address fields
	address.Street = updatedAddress.Street
	address.Number = updatedAddress.Number
	address.ApartmentSuite = updatedAddress.ApartmentSuite

	// Validate the country
	country, err := s.locationRepo.GetCountryByISO3(updatedAddress.City.Country.ISO3)
	if err != nil || country == nil {
		log.Println("country not found:", updatedAddress.City.Country.ISO3)
		return modelsError.ErrCountryNotFound
	}

	// Check if the city exists in the specified country
	city, err := s.locationRepo.GetCityByCountryPostalCode(country.ID, updatedAddress.City.PostalCode)
	if err != nil && !strings.Contains(err.Error(), "record not found") {
		log.Println("error getting city from the database:", err)
		return err
	}

	// Determine if we need to create a new city
	createNewCity := false
	if city == nil || city.ID == 0 {
		createNewCity = true
	} else if city.Name != updatedAddress.City.Name || city.CountryID != country.ID {
		// Even if the postal code matches, if the name or country doesn't, we need a new city
		createNewCity = true
	}

	if createNewCity {
		// Create a new city
		city = &modelsCommon.City{
			Name:       updatedAddress.City.Name,
			PostalCode: updatedAddress.City.PostalCode,
			CountryID:  country.ID,
		}
		err = s.locationRepo.CreateCity(city)
		if err != nil {
			log.Println("error creating city:", err)
			return err
		}
	}

	// Update address with the new city
	address.CityID = city.ID
	address.City = *city

	// Persist updated address
	err = s.locationRepo.UpdateAddress(address)
	if err != nil {
		log.Println("error updating address:", err)
		return err
	}

	log.Println("address updated successfully:", address.ID)
	return nil
}
