package dto

type RegisterAttendeeRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserProfileRequest struct {
	Username    string `json:"username"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	PhoneNumber string `json:"phoneNumber"`
}

type CreateAddressRequest struct {
	Username       string `json:"username"`
	Street         string `json:"street"`
	Number         string `json:"number"`
	ApartmentSuite string `json:"apartmentSuite"`
	City           string `json:"city"`
	PostalCode     string `json:"postalCode"`
	Country        string `json:"country"`
}
