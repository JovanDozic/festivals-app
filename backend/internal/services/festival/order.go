package services

import (
	dtoCommon "backend/internal/dto/common"
	dtoFestival "backend/internal/dto/festival"
	modelsCommon "backend/internal/models/common"
	modelsFestival "backend/internal/models/festival"
	reposCommon "backend/internal/repositories/common"
	reposFestival "backend/internal/repositories/festival"
	servicesCommon "backend/internal/services/common"
	servicesUser "backend/internal/services/user"
	"errors"
	"log"
	"strings"
)

type OrderService interface {
	CreateFestivalTicket(festivalTicket *modelsFestival.FestivalTicket) error
	CreateOrder(order *modelsFestival.Order) error
	CreateFestivalPackage(festivalPackage *modelsFestival.FestivalPackage) error
	CreateFestivalPackageAddon(festivalPackageAddon *modelsFestival.FestivalPackageAddon) error
	GetOrder(username string, orderId uint) (*dtoFestival.OrderDTO, error)
	GetOrdersAttendee(username string) ([]dtoFestival.OrderPreviewDTO, error)
	GetOrdersEmployee(festivalId uint) ([]dtoFestival.OrderPreviewDTO, error)
	GetBraceletOrdersAttendee(username string) ([]dtoFestival.OrderDTO, error)
	IssueBracelet(request *modelsFestival.Bracelet) error
	ActivateBracelet(username string, braceletId uint, userEnteredPIN string) error
	TopUpBracelet(username string, braceletId uint, amount float64) error
	CreateHelpRequest(username string, request dtoFestival.ActivateBraceletHelpRequest) error
	GetHelpRequest(braceletId uint) (*modelsFestival.ActivationHelpRequest, error)
	ApproveHelpRequest(braceletId uint) error
	RejectHelpRequest(braceletId uint) error
	GetShippingLabel(orderId uint) ([]byte, error)
	GetBraceletById(braceletId uint) (*modelsFestival.Bracelet, error)
}

type orderService struct {
	orderRepo       reposFestival.OrderRepo
	itemRepo        reposFestival.ItemRepo
	festivalRepo    reposFestival.FestivalRepo
	userService     servicesUser.UserService
	imageRepo       reposCommon.ImageRepo
	locationService servicesCommon.LocationService
	pdfGenerator    servicesCommon.PDFGenerator
	emailService    servicesCommon.EmailService
}

func NewOrderService(
	or reposFestival.OrderRepo,
	ir reposFestival.ItemRepo,
	fr reposFestival.FestivalRepo,
	us servicesUser.UserService,
	imr reposCommon.ImageRepo,
	ls servicesCommon.LocationService,
	pg servicesCommon.PDFGenerator,
) OrderService {
	return &orderService{
		orderRepo:       or,
		itemRepo:        ir,
		festivalRepo:    fr,
		userService:     us,
		imageRepo:       imr,
		locationService: ls,
		pdfGenerator:    pg,
	}
}

