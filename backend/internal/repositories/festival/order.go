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
	GetOrdersAttendee(username string) ([]models.Order, error)
	GetOrdersEmployee(festivalId uint) ([]models.Order, error)
	GetFestivalTicket(festivalTicketId uint) (*models.FestivalTicket, error)
	GetFestivalPackage(festivalPackageId uint) (*models.FestivalPackage, error)
	CreateBracelet(bracelet *models.Bracelet) error
	GetBraceletByTicketId(festivalTicketId uint) (*models.Bracelet, error)
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

func (r *orderRepo) GetOrdersAttendee(username string) ([]models.Order, error) {
	orders := []models.Order{}
	err := r.db.
		Preload("FestivalTicket").
		Preload("FestivalPackage").
		Preload("User").
		Preload("User.User").
		Preload("FestivalTicket.Item.Item.Festival").
		Joins("JOIN users ON orders.user_id = users.id").
		Where("users.username = ?", username).
		Find(&orders).Error
	return orders, err
}

func (r *orderRepo) GetOrdersEmployee(festivalId uint) ([]models.Order, error) {
	orders := []models.Order{}
	err := r.db.
		Preload("FestivalTicket").
		Preload("FestivalPackage").
		Preload("User").
		Preload("User.User").
		Preload("FestivalTicket.Item.Item.Festival").
		Joins("JOIN festival_tickets ON orders.festival_ticket_id = festival_tickets.id").
		Joins("JOIN items ON festival_tickets.item_id = items.id").
		Joins("JOIN festivals ON items.festival_id = festivals.id").
		Where("festivals.id = ?", festivalId).
		Find(&orders).Error
	return orders, err
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

func (r *orderRepo) CreateBracelet(bracelet *models.Bracelet) error {
	return r.db.Create(bracelet).Error
}

func (r *orderRepo) GetBraceletByTicketId(festivalTicketId uint) (*models.Bracelet, error) {
	bracelet := &models.Bracelet{}
	err := r.db.
		Where("festival_ticket_id = ?", festivalTicketId).
		First(&bracelet).Error
	return bracelet, err
}
