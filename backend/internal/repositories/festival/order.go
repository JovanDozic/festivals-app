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
	GetOrder(orderId uint) (*models.Order, error)
	GetFestivalTicket(festivalTicketId uint) (*models.FestivalTicket, error)
	GetFestivalPackage(festivalPackageId uint) (*models.FestivalPackage, error)
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepo {
	return &orderRepo{db}
}

func (r *orderRepo) CreateFestivalTicket(festivalTicket *models.FestivalTicket) error {
	// this one creates instance of a item - festival ticket (this is connecting order and ticket_type -> item)
	return r.db.Create(festivalTicket).Error
}

func (r *orderRepo) CreateOrder(order *models.Order) error {
	// LAST step for creating any order
	return r.db.Create(order).Error
}

func (r *orderRepo) CreateFestivalPackage(festivalPackage *models.FestivalPackage) error {
	// step 2 for creating package order (step 1 is CreateFestivalTicket)
	return r.db.Create(festivalPackage).Error
}

func (r *orderRepo) CreateFestivalPackageAddon(festivalPackageAddon *models.FestivalPackageAddon) error {
	// step 3 for creating package order (do this per every addon)
	return r.db.Create(festivalPackageAddon).Error
}

func (r *orderRepo) GetOrder(orderId uint) (*models.Order, error) {
	order := &models.Order{}
	err := r.db.
		Preload("FestivalTicket").
		Preload("FestivalPackage").
		Preload("User").
		Preload("User.User").
		First(order, orderId).Error
	return order, err
}

func (r *orderRepo) GetFestivalTicket(festivalTicketId uint) (*models.FestivalTicket, error) {
	festivalTicket := &models.FestivalTicket{}
	err := r.db.Where("id = ?", festivalTicketId).First(&festivalTicket).Error
	return festivalTicket, err
}

func (r *orderRepo) GetFestivalPackage(festivalPackageId uint) (*models.FestivalPackage, error) {
	festivalPackage := &models.FestivalPackage{}
	err := r.db.Where("id = ?", festivalPackageId).First(&festivalPackage).Error
	return festivalPackage, err
}