func (s *orderService) CreateFestivalTicket(festivalTicket *modelsFestival.FestivalTicket) error {

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

func (s *orderService) CreateOrder(order *modelsFestival.Order) error {
	return s.orderRepo.CreateOrder(order)
}

func (s *orderService) CreateFestivalPackage(festivalPackage *modelsFestival.FestivalPackage) error {
	return s.orderRepo.CreateFestivalPackage(festivalPackage)
}

func (s *orderService) CreateFestivalPackageAddon(festivalPackageAddon *modelsFestival.FestivalPackageAddon) error {

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

	if username != "" && order.User.User.Username != username {
		return nil, errors.New("order not found")
	}

	attendee, err := s.userService.GetUserProfile(order.User.User.Username)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	orderDto := &dtoFestival.OrderDTO{
		OrderID:    order.ID,
		Timestamp:  order.CreatedAt,
		TotalPrice: order.TotalAmount,
		Username:   order.User.User.Username,
		Attendee:   attendee,
	}

	if order.FestivalPackage == nil {
		orderDto.OrderType = "TICKET"
	} else {
		orderDto.OrderType = "PACKAGE"
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

	orderDto.Ticket = dtoFestival.ItemResponse{
		ItemId:      ticketItem.ID,
		Name:        ticketItem.Name,
		Price:       0,
		Type:        ticketItem.Type,
		Description: ticketItem.Description,
	}
	orderDto.FestivalTicketId = &festivalTicket.ID

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
				orderDto.TransportAddon = transportAddon
			}

			if addon.Category == "CAMP" {
				campAddon, err := s.itemRepo.GetCampAddon(addon.ItemID)
				if err != nil {
					log.Println("error: ", err)
					return nil, err
				}
				orderDto.CampAddon = campAddon
			}

			if addon.Category == "GENERAL" {
				generalAddon, err := s.itemRepo.GetGeneralAddon(addon.ItemID)
				if err != nil {
					log.Println("error: ", err)
					return nil, err
				}
				orderDto.GeneralAddons = append(orderDto.GeneralAddons, *generalAddon)
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
			AddressId:      &festival.Address.ID,
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

	orderDto.Festival = dtoFestival.FestivalResponse{
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

	bracelet, err := s.orderRepo.GetBraceletByTicketId(order.FestivalTicketID)
	if err != nil && !strings.Contains(err.Error(), "record not found") {
		log.Println("error: ", err)
		return nil, err
	}

	if bracelet != nil && bracelet.ID != 0 {
		orderDto.BraceletStatus = &bracelet.Status
		employee, err := s.userService.GetUserProfileById(bracelet.EmployeeID)
		if err != nil {
			log.Println("error: ", err)
			return nil, err
		}

		orderDto.Bracelet = &dtoFestival.BraceletDTO{
			BraceletID:    bracelet.ID,
			BarcodeNumber: bracelet.BarcodeNumber,
			Status:        bracelet.Status,
			Balance:       bracelet.Balance,
			Employee:      employee,
		}
	} else {
		orderDto.BraceletStatus = nil
		orderDto.Bracelet = nil
	}

	return orderDto, nil
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
			FestivalTicketId: &order.FestivalTicketID,
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

		bracelet, err := s.orderRepo.GetBraceletByTicketId(order.FestivalTicketID)
		if err != nil && !strings.Contains(err.Error(), "record not found") {
			log.Println("error: ", err)
			return nil, err
		}

		if bracelet != nil {
			orderDto.BraceletStatus = &bracelet.Status
			orderDto.BraceletID = &bracelet.ID
		} else {
			orderDto.BraceletStatus = nil
		}

		response = append(response, orderDto)
	}

	return response, nil
}

func (s *orderService) IssueBracelet(request *modelsFestival.Bracelet) error {
	return s.orderRepo.CreateBracelet(request)
}

func (s *orderService) GetBraceletOrdersAttendee(username string) ([]dtoFestival.OrderDTO, error) {

	orders, err := s.orderRepo.GetOrdersAttendee(username)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	var response []dtoFestival.OrderDTO

	for _, order := range orders {

		orderDto := dtoFestival.OrderDTO{
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

		bracelet, err := s.orderRepo.GetBraceletByTicketId(order.FestivalTicketID)
		if err != nil && !strings.Contains(err.Error(), "record not found") {
			log.Println("error: ", err)
			return nil, err
		}

		if bracelet != nil && bracelet.ID != 0 {
			orderDto.BraceletStatus = &bracelet.Status
			employee, err := s.userService.GetUserProfileById(bracelet.EmployeeID)
			if err != nil {
				log.Println("error: ", err)
				return nil, err
			}

			orderDto.Bracelet = &dtoFestival.BraceletDTO{
				BraceletID:    bracelet.ID,
				BarcodeNumber: bracelet.BarcodeNumber,
				Status:        bracelet.Status,
				Balance:       bracelet.Balance,
				Employee:      employee,
			}
		} else {
			orderDto.BraceletStatus = nil
			orderDto.Bracelet = nil
		}

		response = append(response, orderDto)
	}

	return response, nil
}

func (s *orderService) ActivateBracelet(username string, braceletId uint, userEnteredPIN string) error {

	bracelet, err := s.orderRepo.GetBraceletById(braceletId)
	if err != nil {
		return err
	}

	if bracelet.Attendee.User.Username != username {
		return errors.New("bracelet not found or it does not belong to logged in user")
	}

	// only activate if it is issued
	if bracelet.Status != "ISSUED" {
		return errors.New("bracelet is not issued (status: " + bracelet.Status + ")")
	}

	if bracelet.PIN != userEnteredPIN {
		return errors.New("invalid PIN")
	}

	bracelet.Status = "ACTIVATED"

	return s.orderRepo.UpdateBracelet(bracelet)
}

func (s *orderService) TopUpBracelet(username string, braceletId uint, amount float64) error {

	bracelet, err := s.orderRepo.GetBraceletById(braceletId)
	if err != nil {
		return err
	}

	if bracelet.Attendee.User.Username != username {
		return errors.New("bracelet not found or it does not belong to logged in user")
	}

	// only activate if it is issued
	if bracelet.Status != "ACTIVATED" {
		return errors.New("bracelet is not active (status: " + bracelet.Status + ")")
	}

	bracelet.Balance += amount

	return s.orderRepo.UpdateBracelet(bracelet)
}

func (s *orderService) CreateHelpRequest(username string, request dtoFestival.ActivateBraceletHelpRequest) error {

	attendeeId, err := s.userService.GetUserID(username)
	if err != nil {
		return err
	}

	bracelet, err := s.orderRepo.GetBraceletById(request.BraceletId)
	if err != nil {
		return err
	}

	bracelet.Status = "HELP_REQUESTED"

	if err := s.orderRepo.UpdateBracelet(bracelet); err != nil {
		return err
	}

	image := modelsCommon.Image{
		URL: request.ImageURL,
	}

	if err := s.imageRepo.Create(&image); err != nil {
		return err
	}

	helpRequest := modelsFestival.ActivationHelpRequest{
		UserEnteredPIN:     request.PINUser,
		UserEnteredBarcode: request.BarcodeNumberUser,
		IssueDescription:   request.IssueDescription,
		Status:             "OPEN",
		BraceletID:         request.BraceletId,
		AttendeeID:         attendeeId,
		ProofImageID:       image.ID,
	}

	if err := s.orderRepo.CreateHelpRequest(&helpRequest); err != nil {
		return err
	}

	return nil
}

func (s *orderService) GetHelpRequest(braceletId uint) (*modelsFestival.ActivationHelpRequest, error) {
	return s.orderRepo.GetHelpRequest(braceletId)
}

func (s *orderService) ApproveHelpRequest(braceletId uint) error {

	helpRequest, err := s.orderRepo.GetHelpRequest(braceletId)
	if err != nil {
		return err
	}

	bracelet, err := s.orderRepo.GetBraceletById(braceletId)
	if err != nil {
		return err
	}

	bracelet.Status = "ACTIVATED"

	if err := s.orderRepo.UpdateBracelet(bracelet); err != nil {
		return err
	}

	helpRequest.Status = "APPROVED"

	return s.orderRepo.UpdateHelpRequest(helpRequest)
}

func (s *orderService) RejectHelpRequest(braceletId uint) error {

	helpRequest, err := s.orderRepo.GetHelpRequest(braceletId)
	if err != nil {
		return err
	}

	bracelet, err := s.orderRepo.GetBraceletById(braceletId)
	if err != nil {
		return err
	}

	bracelet.Status = "REJECTED"

	if err := s.orderRepo.UpdateBracelet(bracelet); err != nil {
		return err
	}

	helpRequest.Status = "REJECTED"

	return s.orderRepo.UpdateHelpRequest(helpRequest)
}

func (s *orderService) GetShippingLabel(orderId uint) ([]byte, error) {

	order, err := s.GetOrder("", orderId)
	if err != nil {
		return nil, err
	}

	attendee := order.Attendee
	festival := order.Festival

	attendeeAddress, err := s.locationService.GetAddressByID(*attendee.Address.AddressId)
	if err != nil {
		return nil, err
	}

	festivalAddressId, err := s.locationService.GetAddressByID(*festival.Address.AddressId)
	if err != nil {
		return nil, err
	}

	pdfBytes, err := s.pdfGenerator.CreateShippingLabel(festivalAddressId, attendeeAddress, &festival, *attendee)
	if err != nil {
		return nil, err
	}

	return pdfBytes, nil
}

func (s *orderService) GetBraceletById(braceletId uint) (*modelsFestival.Bracelet, error) {
	return s.orderRepo.GetBraceletById(braceletId)
}
