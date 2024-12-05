package handlers

import (
	dtoCommon "backend/internal/dto/common"
	dtoFestival "backend/internal/dto/festival"
	"backend/internal/models"
	modelsCommon "backend/internal/models/common"
	modelsFestival "backend/internal/models/festival"
	services "backend/internal/services/festival"
	"backend/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetIDParamFromRequest(r *http.Request, paramName string) (uint, error) {
	vars := mux.Vars(r)
	idString := vars[paramName]

	if idString == "" {
		return 0, models.ErrBadRequest
	}

	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		return 0, models.ErrBadRequest
	}

	return uint(id), nil
}

func GetParamFromRequest(r *http.Request, paramName string) (string, error) {
	vars := mux.Vars(r)
	paramString := vars[paramName]

	if paramString == "" {
		return "", models.ErrBadRequest
	}

	return paramString, nil
}

func AuthOrganizerForFestival(w http.ResponseWriter, r *http.Request, fs *services.FestivalService) (uint, bool) {
	if !utils.AuthOrganizerRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return 0, false
	}

	festivalId, err := GetIDParamFromRequest(r, "festivalId")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return 0, false
	}

	isOrganizer, err := (*fs).IsOrganizer(utils.GetUsername(r.Context()), festivalId)
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

func AuthOrganizerOrEmployeeForFestival(w http.ResponseWriter, r *http.Request, fs *services.FestivalService) (uint, bool) {

	isOrganizer := utils.AuthOrganizerRole(r.Context())
	isEmployee := utils.AuthEmployeeRole(r.Context())

	if !isOrganizer && !isEmployee {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return 0, false
	}

	festivalId, err := GetIDParamFromRequest(r, "festivalId")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return 0, false
	}

	var ok bool
	var roleName string
	if isOrganizer {
		roleName = "organizer"
		ok, err = (*fs).IsOrganizer(utils.GetUsername(r.Context()), festivalId)
	} else if isEmployee {
		roleName = "employee"
		ok, err = (*fs).IsEmployee(utils.GetUsername(r.Context()), festivalId)
	}

	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 0, false
	} else if !ok {

		log.Printf("error: %s (%s) is not authorized for festival ID: %d", utils.GetUsername(r.Context()), roleName, festivalId)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return 0, false
	}

	return festivalId, true
}

func MapFestivalToResponse(festival modelsFestival.Festival, images []modelsCommon.Image) dtoFestival.FestivalResponse {

	var address *dtoCommon.GetAddressResponse
	if festival.Address != nil {
		address = &dtoCommon.GetAddressResponse{
			Street:         festival.Address.Street,
			Number:         festival.Address.Number,
			ApartmentSuite: festival.Address.ApartmentSuite,
			City:           festival.Address.City.Name,
			PostalCode:     festival.Address.City.PostalCode,
			Country:        festival.Address.City.Country.NiceName,
			CountryISO3:    festival.Address.City.Country.ISO3,
			CountryISO2:    festival.Address.City.Country.ISO,
		}
	} else {
		address = nil
	}

	imageResponses := make([]dtoCommon.GetImageResponse, len(images))
	for i, image := range images {
		imageResponses[i] = dtoCommon.GetImageResponse{
			ID:  image.ID,
			URL: image.URL,
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
