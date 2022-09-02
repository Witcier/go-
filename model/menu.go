package model

import "witcier/go-api/global"

type Menu struct {
	global.Category
	Uri   string `json:"uri" gorm:"comment:uri"`
	Order int    `json:"order" gorm:"default:0;comment:排序"`
	global.Timestamp
}
