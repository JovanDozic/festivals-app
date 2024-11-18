package repositories

import (
	modelsCommon "backend/internal/models/common"

	"gorm.io/gorm"
)

type ImageRepo interface {
	Create(image *modelsCommon.Image) error
	Get(id uint) (*modelsCommon.Image, error)
	GetByFestival(festivalId uint) ([]modelsCommon.Image, error)
}

type imageRepo struct {
	db *gorm.DB
}

func NewImageRepo(db *gorm.DB) ImageRepo {
	return &imageRepo{db}
}

func (r *imageRepo) Create(image *modelsCommon.Image) error {
	return r.db.Create(image).Error
}

func (r *imageRepo) Get(id uint) (*modelsCommon.Image, error) {
	var image modelsCommon.Image
	if err := r.db.First(&image, id).Error; err != nil {
		return nil, err
	}

	return &image, nil
}

func (r *imageRepo) GetByFestival(festivalId uint) ([]modelsCommon.Image, error) {
	var images []modelsCommon.Image
	err := r.db.Table("images").
		Joins("JOIN festival_images ON images.id = festival_images.image_id").
		Where("festival_images.festival_id = ?", festivalId).
		Find(&images).Error
	if err != nil {
		return nil, err
	}

	return images, nil
}
