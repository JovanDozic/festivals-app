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
	Id          uint                           `json:"id"`
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	StartDate   string                         `json:"startDate"`
	EndDate     string                         `json:"endDate"`
	Capacity    int                            `json:"capacity"`
	Address     dtoCommon.UpdateAddressRequest `json:"address"`
}

type AddImageRequest struct {
	ImageUrl string `json:"imageUrl"`
}

type CreateItemRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	AvailableNumber int    `json:"availableNumber"`
	Type            string `json:"type" validate:"oneof=TICKET_TYPE PACKAGE_ADDON"`
}

type CreatePackageAddonRequest struct {
	ItemID   uint   `json:"itemId"`
	Category string `json:"category"`
}

type CreatePriceListItemRequest struct {
	ItemId   uint    `json:"itemId"`
	Price    float64 `json:"price"`
	DateFrom *string `json:"dateFrom"`
	DateTo   *string `json:"dateTo"`
	IsFixed  bool    `json:"isFixed"`
}

type UpdateItemRequest struct {
	Id              uint                         `json:"id"`
	Name            string                       `json:"name"`
	Description     string                       `json:"description"`
	Type            string                       `json:"type"`
	AvailableNumber int                          `json:"availableNumber"`
	RemainingNumber int                          `json:"remainingNumber"`
	PriceListItems  []UpdatePriceListItemRequest `json:"priceListItems"`
}

type UpdatePriceListItemRequest struct {
	Id       uint    `json:"id"`
	Price    float64 `json:"price"`
	IsFixed  bool    `json:"isFixed"`
	DateFrom *string `json:"dateFrom"`
	DateTo   *string `json:"dateTo"`
}
