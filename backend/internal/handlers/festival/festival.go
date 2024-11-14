package handlers

import (
	dtoFestival "backend/internal/dto/festival"
	"backend/internal/models"
	modelsCommon "backend/internal/models/common"
	modelsFestival "backend/internal/models/festival"
	servicesFestival "backend/internal/services/festival"
	"backend/internal/utils"
	"log"
	"net/http"
)

type FestivalHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByOrganizer(w http.ResponseWriter, r *http.Request)
}

type festivalHandler struct {
	festivalService servicesFestival.FestivalService
}

func NewFestivalHandler(festivalService servicesFestival.FestivalService) FestivalHandler {
	return &festivalHandler{
		festivalService: festivalService,
	}
}

func (h *festivalHandler) Create(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthOrganizer(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var input dtoFestival.CreateFestivalRequest
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

	festival := modelsFestival.Festival{
		Name:        input.Name,
		Description: input.Description,
		StartDate:   utils.ParseDate(input.StartDate),
		EndDate:     utils.ParseDate(input.EndDate),
		Capacity:    input.Capacity,
	}

	address := modelsCommon.Address{
		Street:         input.Address.Street,
		Number:         input.Address.Number,
		ApartmentSuite: &input.Address.ApartmentSuite,
		City: modelsCommon.City{
			Name:       input.Address.City,
			PostalCode: input.Address.PostalCode,
			Country: modelsCommon.Country{
				ISO3: input.Address.CountryISO3,
			},
		},
	}

	if err := h.festivalService.Create(&festival, utils.GetUsername(r.Context()), &address); err != nil {
		log.Println("error:", err)
		switch err {
		case models.ErrNotFound:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		case models.ErrCountryNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		case models.ErrUserHasAddress:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"message": "festival created successfully"}, nil)
	log.Println("festival created successfully:", festival.Name, "by", utils.GetUsername(r.Context()))
}

func (h *festivalHandler) GetByOrganizer(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthOrganizer(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	festivals, err := h.festivalService.GetByOrganizer(utils.GetUsername(r.Context()))
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"festivals": festivals}, nil)
	log.Println("festivals retrieved successfully for", utils.GetUsername(r.Context()))
}
