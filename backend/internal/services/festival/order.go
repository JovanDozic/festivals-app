package services

import (
	dtoCommon "backend/internal/dto/common"
	dtoFestival "backend/internal/dto/festival"
	models "backend/internal/models/festival"
	reposFestival "backend/internal/repositories/festival"
	"log"
)

type OrderService interface {
	CreateFestivalTicket(festivalTicket *models.FestivalTicket) error
	CreateOrder(order *models.Order) error
	CreateFestivalPackage(festivalPackage *models.FestivalPackage) error
	CreateFestivalPackageAddon(festivalPackageAddon *models.FestivalPackageAddon) error
	GetOrder(orderId uint) (*dtoFestival.OrderDTO, error)
}

type orderService struct {
	orderRepo    reposFestival.OrderRepo
	itemRepo     reposFestival.ItemRepo
	festivalRepo reposFestival.FestivalRepo
}

func NewOrderService(or reposFestival.OrderRepo, ir reposFestival.ItemRepo, fr reposFestival.FestivalRepo) OrderService {
	return &orderService{
		orderRepo:    or,
		itemRepo:     ir,
		festivalRepo: fr,
	}
}

func (s *orderService) CreateFestivalTicket(festivalTicket *models.FestivalTicket) error {

	item, _, err := s.itemRepo.GetItemAndPriceListItemsIDs(festivalTicket.ItemID)
	if err != nil {
		return nil
	}

	item.RemainingNumber -= 1

	if err := s.itemRepo.UpdateItem(item); err != nil {
		return err
	}

	return s.orderRepo.CreateFestivalTicket(festivalTicket)
}

func (s *orderService) CreateOrder(order *models.Order) error {
	return s.orderRepo.CreateOrder(order)
}

func (s *orderService) CreateFestivalPackage(festivalPackage *models.FestivalPackage) error {
	return s.orderRepo.CreateFestivalPackage(festivalPackage)
}

func (s *orderService) CreateFestivalPackageAddon(festivalPackageAddon *models.FestivalPackageAddon) error {

	item, _, err := s.itemRepo.GetItemAndPriceListItemsIDs(festivalPackageAddon.ItemID)
	if err != nil {
		return nil
	}

	item.RemainingNumber -= 1

	if err := s.itemRepo.UpdateItem(item); err != nil {
		return err
	}

	return s.orderRepo.CreateFestivalPackageAddon(festivalPackageAddon)
}

func (s *orderService) GetOrder(orderId uint) (*dtoFestival.OrderDTO, error) {

	order, err := s.orderRepo.GetOrder(orderId)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	response := &dtoFestival.OrderDTO{
		OrderID:    order.ID,
		Timestamp:  order.CreatedAt,
		TotalPrice: order.TotalAmount,
	}

	if order.FestivalPackage == nil {
		response.OrderType = "TICKET"
	} else {
		response.OrderType = "PACKAGE"
	}

	festivalTicket, err := s.orderRepo.GetFestivalTicket(order.FestivalTicketID)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	ticketItem, _, err := s.itemRepo.GetItemAndPriceListItemsIDs(festivalTicket.ItemID)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	response.Ticket = dtoFestival.ItemResponse{
		ItemId:      ticketItem.ID,
		Name:        ticketItem.Name,
		Price:       0,
		Type:        ticketItem.Type,
		Description: ticketItem.Description,
	}

	// todo: get package

	// now we get festival

	festival, err := s.festivalRepo.GetById(ticketItem.FestivalID)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	var address *dtoCommon.GetAddressResponse
	if festival.Address != nil {
		address = &dtoCommon.GetAddressResponse{
			Street:         festival.Address.Street,
			Number:         festival.Address.Number,
			ApartmentSuite: festival.Address.ApartmentSuite,
			City:           festival.Address.City.Name,
			PostalCode:     festival.Address.City.PostalCode,
			Country:        festival.Address.City.Country.NiceName,
			CountryISO3:    festival.Address.City.Country.ISO3,
			CountryISO2:    festival.Address.City.Country.ISO,
		}
	} else {
		address = nil
	}

	response.Festival = dtoFestival.FestivalResponse{
		ID:          festival.ID,
		Name:        festival.Name,
		Description: festival.Description,
		StartDate:   festival.StartDate,
		EndDate:     festival.EndDate,
		Capacity:    festival.Capacity,
		Status:      festival.Status,
		StoreStatus: festival.StoreStatus,
		Address:     address,
	}

	// todo: get bracelet

	return response, nil
}
