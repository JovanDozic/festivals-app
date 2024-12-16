package repos

import (
	"backend/internal/repos/common"
	"backend/internal/repos/festival"
	"backend/internal/repos/user"

	"gorm.io/gorm"
)

type Repos struct {
	Log         user.LogRepo
	User        user.UserRepo
	UserProfile user.UserProfileRepo
	Location    common.LocationRepo
	Image       common.ImageRepo
	Festival    festival.FestivalRepo
	Item        festival.ItemRepo
	Order       festival.OrderRepo
}

func NewRepositories(db *gorm.DB) *Repos {
	return &Repos{
		Log:         user.NewLogRepo(db),
		User:        user.NewUserRepo(db),
		UserProfile: user.NewUserProfileRepo(db),
		Location:    common.NewLocationRepo(db),
		Image:       common.NewImageRepo(db),
		Festival:    festival.NewFestivalRepo(db),
		Item:        festival.NewItemRepo(db),
		Order:       festival.NewOrderRepository(db),
	}
}
