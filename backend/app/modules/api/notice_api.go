package api

// 接口层
import (
	"nft_object/app/core"
	"nft_object/app/modules/service"
	"nft_object/library/response"
	"nft_object/statusCode"

	"github.com/gogf/gf/net/ghttp"
)

// 公告对外服务变量
var NoticeApi = func() *notice {
	return &notice{
		proxy: service.NoticeImpl(),
		core: core.CoreApi{
			CheckRules: core.CheckRule{
				"QueryNotice": {},
			},
		},
	}
}

// 公告api类
type notice struct {
	core  core.CoreApi
	proxy service.INotice
}

// 获取最新公告
func (a *notice) QueryNotice(r *ghttp.Request) {
	ctx, _, err := a.core.CheckParams(r)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}

	res, isWarning, err := a.proxy.GetNewNotice(ctx)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", core.MapI{"is_warn": isWarning, "data": res})
}

//	获取历史公告
func (a *notice) QueryHistoryNotice(r *ghttp.Request) {
	ctx, _, err := a.core.CheckParams(r)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	res, err := a.proxy.GetHistoryNotice(ctx, r.GetMap())
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", res)
}
