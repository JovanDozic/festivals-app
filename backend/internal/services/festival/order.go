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
}

func NewOrderService(orderRepo reposFestival.OrderRepo) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (s *orderService) CreateFestivalTicket(festivalTicket *models.FestivalTicket) error {
	return s.orderRepo.CreateFestivalTicket(festivalTicket)
}

func (s *orderService) CreateOrder(order *models.Order) error {
	return s.orderRepo.CreateOrder(order)
}
