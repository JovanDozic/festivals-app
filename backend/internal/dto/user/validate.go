package dto

import (
	"backend/internal/models"
	"regexp"
	"time"
)

func (r *LoginRequest) Validate() error {
	if r.Username == "" {
		return models.ErrMissingFields
	}
	if r.Password == "" {
		return models.ErrMissingFields
	}
	return nil
}

func (r *RegisterAttendeeRequest) Validate() error {
	if r.Username == "" {
		return models.ErrMissingFields
	}
	if r.Password == "" {
		return models.ErrMissingFields
	}
	if r.Email == "" {
		return models.ErrMissingFields
	}
	if !isValidEmailFormat(r.Email) {
		return models.ErrInvalidEmailFormat
	}
	return nil
}

func (r *CreateUserProfileRequest) Validate() error {
	if r.FirstName == "" {
		return models.ErrMissingFields
	}
	if r.LastName == "" {
		return models.ErrMissingFields
	}
	if r.DateOfBirth == "" {
		return models.ErrMissingFields
	}
	if r.PhoneNumber == "" {
		return models.ErrMissingFields
	}
	if !isValidDateFormat(r.DateOfBirth) {
		return models.ErrInvalidDateFormat
	}
	return nil
}

func isValidDateFormat(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

func isValidEmailFormat(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(pattern, email)
	if err != nil {
		return false
	}
	return matched
}

func (r *CreateUserAddressRequest) Validate() error {
	if r.Street == "" {
		return models.ErrMissingFields
	}
	if r.Number == "" {
		return models.ErrMissingFields
	}
	if r.City == "" {
		return models.ErrMissingFields
	}
	if r.PostalCode == "" {
		return models.ErrMissingFields
	}
	if r.CountryISO3 == "" {
		return models.ErrMissingFields
	}
	return nil
}
