package handlers

import (
	"backend/internal/config"
	dto "backend/internal/dto/common"
	"backend/internal/utils"
	"log"
	"net/http"
)

type HealthHandler interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
}

type healthHandler struct {
	config *config.Config
}

func NewHealthHandler(config *config.Config) HealthHandler {
	return &healthHandler{
		config: config,
	}
}

func (ch *healthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	data := dto.HealthCheckResponse{
		ServiceName: ch.config.App.Name,
		Status:      "ok",
		Environment: ch.config.App.Env,
		API:         ch.config.App.APIVersion,
		Secure:      false,
	}

	if err := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"healthCheck": data}, nil); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
