package services

import (
	"backend/internal/config"
	"backend/internal/repos"
	"backend/internal/services/common"
	"backend/internal/services/festival"
	"backend/internal/services/user"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Services struct {
	Logger   user.Logger
	Log      user.LogService
	User     user.UserService
	Email    common.EmailService
	Location common.LocationService
	AWS      common.AWSService
	Festival festival.FestivalService
	Item     festival.ItemService
	Order    festival.OrderService
}

func NewServices(cfg *config.Config, repos *repos.Repos, s3Client *s3.Client, s3Presign *s3.PresignClient) *Services {

	userService := user.NewUserService(cfg, repos.User, repos.UserProfile, repos.Location, repos.Image)

	return &Services{
		Logger:   user.NewLogger(repos.Log, repos.User),
		Log:      user.NewLogService(repos.Log),
		User:     user.NewUserService(cfg, repos.User, repos.UserProfile, repos.Location, repos.Image),
		Email:    common.NewEmailService(cfg),
		Location: common.NewLocationService(repos.Location),
		AWS:      common.NewAWSService(s3Client, s3Presign, cfg),
		Festival: festival.NewFestivalService(cfg, repos.Festival, repos.User, repos.Location, repos.Image, repos.Order),
		Item:     festival.NewItemService(cfg, repos.Item, repos.Location, repos.Image),
		Order:    festival.NewOrderService(repos.Order, repos.Item, repos.Festival, userService, repos.Image, repos.Location),
	}
}
