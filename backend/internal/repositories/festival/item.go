package repositories

import (
	"backend/internal/models"
	modelsFestival "backend/internal/models/festival"
	"errors"
	"time"

	"gorm.io/gorm"
)

type ItemRepo interface {
	CreateItem(item *modelsFestival.Item) error
	CreatePriceList(priceList *modelsFestival.PriceList) error
	GetPriceList(festivalId uint) (*modelsFestival.PriceList, error)
	CreatePriceListItem(priceListItem *modelsFestival.PriceListItem) error
	GetCurrentTicketTypes(festivalId uint) ([]modelsFestival.PriceListItem, error)
	GetTicketTypesCount(festivalId uint) (int, error)
	GetItemAndPriceListItemsIDs(itemId uint) (*modelsFestival.Item, []uint, error)
	GetPriceListItemsByIDs(priceListItemIDs []uint) ([]modelsFestival.PriceListItem, error)
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

func (r *itemRepo) GetCurrentTicketTypes(festivalId uint) ([]modelsFestival.PriceListItem, error) {

	// * Get festival price list
	var currentPriceList modelsFestival.PriceList
	err := r.db.
		Where("festival_id = ?", festivalId).
		First(&currentPriceList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrNoPriceListFound
		}
		return nil, err
	}

	// * Get all price list items with preloaded items
	var priceListItems []modelsFestival.PriceListItem
	err = r.db.
		Preload("Item").
		Joins("JOIN items ON price_list_items.item_id = items.id").
		Where("items.type = ?", modelsFestival.ItemTicketType).
		Where("price_list_id = ?", currentPriceList.ID).
		// Where("date_from <= ?", time.Now()).
		// Where("date_to >= ?", time.Now()).
		Find(&priceListItems).Error
	if err != nil {
		return nil, err
	}

	now := time.Now()
	filteredPriceListItems := make([]modelsFestival.PriceListItem, 0)
	for _, pli := range priceListItems {
		if !pli.IsFixed && pli.DateFrom != nil && pli.DateTo != nil {
			if now.Before(*pli.DateFrom) || now.After(*pli.DateTo) {
				continue
			}
		}
		filteredPriceListItems = append(filteredPriceListItems, pli)
	}

	return filteredPriceListItems, nil
}

func (r *itemRepo) GetTicketTypesCount(festivalId uint) (int, error) {

	// todo: make sure this returns only ticket types that have a price with them
	var count int64
	err := r.db.Table("items").
		Where("festival_id = ? AND type = ?", festivalId, modelsFestival.ItemTicketType).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

// this guy returns item info and IDs of all of it's prices
func (r *itemRepo) GetItemAndPriceListItemsIDs(itemId uint) (*modelsFestival.Item, []uint, error) {
	var item modelsFestival.Item
	var priceListItemIDs []uint

	err := r.db.Model(&modelsFestival.PriceListItem{}).
		Select("id").
		Where("item_id = ?", itemId).
		Find(&priceListItemIDs).Error
	if err != nil {
		return nil, nil, err
	}

	if err := r.db.First(&item, itemId).Error; err != nil {
		return nil, nil, err
	}

	return &item, priceListItemIDs, nil
}

// this guy get all prices for given list of IDs
func (r *itemRepo) GetPriceListItemsByIDs(priceListItemIDs []uint) ([]modelsFestival.PriceListItem, error) {
	var priceListItems []modelsFestival.PriceListItem

	err := r.db.
		Where("id IN ?", priceListItemIDs).
		Find(&priceListItems).Error
	if err != nil {
		return nil, err
	}

	return priceListItems, nil
}
