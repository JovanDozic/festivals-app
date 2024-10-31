package services

import (
	"errors"
	"location-service/internal/models"
	"location-service/internal/repos"
	"log"

	pb "location-service/internal/proto/location"

	"github.com/google/uuid"
)

type LocationService interface {
	// Sta nam sve treba:
	// Kreiranje nove
	CreateAddress(req *pb.SaveAddressRequest) (uuid.UUID, error)
	// Updateovanje neke adrese
	// Brisanje adrese
	// Dohvatanje svih drzava
	// Dohvatanje gradova u drzavi
	// Dohvatanje svih gradova
	GetCities() ([]models.City, error)
	// Dohvatanje adresa u gradu
	// Dohvatanje adrese po id-u

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

func (s *locationService) CreateAddress(req *pb.SaveAddressRequest) (uuid.UUID, error) {

	var address models.Address

	// Check if country exists
	countryFromDB, err := s.countryRepo.GetByISO3(req.CountryISO3)
	if err != nil {
		log.Println("country does not exist")
		return uuid.Nil, errors.New("country does not exist")
	}

	// Check if city exists
	// We need to try to find a city in the found country with the given name or postal code,
	// If it's not found, then we create it
	cityFromDB, err := s.cityRepo.GetByCountryAndPostalCode(countryFromDB.CountryID, req.PostalCode)
	if err != nil {
		log.Println("city does not exist, creating new city")
		city := models.City{
			Name:       req.City,
			CountryID:  countryFromDB.CountryID,
			PostalCode: req.PostalCode,
		}
		if err := s.cityRepo.Create(&city); err != nil {
			log.Println("error creating city:", err)
			return uuid.Nil, err
		}
		cityFromDB = &city
	}

	address = models.Address{
		Street:         req.Street,
		Number:         req.Number,
		ApartmentSuite: req.ApartmentSuite,
		CityID:         cityFromDB.CityID,
	}

	err = s.addressRepo.Create(&address)
	if err != nil {
		log.Println("error creating address:", err)
		return uuid.Nil, err
	}

	log.Println("address created successfully:", address.AddressID)
	return address.AddressID, err
}

func (s *locationService) GetCities() ([]models.City, error) {
	return s.cityRepo.GetAll()
}
