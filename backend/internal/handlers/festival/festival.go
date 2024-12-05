package handlers

import (
	dtoFestival "backend/internal/dto/festival"
	"backend/internal/models"
	modelsCommon "backend/internal/models/common"
	modelsFestival "backend/internal/models/festival"
	servicesCommon "backend/internal/services/common"
	servicesFestival "backend/internal/services/festival"
	"backend/internal/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type FestivalHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByOrganizer(w http.ResponseWriter, r *http.Request)
	GetByEmployee(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	PublishFestival(w http.ResponseWriter, r *http.Request)
	CancelFestival(w http.ResponseWriter, r *http.Request)
	CompleteFestival(w http.ResponseWriter, r *http.Request)
	OpenStore(w http.ResponseWriter, r *http.Request)
	CloseStore(w http.ResponseWriter, r *http.Request)
	AddImage(w http.ResponseWriter, r *http.Request)
	RemoveImage(w http.ResponseWriter, r *http.Request)
	GetImages(w http.ResponseWriter, r *http.Request)
	Employ(w http.ResponseWriter, r *http.Request)
	Fire(w http.ResponseWriter, r *http.Request)
	GetEmployeeCount(w http.ResponseWriter, r *http.Request)
}

type festivalHandler struct {
	festivalService servicesFestival.FestivalService
	locationService servicesCommon.LocationService
}

func NewFestivalHandler(festivalService servicesFestival.FestivalService, locationService servicesCommon.LocationService) FestivalHandler {
	return &festivalHandler{
		festivalService: festivalService,
		locationService: locationService,
	}
}

func (h *festivalHandler) Create(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthOrganizerRole(r.Context()) {
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

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"message": "festival created successfully", "id": festival.ID}, nil)
	log.Println("festival created successfully:", festival.Name, "by", utils.GetUsername(r.Context()))
}

func (h *festivalHandler) GetByOrganizer(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthOrganizerRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	festivals, err := h.festivalService.GetByOrganizer(utils.GetUsername(r.Context()))
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	festivalsResponse := dtoFestival.FestivalsResponse{
		Festivals: make([]dtoFestival.FestivalResponse, len(festivals)),
	}

	for i, festival := range festivals {
		images, err := h.festivalService.GetImages(festival.ID)
		if err != nil {
			log.Println("error:", err)
			continue
		}
		festivalsResponse.Festivals[i] = MapFestivalToResponse(festival, images)
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"festivals": festivalsResponse.Festivals}, nil)
	log.Println("festivals retrieved successfully for organizer", utils.GetUsername(r.Context()))
}

func (h *festivalHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthAttendeeRole(r.Context()) {
		log.Println("error:", models.ErrUnauthorized)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	festivals, err := h.festivalService.GetAllPublic()
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	festivalsResponse := dtoFestival.FestivalsResponse{
		Festivals: make([]dtoFestival.FestivalResponse, len(festivals)),
	}

	for i, festival := range festivals {
		images, err := h.festivalService.GetImages(festival.ID)
		if err != nil {
			log.Println("error:", err)
			continue
		}
		festivalsResponse.Festivals[i] = MapFestivalToResponse(festival, images)
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"festivals": festivalsResponse.Festivals}, nil)
	log.Println("all festivals retrieved successfully")
}

func (h *festivalHandler) GetById(w http.ResponseWriter, r *http.Request) {

	if !utils.Auth(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	festivalId := vars["festivalId"]

	if festivalId == "" {
		log.Println("error:", models.ErrBadRequest)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	festival, err := h.festivalService.GetById(utils.ToUint(festivalId))
	if err != nil {
		log.Println("error:", err)
		switch err {
		case models.ErrNotFound:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	images, err := h.festivalService.GetImages(festival.ID)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	festivalResponse := MapFestivalToResponse(*festival, images)

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"festival": festivalResponse}, nil)
	log.Println("festival retrieved successfully:", festival.Name)
}

func (h *festivalHandler) Update(w http.ResponseWriter, r *http.Request) {

	festivalId, ok := AuthOrganizerForFestival(w, r, &h.festivalService)
	if !ok {
		return
	}

	var input dtoFestival.UpdateFestivalRequest
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

	if err := h.festivalService.Update(festivalId, &festival); err != nil {
		log.Println("error:", err)
		switch err {
		case models.ErrNotFound:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		case models.ErrCountryNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	addressId, err := h.festivalService.GetAddressID(festivalId)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
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

	if err := h.locationService.UpdateAddress(addressId, &address); err != nil {
		log.Println("error:", err)
		switch err {
		case models.ErrNotFound:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		case models.ErrCountryNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "festival updated successfully"}, nil)
	log.Println("festival updated successfully:", festival.Name, "by", utils.GetUsername(r.Context()))
}

func (h *festivalHandler) Delete(w http.ResponseWriter, r *http.Request) {

	festivalId, ok := AuthOrganizerForFestival(w, r, &h.festivalService)
	if !ok {
		return
	}

	if err := h.festivalService.Delete(festivalId); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "festival deleted successfully"}, nil)
	log.Println("festival deleted successfully:", festivalId)
}

func (h *festivalHandler) PublishFestival(w http.ResponseWriter, r *http.Request) {

	festivalId, ok := AuthOrganizerForFestival(w, r, &h.festivalService)
	if !ok {
		return
	}

	if err := h.festivalService.PublishFestival(festivalId); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "festival published successfully"}, nil)
	log.Println("festival published successfully:", festivalId)
}

