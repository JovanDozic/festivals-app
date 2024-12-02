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

type TransportAddonDTO struct {
	PriceListItemID          uint       `json:"priceListItemId"`
	PriceListID              uint       `json:"priceListId"`
	ItemID                   uint       `json:"itemId"`
	ItemName                 string     `json:"itemName"`
	ItemDescription          string     `json:"itemDescription"`
	ItemType                 string     `json:"itemType"`
	ItemAvailableNumber      int        `json:"itemAvailableNumber"`
	ItemRemainingNumber      int        `json:"itemRemainingNumber"`
	PriceListItemDateFrom    *time.Time `json:"dateFrom"`
	PriceListItemDateTo      *time.Time `json:"dateTo"`
	PriceListItemIsFixed     bool       `json:"isFixed"`
	Price                    float64    `json:"price"`
	PackageAddonCategory     string     `json:"packageAddonCategory"`
	TransportType            string     `json:"transportType"`
	DepartureTime            time.Time  `json:"departureTime"`
	ArrivalTime              time.Time  `json:"arrivalTime"`
	ReturnDepartureTime      time.Time  `json:"returnDepartureTime"`
	ReturnArrivalTime        time.Time  `json:"returnArrivalTime"`
	DepartureCityID          uint       `json:"departureCityId"`
	DepartureCityName        string     `json:"departureCityName"`
	DeparturePostalCode      string     `json:"departurePostalCode"`
	DepartureCountryISO3     string     `json:"departureCountryISO3"`
	DepartureCountryISO      string     `json:"departureCountryISO"`
	DepartureCountryNiceName string     `json:"departureCountryNiceName"`
	ArrivalCityID            uint       `json:"arrivalCityId"`
	ArrivalCityName          string     `json:"arrivalCityName"`
	ArrivalPostalCode        string     `json:"arrivalPostalCode"`
	ArrivalCountryISO3       string     `json:"arrivalCountryISO3"`
	ArrivalCountryISO        string     `json:"arrivalCountryISO"`
	ArrivalCountryNiceName   string     `json:"arrivalCountryNiceName"`
}

type GeneralAddonDTO struct {
	PriceListItemID       uint       `json:"priceListItemId"`
	PriceListID           uint       `json:"priceListId"`
	ItemID                uint       `json:"itemId"`
	ItemName              string     `json:"itemName"`
	ItemDescription       string     `json:"itemDescription"`
	ItemType              string     `json:"itemType"`
	ItemAvailableNumber   int        `json:"itemAvailableNumber"`
	ItemRemainingNumber   int        `json:"itemRemainingNumber"`
	PriceListItemDateFrom *time.Time `json:"dateFrom"`
	PriceListItemDateTo   *time.Time `json:"dateTo"`
	PriceListItemIsFixed  bool       `json:"isFixed"`
	Price                 float64    `json:"price"`
	PackageAddonCategory  string     `json:"packageAddonCategory"`
}
