package dto

import (
	dtoCommon "backend/internal/dto/common"
)

type CreateFestivalRequest struct {
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	StartDate   string                         `json:"startDate"`
	EndDate     string                         `json:"endDate"`
	Capacity    int                            `json:"capacity"`
	Address     dtoCommon.CreateAddressRequest `json:"address"`
	// Status is defaulting to PRIVATE
	// StoreStatus is defaulting to CLOSED
}

type UpdateFestivalRequest struct {
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	StartDate   string                         `json:"startDate"`
	EndDate     string                         `json:"endDate"`
	Capacity    int                            `json:"capacity"`
	Address     dtoCommon.CreateAddressRequest `json:"address"`
}

type AddImageRequest struct {
	ImageUrl string `json:"imageUrl"`
}
