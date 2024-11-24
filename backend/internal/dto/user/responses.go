package dto

import dto "backend/internal/dto/common"

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
