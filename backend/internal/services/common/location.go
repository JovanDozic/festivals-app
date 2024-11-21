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

	if address.Street != updatedAddress.Street || address.Number != updatedAddress.Number || *address.ApartmentSuite != *updatedAddress.ApartmentSuite {
		address.Street = updatedAddress.Street
		address.Number = updatedAddress.Number
		address.ApartmentSuite = updatedAddress.ApartmentSuite

		err = s.addressRepo.Update(address)
		if err != nil {
			log.Println("error updating address:", err)
			return err
		}
	}

	if address.City.PostalCode != updatedAddress.City.PostalCode {
		city, err := s.cityRepo.GetByPostalCode(updatedAddress.City.PostalCode)
		if err != nil && !strings.Contains(err.Error(), "record not found") {
			log.Println("error getting city from the database:", err)
			return err
		}

		// If the city does not exist, create one
		if city == nil || city.ID == 0 {
			city = &modelsCommon.City{
				PostalCode: updatedAddress.City.PostalCode,
				Name:       updatedAddress.City.Name,
			}

			err = s.cityRepo.Create(city)
			if err != nil {
				log.Println("error creating city:", err)
				return err
			}
		}

		// Update the address with the new city
		address.CityID = city.ID
		address.City = *city

		err = s.addressRepo.Update(address)
		if err != nil {
			log.Println("error updating address:", err)
			return err
		}
	}

	// Check if country changed
	if address.City.Country.ISO3 != updatedAddress.City.Country.ISO3 {
		// Get the country from the database
		country, err := s.countryRepo.GetByISO3(updatedAddress.City.Country.ISO3)
		if err != nil {
			log.Println("error getting country from the database:", err)
			return err
		}

		// If the country does not exist, return an error
		if country == nil {
			log.Println("country '", updatedAddress.City.Country.ISO3, "' does not exist")
			return modelsError.ErrCountryNotFound
		}

		// Check if the postal code exists in the country
		city, err := s.cityRepo.GetByCountryPostalCode(country.ISO, updatedAddress.City.PostalCode)
		if err != nil && !strings.Contains(err.Error(), "record not found") {
			log.Println("error getting city from the database:", err)
			return err
		}

		// If the city does not exist, create one
		if city == nil || city.ID == 0 {
			city = &modelsCommon.City{
				PostalCode: updatedAddress.City.PostalCode,
				Name:       updatedAddress.City.Name,
				CountryID:  country.ID,
				Country:    *country,
			}

			err = s.cityRepo.Create(city)
			if err != nil {
				log.Println("error creating city:", err)
				return err
			}
		}

		// Update the address with the new city
		address.CityID = city.ID
		address.City = *city

		err = s.addressRepo.Update(address)
		if err != nil {
			log.Println("error updating address:", err)
			return err
		}
	}

	log.Println("address updated successfully:", address.ID)
	return nil
}
