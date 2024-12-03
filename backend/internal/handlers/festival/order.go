package handlers

import (
	dtoFestival "backend/internal/dto/festival"
	models "backend/internal/models/festival"
	servicesFestival "backend/internal/services/festival"
	servicesUser "backend/internal/services/user"
	"backend/internal/utils"
	"log"
	"net/http"
)

type OrderHandler interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
}

type orderHandler struct {
	orderService servicesFestival.OrderService
	userService  servicesUser.UserService
}

func NewOrderHandler(os servicesFestival.OrderService, us servicesUser.UserService) OrderHandler {
	return &orderHandler{orderService: os, userService: us}
}

func (h *orderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthAttendeeRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var input dtoFestival.CreateTicketOrderRequest
	if err := utils.ReadJSON(w, r, &input); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	festivalTicket := models.FestivalTicket{
		ItemID: input.TicketTypeId,
	}

	if err := h.orderService.CreateFestivalTicket(&festivalTicket); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	username := utils.GetUsername(r.Context())
	attendeeId, err := h.userService.GetUserID(username)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	order := models.Order{
		FestivalTicketID: festivalTicket.ID,
		TotalAmount:      float64(input.TotalPrice),
		UserID:           attendeeId,
	}

	if err := h.orderService.CreateOrder(&order); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"orderId": order.ID}, nil)
	log.Println("order created", order.ID)
}
