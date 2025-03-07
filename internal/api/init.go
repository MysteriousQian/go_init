package api

import (
	"go_server/internal/api/user"
	"go_server/internal/handler/network/server"
	"go_server/pkg/util/log"

	"github.com/spf13/viper"
)

const (
	userPath = "/user" // 用户路径
)

// 用户路由
var userRouter = []server.Router{
	{
		RequestType: "POST",
		Path:        userPath + "/login",
		Handler:     user.Login,
	},
	{
		RequestType: "GET",
		Path:        userPath + "/logout",
		Handler:     user.LoginOut,
		JwtEnabled:  true,
	},
}

func mergeRouter(router ...[]server.Router) []server.Router {
	var routers []server.Router
	for _, r := range router {
		routers = append(routers, r...)
	}
	return routers
}

/*
初始化路由和Web服务监听
*/
func Setup() {
	port := viper.GetInt("web.port")
	server.Stop()
	go func() {
		server.InitGinEngine(
			viper.GetString("web.mode"),
			mergeRouter(
				userRouter,
			),
			viper.GetBool("web.recordLog"),
			viper.GetBool("web.recovery"),
			viper.GetBool("web.allowCors"),
			port,
			viper.GetInt("web.readTimeout"),
			viper.GetInt("web.weiteTimeout"),
		)
		err := server.Run()
		if err != nil && err.Error() != "http: Server closed" {
			log.Fatalln("接口服务启动失败: %v", err)
		}
	}()
	log.Info("接口服务已启动,端口号:[%d]", port)
}
