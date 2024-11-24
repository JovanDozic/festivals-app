package repositories

import (
	models "backend/internal/models/festival"

	"gorm.io/gorm"
)

type FestivalRepo interface {
	Create(festival *models.Festival, organizerId uint) error
	GetByOrganizer(organizerId uint) ([]models.Festival, error)
	GetAll() ([]models.Festival, error)
	GetById(festivalId uint) (*models.Festival, error)
	Update(festival *models.Festival) error
	Delete(festivalId uint) error
	IsOrganizer(festivalId uint, organizerId uint) (bool, error)
	AddImage(festivalId uint, imageId uint) error
	Employ(festivalId uint, employeeId uint) error
	GetEmployeeCount(festivalId uint) (int, error)
}

type festivalRepo struct {
	db *gorm.DB
}

func NewFestivalRepo(db *gorm.DB) FestivalRepo {
	return &festivalRepo{db}
}

func (r *festivalRepo) Create(festival *models.Festival, organizerId uint) error {

	err := r.db.Create(festival).Error
	if err != nil {
		return err
	}

	festivalOrganizer := &models.FestivalOrganizer{
		FestivalID: festival.ID,
		UserID:     organizerId,
	}

	err = r.db.Create(festivalOrganizer).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *festivalRepo) GetByOrganizer(organizerId uint) ([]models.Festival, error) {
	var festivals []models.Festival
	err := r.db.
		Preload("Address").
		Preload("Address.City").
		Preload("Address.City.Country").
		Joins("JOIN festival_organizers ON festivals.id = festival_organizers.festival_id").
		Where("festival_organizers.user_id = ?", organizerId).
		Find(&festivals).Error
	if err != nil {
		return nil, err
	}

	return festivals, nil
}

func (r *festivalRepo) GetAll() ([]models.Festival, error) {

	var festivals []models.Festival
	err := r.db.Find(&festivals).Error
	if err != nil {
		return nil, err
	}

	return festivals, nil
}

func (r *festivalRepo) GetById(festivalId uint) (*models.Festival, error) {
	var festival models.Festival
	err := r.db.Preload("Address").
		Preload("Address.City").
		Preload("Address.City.Country").
		First(&festival, festivalId).Error
	if err != nil {
		return nil, err
	}

	return &festival, nil
}

func (r *festivalRepo) Update(festival *models.Festival) error {

	err := r.db.Save(festival).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *festivalRepo) Delete(festivalId uint) error {

	err := r.db.Delete(&models.Festival{}, festivalId).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *festivalRepo) IsOrganizer(festivalId uint, organizerId uint) (bool, error) {

	var count int64
	err := r.db.Table("festival_organizers").
		Where("festival_id = ? AND user_id = ?", festivalId, organizerId).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *festivalRepo) AddImage(festivalId uint, imageId uint) error {

	festivalImage := &models.FestivalImage{
		FestivalID: festivalId,
		ImageID:    imageId,
	}

	err := r.db.Create(festivalImage).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *festivalRepo) Employ(festivalId uint, employeeId uint) error {

	festivalEmployee := &models.FestivalEmployee{
		FestivalID: festivalId,
		UserID:     employeeId,
	}

	err := r.db.Create(festivalEmployee).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *festivalRepo) GetEmployeeCount(festivalId uint) (int, error) {

	var count int64
	err := r.db.Table("festival_employees").
		Where("festival_id = ?", festivalId).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}
