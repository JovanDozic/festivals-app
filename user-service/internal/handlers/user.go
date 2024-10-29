package handlers

import (
	"user-service/internal/dto"
	"user-service/internal/models"
	"user-service/internal/services"
	"user-service/internal/utils"
	"log"
	"net/http"
)

type UserHandler interface {
	RegisterAttendee(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
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
