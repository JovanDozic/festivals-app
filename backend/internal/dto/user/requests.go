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

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type UpdateUserProfileRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth" validate:"datetime=2006-01-02"`
	PhoneNumber string `json:"phoneNumber"`
}

type UpdateUserEmailRequest struct {
	Email string `json:"email"`
}
