package dto

type HealthCheckResponse struct {
	ServiceName string `json:"service_name"`
	Status      string `json:"status"`
	Environment string `json:"environment"`
	API         string `json:"API"`
	Secure      bool   `json:"secure"`
}
