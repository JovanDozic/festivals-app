package handlers

import (
	"backend/internal/handlers/common"
	"backend/internal/handlers/festival"
	"backend/internal/handlers/user"
	"backend/internal/services"
)

type Handlers struct {
	AWS      common.AWSHandler
	Log      user.LogHandler
	User     user.UserHandler
	Festival festival.FestivalHandler
	Item     festival.ItemHandler
	Order    festival.OrderHandler
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		AWS:      common.NewAWSHandler(services.Logger, services.AWS, services.Festival),
		Log:      user.NewLogHandler(services.Log),
		User:     user.NewUserHandler(services.Logger, services.User, services.Location),
		Festival: festival.NewFestivalHandler(services.Logger, services.Festival, services.Location),
		Item:     festival.NewItemHandler(services.Logger, services.Item, services.Festival),
		Order:    festival.NewOrderHandler(services.Logger, services.Order, services.User, services.Email),
	}
}
