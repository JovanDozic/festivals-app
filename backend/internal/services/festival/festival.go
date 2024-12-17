package festival

import (
	"backend/internal/config"
	modelsError "backend/internal/models"
	modelsCommon "backend/internal/models/common"
	modelsFestival "backend/internal/models/festival"
	"backend/internal/repos/common"
	"backend/internal/repos/festival"
	"backend/internal/repos/user"
	"log"
	"strings"
)

type FestivalService interface {
	Create(festival *modelsFestival.Festival, username string, address *modelsCommon.Address) error
	GetByOrganizer(username string) ([]modelsFestival.Festival, error)
	GetByOrganizerById(id uint) ([]modelsFestival.Festival, error)
	GetByEmployee(username string) ([]modelsFestival.Festival, error)
	GetAll() ([]modelsFestival.Festival, error)
	GetAllPublic() ([]modelsFestival.Festival, error)
	GetById(festivalId uint) (*modelsFestival.Festival, error)
	Update(festivalId uint, updatedFestival *modelsFestival.Festival) error
	Delete(festivalId uint) error
	PublishFestival(festivalId uint) error
	CancelFestival(festivalId uint) error
	CompleteFestival(festivalId uint) error
	OpenStore(festivalId uint) error
	CloseStore(festivalId uint) error
	IsOrganizer(username string, festivalId uint) (bool, error)
	IsEmployee(username string, festivalId uint) (bool, error)
	GetImages(festivalId uint) ([]modelsCommon.Image, error)
	AddImage(festivalId uint, image *modelsCommon.Image) error
	RemoveImage(festivalId uint, imageId uint) error
	GetAddressID(festivalId uint) (uint, error)
	Employ(festivalId uint, employeeId uint) error
	Fire(festivalId uint, employeeId uint) error
	GetEmployeeCount(festivalId uint) (int, error)
	GetAttendeeCount() (int, error)
	GetFestivalCount() (int, error)
}

type festivalService struct {
	config       *config.Config
	festivalRepo festival.FestivalRepo
	userRepo     user.UserRepo
	locationRepo common.LocationRepo
	imageRepo    common.ImageRepo
	orderRepo    festival.OrderRepo
}

func NewFestivalService(
	cfg *config.Config,
	fr festival.FestivalRepo,
	ur user.UserRepo,
	lr common.LocationRepo,
	ir common.ImageRepo,
	or festival.OrderRepo,
) FestivalService {
	return &festivalService{
		config:       cfg,
		festivalRepo: fr,
		userRepo:     ur,
		locationRepo: lr,
		imageRepo:    ir,
		orderRepo:    or,
	}
}

func (s *festivalService) Create(festival *modelsFestival.Festival, username string, address *modelsCommon.Address) error {

	if err := festival.Validate(); err != nil {
		return err
	}

	if err := address.Validate(); err != nil {
		return err
	}

	err := s.locationRepo.CreateAddress(address)
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

func (s *festivalService) GetByOrganizer(username string) ([]modelsFestival.Festival, error) {

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	festivals, err := s.festivalRepo.GetByOrganizer(user.ID)
	if err != nil {
		return nil, err
	}

	return festivals, nil
}

func (s *festivalService) GetByOrganizerById(id uint) ([]modelsFestival.Festival, error) {

	user, err := s.userRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	festivals, err := s.festivalRepo.GetByOrganizer(user.ID)
	if err != nil {
		return nil, err
	}

	return festivals, nil
}

func (s *festivalService) GetByEmployee(username string) ([]modelsFestival.Festival, error) {

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	festivals, err := s.festivalRepo.GetByEmployee(user.ID)
	if err != nil {
		return nil, err
	}

	return festivals, nil
}

func (s *festivalService) GetAllPublic() ([]modelsFestival.Festival, error) {

	festivals, err := s.festivalRepo.GetAll()
	if err != nil {
		return nil, err
	}

	filteredFestivals := make([]modelsFestival.Festival, 0)
	for _, festival := range festivals {
		if festival.Status != "PRIVATE" && festival.Status != "CANCELLED" {
			filteredFestivals = append(filteredFestivals, festival)
		}
	}
	festivals = filteredFestivals

	return festivals, nil
}

func (s *festivalService) GetAll() ([]modelsFestival.Festival, error) {

	festivals, err := s.festivalRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return festivals, nil
}

func (s *festivalService) GetById(festivalId uint) (*modelsFestival.Festival, error) {

	festival, err := s.festivalRepo.GetById(festivalId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, modelsError.ErrNotFound
		}
		return nil, err
	}

	return festival, nil
}

func (s *festivalService) Update(festivalId uint, updatedFestival *modelsFestival.Festival) error {

	if err := updatedFestival.Validate(); err != nil {
		return err
	}

	festival, err := s.festivalRepo.GetById(festivalId)
	if err != nil {
		return err
	}

	festival.Name = updatedFestival.Name
	festival.Description = updatedFestival.Description
	festival.StartDate = updatedFestival.StartDate
	festival.EndDate = updatedFestival.EndDate
	festival.Capacity = updatedFestival.Capacity

	if err := s.festivalRepo.Update(festival); err != nil {
		return err
	}

	return nil
}

func (s *festivalService) Delete(festivalId uint) error {

	if err := s.festivalRepo.Delete(festivalId); err != nil {
		return err
	}

	return nil
}

