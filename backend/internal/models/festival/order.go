package models

import (
	modelsCommon "backend/internal/models/common"
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
	FestivalPackageID *uint
	FestivalPackage   *FestivalPackage
}

type Bracelet struct {
	gorm.Model
	PIN              string
	BarcodeNumber    string
	Balance          float64
	FestivalTicketID uint
	FestivalTicket   FestivalTicket
	AttendeeID       uint
	Attendee         modelsUser.Attendee `gorm:"foreignKey:AttendeeID"`
	// * this does not need to be nullable because employee is issuing the bracelet to X attendee who has Y ticket, so we have all ot the needed IDs
	EmployeeID uint
	Employee   modelsUser.Employee `gorm:"foreignKey:EmployeeID"`
}

type ActivationHelpRequest struct {
	gorm.Model
	UserEnteredPIN   string
	IssueDescription string
	ProofImageID     uint
	ProofImage       modelsCommon.Image `gorm:"foreignKey:ProofImageID"`
	Status           string             // Status: OPEN, ACCEPTED, REJECTED
	BraceletID       uint
	Bracelet         Bracelet
	AttendeeID       uint
	Attendee         modelsUser.Attendee `gorm:"foreignKey:AttendeeID"`
	EmployeeID       *uint
	Employee         *modelsUser.Employee `gorm:"foreignKey:EmployeeID"`
}
