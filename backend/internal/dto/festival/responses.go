package dto

import (
	dtoCommon "backend/internal/dto/common"
	"time"
)

type FestivalsResponse struct {
	Festivals []FestivalResponse `json:"festivals"`
}

type FestivalResponse struct {
	ID          uint                          `json:"id"`
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	StartDate   time.Time                     `json:"startDate"`
	EndDate     time.Time                     `json:"endDate"`
	Capacity    int                           `json:"capacity"`
	Status      string                        `json:"status"`
	StoreStatus string                        `json:"storeStatus"`
	Address     *dtoCommon.GetAddressResponse `json:"address"`
	Images      []dtoCommon.GetImageResponse  `json:"images"`
}
