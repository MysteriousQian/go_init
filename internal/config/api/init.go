package ad

import (
	"go_server/internal/api"
	"go_server/internal/config"
	"go_server/internal/db/core/gorm"
	"go_server/internal/db/models"
	"go_server/pkg/util/log"
)

/*
初始化配置
*/
func init() {
	config.LoadConfig("config", "yaml", ".")
	loggingInit()
	databaseInit()
	clientApiInit()
	config.WatchingConfig()
}

/*
初始化日志配置
*/
func loggingInit() {
	err := log.Setup()
	if err != nil {
		log.Fatalln("程序启动失败:%s", err.Error())
	}
	config.SetWatching("logging", func(oldWebConfig, newWebConfig interface{}) {
		log.Info("日志配置发生变化,将应用最新配置:\n旧配置:%+v,\n新配置:%+v", oldWebConfig, newWebConfig)
		err := log.Setup()
		if err != nil {
			log.Fatalln("更新日志配置失败:%s", err.Error())
		}
	}, nil)
}

/*
初始化接口服务
*/
func clientApiInit() {
	api.Setup()
	// stripehandler.Setup()
	// stripehandler.Init()
	config.SetWatching("web", func(oldWebConfig, newWebConfig interface{}) {
		log.Info("接口服务配置发生变化,将重启接口服务:\n旧配置:%+v,\n新配置:%+v", oldWebConfig, newWebConfig)
		api.Setup()
	}, nil)
}

/*
初始化数据库
*/
func databaseInit() {
	gorm.Setup()
	models.AutoMigrateAllModels()
	config.SetWatching("database", func(oldDatabaseConfig, newDatabaseConfig interface{}) {
		log.Info("数据库配置发生变化,将重新连接数据库:\n旧配置:%+v,\n新配置:%+v", oldDatabaseConfig, newDatabaseConfig)
		gorm.Setup()
	}, nil)
}
