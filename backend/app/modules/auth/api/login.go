package api

// 权限校验模块

import (
	adminSer "nft_object/app/modules/admin/service"
	"nft_object/app/modules/auth/service"

	// 用户管理模块的接口

	"nft_object/library/response"
	"nft_object/statusCode"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 登录验证
var Login = loginApi{}

type loginApi struct{}

// 退出登录
func (a *loginApi) LoginOut(r *ghttp.Request) {

	r.Session.Clear()
	host := r.GetHost()
	// 跳转
	// client_url := "http://" + host + "/#/dashboard"
	response.Json(r, statusCode.SUCCESS, "success", host)
}

// 用户信息
func (a *loginApi) UserInfo(r *ghttp.Request) {
	res := service.Login.UserInfo(r)
	if g.IsEmpty(res) {
		response.Json(r, statusCode.ERROR_NO_LOGIN, "登陆失败")
	}
	response.Json(r, statusCode.SUCCESS, "ok", res)
}

func (a *loginApi) Login(r *ghttp.Request) {
	pwd := r.GetString("password")
	login_name := r.GetString("username")
	if len(pwd) < 6 || len(pwd) > 16 {
		response.Json(r, statusCode.ERROR_PARAMS, "请输入6-16位密码")
	}
	if len(login_name) < 6 || len(login_name) > 16 {
		response.Json(r, statusCode.ERROR_PARAMS, "请输入6-16位账户")
	}
	// 获取到登录信息

	adminInfo, err := service.Login.Login(r.Context(), login_name, pwd)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", adminInfo)

}

// 注册
func (a *loginApi) Register(r *ghttp.Request) {
	pwd := r.GetString("password")
	login_name := r.GetString("username")
	if len(pwd) < 6 || len(pwd) > 16 {
		response.Json(r, statusCode.ERROR_PARAMS, "请输入6-16位密码")
	}
	if len(login_name) < 6 || len(login_name) > 16 {
		response.Json(r, statusCode.ERROR_PARAMS, "请输入6-16位账户")
	}
	err := adminSer.Admin.Register(r.Context(), login_name, pwd)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok")
}
