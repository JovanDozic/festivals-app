package dto

import (
	dto "backend/internal/dto/common"
	"time"
)

type HealthCheckResponse struct {
	ServiceName string `json:"service_name"`
	Status      string `json:"status"`
	Environment string `json:"environment"`
	API         string `json:"API"`
	Secure      bool   `json:"secure"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type GetUserProfileResponse struct {
	Username    string                  `json:"username"`
	Email       string                  `json:"email"`
	Role        string                  `json:"role"`
	FirstName   string                  `json:"firstName"`
	LastName    string                  `json:"lastName"`
	DateOfBirth string                  `json:"dateOfBirth"`
	PhoneNumber string                  `json:"phoneNumber"`
	Address     *dto.GetAddressResponse `json:"address"`
	ImageURL    *string                 `json:"imageURL"`
}

type CreateStaffResponse struct {
	Username string `json:"username"`
	UserId   uint   `json:"userId"`
}

type GetEmployeesResponse struct {
	FestivalId uint               `json:"festivalId"`
	Employees  []EmployeeResponse `json:"employees"`
}

type EmployeeResponse struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	PhoneNumber string `json:"phoneNumber"`
}

type UserListResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
}

type LogResponse struct {
	ID        uint      `json:"id"`
	Username  *string   `json:"username"`
	Role      *string   `json:"role"`
	Message   string    `json:"message"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
}
