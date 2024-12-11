package models

import (
	modelsCommon "backend/internal/models/common"
	"time"

	"gorm.io/gorm"
)

const (
	ItemTicketType   = "TICKET_TYPE"
	ItemPackageAddon = "PACKAGE_ADDON"

	PackageAddonGeneral   = "GENERAL"
	PackageAddonCamp      = "CAMP"
	PackageAddonTransport = "TRANSPORT"
)

type PriceList struct {
	gorm.Model
	FestivalID uint
	Festival   Festival
}

type Item struct {
	gorm.Model
	FestivalID      uint
	Festival        Festival
	Name            string
	Type            string // Type: TICKET_TYPE, PACKAGE_ADDON
	Description     string
	AvailableNumber int
	RemainingNumber int
}

type PriceListItem struct {
	gorm.Model
	PriceListID uint
	PriceList   PriceList
	ItemID      uint
	Item        Item
	DateFrom    *time.Time
	DateTo      *time.Time
	IsFixed     bool
	Price       float64
}

type TicketType struct {
	ItemID uint `gorm:"primaryKey"`
	Item   Item
}

type PackageAddon struct {
	ItemID   uint `gorm:"primaryKey"`
	Item     Item
	Category string // Category: CAMP_ADDON, TRANSPORT_ADDON, CUSTOM_ADDON
}

type PackageAddonImage struct {
	ItemID  uint
	Item    PackageAddon
	ImageID uint
	Image   modelsCommon.Image
}

type CampAddon struct {
	ItemID   uint `gorm:"primaryKey"`
	Item     PackageAddon
	CampName string
}

type CampEquipment struct {
	gorm.Model
	ItemID uint
	Item   CampAddon
	Name   string
}

type TransportAddon struct {
	ItemID        uint `gorm:"primaryKey"`
	Item          PackageAddon
	TransportType string // TransportType: BUS, TRAIN, PLANE

	DepartureTime       time.Time
	ArrivalTime         time.Time
	ReturnDepartureTime time.Time
	ReturnArrivalTime   time.Time

	DepartureCityID uint
	DepartureCity   modelsCommon.City `gorm:"foreignKey:DepartureCityID"`
	ArrivalCityID   uint
	ArrivalCity     modelsCommon.City `gorm:"foreignKey:ArrivalCityID"`
}
