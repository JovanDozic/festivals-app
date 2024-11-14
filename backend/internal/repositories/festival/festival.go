package repositories

import (
	models "backend/internal/models/festival"

	"gorm.io/gorm"
)

type FestivalRepo interface {
	Create(festival *models.Festival, organizerId uint) error
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
