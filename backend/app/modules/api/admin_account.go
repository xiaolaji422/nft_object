package api

import (
	"nft_object/app/core"
	"nft_object/app/modules/service"
	"nft_object/library/response"
	"nft_object/statusCode"

	"github.com/gogf/gf/net/ghttp"
)

// 公告对外服务变量
var AccountApi = func() *account {
	return &account{
		proxy: service.AccountImpl(),
		core: core.CoreApi{
			CheckRules: core.CheckRule{
				"Save": {"account@required#账户不能为空", "info@required#额外信息不能为空"},
			},
		},
	}
}

// 公告api类
type account struct {
	core  core.CoreApi
	proxy service.IAccount
}

// 获取最新公告
func (a *account) Save(r *ghttp.Request) {
	ctx, _, err := a.core.CheckParams(r)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}

	err = a.proxy.Save(ctx, r.GetString("account"), r.GetString("info"))
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok")
}

// 获取商品价格列表
func (a *account) List(r *ghttp.Request) {
	ctx, _, err := a.core.CheckParams(r)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}

	res, err := a.proxy.List(ctx)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", res)
}
