package dto

type CreateAddressRequest struct {
	Street         string `json:"street"`
	Number         string `json:"number"`
	ApartmentSuite string `json:"apartment_suite"`
	PostalCode     string `json:"postal_code"`
	City           string `json:"city"`
	Country        string `json:"country"`
}
