package global

import (
	"time"

	"gorm.io/gorm"
)

type ModelID struct {
	ID uint `json:"id" gorm:"primarykey"`
}

type Timestamp struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type Category struct {
	ModelID
	Name     string `json:"name" gorm:"not null;unique;comment:名称"`
	ParentID uint   `json:"parentID" gorm:"not null;default:0"`
}
