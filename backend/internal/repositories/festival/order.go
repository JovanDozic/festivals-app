package repositories

import (
	models "backend/internal/models/festival"

	"gorm.io/gorm"
)

type OrderRepo interface {
	CreateFestivalTicket(festivalTicket *models.FestivalTicket) error
	CreateOrder(order *models.Order) error
	CreateFestivalPackage(festivalPackage *models.FestivalPackage) error
	CreateFestivalPackageAddon(festivalPackageAddon *models.FestivalPackageAddon) error
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

// LAST step for creating any order
func (r *orderRepo) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

// step 2 for creating package order (step 1 is CreateFestivalTicket)
func (r *orderRepo) CreateFestivalPackage(festivalPackage *models.FestivalPackage) error {
	return r.db.Create(festivalPackage).Error
}

// step 3 for creating package order (do this per every addon)
func (r *orderRepo) CreateFestivalPackageAddon(festivalPackageAddon *models.FestivalPackageAddon) error {
	return r.db.Create(festivalPackageAddon).Error
}
