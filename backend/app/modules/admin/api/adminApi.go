package api

import (
	"nft_object/app/modules/admin/service"
	"nft_object/library/response"
	"nft_object/statusCode"
	"strings"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 用户定制权限接口
var AdminApi = adminApiApi{}

type adminApiApi struct{}

// 启用、禁用记录
func (a *adminApiApi) EnableRecord(r *ghttp.Request) {
	id := r.Get("id")
	rows, err := service.AdminApi.EnableRecord(r.Context(), gconv.Int(id), gconv.Int(r.Get("status")))
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", rows)
}

// 给用户添加角色
func (a *adminApiApi) AddApis(r *ghttp.Request) {
	apiIdsReq := gconv.String(r.Get("apiIds"))
	adminIdsReq := gconv.String(r.Get("admin_ids"))
	admin_ids := strings.Split(adminIdsReq, ",")
	apiIds := strings.Split(apiIdsReq, ",")

	successNum, errNum, err := service.AdminApi.AddAdminApis(r.Context(), admin_ids, apiIds, gconv.Int(r.Get("enabled")))
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", successNum, errNum)
}