func (s *festivalService) PublishFestival(festivalId uint) error {

	festival, err := s.festivalRepo.GetById(festivalId)
	if err != nil {
		return err
	}

	festival.Status = "PUBLIC"

	if err := s.festivalRepo.Update(festival); err != nil {
		return err
	}

	return nil
}

func (s *festivalService) CancelFestival(festivalId uint) error {

	festival, err := s.festivalRepo.GetById(festivalId)
	if err != nil {
		return err
	}

	festival.Status = "CANCELLED"

	if err := s.festivalRepo.Update(festival); err != nil {
		return err
	}

	bracelets, err := s.orderRepo.GetBraceletsByFestival(festivalId)
	if err != nil {
		log.Println(err)
	}

	for _, bracelet := range bracelets {
		bracelet.Status = "DEACTIVATED"
		if err := s.orderRepo.UpdateBracelet(&bracelet); err != nil {
			log.Println(err)
		}
	}

	return nil
}

func (s *festivalService) CompleteFestival(festivalId uint) error {

	festival, err := s.festivalRepo.GetById(festivalId)
	if err != nil {
		return err
	}

	festival.Status = "COMPLETED"

	if err := s.festivalRepo.Update(festival); err != nil {
		return err
	}

	bracelets, err := s.orderRepo.GetBraceletsByFestival(festivalId)
	if err != nil {
		log.Println(err)
	}

	for _, bracelet := range bracelets {
		bracelet.Status = "DEACTIVATED"
		if err := s.orderRepo.UpdateBracelet(&bracelet); err != nil {
			log.Println(err)
		}
	}

	return nil
}

func (s *festivalService) OpenStore(festivalId uint) error {

	festival, err := s.festivalRepo.GetById(festivalId)
	if err != nil {
		return err
	}

	festival.StoreStatus = "OPEN"

	if err := s.festivalRepo.Update(festival); err != nil {
		return err
	}

	return nil
}

func (s *festivalService) CloseStore(festivalId uint) error {

	festival, err := s.festivalRepo.GetById(festivalId)
	if err != nil {
		return err
	}

	festival.StoreStatus = "CLOSED"

	if err := s.festivalRepo.Update(festival); err != nil {
		return err
	}

	return nil
}

func (s *festivalService) IsOrganizer(username string, festivalId uint) (bool, error) {

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return false, err
	}

	isOrganizer, err := s.festivalRepo.IsOrganizer(festivalId, user.ID)
	if err != nil {
		return false, err
	}

	return isOrganizer, nil
}

func (s *festivalService) IsEmployee(username string, festivalId uint) (bool, error) {

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return false, err
	}

	isEmployee, err := s.festivalRepo.IsEmployee(festivalId, user.ID)
	if err != nil {
		return false, err
	}

	return isEmployee, nil
}

func (s *festivalService) GetImages(festivalId uint) ([]modelsCommon.Image, error) {

	images, err := s.imageRepo.GetByFestival(festivalId)
	if err != nil {
		return nil, err
	}

	return images, nil
}

func (s *festivalService) AddImage(festivalId uint, image *modelsCommon.Image) error {

	if err := image.Validate(); err != nil {
		return err
	}

	if err := s.imageRepo.Create(image); err != nil {
		return err
	}

	if err := s.festivalRepo.AddImage(festivalId, image.ID); err != nil {
		return err
	}

	return nil
}

func (s *festivalService) RemoveImage(festivalId uint, imageId uint) error {

	if err := s.festivalRepo.RemoveImage(festivalId, imageId); err != nil {
		return err
	}

	return nil
}

func (s *festivalService) GetAddressID(festivalId uint) (uint, error) {

	festival, err := s.festivalRepo.GetById(festivalId)
	if err != nil {
		return 0, err
	}

	address, err := s.locationRepo.GetAddressByID(festival.AddressID)
	if err != nil {
		return 0, err
	}

	return address.ID, nil
}

func (s *festivalService) Employ(festivalId uint, employeeId uint) error {

	if err := s.festivalRepo.Employ(festivalId, employeeId); err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate key value"):
			return modelsError.ErrDuplicateUser
		case strings.Contains(err.Error(), "violates foreign key constraint"):
			return modelsError.ErrUserNotFound
		case strings.Contains(err.Error(), "foreign key constraint"):
			return modelsError.ErrRoleNotFound
		default:
			return err
		}
	}

	return nil
}

func (s *festivalService) Fire(festivalId uint, employeeId uint) error {

	if err := s.festivalRepo.Fire(festivalId, employeeId); err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate key value"):
			return modelsError.ErrDuplicateUser
		case strings.Contains(err.Error(), "violates foreign key constraint"):
			return modelsError.ErrUserNotFound
		case strings.Contains(err.Error(), "foreign key constraint"):
			return modelsError.ErrRoleNotFound
		default:
			return err
		}
	}

	return nil
}

func (s *festivalService) GetEmployeeCount(festivalId uint) (int, error) {

	count, err := s.festivalRepo.GetEmployeeCount(festivalId)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *festivalService) GetAttendeeCount() (int, error) {

	count, err := s.userRepo.GetAttendeeCount()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *festivalService) GetFestivalCount() (int, error) {

	count, err := s.festivalRepo.GetFestivalCount()
	if err != nil {
		return 0, err
	}

	return count, nil
}
