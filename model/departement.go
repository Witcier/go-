package model

import (
	"witcier/go-api/global"
)

type Department struct {
	global.ModelID
	Name   string `json:"name" gorm:"not null;unique;comment:部门名"`
	Remark string `json:"remark" gorm:"not null;comment:备注"`
	global.Timestamp
}
