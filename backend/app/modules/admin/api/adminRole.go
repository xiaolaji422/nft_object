package api

import (
	"nft_object/app/core"
	"nft_object/app/modules/admin/service"
	"nft_object/library/response"
	"nft_object/statusCode"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 账户角色接口
var AdminRole = adminRoleApi{}

type adminRoleApi struct {
	core.BaseApi
}

// 启用、禁用接口
func (a *adminRoleApi) EnableRecord(r *ghttp.Request) {
	id := r.Get("id")
	enabled := r.Get("enabled")
	rows, err := service.AdminRole.EnableRecord(r.Context(), gconv.Int(id), gconv.Int(enabled))
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", rows)
}

// 账户绑定角色
func (a *adminRoleApi) AddRole(r *ghttp.Request) {
	roleId := gconv.Int(r.Get("role_id"))
	adminIdsReq := gconv.String(r.Get("admin_ids"))
	admin_ids := strings.Split(adminIdsReq, ",")
	successNum, err := service.AdminRole.AddAdminRole(r.Context(), admin_ids, roleId)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", g.Map{"successNum": successNum})
}

// 账户删除角色
func (a *adminRoleApi) DelRole(r *ghttp.Request) {
	id := gconv.Int(r.Get("id"))
	if id <= 0 {
		response.Json(r, statusCode.ERROR, "id不能为空")
	}
	// 数据校验
	info, err := service.AdminRole.DelAdminRole(r.Context(), int(id))
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", info)
}
