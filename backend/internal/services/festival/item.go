package services

import (
	"backend/internal/config"
	modelsFestival "backend/internal/models/festival"
	reposFestival "backend/internal/repositories/festival"
	"strings"
)

type ItemService interface {
	CreateItem(item *modelsFestival.Item) error
	CreatePriceListItem(festivalId, itemId uint, priceListItem *modelsFestival.PriceListItem) error
	GetCurrentTicketTypes(festivalId uint) ([]modelsFestival.PriceListItem, error)
	GetTicketTypesCount(festivalId uint) (int, error)
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
