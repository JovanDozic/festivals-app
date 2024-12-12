package handlers

import (
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
