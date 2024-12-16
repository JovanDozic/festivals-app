package handlers

import (
	dto "backend/internal/dto/common"
	modelsUser "backend/internal/models/user"

	"backend/internal/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetIDParamFromRequest(r *http.Request, paramName string) (uint, error) {
	vars := mux.Vars(r)
	idString := vars[paramName]

	if idString == "" {
		return 0, models.ErrBadRequest
	}

	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		return 0, models.ErrBadRequest
	}

	return uint(id), nil
}

func GetParamFromRequest(r *http.Request, paramName string) (string, error) {
	vars := mux.Vars(r)
	paramString := vars[paramName]

	if paramString == "" {
		return "", models.ErrBadRequest
	}

	return paramString, nil
}

func mapLogsToResponses(logs []modelsUser.Log) []dto.LogResponse {
	responses := make([]dto.LogResponse, len(logs))
	for i, log := range logs {
		response := dto.LogResponse{
			ID:        log.ID,
			Message:   log.Description,
			CreatedAt: log.CreatedAt,
			Type:      log.Type,
		}
		if log.User != nil {
			response.Username = &log.User.Username
		}
		responses[i] = response
	}
	return responses
}
