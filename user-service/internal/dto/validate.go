package dto

import "user-service/internal/models"

func (r *LoginRequest) Validate() error {
	if r.Username == "" {
		return models.ErrEmptyUsername
	}
	if r.Password == "" {
		return models.ErrEmptyPassword
	}
	return nil
}