func (h *festivalHandler) CancelFestival(w http.ResponseWriter, r *http.Request) {

	festivalId, ok := AuthOrganizerForFestival(w, r, &h.festivalService)
	if !ok {
		return
	}

	if err := h.festivalService.CancelFestival(festivalId); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "festival cancelled successfully"}, nil)
	log.Println("festival cancelled successfully:", festivalId)
}

func (h *festivalHandler) CompleteFestival(w http.ResponseWriter, r *http.Request) {

	festivalId, ok := AuthOrganizerForFestival(w, r, &h.festivalService)
	if !ok {
		return
	}

	if err := h.festivalService.CompleteFestival(festivalId); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "festival completed successfully"}, nil)
	log.Println("festival completed successfully:", festivalId)
}

func (h *festivalHandler) OpenStore(w http.ResponseWriter, r *http.Request) {

	festivalId, ok := AuthOrganizerForFestival(w, r, &h.festivalService)
	if !ok {
		return
	}

	if err := h.festivalService.OpenStore(festivalId); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "store opened successfully"}, nil)
	log.Println("store opened successfully for festival:", festivalId)
}

func (h *festivalHandler) CloseStore(w http.ResponseWriter, r *http.Request) {

	festivalId, ok := AuthOrganizerForFestival(w, r, &h.festivalService)
	if !ok {
		return
	}

	if err := h.festivalService.CloseStore(festivalId); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "store closed successfully"}, nil)
	log.Println("store closed successfully for festival:", festivalId)
}

func (h *festivalHandler) AddImage(w http.ResponseWriter, r *http.Request) {

	festivalId, ok := AuthOrganizerForFestival(w, r, &h.festivalService)
	if !ok {
		return
	}

	var input dtoFestival.AddImageRequest
	if err := utils.ReadJSON(w, r, &input); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.festivalService.AddImage(festivalId, &modelsCommon.Image{
		URL: input.ImageUrl,
	}); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"message": "image added successfully"}, nil)
	log.Println("image added successfully for festival:", festivalId)
}

func (h *festivalHandler) RemoveImage(w http.ResponseWriter, r *http.Request) {

	festivalId, ok := AuthOrganizerForFestival(w, r, &h.festivalService)
	if !ok {
		return
	}

	imageId, err := GetIDParamFromRequest(r, "imageId")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.festivalService.RemoveImage(festivalId, imageId); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"message": "image removed successfully"}, nil)
	log.Println("image removed successfully for festival:", festivalId)
}

func (h *festivalHandler) GetImages(w http.ResponseWriter, r *http.Request) {

	if !utils.Auth(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	festivalId, err := GetIDParamFromRequest(r, "festivalId")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	images, err := h.festivalService.GetImages(festivalId)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"images": images}, nil)
	log.Println("images retrieved successfully for festival:", festivalId)
}

func (h *festivalHandler) Employ(w http.ResponseWriter, r *http.Request) {

	festivalId, ok := AuthOrganizerForFestival(w, r, &h.festivalService)
	if !ok {
		return
	}

	employeeId, err := GetIDParamFromRequest(r, "employeeId")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.festivalService.Employ(festivalId, employeeId); err != nil {
		log.Println("error:", err)
		switch err {
		case models.ErrNotFound:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		case models.ErrEmployeeAlreadyEmployed:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"message": "employee employed successfully"}, nil)
	log.Println("employee employed successfully for festival:", festivalId)
}

func (h *festivalHandler) Fire(w http.ResponseWriter, r *http.Request) {

	festivalId, ok := AuthOrganizerForFestival(w, r, &h.festivalService)
	if !ok {
		return
	}

	employeeId, err := GetIDParamFromRequest(r, "employeeId")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.festivalService.Fire(festivalId, employeeId); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "employee fired successfully"}, nil)
	log.Println("employee fired successfully for festival:", festivalId)
}

func (h *festivalHandler) GetEmployeeCount(w http.ResponseWriter, r *http.Request) {

	festivalId, ok := AuthOrganizerOrEmployeeForFestival(w, r, &h.festivalService)
	if !ok {
		return
	}

	count, err := h.festivalService.GetEmployeeCount(festivalId)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, dtoFestival.FestivalPropCountResponse{
		FestivalId: festivalId,
		Count:      count,
	}, nil)
	log.Println("employee count retrieved successfully for festival:", festivalId)
}

func (h *festivalHandler) GetByEmployee(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthEmployeeRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	festivals, err := h.festivalService.GetByEmployee(utils.GetUsername(r.Context()))
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	festivalsResponse := dtoFestival.FestivalsResponse{
		Festivals: make([]dtoFestival.FestivalResponse, len(festivals)),
	}

	for i, festival := range festivals {
		images, err := h.festivalService.GetImages(festival.ID)
		if err != nil {
			log.Println("error:", err)
			continue
		}
		festivalsResponse.Festivals[i] = MapFestivalToResponse(festival, images)
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"festivals": festivalsResponse.Festivals}, nil)
	log.Println("festivals retrieved successfully for employee", utils.GetUsername(r.Context()))
}
