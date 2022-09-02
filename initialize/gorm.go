package initialize

import (
	"os"
	"witcier/go-api/global"
	"witcier/go-api/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
		model.Department{},
		model.Position{},
		model.Menu{},
	)

	if err != nil {
		global.Log.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	global.Log.Info("register table success")
}
