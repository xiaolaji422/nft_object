package api

import (
	"nft_object/app/core"
	"nft_object/app/modules/service"
	"nft_object/library/response"
	"nft_object/statusCode"

	"github.com/gogf/gf/net/ghttp"
)

// 接口层

// 公告对外服务变量
var SendMsgApi = func() *send_msg {
	return &send_msg{
		proxy: service.SendMsgImpl(),
		core: core.CoreApi{
			CheckRules: core.CheckRule{
				"SendMsg": {"userid@required#userid不能为空", "msg@required#msg不能为空"},
			},
		},
	}
}

// 公告api类
type send_msg struct {
	core  core.CoreApi
	proxy service.ISendMsg
}

func (a *send_msg) SendMsg(r *ghttp.Request) {
	ctx, _, err := a.core.CheckParams(r)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}

	msg := service.NewMessage("API", r.GetString("msg"))
	err = a.proxy.SendMsg(ctx, r.GetString("userid"), msg)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok")
}
