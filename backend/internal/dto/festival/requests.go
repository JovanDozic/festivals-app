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
}

type UpdateFestivalRequest struct {
	ID          uint                           `json:"id"`
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
	ItemID   uint    `json:"itemId"`
	Price    float64 `json:"price"`
	DateFrom *string `json:"dateFrom"`
	DateTo   *string `json:"dateTo"`
	IsFixed  bool    `json:"isFixed"`
}

type UpdateItemRequest struct {
	ID              uint                         `json:"id"`
	Name            string                       `json:"name"`
	Description     string                       `json:"description"`
	Type            string                       `json:"type"`
	AvailableNumber int                          `json:"availableNumber"`
	RemainingNumber int                          `json:"remainingNumber"`
	PriceListItems  []UpdatePriceListItemRequest `json:"priceListItems"`
}

type UpdatePriceListItemRequest struct {
	ID       uint    `json:"id"`
	Price    float64 `json:"price"`
	IsFixed  bool    `json:"isFixed"`
	DateFrom *string `json:"dateFrom"`
	DateTo   *string `json:"dateTo"`
}

type CityRequest struct {
	Name        string `json:"name"`
	PostalCode  string `json:"postalCode"`
	CountryISO3 string `json:"countryISO3"`
}

type CreateTransportPackageAddonRequest struct {
	ItemID        uint   `json:"itemId"`
	TransportType string `json:"transportType" validate:"oneof=BUS TRAIN PLANE"`

	DepartureCity CityRequest `json:"departureCity"`
	ArrivalCity   CityRequest `json:"arrivalCity"`

	DepartureTime string `json:"departureTime"`
	ArrivalTime   string `json:"arrivalTime"`

	ReturnDepartureTime string `json:"returnDepartureTime"`
	ReturnArrivalTime   string `json:"returnArrivalTime"`
}

type CreateCampPackageAddonRequest struct {
	ItemID        uint               `json:"itemId"`
	CampName      string             `json:"campName"`
	ImageURL      string             `json:"imageUrl"`
	EquipmentList []EquipmentRequest `json:"equipmentList"`
}

type EquipmentRequest struct {
	Name string `json:"name"`
}

type CreateTicketOrderRequest struct {
	TicketTypeId uint `json:"ticketTypeId"`
	TotalPrice   int  `json:"totalPrice"`
}
