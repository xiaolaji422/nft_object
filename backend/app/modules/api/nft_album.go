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
var NFTAlbumApi = func() *album {
	return &album{
		proxy: service.NFTAlbumImpl(),
		core: core.CoreApi{
			CheckRules: core.CheckRule{
				"QueryNotice": {},
			},
		},
	}
}

// 公告api类
type album struct {
	core  core.CoreApi
	proxy service.INFTAlbum
}

// 获取最新公告
func (a *album) Search(r *ghttp.Request) {
	ctx, _, err := a.core.CheckParams(r)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}

	res, err := a.proxy.Search(ctx, 1, r.GetString("keyword"))
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", res)
}
