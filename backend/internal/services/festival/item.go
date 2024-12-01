package services

import (
	"backend/internal/config"
	dto "backend/internal/dto/festival"
	modelsFestival "backend/internal/models/festival"
	reposFestival "backend/internal/repositories/festival"
	"backend/internal/utils"
	"errors"
	"log"
	"strings"
)

type ItemService interface {
	CreateItem(item *modelsFestival.Item) error
	UpdatePackageAddonCategory(packageAddon *modelsFestival.PackageAddon) error
	CreatePriceListItem(festivalId, itemId uint, priceListItem *modelsFestival.PriceListItem) error
	GetCurrentTicketTypes(festivalId uint) ([]modelsFestival.PriceListItem, error)
	GetTicketTypesCount(festivalId uint) (int, error)
	GetTicketTypes(itemId uint) (*dto.GetItemResponse, error)
	UpdateItemAndPrices(request dto.UpdateItemRequest) error
	DeleteTicketType(itemId uint) error
	GetCurrentPackageAddons(festivalId uint, category string) ([]modelsFestival.PriceListItem, error)
}

type itemService struct {
	config   *config.Config
	itemRepo reposFestival.ItemRepo
}

func NewItemService(
	config *config.Config,
	itemRepo reposFestival.ItemRepo,
) ItemService {
	return &itemService{
		config:   config,
		itemRepo: itemRepo,
	}
}

func (s *itemService) CreateItem(item *modelsFestival.Item) error {
	err := s.itemRepo.CreateItem(item)
	if err != nil {
		return err
	}

	return nil
}

func (s *itemService) UpdatePackageAddonCategory(packageAddon *modelsFestival.PackageAddon) error {
	err := s.itemRepo.UpdatePackageAddonCategory(packageAddon)
	if err != nil {
		return err
	}
	return nil
}

func (s *itemService) CreatePriceListItem(festivalId, itemId uint, priceListItem *modelsFestival.PriceListItem) error {
	priceList, err := s.itemRepo.GetPriceList(festivalId)
	if err != nil && !strings.Contains(err.Error(), "record not found") {
		return err
	}

	if priceList == nil {
		priceList = &modelsFestival.PriceList{
			FestivalID: festivalId,
		}
		err = s.itemRepo.CreatePriceList(priceList)
		if err != nil {
			return err
		}
	}

	priceListItem.PriceListID = priceList.ID

	err = s.itemRepo.CreatePriceListItem(priceListItem)
	if err != nil {
		return err
	}

	return nil
}

func (s *itemService) GetCurrentTicketTypes(festivalId uint) ([]modelsFestival.PriceListItem, error) {
	return s.itemRepo.GetCurrentTicketTypes(festivalId)
}

func (s *itemService) GetCurrentPackageAddons(festivalId uint, category string) ([]modelsFestival.PriceListItem, error) {

	if category == "" ||
		(strings.ToUpper(category) != modelsFestival.PackageAddonGeneral &&
			strings.ToUpper(category) != modelsFestival.PackageAddonCamp &&
			strings.ToUpper(category) != modelsFestival.PackageAddonTransport) {
		return nil, errors.New("invalid category")
	}

	return s.itemRepo.GetCurrentPackageAddons(festivalId, strings.ToUpper(category))
}

func (s *itemService) GetTicketTypesCount(festivalId uint) (int, error) {
	return s.itemRepo.GetTicketTypesCount(festivalId)
}

func (s *itemService) GetTicketTypes(itemId uint) (*dto.GetItemResponse, error) {

	// we need to do the mapping here as repositories needs to be called multiple times and we do not have a model that supports this type of the response, only DTO does

	item, priceIds, err := s.itemRepo.GetItemAndPriceListItemsIDs(itemId)
	if err != nil {
		return nil, err
	}

	priceListItems, err := s.itemRepo.GetPriceListItemsByIDs(priceIds)
	if err != nil {
		return nil, err
	}

	response := dto.GetItemResponse{
		Id:              item.ID,
		Name:            item.Name,
		Description:     item.Description,
		Type:            item.Type,
		AvailableNumber: item.AvailableNumber,
		RemainingNumber: item.RemainingNumber,
		PriceListItems:  make([]dto.PriceListItemResponse, len(priceListItems)),
	}

	for i, priceListItem := range priceListItems {
		response.PriceListItems[i] = dto.PriceListItemResponse{
			Id:       priceListItem.ID,
			Price:    priceListItem.Price,
			IsFixed:  priceListItem.IsFixed,
			DateFrom: priceListItem.DateFrom,
			DateTo:   priceListItem.DateTo,
		}
	}

	return &response, nil
}

func (s *itemService) UpdateItemAndPrices(request dto.UpdateItemRequest) error {

	itemId := request.Id
	var priceListItemIds []uint
	for _, priceListItem := range request.PriceListItems {
		priceListItemIds = append(priceListItemIds, priceListItem.Id)
	}

	itemDb, priceIdsDb, err := s.itemRepo.GetItemAndPriceListItemsIDs(itemId)
	if err != nil {
		return err
	}

	priceListItemsDb, err := s.itemRepo.GetPriceListItemsByIDs(priceIdsDb)
	if err != nil {
		return err
	}

	log.Printf("itemDb: %+v", itemDb)
	log.Printf("priceListItemIds: %+v", priceListItemIds)
	log.Printf("priceIdsDb: %+v", priceIdsDb)

	// update item
	itemDb.Name = request.Name
	itemDb.Description = request.Description
	itemDb.AvailableNumber = request.AvailableNumber

	err = s.itemRepo.UpdateItem(itemDb)
	if err != nil {
		return err
	}

	// update prices
	for _, priceListItem := range priceListItemsDb {
		for _, priceListItemRequest := range request.PriceListItems {
			if priceListItem.ID == priceListItemRequest.Id {
				priceListItem.Price = priceListItemRequest.Price
				priceListItem.IsFixed = priceListItemRequest.IsFixed
				priceListItem.DateFrom = utils.ParseDateNil(priceListItemRequest.DateFrom)
				priceListItem.DateTo = utils.ParseDateNil(priceListItemRequest.DateTo)

				err = s.itemRepo.UpdatePriceListItem(&priceListItem)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (s *itemService) DeleteTicketType(itemId uint) error {
	return s.itemRepo.DeleteTicketType(itemId)
}
