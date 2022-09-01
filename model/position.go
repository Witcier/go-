package model

import (
	"witcier/go-api/global"
)

type Position struct {
	global.ModelID
	Name     string `json:"name" gorm:"not null;unique;comment:职位名"`
	Remark   string `json:"remark" gorm:"not null;comment:备注"`
	ParentID uint   `json:"parentID" gorm:"not null;default:0;comment:备注"`
	global.Timestamp
}
