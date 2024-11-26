package handlers

import (
	dtoFestival "backend/internal/dto/festival"
	modelsFestival "backend/internal/models/festival"
	servicesFestival "backend/internal/services/festival"
	"backend/internal/utils"
	"log"
	"net/http"
)

type ItemHandler interface {
	CreateItem(w http.ResponseWriter, r *http.Request)
	CreatePriceListItem(w http.ResponseWriter, r *http.Request)
	GetCurrentTicketTypes(w http.ResponseWriter, r *http.Request)
}

type itemHandler struct {
	itemService     servicesFestival.ItemService
	festivalService servicesFestival.FestivalService
}

func NewItemHandler(
	itemService servicesFestival.ItemService,
	festivalService servicesFestival.FestivalService,
) ItemHandler {
	return &itemHandler{
		itemService:     itemService,
		festivalService: festivalService,
	}
}

func (h *itemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {

	festivalId, ok := h.authorizeOrganizerForFestival(w, r)
	if !ok {
		return
	}

	var input dtoFestival.CreateItemRequest
	if err := utils.ReadJSON(w, r, &input); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	item := modelsFestival.Item{
		Name:            input.Name,
		Description:     input.Description,
		FestivalID:      festivalId,
		Type:            input.Type,
		AvailableNumber: input.AvailableNumber,
		RemainingNumber: input.AvailableNumber,
	}

	if err := h.itemService.CreateItem(&item); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"itemId": item.ID}, nil)
	log.Println("item created:", input.Name)
}

func (h *itemHandler) CreatePriceListItem(w http.ResponseWriter, r *http.Request) {

	festivalId, ok := h.authorizeOrganizerForFestival(w, r)
	if !ok {
		return
	}

	var input dtoFestival.CreatePriceListItemRequest
	if err := utils.ReadJSON(w, r, &input); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	priceListItem := modelsFestival.PriceListItem{
		ItemID:   input.ItemId,
		Price:    input.Price,
		IsFixed:  input.IsFixed,
		DateFrom: utils.ParseDateNil(input.DateFrom),
		DateTo:   utils.ParseDateNil(input.DateTo),
	}

	if err := h.itemService.CreatePriceListItem(festivalId, input.ItemId, &priceListItem); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"priceListItemId": priceListItem.ID}, nil)
	log.Println("price list item created:", priceListItem.ID)
}

func (h *itemHandler) GetCurrentTicketTypes(w http.ResponseWriter, r *http.Request) {

	festivalId, ok := h.authorizeOrganizerForFestival(w, r)
	if !ok {
		return
	}

	priceListItems, err := h.itemService.GetCurrentTicketTypes(festivalId)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var currentTicketTypes []dtoFestival.GetCurrentTicketTypesResponse
	for _, priceListItem := range priceListItems {
		currentTicketTypes = append(currentTicketTypes, dtoFestival.GetCurrentTicketTypesResponse{
			ItemId:          priceListItem.ItemID,
			PriceListItemId: priceListItem.ID,
			Name:            priceListItem.Item.Name,
			Description:     priceListItem.Item.Description,
			AvailableNumber: priceListItem.Item.AvailableNumber,
			RemainingNumber: priceListItem.Item.RemainingNumber,
			Price:           priceListItem.Price,
			IsFixed:         priceListItem.IsFixed,
			DateFrom:        priceListItem.DateFrom,
			DateTo:          priceListItem.DateTo,
		})
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"ticketTypes": currentTicketTypes}, nil)
	log.Println("current ticket types retrieved")
}
