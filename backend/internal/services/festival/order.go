package services

import (
	models "backend/internal/models/festival"
	reposFestival "backend/internal/repositories/festival"
)

type OrderService interface {
	CreateFestivalTicket(festivalTicket *models.FestivalTicket) error
	CreateOrder(order *models.Order) error
}

type orderService struct {
	orderRepo reposFestival.OrderRepo
	itemRepo  reposFestival.ItemRepo
}

func NewOrderService(or reposFestival.OrderRepo, ir reposFestival.ItemRepo) OrderService {
	return &orderService{
		orderRepo: or,
		itemRepo:  ir,
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
