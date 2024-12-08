package dto

import (
	dtoCommon "backend/internal/dto/common"
	dtoUser "backend/internal/dto/user"
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

type EquipmentResponse struct {
	Name string `json:"name"`
}

type CampAddonDTO struct {
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
	CampName              string     `json:"campName"`
	ImageURL              string     `json:"imageUrl"`
	EquipmentNames        string     `json:"equipmentNames"`
}

type OrderDTO struct {
	OrderID          uint                            `json:"orderId"`
	OrderType        string                          `json:"orderType"` // this is like ticket or package
	Timestamp        time.Time                       `json:"timestamp"`
	TotalPrice       float64                         `json:"totalPrice"`
	Ticket           ItemResponse                    `json:"ticket"`
	TransportAddon   *TransportAddonDTO              `json:"transportAddon"`
	CampAddon        *CampAddonDTO                   `json:"campAddon"`
	GeneralAddons    []GeneralAddonDTO               `json:"generalAddons"`
	Festival         FestivalResponse                `json:"festival"`
	Username         string                          `json:"username"`
	Attendee         *dtoUser.GetUserProfileResponse `json:"attendee"`
	BraceletStatus   *string                         `json:"braceletStatus"`
	FestivalTicketId *uint                           `json:"festivalTicketId"`
	Bracelet         *BraceletDTO                    `json:"bracelet"`
}

type OrderPreviewDTO struct {
	OrderID          uint                            `json:"orderId"`
	OrderType        string                          `json:"orderType"`
	Timestamp        time.Time                       `json:"timestamp"`
	TotalPrice       float64                         `json:"totalPrice"`
	Festival         FestivalResponse                `json:"festival"`
	Username         string                          `json:"username"`
	Attendee         *dtoUser.GetUserProfileResponse `json:"attendee"`
	BraceletStatus   *string                         `json:"braceletStatus"`
	BraceletID       *uint                           `json:"braceletId"`
	FestivalTicketId *uint                           `json:"festivalTicketId"`
}

type BraceletDTO struct {
	BraceletID    uint                            `json:"braceletId"`
	BarcodeNumber string                          `json:"barcodeNumber"`
	Balance       float64                         `json:"balance"`
	Status        string                          `json:"status"`
	Employee      *dtoUser.GetUserProfileResponse `json:"employee"`
	PIN           *string                         `json:"pin"`
}

type ActivationHelpRequestDTO struct {
	ActivationHelpRequestID uint                           `json:"activationHelpRequestId"`
	UserEnteredPIN          string                         `json:"userEnteredPIN"`
	UserEnteredBarcode      string                         `json:"userEnteredBarcode"`
	IssueDescription        string                         `json:"issueDescription"`
	ImageURL                string                         `json:"imageUrl"`
	Status                  string                         `json:"status"`
	Bracelet                BraceletDTO                    `json:"bracelet"`
	Attendee                dtoUser.GetUserProfileResponse `json:"attendee"`
	Employee                dtoUser.GetUserProfileResponse `json:"employee"`
}
