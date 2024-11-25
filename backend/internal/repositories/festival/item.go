package repositories

import (
	modelsFestival "backend/internal/models/festival"
	"errors"

	"gorm.io/gorm"
)

type ItemRepo interface {
	CreateItem(item *modelsFestival.Item) error
	CreatePriceList(priceList *modelsFestival.PriceList) error
	GetPriceList(festivalId uint) (*modelsFestival.PriceList, error)
	CreatePriceListItem(priceListItem *modelsFestival.PriceListItem) error
}

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) ItemRepo {
	return &itemRepo{db}
}

func (r *itemRepo) CreateItem(item *modelsFestival.Item) error {

	err := r.db.Create(item).Error
	if err != nil {
		return err
	}

	if item.Type == modelsFestival.ItemTicketType {
		ticketType := modelsFestival.TicketType{
			ItemID: item.ID,
			Item:   *item,
		}
		err = r.db.Create(&ticketType).Error
		if err != nil {
			return err
		}
		return nil
	} else if item.Type == modelsFestival.ItemPackageAddon {
		packageAddon := modelsFestival.PackageAddon{
			ItemID: item.ID,
			Item:   *item,
		}
		err = r.db.Create(&packageAddon).Error
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("invalid item type")
}

func (r *itemRepo) CreatePriceList(priceList *modelsFestival.PriceList) error {
	err := r.db.Create(priceList).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *itemRepo) GetPriceList(festivalId uint) (*modelsFestival.PriceList, error) {
	var priceList modelsFestival.PriceList
	err := r.db.Where("festival_id = ?", festivalId).First(&priceList).Error
	if err != nil {
		return nil, err
	}

	return &priceList, nil
}

func (r *itemRepo) CreatePriceListItem(priceListItem *modelsFestival.PriceListItem) error {
	err := r.db.Create(priceListItem).Error
	if err != nil {
		return err
	}

	return nil
}
