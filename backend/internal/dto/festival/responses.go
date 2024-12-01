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

type FestivalPropCountResponse struct {
	FestivalId uint `json:"festivalId"`
	Count      int  `json:"count"`
}

// this one returns only the current price
type ItemResponse struct {
	ItemId          uint       `json:"itemId"`
	PriceListItemId uint       `json:"priceListItemId"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	Type            string     `json:"type"`
	AvailableNumber int        `json:"availableNumber"`
	RemainingNumber int        `json:"remainingNumber"`
	Price           float64    `json:"price"`
	IsFixed         bool       `json:"isFixed"`
	DateFrom        *time.Time `json:"dateFrom"`
	DateTo          *time.Time `json:"dateTo"`
}

type GetItemsResponse struct {
	FestivalId uint           `json:"festivalId"`
	Items      []ItemResponse `json:"items"`
}

// this one returns all prices
type GetItemResponse struct {
	Id              uint                    `json:"id"`
	Name            string                  `json:"name"`
	Description     string                  `json:"description"`
	Type            string                  `json:"type"`
	AvailableNumber int                     `json:"availableNumber"`
	RemainingNumber int                     `json:"remainingNumber"`
	PriceListItems  []PriceListItemResponse `json:"priceListItems"`
}

type PriceListItemResponse struct {
	Id       uint       `json:"id"`
	Price    float64    `json:"price"`
	IsFixed  bool       `json:"isFixed"`
	DateFrom *time.Time `json:"dateFrom"`
	DateTo   *time.Time `json:"dateTo"`
}

type GetPackageAddonsResponse struct {
	FestivalId uint           `json:"festivalId"`
	Category   string         `json:"category"`
	Items      []ItemResponse `json:"items"`
}
