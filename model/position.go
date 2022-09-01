package model

import (
	"witcier/go-api/global"
)

type Position struct {
	global.Category
	Remark string `json:"remark" gorm:"not null;comment:备注"`
	global.Timestamp
}
