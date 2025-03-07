package user

import (
	"fmt"
	"go_server/internal/handler/network/server"
	"go_server/internal/services"
	"go_server/pkg/util/log"
)

// 登录
func Login(resp server.Response) {
	param := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	err := resp.Json(&param)
	if err != nil || param.Username == "" || param.Password == "" {
		resp.Failed("param error")
		return
	}
	// 账号密码登录
	user, err := services.Login(param.Username, param.Password)
	if err != nil {
		resp.Failed(fmt.Sprintf("%v", err))
		return
	}
	token, err := server.GenerateToken(fmt.Sprintf("%d", user.Id))
	if err != nil {
		log.Error("GenerateToken error:%v", err)
		resp.Failed("auth failed, please contact the administrator")
		return
	}
	resp.Res["auth_token"] = token
	resp.Res["user_id"] = user.Id
	resp.Res["user_name"] = user.Name
	resp.Res["create_time"] = user.CreateTime
	resp.Success("operate success")
}

// 账号退出登陆
func LoginOut(resp server.Response) {
	uid := resp.GetUserID("user_id")
	if uid <= 0 {
		resp.Failed("param error")
		return
	}
	log.Info("userId:%d login out", uid)
	resp.Success("operate success")
}
