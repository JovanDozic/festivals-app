package dto

type HealthCheckResponse struct {
	ServiceName string `json:"service_name"`
	Status      string `json:"status"`
	Environment string `json:"environment"`
	API         string `json:"API"`
	Secure      bool   `json:"secure"`
}

type GetAddressResponse struct {
	Street         string  `json:"street"`
	Number         string  `json:"number"`
	ApartmentSuite *string `json:"apartmentSuite"`
	City           string  `json:"city"`
	PostalCode     string  `json:"postalCode"`
	Country        string  `json:"country"`
	CountryISO3    string  `json:"countryISO3"`
}

type GetImageResponse struct {
	Url string `json:"url"`
}
