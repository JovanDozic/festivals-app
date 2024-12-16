package handlers

import (
	dto "backend/internal/dto/common"
	servicesCommon "backend/internal/services/common"
	servicesFestival "backend/internal/services/festival"
	sevicesUser "backend/internal/services/user"
	"backend/internal/utils"
	"log"
	"net/http"
)

type AWSHandler interface {
	GetPresignedURL(w http.ResponseWriter, r *http.Request)
}

type awsHandler struct {
	log             sevicesUser.Logger
	awsService      servicesCommon.AWSService
	festivalService servicesFestival.FestivalService
}

func NewAWSHandler(lg sevicesUser.Logger, as servicesCommon.AWSService, fs servicesFestival.FestivalService) AWSHandler {
	return &awsHandler{
		awsService:      as,
		festivalService: fs,
		log:             lg,
	}
}

func (h *awsHandler) GetPresignedURL(w http.ResponseWriter, r *http.Request) {

	if !utils.Auth(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var input dto.GetPresignedURLRequest
	if err := utils.ReadJSON(w, r, &input); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	response, err := h.awsService.GetPresignedURL(&input)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, response, nil)
	h.log.Info("retrieved presigned upload image URL", r.Context())
}
