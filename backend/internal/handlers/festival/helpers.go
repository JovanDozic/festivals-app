package handlers

import (
	dtoCommon "backend/internal/dto/common"
	dtoFestival "backend/internal/dto/festival"
	"backend/internal/models"
	modelsCommon "backend/internal/models/common"
	modelsFestival "backend/internal/models/festival"
	"backend/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getFestivalIDFromRequest(r *http.Request) (uint, error) {
	vars := mux.Vars(r)
	festivalIdString := vars["festivalId"]

	if festivalIdString == "" {
		return 0, models.ErrBadRequest
	}

	festivalId, err := strconv.ParseUint(festivalIdString, 10, 32)
	if err != nil {
		return 0, models.ErrBadRequest
	}

	return uint(festivalId), nil
}

func (h *festivalHandler) authorizeOrganizerForFestival(w http.ResponseWriter, r *http.Request) (uint, bool) {
	if !utils.AuthOrganizerRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return 0, false
	}

	festivalId, err := getFestivalIDFromRequest(r)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return 0, false
	}

	isOrganizer, err := h.festivalService.IsOrganizer(utils.GetUsername(r.Context()), festivalId)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 0, false
	} else if !isOrganizer {
		log.Printf("error: organizer %s is not authorized for festival ID: %d", utils.GetUsername(r.Context()), festivalId)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return 0, false
	}

	return festivalId, true
}

func mapFestivalToResponse(festival modelsFestival.Festival, images []modelsCommon.Image) dtoFestival.FestivalResponse {

	var address *dtoCommon.GetAddressResponse
	if festival.Address != nil {
		address = &dtoCommon.GetAddressResponse{
			Street:         festival.Address.Street,
			Number:         festival.Address.Number,
			ApartmentSuite: festival.Address.ApartmentSuite,
			City:           festival.Address.City.Name,
			PostalCode:     festival.Address.City.PostalCode,
			Country:        festival.Address.City.Country.Name,
		}
	} else {
		address = nil
	}

	imageResponses := make([]dtoCommon.GetImageResponse, len(images))
	for i, image := range images {
		imageResponses[i] = dtoCommon.GetImageResponse{
			Url: image.URL,
		}
	}

	return dtoFestival.FestivalResponse{
		ID:          festival.ID,
		Name:        festival.Name,
		Description: festival.Description,
		StartDate:   festival.StartDate,
		EndDate:     festival.EndDate,
		Capacity:    festival.Capacity,
		Status:      festival.Status,
		StoreStatus: festival.StoreStatus,
		Address:     address,
		Images:      imageResponses,
	}
}