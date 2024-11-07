package dto

type HealthCheckResponse struct {
	ServiceName string `json:"service_name"`
	Status      string `json:"status"`
	Environment string `json:"environment"`
	API         string `json:"API"`
	Secure      bool   `json:"secure"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type GetUserProfileResponse struct {
	Username    string              `json:"username"`
	Email       string              `json:"email"`
	Role        string              `json:"role"`
	FirstName   string              `json:"firstName"`
	LastName    string              `json:"lastName"`
	DateOfBirth string              `json:"dateOfBirth"`
	PhoneNumber string              `json:"phoneNumber"`
	Address     *GetAddressResponse `json:"address"`
}

type GetAddressResponse struct {
	Street         string `json:"street"`
	Number         string `json:"number"`
	ApartmentSuite string `json:"apartmentSuite"`
	City           string `json:"city"`
	PostalCode     string `json:"postalCode"`
	Country        string `json:"country"`
}
