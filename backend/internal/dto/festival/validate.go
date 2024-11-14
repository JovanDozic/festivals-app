package dto

import "backend/internal/models"

func (f *CreateFestivalRequest) Validate() error {
	if f.Name == "" {
		return models.ErrMissingFields
	}
	if f.Description == "" {
		return models.ErrMissingFields
	}
	if f.StartDate == "" {
		return models.ErrMissingFields
	}
	if f.EndDate == "" {
		return models.ErrMissingFields
	}
	if f.Address.Validate() != nil {
		return models.ErrMissingFields
	}
	return nil
}
