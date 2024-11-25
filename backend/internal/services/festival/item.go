package services

import (
	"backend/internal/config"
	modelsFestival "backend/internal/models/festival"
	reposFestival "backend/internal/repositories/festival"
)

type ItemService interface {
	CreateItem(item *modelsFestival.Item) error
	CreatePriceListItem(festivalId, itemId uint, priceListItem *modelsFestival.PriceListItem) error
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
	if err != nil {
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
