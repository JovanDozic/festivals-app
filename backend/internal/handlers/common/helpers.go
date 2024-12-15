package handlers

import (
	dto "backend/internal/dto/common"
	modelsCommon "backend/internal/models"
	modelsUser "backend/internal/models/user"
	"net/http"

	"github.com/gorilla/mux"
)

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

func GetParamFromRequest(r *http.Request, paramName string) (string, error) {
	vars := mux.Vars(r)
	paramString := vars[paramName]

	if paramString == "" {
		return "", modelsCommon.ErrBadRequest
	}

	return paramString, nil
}
