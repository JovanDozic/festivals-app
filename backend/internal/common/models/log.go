package common

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Log struct {
	LogID       uuid.UUID       `json:"logId" gorm:"column:log_id;primaryKey;type:uuid"`
	Type        string          `json:"type" gorm:"column:type;size:100;not null"`
	Timestamp   time.Time       `json:"timestamp" gorm:"column:timestamp;not null"`
	Description *string         `json:"description" gorm:"column:description"`
	Data        *datatypes.JSON `json:"data" gorm:"column:data;type:jsonb"`
	UserID      *uuid.UUID      `json:"userId" gorm:"column:user_id;type:uuid"`
}
