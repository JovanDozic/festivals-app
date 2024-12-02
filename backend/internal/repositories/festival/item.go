package repositories

import (
	"backend/internal/models"
	modelsFestival "backend/internal/models/festival"
	"backend/internal/utils"
	"errors"
	"time"

	"gorm.io/gorm"
)

type ItemRepo interface {
	CreateItem(item *modelsFestival.Item) error
	UpdatePackageAddonCategory(packageAddon *modelsFestival.PackageAddon) error
	CreatePriceList(priceList *modelsFestival.PriceList) error
	GetPriceList(festivalId uint) (*modelsFestival.PriceList, error)
	CreatePriceListItem(priceListItem *modelsFestival.PriceListItem) error
	GetCurrentTicketTypes(festivalId uint) ([]modelsFestival.PriceListItem, error)
	GetTicketTypesCount(festivalId uint) (int, error)
	GetPackageAddonsCount(festivalId uint, category string) (int, error)
	GetItemAndPriceListItemsIDs(itemId uint) (*modelsFestival.Item, []uint, error)
	GetPriceListItemsByIDs(priceListItemIDs []uint) ([]modelsFestival.PriceListItem, error)
	UpdateItem(item *modelsFestival.Item) error
	UpdatePriceListItem(priceListItem *modelsFestival.PriceListItem) error
	DeleteTicketType(itemID uint) error
	GetCurrentPackageAddons(festivalId uint, category string) ([]modelsFestival.PriceListItem, error)
	CreateTransportPackageAddon(transportAddon *modelsFestival.TransportAddon) error
	CreateCampPackageAddon(campAddon *modelsFestival.CampAddon) error
	CreateCampEquipment(campEquipment *modelsFestival.CampEquipment) error
	CreatePackageAddonImage(packageAddonImage *modelsFestival.PackageAddonImage) error
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

func (r *itemRepo) UpdatePackageAddonCategory(packageAddon *modelsFestival.PackageAddon) error {
	return r.db.Save(packageAddon).Error
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
		Order("items.id").
		Find(&priceListItems).Error
	if err != nil {
		return nil, err
	}

	today := utils.StripTime(time.Now())

	filteredPriceListItems := make([]modelsFestival.PriceListItem, 0, len(priceListItems))
	for _, pli := range priceListItems {
		if !pli.IsFixed && pli.DateFrom != nil && pli.DateTo != nil {
			dateFrom := utils.StripTime(*pli.DateFrom)
			dateTo := utils.StripTime(*pli.DateTo)
			if !utils.IsDateInRange(today, dateFrom, dateTo) {
				continue
			}
		}
		filteredPriceListItems = append(filteredPriceListItems, pli)
	}

	return filteredPriceListItems, nil
}

func (r *itemRepo) GetCurrentPackageAddons(festivalId uint, category string) ([]modelsFestival.PriceListItem, error) {

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
		Joins("JOIN package_addons ON items.id = package_addons.item_id").
		Where("package_addons.category = ?", category).
		Where("items.type = ?", modelsFestival.ItemPackageAddon).
		Where("price_list_id = ?", currentPriceList.ID).
		Order("items.id").
		Find(&priceListItems).Error
	if err != nil {
		return nil, err
	}

	today := utils.StripTime(time.Now())

	filteredPriceListItems := make([]modelsFestival.PriceListItem, 0, len(priceListItems))
	for _, pli := range priceListItems {
		if !pli.IsFixed && pli.DateFrom != nil && pli.DateTo != nil {
			dateFrom := utils.StripTime(*pli.DateFrom)
			dateTo := utils.StripTime(*pli.DateTo)
			if !utils.IsDateInRange(today, dateFrom, dateTo) {
				continue
			}
		}
		filteredPriceListItems = append(filteredPriceListItems, pli)
	}

	return filteredPriceListItems, nil
}

func (r *itemRepo) GetTicketTypesCount(festivalId uint) (int, error) {
	var count int64
	err := r.db.Table("items").
		Joins("JOIN price_list_items ON items.id = price_list_items.item_id").
		Where("items.festival_id = ? AND items.type = ? AND items.deleted_at IS NULL", festivalId, modelsFestival.ItemTicketType).
		Select("COUNT(DISTINCT items.id)").
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *itemRepo) GetPackageAddonsCount(festivalId uint, category string) (int, error) {
	var count int64
	err := r.db.Table("items").
		Joins("JOIN price_list_items ON items.id = price_list_items.item_id").
		Joins("JOIN package_addons ON items.id = package_addons.item_id").
		Where("items.festival_id = ? AND items.type = ? AND items.deleted_at IS NULL AND package_addons.category = ?", festivalId, modelsFestival.ItemPackageAddon, category).
		Select("COUNT(DISTINCT items.id)").
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
		Order("date_from").
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
		Order("date_from").
		Find(&priceListItems).Error
	if err != nil {
		return nil, err
	}

	return priceListItems, nil
}

func (r *itemRepo) UpdateItem(item *modelsFestival.Item) error {
	err := r.db.Save(item).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *itemRepo) UpdatePriceListItem(priceListItem *modelsFestival.PriceListItem) error {
	err := r.db.Save(priceListItem).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *itemRepo) DeleteTicketType(itemID uint) error {

	return r.db.Transaction((func(tx *gorm.DB) error {
		if err := tx.Where("item_id = ?", itemID).Delete(&modelsFestival.PriceListItem{}).Error; err != nil {
			return err
		}

		if err := tx.Where("item_id = ?", itemID).Delete(&modelsFestival.TicketType{}).Error; err != nil {
			return err
		}

		if err := tx.Delete(&modelsFestival.Item{}, itemID).Error; err != nil {
			return err
		}

		return nil
	}))
}

func (r *itemRepo) CreateTransportPackageAddon(transportAddon *modelsFestival.TransportAddon) error {
	return r.db.Create(transportAddon).Error
}

func (r *itemRepo) CreateCampPackageAddon(campAddon *modelsFestival.CampAddon) error {
	return r.db.Create(campAddon).Error
}

func (r *itemRepo) CreateCampEquipment(campEquipment *modelsFestival.CampEquipment) error {
	return r.db.Create(campEquipment).Error
}

func (r *itemRepo) CreatePackageAddonImage(packageAddonImage *modelsFestival.PackageAddonImage) error {
	return r.db.Create(packageAddonImage).Error
}
