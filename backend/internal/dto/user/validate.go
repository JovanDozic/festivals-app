package dto

import (
	"backend/internal/models"
	"backend/internal/utils"
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

func (r *RegisterUserRequest) Validate() error {
	if r.Username == "" {
		return models.ErrMissingFields
	}
	if r.Password == "" {
		return models.ErrMissingFields
	}
	if r.Email == "" {
		return models.ErrMissingFields
	}
	if !utils.IsEmailValid(r.Email) {
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
	if !utils.IsDateValid(r.DateOfBirth) {
		return models.ErrInvalidDateFormat
	}
	return nil
}

func (r *ChangePasswordRequest) Validate() error {
	if r.OldPassword == "" {
		return models.ErrMissingFields
	}
	if r.NewPassword == "" {
		return models.ErrMissingFields
	}
	return nil
}

func (u *UpdateUserProfileRequest) Validate() error {
	if u.FirstName == "" {
		return models.ErrMissingFields
	}
	if u.LastName == "" {
		return models.ErrMissingFields
	}
	if u.DateOfBirth == "" {
		return models.ErrMissingFields
	}
	if u.PhoneNumber == "" {
		return models.ErrMissingFields
	}
	if !utils.IsDateValid(u.DateOfBirth) {
		return models.ErrInvalidDateFormat
	}
	return nil
}

func (u *UpdateUserEmailRequest) Validate() error {
	if u.Email == "" {
		return models.ErrMissingFields
	}
	if !utils.IsEmailValid(u.Email) {
		return models.ErrInvalidEmailFormat
	}
	return nil
}

func (r *CreateStaffRequest) Validate() error {
	if r.Username == "" {
		return models.ErrMissingFields
	}
	if r.Password == "" {
		return models.ErrMissingFields
	}
	if r.Email == "" {
		return models.ErrMissingFields
	}
	if !utils.IsEmailValid(r.Email) {
		return models.ErrInvalidEmailFormat
	}
	if err := r.UserProfile.Validate(); err != nil {
		return err
	}
	return nil
}

func (r *UpdateStaffEmailRequest) Validate() error {
	if r.Username == "" {
		return models.ErrMissingFields
	}
	if r.Email == "" {
		return models.ErrMissingFields
	}
	if !utils.IsEmailValid(r.Email) {
		return models.ErrInvalidEmailFormat
	}
	return nil
}

func (r *UpdateStaffProfileRequest) Validate() error {
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
	if !utils.IsDateValid(r.DateOfBirth) {
		return models.ErrInvalidDateFormat
	}
	return nil
}

func (r *UpdateProfilePhotoRequest) Validate() error {
	if r.URL == "" {
		return models.ErrMissingFields
	}
	return nil
}
