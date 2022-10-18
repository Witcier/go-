package model

import "witcier/go-api/global"

type Role struct {
	global.ModelID
	Name   string `json:"name" gorm:"not null;unique;comment:用户组名"`
	Remark string `json:"remark" gorm:"comment:备注"`
	Menus  []Menu `gorm:"many2many:role_menus;"`
	global.Timestamp
}
