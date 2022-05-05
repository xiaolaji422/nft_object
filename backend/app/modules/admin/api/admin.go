package api

import (
	"nft_object/app/core"
	"nft_object/app/modules/admin/service"

	"nft_object/library/response"
	"nft_object/statusCode"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 账户接口
var Admin = adminApi{}

type adminApi struct {
	core.BaseApi
}

// 账户列表
func (a *adminApi) Items(r *ghttp.Request) {
	params := r.GetMap()
	// 删除角色搜索
	if _, ok := params["role_id"]; ok {
		role_id := params["role_id"]
		delete(params, "role_id")
		// 获取roleId下面的adminId
		adminIds, err := service.AdminRole.GetAdminIdsByRoleId(r.Context(), g.Map{
			"role_id": role_id,
		})
		if err != nil {
			response.Json(r, statusCode.ERROR, err.Error())
		}

		if len(adminIds) <= 0 {
			// 没有查找到结果
			response.Json(r, statusCode.SUCCESS, "ok", new(core.PageReult))
		}
		params["id"] = g.Map{"value": adminIds, "operator": "in"}
	}
	list, err := service.Admin.Items(r.Context(), params)

	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", list)
}

// 账户详情
func (a *adminApi) Detail(r *ghttp.Request) {
	list, err := service.Admin.Info(r.Context(), r.GetMap())
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", list)
}

// 用户角色
func (a *adminApi) Roles(r *ghttp.Request) {
	admin_id := gconv.Int(r.Get("admin_id"))
	list, err := service.AdminRole.AdminRoleListInfo(r.Context(), admin_id)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", list)
}

// 用户接口
func (a *adminApi) AdminApiArray(r *ghttp.Request) {
	admin_id := gconv.Int(r.Get("admin_id"))
	list, err := service.AdminApi.AdminApiListByAdmin(r.Context(), admin_id)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", list)
}

// 获取该用户的所有的接口权限
func (a *adminApi) AdminApis(r *ghttp.Request) {
	admin_id := gconv.Int(r.Get("admin_id"))
	info, err := service.Admin.AdminAllApi(r.Context(), admin_id)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", info)
}
