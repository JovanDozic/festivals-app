package dto

import (
	"backend/internal/models"
	modelsFestival "backend/internal/models/festival"
	"errors"
)

func (f *CreateFestivalRequest) Validate() error {
	if f.Name == "" {
		return models.ErrMissingFields
	}
	if f.Description == "" {
		return models.ErrMissingFields
	}
	if f.StartDate == "" {
		return models.ErrMissingFields
	}
	if f.EndDate == "" {
		return models.ErrMissingFields
	}
	if f.Address.Validate() != nil {
		return models.ErrMissingFields
	}
	return nil
}

func (f *UpdateFestivalRequest) Validate() error {
	// TODO: implement validation
	return nil
}

func (f *CreateItemRequest) Validate() error {
	if f.Name == "" {
		return models.ErrMissingFields
	}
	if f.Description == "" {
		return models.ErrMissingFields
	}
	if f.AvailableNumber == 0 {
		return models.ErrMissingFields
	}
	if f.Type == "" {
		return models.ErrMissingFields
	}
	return nil
}

func (f *CreatePriceListItemRequest) Validate() error {
	if f.ItemID == 0 {
		return models.ErrMissingFields
	}
	if f.Price == 0 {
		return models.ErrMissingFields
	}
	if !f.IsFixed {
		if f.DateFrom == nil {
			return models.ErrMissingFields
		}
		if f.DateTo == nil {
			return models.ErrMissingFields
		}
	}
	return nil
}

func (f *CreatePackageAddonRequest) Validate() error {
	if f.ItemID == 0 {
		return models.ErrMissingFields
	}
	if f.Category == "" {
		return models.ErrMissingFields
	}
	if f.Category != modelsFestival.PackageAddonGeneral && f.Category != modelsFestival.PackageAddonCamp && f.Category != modelsFestival.PackageAddonTransport {
		return models.ErrInvalidFields
	}
	return nil
}

func (f *CreateTransportPackageAddonRequest) Validate() error {
	if f.ItemID == 0 {
		return models.ErrMissingFields
	}
	if f.ArrivalCity.Name == "" || f.ArrivalCity.CountryISO3 == "" || f.ArrivalCity.PostalCode == "" {
		return models.ErrMissingFields
	}
	if f.DepartureCity.Name == "" || f.DepartureCity.CountryISO3 == "" || f.DepartureCity.PostalCode == "" {
		return models.ErrMissingFields
	}
	if f.ArrivalTime == "" {
		return models.ErrMissingFields
	}
	if f.DepartureTime == "" {
		return models.ErrMissingFields
	}
	if f.ReturnArrivalTime == "" {
		return models.ErrMissingFields
	}
	if f.ReturnDepartureTime == "" {
		return models.ErrMissingFields
	}
	return nil
}

func (f *CreateCampPackageAddonRequest) Validate() error {
	if f.ItemID == 0 {
		return models.ErrMissingFields
	}
	if f.CampName == "" {
		return models.ErrMissingFields
	}
	if f.ImageURL == "" {
		return models.ErrMissingFields
	}
	if len(f.EquipmentList) == 0 {
		return models.ErrMissingFields
	}
	return nil
}

func (f *CreateTicketOrderRequest) Validate() error {
	if f.TicketTypeId == 0 {
		return models.ErrMissingFields
	}
	if f.TotalPrice == 0 {
		return models.ErrMissingFields
	}
	return nil
}

func (f *CreatePackageOrderRequest) Validate() error {
	if f.TicketTypeId == 0 {
		return models.ErrMissingFields
	}
	if f.TotalPrice == 0 {
		return models.ErrMissingFields
	}
	return nil
}

func (f *IssueBraceletRequest) Validate() error {
	if f.PIN == "" {
		return models.ErrMissingFields
	}
	if f.BarcodeNumber == "" {
		return models.ErrMissingFields
	}
	if f.AttendeeUsername == "" {
		return models.ErrMissingFields
	}
	if f.FestivalTicketId == 0 {
		return models.ErrMissingFields
	}
	if f.OrderId == 0 {
		return models.ErrMissingFields
	}
	return nil
}

func (f *ActivateBraceletRequest) Validate() error {
	if f.BraceletId == 0 {
		return models.ErrMissingFields
	}
	if f.PIN == "" {
		return models.ErrMissingFields
	}
	return nil
}

func (f *TopUpBraceletRequest) Validate() error {
	if f.BraceletId == 0 {
		return models.ErrMissingFields
	}
	if f.Amount < 1 {
		return errors.New("amount must be greater than 1")
	}
	return nil
}

func (f *ActivateBraceletHelpRequest) Validate() error {
	if f.BraceletId == 0 {
		return models.ErrMissingFields
	}
	if f.BarcodeNumberUser == "" {
		return models.ErrMissingFields
	}
	if f.PINUser == "" {
		return models.ErrMissingFields
	}
	if f.IssueDescription == "" {
		return models.ErrMissingFields
	}
	return nil
}
