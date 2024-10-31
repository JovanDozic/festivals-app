package handlers

import (
	"log"
	"net/http"
	"user-service/internal/dto"
	"user-service/internal/models"
	"user-service/internal/services"
	"user-service/internal/utils"
)

type UserHandler interface {
	RegisterAttendee(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	TryAddress(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) UserHandler {
	return &userHandler{service: service}
}

func (h *userHandler) RegisterAttendee(w http.ResponseWriter, r *http.Request) {

	var input dto.RegisterAttendeeRequest
	if err := utils.ReadJSON(w, r, &input); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// todo: validate input

	user := models.User{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
		Role:     "ATTENDEE",
	}

	err := h.service.Create(&user)
	if err != nil {
		log.Println("error:", err)
		switch err {
		case models.ErrDuplicateUser:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Attendee registered successfully"))
	log.Println("attendee registered:", input.Username)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input dto.LoginRequest
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

	token, err := h.service.Login(input.Username, input.Password)
	if err != nil {
		log.Println("error:", err)
		switch err {
		case models.ErrNotFound:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		case models.ErrInvalidPassword:
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"token": token}, nil); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *userHandler) TryAddress(w http.ResponseWriter, r *http.Request) {

	log.Println("try address")

	var input dto.SaveAddressRequest
	if err := utils.ReadJSON(w, r, &input); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	addressId, err := h.service.SaveAddress(input.Street, input.Number, input.ApartmentSuite, input.City, input.PostalCode, input.CountryISO3)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	log.Println("address saved:", addressId)

	if err := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"address_id": addressId}, nil); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
