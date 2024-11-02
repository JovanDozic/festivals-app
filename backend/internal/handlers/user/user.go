package handlers

import (
	dtoUser "backend/internal/dto/user"
	"backend/internal/models"
	modelsUser "backend/internal/models/user"
	servicesUser "backend/internal/services/user"
	"backend/internal/utils"
	"log"
	"net/http"
)

type UserHandler interface {
	RegisterAttendee(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	CreateUserProfile(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService servicesUser.UserService
}

func NewUserHandler(us servicesUser.UserService) UserHandler {
	return &userHandler{userService: us}
}

func (h *userHandler) RegisterAttendee(w http.ResponseWriter, r *http.Request) {

	var input dtoUser.RegisterAttendeeRequest
	if err := utils.ReadJSON(w, r, &input); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// todo: validate input

	user := modelsUser.User{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
		Role:     "ATTENDEE",
	}

	err := h.userService.Create(&user)
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
	var input dtoUser.LoginRequest
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

	token, err := h.userService.Login(input.Username, input.Password)
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

func (h *userHandler) CreateUserProfile(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthAttendee(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var input dtoUser.CreateUserProfileRequest
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

	profile := modelsUser.UserProfile{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		DateOfBirth: utils.ParseDate(input.DateOfBirth),
		PhoneNumber: &input.PhoneNumber,
	}
	username := utils.GetUsername(r.Context())

	if err := h.userService.CreateUserProfile(username, &profile); err != nil {
		log.Println("error:", err)
		switch err {
		case models.ErrNotFound:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		case models.ErrUserHasProfile:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("profile created successfully"))
	log.Println("profile created successfully for user:", username)
}
