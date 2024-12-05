package services

import (
	dtoCommon "backend/internal/dto/common"
	dtoFestival "backend/internal/dto/festival"
	models "backend/internal/models/festival"
	reposFestival "backend/internal/repositories/festival"
	servicesUser "backend/internal/services/user"
	"errors"
	"log"
)

type OrderService interface {
	CreateFestivalTicket(festivalTicket *models.FestivalTicket) error
	CreateOrder(order *models.Order) error
	CreateFestivalPackage(festivalPackage *models.FestivalPackage) error
	CreateFestivalPackageAddon(festivalPackageAddon *models.FestivalPackageAddon) error
	GetOrder(username string, orderId uint) (*dtoFestival.OrderDTO, error)
	GetOrdersAttendee(username string) ([]dtoFestival.OrderPreviewDTO, error)
	GetOrdersEmployee(festivalId uint) ([]dtoFestival.OrderPreviewDTO, error)
}

type orderService struct {
	orderRepo    reposFestival.OrderRepo
	itemRepo     reposFestival.ItemRepo
	festivalRepo reposFestival.FestivalRepo
	userService  servicesUser.UserService
}

func NewOrderService(or reposFestival.OrderRepo, ir reposFestival.ItemRepo, fr reposFestival.FestivalRepo, us servicesUser.UserService) OrderService {
	return &orderService{
		orderRepo:    or,
		itemRepo:     ir,
		festivalRepo: fr,
		userService:  us,
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

func (s *orderService) GetOrder(username string, orderId uint) (*dtoFestival.OrderDTO, error) {

	order, err := s.orderRepo.GetOrder(orderId)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	if order.User.User.Username != username {
		return nil, errors.New("order not found")
	}

	response := &dtoFestival.OrderDTO{
		OrderID:    order.ID,
		Timestamp:  order.CreatedAt,
		TotalPrice: order.TotalAmount,
		Username:   order.User.User.Username,
	}

	if order.FestivalPackage == nil {
		response.OrderType = "TICKET"
	} else {
		response.OrderType = "PACKAGE"
	}

	// * get ticket

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

	// * get package

	if order.FestivalPackage != nil {
		festivalPackage, err := s.orderRepo.GetFestivalPackage(*order.FestivalPackageID)
		if err != nil {
			log.Println("error: ", err)
			return nil, err
		}

		packageAddons, err := s.itemRepo.GetAddonsFromPackage(festivalPackage.ID)
		if err != nil {
			log.Println("error: ", err)
			return nil, err
		}

		for _, addon := range packageAddons {

			if addon.Category == "TRANSPORT" {
				transportAddon, err := s.itemRepo.GetTransportAddon(addon.ItemID)
				if err != nil {
					log.Println("error: ", err)
					return nil, err
				}
				response.TransportAddon = transportAddon
			}

			if addon.Category == "CAMP" {
				campAddon, err := s.itemRepo.GetCampAddon(addon.ItemID)
				if err != nil {
					log.Println("error: ", err)
					return nil, err
				}
				response.CampAddon = campAddon
			}

			if addon.Category == "GENERAL" {
				generalAddon, err := s.itemRepo.GetGeneralAddon(addon.ItemID)
				if err != nil {
					log.Println("error: ", err)
					return nil, err
				}
				response.GeneralAddons = append(response.GeneralAddons, *generalAddon)
			}
		}
	}

	// * now we get festival

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

func (s *orderService) GetOrdersAttendee(username string) ([]dtoFestival.OrderPreviewDTO, error) {

	orders, err := s.orderRepo.GetOrdersAttendee(username)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	var response []dtoFestival.OrderPreviewDTO

	for _, order := range orders {

		orderDto := dtoFestival.OrderPreviewDTO{
			OrderID:    order.ID,
			Timestamp:  order.CreatedAt,
			TotalPrice: order.TotalAmount,
			Username:   order.User.User.Username,
			Festival: dtoFestival.FestivalResponse{
				ID:        order.FestivalTicket.Item.Item.Festival.ID,
				Name:      order.FestivalTicket.Item.Item.Festival.Name,
				StartDate: order.FestivalTicket.Item.Item.Festival.StartDate,
				EndDate:   order.FestivalTicket.Item.Item.Festival.EndDate,
			},
		}

		if order.FestivalPackage == nil {
			orderDto.OrderType = "TICKET"
		} else {
			orderDto.OrderType = "PACKAGE"
		}

		response = append(response, orderDto)
	}

	return response, nil
}

func (s *orderService) GetOrdersEmployee(festivalId uint) ([]dtoFestival.OrderPreviewDTO, error) {

	orders, err := s.orderRepo.GetOrdersEmployee(festivalId)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	var response []dtoFestival.OrderPreviewDTO

	for _, order := range orders {

		orderDto := dtoFestival.OrderPreviewDTO{
			OrderID:    order.ID,
			Timestamp:  order.CreatedAt,
			TotalPrice: order.TotalAmount,
			Username:   order.User.User.Username,
			Festival: dtoFestival.FestivalResponse{
				ID:        order.FestivalTicket.Item.Item.Festival.ID,
				Name:      order.FestivalTicket.Item.Item.Festival.Name,
				StartDate: order.FestivalTicket.Item.Item.Festival.StartDate,
				EndDate:   order.FestivalTicket.Item.Item.Festival.EndDate,
			},
		}

		if order.FestivalPackage == nil {
			orderDto.OrderType = "TICKET"
		} else {
			orderDto.OrderType = "PACKAGE"
		}

		attendee, err := s.userService.GetUserProfile(order.User.User.Username)
		if err != nil {
			log.Println("error: ", err)
			return nil, err
		}

		orderDto.Attendee = attendee

		response = append(response, orderDto)
	}

	// todo: append bracelet status

	return response, nil
}
