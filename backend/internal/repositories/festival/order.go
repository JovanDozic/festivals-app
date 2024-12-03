package repositories

import (
	models "backend/internal/models/festival"

	"gorm.io/gorm"
)

type OrderRepo interface {
	CreateFestivalTicket(festivalTicket *models.FestivalTicket) error
	CreateOrder(order *models.Order) error
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepo {
	return &orderRepo{db}
}

// this one creates instance of a item - festival ticket (this is connecting order and ticket_type -> item)
func (r *orderRepo) CreateFestivalTicket(festivalTicket *models.FestivalTicket) error {
	return r.db.Create(festivalTicket).Error
}

func (r *orderRepo) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}
