package handlers

import (
	"auth-service/internal/dto"
	"auth-service/internal/models"
	"auth-service/internal/services"
	"auth-service/internal/utils"
	"log"
	"net/http"
)

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	RegisterAttendee(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) UserHandler {
	return &userHandler{service: service}
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {

	var input dto.CreateUserRequest
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
		Role:     input.Role,
	}

	err := h.service.Create(&user)
	if err != nil {
		log.Println("error:", err)
		switch err {
		case models.ErrDuplicateUsername:
			http.Error(w, err.Error(), http.StatusConflict)
		case models.ErrRoleNotFound:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
	log.Println("user created:", input.Username)
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
		case models.ErrDuplicateUsername:
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
