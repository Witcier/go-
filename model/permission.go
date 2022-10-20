package model

import "witcier/go-api/global"

type Permission struct {
	global.Category
	Rule string `json:"rule" gorm:"not null;comment:权限规则"`
	global.Timestamp
}
