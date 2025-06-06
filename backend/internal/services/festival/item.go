package festival

import (
	"backend/internal/config"
	dto "backend/internal/dto/festival"
	modelsCommon "backend/internal/models/common"
	modelsFestival "backend/internal/models/festival"
	"backend/internal/repos/common"
	repos "backend/internal/repos/festival"
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
	GetPackageAddonsCount(festivalId uint, category string) (int, error)
	GetAllPackageAddonsCount(festivalId uint) (int, error)
	GetTicketTypes(itemId uint) (*dto.GetItemResponse, error)
	UpdateItemAndPrices(request dto.UpdateItemRequest) error
	DeleteTicketType(itemId uint) error
	GetCurrentPackageAddons(festivalId uint, category string) ([]modelsFestival.PriceListItem, error)
	CreateTransportPackageAddon(request dto.CreateTransportPackageAddonRequest) error
	CreateCampPackageAddon(request dto.CreateCampPackageAddonRequest) error
	GetTransportAddons(festivalId uint) ([]dto.TransportAddonDTO, error)
	GetGeneralAddons(festivalId uint) ([]dto.GeneralAddonDTO, error)
	GetCampAddons(festivalId uint) ([]dto.CampAddonDTO, error)
	GetAvailableDepartureCountries(festivalId uint) ([]modelsCommon.Country, error)
}

type itemService struct {
	config       *config.Config
	itemRepo     repos.ItemRepo
	locationRepo common.LocationRepo
	imageRepo    common.ImageRepo
}

func NewItemService(
	cfg *config.Config,
	ir repos.ItemRepo,
	lr common.LocationRepo,
	img common.ImageRepo,
) ItemService {
	return &itemService{
		config:       cfg,
		itemRepo:     ir,
		locationRepo: lr,
		imageRepo:    img,
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

	if priceList == nil || priceList.ID == 0 {
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

func (s *itemService) GetPackageAddonsCount(festivalId uint, category string) (int, error) {
	if category == "" ||
		(strings.ToUpper(category) != modelsFestival.PackageAddonGeneral &&
			strings.ToUpper(category) != modelsFestival.PackageAddonCamp &&
			strings.ToUpper(category) != modelsFestival.PackageAddonTransport) {
		return 0, errors.New("invalid category")
	}

	return s.itemRepo.GetPackageAddonsCount(festivalId, strings.ToUpper(category))
}

func (s *itemService) GetAllPackageAddonsCount(festivalId uint) (int, error) {
	sum := 0
	for _, category := range []string{modelsFestival.PackageAddonGeneral, modelsFestival.PackageAddonCamp, modelsFestival.PackageAddonTransport} {
		count, err := s.itemRepo.GetPackageAddonsCount(festivalId, strings.ToUpper(category))
		if err != nil {
			return -1, err
		}
		sum += count
	}

	return sum, nil
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

	itemId := request.ID

	itemDb, priceIdsDb, err := s.itemRepo.GetItemAndPriceListItemsIDs(itemId)
	if err != nil {
		return err
	}

	priceListItemsDb, err := s.itemRepo.GetPriceListItemsByIDs(priceIdsDb)
	if err != nil {
		return err
	}

	// check if available number changed, if so, update remaining number
	if itemDb.AvailableNumber != request.AvailableNumber {
		itemDb.RemainingNumber += request.AvailableNumber - itemDb.AvailableNumber
	}

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
			if priceListItem.ID == priceListItemRequest.ID {
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

func (s *itemService) CreateTransportPackageAddon(request dto.CreateTransportPackageAddonRequest) error {

	departureCity := &modelsCommon.City{
		Name:       request.DepartureCity.Name,
		PostalCode: request.DepartureCity.PostalCode,
	}

	departureCountry := &modelsCommon.Country{
		ISO3: request.DepartureCity.CountryISO3,
	}

	arrivalCity := &modelsCommon.City{
		Name:       request.ArrivalCity.Name,
		PostalCode: request.ArrivalCity.PostalCode,
	}

	arrivalCountry := &modelsCommon.Country{
		ISO3: request.ArrivalCity.CountryISO3,
	}

	err := s.locationRepo.FillCityAndCountryModels(departureCity, departureCountry)
	if err != nil {
		log.Println("error getting city and country:", err)
		return err
	}

	err = s.locationRepo.FillCityAndCountryModels(arrivalCity, arrivalCountry)
	if err != nil {
		log.Println("error getting city and country:", err)
		return err
	}

	transportAddon := &modelsFestival.TransportAddon{
		ItemID:              request.ItemID,
		DepartureCityID:     departureCity.ID,
		ArrivalCityID:       arrivalCity.ID,
		DepartureTime:       utils.ParseDateTime(request.DepartureTime),
		ArrivalTime:         utils.ParseDateTime(request.ArrivalTime),
		ReturnDepartureTime: utils.ParseDateTime(request.ReturnDepartureTime),
		ReturnArrivalTime:   utils.ParseDateTime(request.ReturnArrivalTime),
		TransportType:       request.TransportType,
	}

	err = s.itemRepo.CreateTransportPackageAddon(transportAddon)
	if err != nil {
		log.Println("error creating transport package addon:", err)
		return err
	}

	return nil
}

func (s *itemService) CreateCampPackageAddon(request dto.CreateCampPackageAddonRequest) error {

	// we need to create camp package addon
	campAddon := &modelsFestival.CampAddon{
		ItemID:   request.ItemID,
		CampName: request.CampName,
	}

	err := s.itemRepo.CreateCampPackageAddon(campAddon)
	if err != nil {
		log.Println("error creating camp package addon:", err)
		return err
	}

	// then we need to add equipment list items
	for _, equipment := range request.EquipmentList {
		campEquipment := &modelsFestival.CampEquipment{
			ItemID: campAddon.ItemID,
			Name:   equipment.Name,
		}

		err := s.itemRepo.CreateCampEquipment(campEquipment)
		if err != nil {
			log.Println("error creating camp equipment:", err)
			continue
		}
	}

	// then we need to save image
	image := &modelsCommon.Image{
		URL: request.ImageURL,
	}

	err = s.imageRepo.Create(image)
	if err != nil {
		log.Println("error creating image:", err)
		return err
	}

	packageAddonImage := &modelsFestival.PackageAddonImage{
		ItemID:  campAddon.ItemID,
		ImageID: image.ID,
	}

	err = s.itemRepo.CreatePackageAddonImage(packageAddonImage)
	if err != nil {
		log.Println("error creating package addon image:", err)
		return err
	}

	return nil
}

func (s *itemService) GetTransportAddons(festivalId uint) ([]dto.TransportAddonDTO, error) {
	transportAddons, err := s.itemRepo.GetTransportAddons(festivalId)
	if err != nil {
		return nil, err
	}
	return transportAddons, nil
}

func (s *itemService) GetGeneralAddons(festivalId uint) ([]dto.GeneralAddonDTO, error) {
	campAddons, err := s.itemRepo.GetGeneralAddons(festivalId)
	if err != nil {
		return nil, err
	}
	return campAddons, nil
}

func (s *itemService) GetCampAddons(festivalId uint) ([]dto.CampAddonDTO, error) {
	campAddons, err := s.itemRepo.GetCampAddons(festivalId)
	if err != nil {
		return nil, err
	}

	return campAddons, nil
}

func (s *itemService) GetAvailableDepartureCountries(festivalId uint) ([]modelsCommon.Country, error) {
	return s.itemRepo.GetAvailableDepartureCountries(festivalId)
}
