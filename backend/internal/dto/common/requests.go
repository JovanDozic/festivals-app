package dto

type CreateAddressRequest struct {
	Street         string `json:"street"`
	Number         string `json:"number"`
	ApartmentSuite string `json:"apartmentSuite"`
	City           string `json:"city"`
	PostalCode     string `json:"postalCode"`
	CountryISO3    string `json:"countryISO3"`
}

type UpdateAddressRequest struct {
	ID             uint   `json:"id"`
	Street         string `json:"street"`
	Number         string `json:"number"`
	ApartmentSuite string `json:"apartmentSuite"`
	City           string `json:"city"`
	PostalCode     string `json:"postalCode"`
	CountryISO3    string `json:"countryISO3"`
}

type GetPresignedURLRequest struct {
	Filename string `json:"filename"`
	FileType string `json:"fileType"`
}
