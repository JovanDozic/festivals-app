package models

import (
	modelsUser "backend/internal/models/user"

	"gorm.io/gorm"
)

type FestivalTicket struct {
	gorm.Model
	ItemID uint
	Item   TicketType
}

type FestivalPackage struct {
	gorm.Model
}

type FestivalPackageAddon struct {
	FestivalPackageID uint
	FestivalPackage   FestivalPackage
	ItemID            uint
	Item              PackageAddon
}

type Order struct {
	gorm.Model
	TotalAmount       float64
	UserID            uint
	User              modelsUser.Attendee
	FestivalTicketID  uint
	FestivalTicket    FestivalTicket
	FestivalPackageID uint
	FestivalPackage   FestivalPackage
}
