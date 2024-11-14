package dto

type CreateAddressRequest struct {
	Street         string `json:"street"`
	Number         string `json:"number"`
	ApartmentSuite string `json:"apartmentSuite"`
	City           string `json:"city"`
	PostalCode     string `json:"postalCode"`
	CountryISO3    string `json:"countryISO3"`
}
