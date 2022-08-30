package model

import (
	"time"
	"witcier/go-api/global"
)

type User struct {
	global.ModelID
	Username      string    `json:"username" gorm:"not null;unique;comment:用户名"`
	Password      string    `json:"-" gorm:"not null;comment:用户密码"`
	RealName      string    `json:"realName" gorm:"not null;comment:真实姓名"`
	Email         string    `json:"email" gorm:"not null;unique;comment:邮箱"`
	Phone         string    `json:"phone" gorm:"not null;unique;comment:国内手机号"`
	Sex           int8      `json:"sex" gorm:"not null;default:0;comment:性别"`
	Birth         int       `json:"birth" gorm:"not null;comment:生日"`
	EntryTime     int       `json:"entryTime" gorm:"not null;comment:入职日期"`
	Status        bool      `json:"status" gorm:"not null;default:1;comment:状态"`
	IsAdmin       bool      `json:"isAdmin" gorm:"not null;default:0;comment:是否管理员"`
	DepartmentID  uint      `json:"departmentID" gorm:"not null;default:0;comment:部门ID"`
	PositionID    uint      `json:"positionID" gorm:"not null;default:0;comment:职位ID"`
	LastLoginTime time.Time `json:"lastLoginTime" gorm:"default:null;comment:最后登陆时间"`
	LastLoginIp   string    `json:"lastLoginIp" gorm:"default:null;comment:最后登陆 Ip"`
	global.Timestamp
}
