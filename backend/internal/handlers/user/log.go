package user

import (
	modelsUser "backend/internal/models/user"
	servicesUser "backend/internal/services/user"
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
	logService servicesUser.LogService
}

func NewLogHandler(ls servicesUser.LogService) LogHandler {
	return &logHandler{
		logService: ls,
	}
}

func (h *logHandler) GetAllLogs(w http.ResponseWriter, r *http.Request) {

	ok, role := utils.AuthGetRole(r.Context())
	if !ok {
		return
	}

	var logs []modelsUser.Log
	var err error
	switch *role {
	case modelsUser.RoleAdmin:
		logs, err = h.logService.GetAll()
	case modelsUser.RoleOrganizer:
		logs, err = h.logService.GetByUserRoles([]modelsUser.UserRole{modelsUser.RoleOrganizer, modelsUser.RoleEmployee, modelsUser.RoleAttendee})
	case modelsUser.RoleEmployee:
		logs, err = h.logService.GetByUserRoles([]modelsUser.UserRole{modelsUser.RoleEmployee, modelsUser.RoleAttendee})
	}
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	response := mapLogsToResponses(logs)

	utils.WriteJSON(w, http.StatusOK, response, nil)
	log.Printf("logs retrieved for: %s (%v)", utils.GetUsername(r.Context()), role)
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

	logs, err := h.logService.GetByUserRole(modelsUser.UserRole(strings.ToUpper(role)))
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	response := mapLogsToResponses(logs)

	utils.WriteJSON(w, http.StatusOK, response, nil)
	log.Println("logs retrieved for admin:" + utils.GetUsername(r.Context()))
}
