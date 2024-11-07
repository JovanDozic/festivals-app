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
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"` // yyyy-mm-dd
	PhoneNumber string `json:"phoneNumber"`
}

type CreateUserAddressRequest struct {
	Street         string `json:"street"`
	Number         string `json:"number"`
	ApartmentSuite string `json:"apartmentSuite"`
	City           string `json:"city"`
	PostalCode     string `json:"postalCode"`
	CountryISO3    string `json:"countryISO3"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
