package models

import "go_server/internal/db/core/gorm"

var DB = gorm.MasterDb

/*
注册模型
这里的所有模型都会在程序启动的时候自动迁移
*/
func AutoMigrateAllModels() error {
	return gorm.MasterDb.AutoMigrate(
		&User{},
	) // 自动迁移数据库
}
