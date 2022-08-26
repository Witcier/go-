package global

import (
	"time"

	"gorm.io/gorm"
)

type ID struct {
	ID uint `gorm:"primarykey"`
}

type Timestamp struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
