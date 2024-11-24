package models

import (
	"time"

	modelsCommon "backend/internal/models/common"
	modelsUser "backend/internal/models/user"

	"gorm.io/gorm"
)

type Festival struct {
	gorm.Model
	Name        string                `json:"name"`
	Description string                `json:"description"`
	StartDate   time.Time             `json:"start_date"`
	EndDate     time.Time             `json:"end_date"`
	Capacity    int                   `json:"capacity"`
	Status      string                `json:"status"`      // Status: PRIVATE (default one), PUBLIC, CANCELLED, COMPLETED
	StoreStatus string                `json:"storeStatus"` // StoreStatus: OPEN, CLOSED
	AddressID   uint                  `json:"address_id"`
	Address     *modelsCommon.Address `json:"address"`
}

type FestivalOrganizer struct {
	FestivalID uint                 `json:"festivalId"`
	Festival   Festival             `json:"festival"`
	UserID     uint                 `json:"userId"`
	User       modelsUser.Organizer `json:"user"`
}

type FestivalEmployee struct {
	FestivalID uint
	Festival   Festival
	UserID     uint
	User       modelsUser.Employee
}

type FestivalImage struct {
	FestivalID uint               `json:"festivalId"`
	Festival   Festival           `json:"festival"`
	ImageID    uint               `json:"imageId"`
	Image      modelsCommon.Image `json:"image"`
}
