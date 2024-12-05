package handlers

import (
	dtoCommon "backend/internal/dto/common"
	dtoUser "backend/internal/dto/user"
	"backend/internal/models"
	modelsCommon "backend/internal/models/common"
	modelsUser "backend/internal/models/user"
	servicesCommon "backend/internal/services/common"
	servicesUser "backend/internal/services/user"
	"backend/internal/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler interface {
	RegisterAttendee(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	GetUserProfile(w http.ResponseWriter, r *http.Request)
	CreateUserProfile(w http.ResponseWriter, r *http.Request)
	CreateUserAddress(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
	UpdateUserProfile(w http.ResponseWriter, r *http.Request)
	UpdateUserEmail(w http.ResponseWriter, r *http.Request)
	UpdateUserAddress(w http.ResponseWriter, r *http.Request)
	CreateEmployee(w http.ResponseWriter, r *http.Request)
	GetFestivalEmployees(w http.ResponseWriter, r *http.Request)
	GetEmployeesNotOnFestival(w http.ResponseWriter, r *http.Request)
	UpdateStaffEmail(w http.ResponseWriter, r *http.Request)
	UpdateStaffProfile(w http.ResponseWriter, r *http.Request)
	UpdateProfilePhoto(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService     servicesUser.UserService
	locationService servicesCommon.LocationService
}

func NewUserHandler(us servicesUser.UserService, ls servicesCommon.LocationService) UserHandler {
	return &userHandler{userService: us, locationService: ls}
}

func (h *userHandler) RegisterAttendee(w http.ResponseWriter, r *http.Request) {

	var input dtoUser.RegisterUserRequest
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

	err := h.userService.CreateUser(&user)
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

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"message": "attendee registered successfully"}, nil)
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

func (h *userHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {

	if !utils.Auth(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	username := utils.GetUsername(r.Context())

	data, err := h.userService.GetUserProfile(username)
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

	if err := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"userProfile": data}, nil); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *userHandler) CreateUserProfile(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthAttendeeRole(r.Context()) {
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
		PhoneNumber: input.PhoneNumber,
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

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"message": "profile created successfully"}, nil)
	log.Println("profile created successfully for user:", username)
}

func (h *userHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {

	if !utils.Auth(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var input dtoUser.ChangePasswordRequest
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

	username := utils.GetUsername(r.Context())

	if err := h.userService.ChangePassword(username, input.OldPassword, input.NewPassword); err != nil {
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

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "password changed successfully"}, nil)
	log.Println("password changed successfully for user:", username)
}

func (h *userHandler) CreateUserAddress(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthAttendeeRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var input dtoCommon.CreateAddressRequest
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

	address := modelsCommon.Address{
		Street:         input.Street,
		Number:         input.Number,
		ApartmentSuite: &input.ApartmentSuite,
		City: modelsCommon.City{
			Name:       input.City,
			PostalCode: input.PostalCode,
			Country: modelsCommon.Country{
				ISO3: input.CountryISO3,
			},
		},
	}

	username := utils.GetUsername(r.Context())

	if err := h.userService.CreateUserAddress(username, &address); err != nil {
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

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"message": "address created successfully"}, nil)
	log.Println("address created successfully for user:", username)
}

func (h *userHandler) UpdateUserEmail(w http.ResponseWriter, r *http.Request) {

	if !utils.Auth((r.Context())) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var input dtoUser.UpdateUserEmailRequest
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

	username := utils.GetUsername(r.Context())

	if err := h.userService.UpdateUserEmail(username, input.Email); err != nil {
		log.Println("error:", err)
		switch err {
		case models.ErrNotFound:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		case models.ErrDuplicateEmail:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "email updated successfully"}, nil)
	log.Println("email updated successfully for user:", username)
}

func (h *userHandler) UpdateUserAddress(w http.ResponseWriter, r *http.Request) {

	if !utils.Auth((r.Context())) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var input dtoCommon.UpdateAddressRequest
	if err := utils.ReadJSON(w, r, &input); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	username := utils.GetUsername(r.Context())

	addressId, err := h.userService.GetAddressID(username)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	address := modelsCommon.Address{
		Street:         input.Street,
		Number:         input.Number,
		ApartmentSuite: &input.ApartmentSuite,
		City: modelsCommon.City{
			Name:       input.City,
			PostalCode: input.PostalCode,
			Country: modelsCommon.Country{
				ISO3: input.CountryISO3,
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

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "user address updated successfully"}, nil)
	log.Println("user address updated successfully:", utils.GetUsername(r.Context()))
}

func (h *userHandler) UpdateStaffEmail(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthOrganizerRole((r.Context())) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var input dtoUser.UpdateStaffEmailRequest
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

	if err := h.userService.UpdateUserEmail(input.Username, input.Email); err != nil {
		log.Println("error:", err)
		switch err {
		case models.ErrNotFound:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		case models.ErrDuplicateEmail:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "email updated successfully"}, nil)
	log.Println("email updated successfully for user:", input.Username)
}

// ! This method only allows user that is logged in to change their profile
func (h *userHandler) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {

	if !utils.Auth(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var input dtoUser.UpdateUserProfileRequest
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

	username := utils.GetUsername(r.Context())

	updatedProfile := modelsUser.UserProfile{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		DateOfBirth: utils.ParseDate(input.DateOfBirth),
		PhoneNumber: input.PhoneNumber,
	}

	if err := h.userService.UpdateUserProfile(username, &updatedProfile); err != nil {
		log.Println("error:", err)
		switch err {
		case models.ErrNotFound:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "profile updated successfully"}, nil)
	log.Println("profile updated successfully for user:", username)
}

func (h *userHandler) UpdateStaffProfile(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthOrganizerRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var input dtoUser.UpdateStaffProfileRequest
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

	updatedProfile := modelsUser.UserProfile{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		DateOfBirth: utils.ParseDate(input.DateOfBirth),
		PhoneNumber: input.PhoneNumber,
	}

	if err := h.userService.UpdateUserProfile(input.Username, &updatedProfile); err != nil {
		log.Println("error:", err)
		switch err {
		case models.ErrNotFound:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "profile updated successfully"}, nil)
	log.Println("profile updated successfully for user:", input.Username)
}

func (h *userHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthOrganizerRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var input dtoUser.CreateStaffRequest
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

	user := modelsUser.User{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
		Role:     string(modelsUser.RoleEmployee),
	}

	if err := h.userService.CreateUser(&user); err != nil {
		log.Println("error:", err)
		switch err {
		case models.ErrDuplicateUser:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	if err := h.userService.CreateUserProfile(input.Username, &modelsUser.UserProfile{
		FirstName:   input.UserProfile.FirstName,
		LastName:    input.UserProfile.LastName,
		DateOfBirth: utils.ParseDate(input.UserProfile.DateOfBirth),
		PhoneNumber: input.UserProfile.PhoneNumber,
	}); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"employee": dtoUser.CreateStaffResponse{
		Username: user.Username,
		UserId:   user.ID,
	}}, nil)
	log.Println("employee created:", input.Username)
}

func (h *userHandler) GetFestivalEmployees(w http.ResponseWriter, r *http.Request) {

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

	employees, err := h.userService.GetFestivalEmployees(utils.ToUint(festivalId))
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response := dtoUser.GetEmployeesResponse{
		FestivalId: utils.ToUint(festivalId),
		Employees:  make([]dtoUser.EmployeeResponse, len(employees)),
	}
	for i, employee := range employees {
		response.Employees[i] = dtoUser.EmployeeResponse{
			ID:          employee.User.ID,
			Username:    employee.User.Username,
			Email:       employee.User.Email,
			FirstName:   employee.FirstName,
			LastName:    employee.LastName,
			DateOfBirth: employee.DateOfBirth.Format("2006-01-02"),
			PhoneNumber: employee.PhoneNumber,
		}
	}

	utils.WriteJSON(w, http.StatusOK, response, nil)
	log.Println("employees retrieved successfully for festival:", festivalId)
}

func (h *userHandler) GetEmployeesNotOnFestival(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthOrganizerRole(r.Context()) {
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

	employees, err := h.userService.GetEmployeesNotOnFestival(utils.ToUint(festivalId))
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	employeesResponse := dtoUser.GetEmployeesResponse{
		FestivalId: utils.ToUint(festivalId),
		Employees:  make([]dtoUser.EmployeeResponse, len(employees)),
	}
	for i, employee := range employees {
		employeesResponse.Employees[i] = dtoUser.EmployeeResponse{
			ID:          employee.User.ID,
			Username:    employee.User.Username,
			Email:       employee.User.Email,
			FirstName:   employee.FirstName,
			LastName:    employee.LastName,
			DateOfBirth: employee.DateOfBirth.Format("2006-01-02"),
			PhoneNumber: employee.PhoneNumber,
		}
	}

	utils.WriteJSON(w, http.StatusOK, employeesResponse, nil)
	log.Println("employees not on festival retrieved successfully for festival:", festivalId)
}

func (h *userHandler) UpdateProfilePhoto(w http.ResponseWriter, r *http.Request) {

	if !utils.Auth(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var input dtoUser.UpdateProfilePhotoRequest
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

	username := utils.GetUsername(r.Context())
	image := modelsCommon.Image{
		URL: input.ImageURL,
	}

	if err := h.userService.UpdateProfilePhoto(username, &image); err != nil {
		log.Println("error:", err)
		switch err {
		case models.ErrNotFound:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "profile photo updated successfully"}, nil)
	log.Println("profile photo updated successfully for user:", username)
}
