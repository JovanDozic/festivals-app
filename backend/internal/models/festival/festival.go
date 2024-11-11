package models

import (
	"time"

	modelsCommon "backend/internal/models/common"
	modelsUser "backend/internal/models/user"

	"gorm.io/gorm"
)

type Festival struct {
	gorm.Model
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	Capacity    int
	Status      string // Status: ACTIVE, PRIVATE, CANCELLED, COMPLETED
	StoreStatus string // StoreStatus: OPEN, CLOSED
	AddressID   uint
	Address     *modelsCommon.Address
}

type FestivalOrganizer struct {
	FestivalID  uint
	Festival    Festival
	OrganizerID uint
	Organizer   modelsUser.User
}

type FestivalEmployee struct {
	FestivalID uint
	Festival   Festival
	EmployeeID uint
	Employee   modelsUser.User
}

type FestivalImage struct {
	FestivalID uint
	Festival   Festival
	ImageID    uint
	Image      modelsCommon.Image
}
