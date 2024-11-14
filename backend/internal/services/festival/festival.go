package services

import (
	"backend/internal/config"
	modelsError "backend/internal/models"
	modelsCommon "backend/internal/models/common"
	modelsFestival "backend/internal/models/festival"
	reposFestival "backend/internal/repositories/festival"
	reposUser "backend/internal/repositories/user"
	servicesCommon "backend/internal/services/common"
	"strings"
)

type FestivalService interface {
	Create(festival *modelsFestival.Festival, username string, address *modelsCommon.Address) error
}

type festivalService struct {
	config          *config.Config
	festivalRepo    reposFestival.FestivalRepo
	userRepo        reposUser.UserRepo
	locationService servicesCommon.LocationService
}

func NewFestivalService(
	config *config.Config,
	festivalRepo reposFestival.FestivalRepo,
	userRepo reposUser.UserRepo,
	locationService servicesCommon.LocationService,
) FestivalService {
	return &festivalService{
		config:          config,
		festivalRepo:    festivalRepo,
		userRepo:        userRepo,
		locationService: locationService,
	}
}

func (s *festivalService) Create(festival *modelsFestival.Festival, username string, address *modelsCommon.Address) error {

	if err := festival.Validate(); err != nil {
		return err
	}

	if err := address.Validate(); err != nil {
		return err
	}

	err := s.locationService.CreateAddress(address)
	if err != nil {
		return err
	}

	festival.AddressID = address.ID

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return err
	}

	festival.Status = "PRIVATE"
	festival.StoreStatus = "CLOSED"

	if err := s.festivalRepo.Create(festival, user.ID); err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate key value"):
			return modelsError.ErrDuplicateUser
		case strings.Contains(err.Error(), "foreign key constraint"):
			return modelsError.ErrRoleNotFound
		default:
			return err
		}
	}

	return nil
}
