package handlers

import (
	"festival-service/internal/config"
	"festival-service/internal/dto"
	"festival-service/internal/utils"
	"log"
	"net/http"
)

type CommonHandler interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
}

type commonHandler struct {
	config *config.Config
}

func NewCommonHandler(config *config.Config) CommonHandler {
	return &commonHandler{
		config: config,
	}
}

func (ch *commonHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
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
