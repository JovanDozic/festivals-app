package dto

import (
	dtoCommon "backend/internal/dto/common"
)

type CreateFestivalRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Capacity    int    `json:"capacity"`
	// Status is defaulting to PRIVATE
	// StoreStatus is defaulting to CLOSED
	Address dtoCommon.CreateAddressRequest `json:"address"`
}
