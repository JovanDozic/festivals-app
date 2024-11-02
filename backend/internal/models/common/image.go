package models

import "github.com/google/uuid"

type Image struct {
	ImageID uuid.UUID `json:"imageId" gorm:"column:image_id;primaryKey;type:uuid"`
	URL     string    `json:"url" gorm:"column:url;size:255;not null"`
}
