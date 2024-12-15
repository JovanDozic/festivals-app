package handlers

import (
	models "backend/internal/models/user"
	services "backend/internal/services/common"
	"backend/internal/utils"
	"log"
	"net/http"
	"strings"
)

type LogHandler interface {
	GetAllLogs(w http.ResponseWriter, r *http.Request)
	GetLogsByRole(w http.ResponseWriter, r *http.Request)
}

type logHandler struct {
	logService services.LogService
}

func NewLogHandler(ls services.LogService) LogHandler {
	return &logHandler{
		logService: ls,
	}
}

func (h *logHandler) GetAllLogs(w http.ResponseWriter, r *http.Request) {

	ok := utils.AuthStaffRole(r.Context())
	if !ok {
		return
	}

	logs, err := h.logService.GetAll()
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	response := mapLogsToResponses(logs)

	utils.WriteJSON(w, http.StatusOK, response, nil)
	log.Println("logs retrieved for admin:" + utils.GetUsername(r.Context()))
}

func (h *logHandler) GetLogsByRole(w http.ResponseWriter, r *http.Request) {

	ok := utils.AuthAdminRole(r.Context())
	if !ok {
		return
	}

	role, err := GetParamFromRequest(r, "role")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	logs, err := h.logService.GetByUserRole(models.UserRole(strings.ToUpper(role)))
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	response := mapLogsToResponses(logs)

	utils.WriteJSON(w, http.StatusOK, response, nil)
	log.Println("logs retrieved for admin:" + utils.GetUsername(r.Context()))
}
