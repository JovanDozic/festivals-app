package services

import (
	"backend/internal/config"
	dto "backend/internal/dto/festival"
	modelsFestival "backend/internal/models/festival"
	reposFestival "backend/internal/repositories/festival"
	"strings"
)

type ItemService interface {
	CreateItem(item *modelsFestival.Item) error
	CreatePriceListItem(festivalId, itemId uint, priceListItem *modelsFestival.PriceListItem) error
	GetCurrentTicketTypes(festivalId uint) ([]modelsFestival.PriceListItem, error)
	GetTicketTypesCount(festivalId uint) (int, error)
	GetTicketTypes(itemId uint) (*dto.GetItemResponse, error)
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
