package dto

import "backend/internal/models"

func (r *CreateAddressRequest) Validate() error {
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
