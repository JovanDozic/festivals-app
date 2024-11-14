package repositories

import (
	models "backend/internal/models/festival"

	"gorm.io/gorm"
)

type FestivalRepo interface {
	Create(festival *models.Festival, organizerId uint) error
	GetByOrganizer(organizerId uint) ([]models.Festival, error)
	GetAll() ([]models.Festival, error)
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
	err := r.db.Table("festivals").
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
