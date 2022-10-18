package main

import (
	"witcier/go-api/core"
	"witcier/go-api/global"
	"witcier/go-api/initialize"

	"go.uber.org/zap"
)

func main() {
	global.Viper = core.Viper()
	global.Log = core.Zap()
	zap.ReplaceGlobals(global.Log)
	global.DB = initialize.Gorm()

	global.Trans = initialize.ValidateTrans("en")

	// 数据库迁移
	if global.DB != nil {
		initialize.RegisterTables(global.DB)
		db, _ := global.DB.DB()

		defer db.Close()
	}

	initialize.GenerateTokenForTest()

	core.RunServer()
}
