package handlers

import (
	dto "backend/internal/dto/common"
	servicesCommon "backend/internal/services/common"
	servicesFestival "backend/internal/services/festival"
	"backend/internal/utils"
	"log"
	"net/http"
)

type AWSHandler interface {
	GetPresignedURL(w http.ResponseWriter, r *http.Request)
}

type awsHandler struct {
	awsService      servicesCommon.AWSService
	festivalService servicesFestival.FestivalService
}

func NewAWSHandler(awsService servicesCommon.AWSService, festivalService servicesFestival.FestivalService) AWSHandler {
	return &awsHandler{
		awsService:      awsService,
		festivalService: festivalService,
	}
}

func (h *awsHandler) GetPresignedURL(w http.ResponseWriter, r *http.Request) {

	festivalId, ok := h.authorizeOrganizerForFestival(w, r)
	if !ok {
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
	log.Println("retrieved presigned upload URL for festival:", festivalId)
}
