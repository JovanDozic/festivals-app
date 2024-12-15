package dto

import "time"

type HealthCheckResponse struct {
	ServiceName string `json:"service_name"`
	Status      string `json:"status"`
	Environment string `json:"environment"`
	API         string `json:"API"`
	Secure      bool   `json:"secure"`
}

type GetAddressResponse struct {
	AddressId      *uint   `json:"addressId"`
	Street         string  `json:"street"`
	Number         string  `json:"number"`
	ApartmentSuite *string `json:"apartmentSuite"`
	City           string  `json:"city"`
	PostalCode     string  `json:"postalCode"`
	Country        string  `json:"country"`
	CountryISO3    string  `json:"countryISO3"`
	CountryISO2    string  `json:"countryISO2"`
	NiceName       *string `json:"niceName"`
}

type GetImageResponse struct {
	ID  uint   `json:"id"`
	URL string `json:"url"`
}

type GetPresignedURLResponse struct {
	UploadURL string `json:"uploadURL"`
	ImageURL  string `json:"imageURL"`
}

type CountryResponse struct {
	ID       uint   `json:"id"`
	NiceName string `json:"niceName"`
	ISO      string `json:"iso"`
	ISO3     string `json:"iso3"`
}

type LogResponse struct {
	ID        uint      `json:"id"`
	Username  *string   `json:"username"`
	Message   string    `json:"message"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
}
