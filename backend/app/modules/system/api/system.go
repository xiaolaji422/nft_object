package api

import (
	"nft_object/library/response"
	"nft_object/statusCode"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

/*
 * @desc: 系统管理的controller,处理接口请求，进行数据校验和请求转发
 */

// 敏感词过滤器
var System = systemApi{}

type systemApi struct{}

// 发版之后验证代码是否更新
func (a *systemApi) GetVersion(r *ghttp.Request) {
	response.JsonExit(r, statusCode.SUCCESS, "ok", g.Map{"version": "1.0.7"})
}
