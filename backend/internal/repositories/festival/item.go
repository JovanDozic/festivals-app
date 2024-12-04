package repositories

import (
	dtoFestival "backend/internal/dto/festival"
	"backend/internal/models"
	modelsCommon "backend/internal/models/common"
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
	GetTransportAddons(festivalId uint) ([]dtoFestival.TransportAddonDTO, error)
	GetTransportAddon(itemId uint) (*dtoFestival.TransportAddonDTO, error)
	GetGeneralAddons(festivalId uint) ([]dtoFestival.GeneralAddonDTO, error)
	GetCampAddons(festivalId uint) ([]dtoFestival.CampAddonDTO, error)
	GetCampEquipment(itemId uint) ([]modelsFestival.CampEquipment, error)
	GetAvailableDepartureCountries(festivalId uint) ([]modelsCommon.Country, error)
	GetAddonsFromPackage(festivalPackageId uint) ([]modelsFestival.PackageAddon, error)
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
		Order("price").
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

func (r *itemRepo) GetTransportAddons(festivalId uint) ([]dtoFestival.TransportAddonDTO, error) {
	var transportAddons []dtoFestival.TransportAddonDTO

	err := r.db.
		Table("price_list_items pli").
		Select(`
			pli.id as price_list_item_id,
			pli.price_list_id,
			pli.item_id,
			i.name as item_name,
			i.description as item_description,
			i.type as item_type,
			i.available_number as item_available_number,
			i.remaining_number as item_remaining_number,
			pli.date_from as date_from,
			pli.date_to as date_to,
			pli.is_fixed as is_fixed,
			pli.price as price,
			pa.category as package_addon_category,
			ta.transport_type,
			ta.departure_time,
			ta.arrival_time,
			ta.return_departure_time,
			ta.return_arrival_time,
			cd.id as departure_city_id,
			cd.name as departure_city_name,
			cd.postal_code as departure_postal_code,
			ccd.iso3 as departure_country_iso3,
			ccd.iso as departure_country_iso,
			ccd.nice_name as departure_country_nice_name,
			ca.id as arrival_city_id,
			ca.name as arrival_city_name,
			ca.postal_code as arrival_postal_code,
			cca.iso3 as arrival_country_iso3,
			cca.iso as arrival_country_iso,
			cca.nice_name as arrival_country_nice_name
		`).
		Joins("JOIN items i ON pli.item_id = i.id").
		Joins("JOIN package_addons pa ON i.id = pa.item_id").
		Joins("JOIN transport_addons ta ON pa.item_id = ta.item_id").
		Joins("JOIN cities cd ON ta.departure_city_id = cd.id").
		Joins("JOIN countries ccd ON cd.country_id = ccd.id").
		Joins("JOIN cities ca ON ta.arrival_city_id = ca.id").
		Joins("JOIN countries cca ON ca.country_id = cca.id").
		Where("i.festival_id = ?", festivalId).
		Scan(&transportAddons).Error

	if err != nil {
		return nil, err
	}

	return transportAddons, nil
}

func (r *itemRepo) GetTransportAddon(itemId uint) (*dtoFestival.TransportAddonDTO, error) {
	var transportAddon dtoFestival.TransportAddonDTO

	err := r.db.
		Table("price_list_items pli").
		Select(`
			pli.id as price_list_item_id,
			pli.price_list_id,
			pli.item_id,
			i.name as item_name,
			i.description as item_description,
			i.type as item_type,
			i.available_number as item_available_number,
			i.remaining_number as item_remaining_number,
			pli.date_from as date_from,
			pli.date_to as date_to,
			pli.is_fixed as is_fixed,
			pli.price as price,
			pa.category as package_addon_category,
			ta.transport_type,
			ta.departure_time,
			ta.arrival_time,
			ta.return_departure_time,
			ta.return_arrival_time,
			cd.id as departure_city_id,
			cd.name as departure_city_name,
			cd.postal_code as departure_postal_code,
			ccd.iso3 as departure_country_iso3,
			ccd.iso as departure_country_iso,
			ccd.nice_name as departure_country_nice_name,
			ca.id as arrival_city_id,
			ca.name as arrival_city_name,
			ca.postal_code as arrival_postal_code,
			cca.iso3 as arrival_country_iso3,
			cca.iso as arrival_country_iso,
			cca.nice_name as arrival_country_nice_name
		`).
		Joins("JOIN items i ON pli.item_id = i.id").
		Joins("JOIN package_addons pa ON i.id = pa.item_id").
		Joins("JOIN transport_addons ta ON pa.item_id = ta.item_id").
		Joins("JOIN cities cd ON ta.departure_city_id = cd.id").
		Joins("JOIN countries ccd ON cd.country_id = ccd.id").
		Joins("JOIN cities ca ON ta.arrival_city_id = ca.id").
		Joins("JOIN countries cca ON ca.country_id = cca.id").
		Where("pli.item_id = ?", itemId).
		Scan(&transportAddon).Error
	if err != nil {
		return nil, err
	}

	return &transportAddon, nil
}

func (r *itemRepo) GetGeneralAddons(festivalId uint) ([]dtoFestival.GeneralAddonDTO, error) {
	var generalAddons []dtoFestival.GeneralAddonDTO

	err := r.db.
		Table("price_list_items pli").
		Select(`
			pli.id as price_list_item_id,
			pli.price_list_id,
			pli.item_id,
			i.name as item_name,
			i.description as item_description,
			i.type as item_type,
			i.available_number as item_available_number,
			i.remaining_number as item_remaining_number,
			pli.date_from as date_from,
			pli.date_to as date_to,
			pli.is_fixed as is_fixed,
			pli.price as price,
			pa.category as package_addon_category
		`).
		Joins("JOIN items i ON pli.item_id = i.id").
		Joins("JOIN package_addons pa ON i.id = pa.item_id").
		Where("i.festival_id = ?", festivalId).
		Where("pa.category = ?", modelsFestival.PackageAddonGeneral).
		Scan(&generalAddons).Error

	if err != nil {
		return nil, err
	}

	return generalAddons, nil
}

func (r *itemRepo) GetCampAddons(festivalId uint) ([]dtoFestival.CampAddonDTO, error) {
	var campAddons []dtoFestival.CampAddonDTO

	err := r.db.
		Table("price_list_items pli").
		Select(`
			pli.id as price_list_item_id,
			pli.price_list_id,
			pli.item_id,
			i.name as item_name,
			i.description as item_description,
			i.type as item_type,
			i.available_number as item_available_number,
			i.remaining_number as item_remaining_number,
			pli.date_from as date_from,
			pli.date_to as date_to,
			pli.is_fixed as is_fixed,
			pli.price as price,
			pa.category as package_addon_category,
			ca.camp_name as camp_name,
			im.url as image_url,
			STRING_AGG(ce.name, ', ') as equipment_names
		`).
		Joins("JOIN items i ON pli.item_id = i.id").
		Joins("JOIN package_addons pa ON i.id = pa.item_id").
		Joins("JOIN camp_addons ca ON pa.item_id = ca.item_id").
		Joins("JOIN package_addon_images pai ON pai.item_id = i.id").
		Joins("JOIN images im ON pai.image_id = im.id").
		Joins("LEFT JOIN camp_equipments ce ON ce.item_id = i.id").
		Where("i.festival_id = ?", festivalId).
		Group("pli.id, pli.price_list_id, pli.item_id, i.name, i.description, i.type, i.available_number, i.remaining_number, pli.date_from, pli.date_to, pli.is_fixed, pli.price, pa.category, ca.camp_name, im.url").
		Scan(&campAddons).Error

	if err != nil {
		return nil, err
	}

	return campAddons, nil
}

func (r *itemRepo) GetCampEquipment(itemId uint) ([]modelsFestival.CampEquipment, error) {
	var campEquipment []modelsFestival.CampEquipment

	err := r.db.
		Where("item_id = ?", itemId).
		Find(&campEquipment).Error

	if err != nil {
		return nil, err
	}

	return campEquipment, nil
}

func (r *itemRepo) GetAvailableDepartureCountries(festivalId uint) ([]modelsCommon.Country, error) {
	var countries []modelsCommon.Country

	err := r.db.
		Distinct("ccd.*").
		Table("countries ccd").
		Joins("JOIN cities cd ON ccd.id = cd.country_id").
		Joins("JOIN transport_addons ta ON ta.departure_city_id = cd.id").
		Joins("JOIN items i ON ta.item_id = i.id").
		Where("i.festival_id = ?", festivalId).
		Scan(&countries).Error

	if err != nil {
		return nil, err
	}

	return countries, nil
}

func (r *itemRepo) GetAddonsFromPackage(festivalPackageId uint) ([]modelsFestival.PackageAddon, error) {
	var packageAddons []modelsFestival.PackageAddon

	err := r.db.
		Table("package_addons as pa").
		Joins("JOIN festival_package_addons as fpa ON pa.item_id = fpa.item_id").
		Where("fpa.festival_package_id = ?", festivalPackageId).
		Find(&packageAddons).Error

	if err != nil {
		return nil, err
	}

	return packageAddons, nil
}
