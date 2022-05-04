package api

import (
	"nft_object/app/core"
	"nft_object/app/modules/admin/service"
	"nft_object/app/modules/admin/validate"
	"nft_object/library/response"
	"nft_object/statusCode"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

// 角色管理
var Role = roleApi{}

type roleApi struct {
	core.BaseApi
	addValidate        *validate.RoleAddValidate
	editValidate       *validate.RoleEditValidate
	roleApiAddValidate *validate.RoleApiAddValidate
}

//列表(所有)
func (a *roleApi) RoleAll(r *ghttp.Request) {
	list, err := service.Role.List(r.Context(), r.GetMap())
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", list)
}

// 列表(分页)
func (a *roleApi) Items(r *ghttp.Request) {
	params := r.GetMap()
	list, err := service.Role.Items(r.Context(), params)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", list)
}

/**
 * @description  : 详情
 * @param         {*ghttp.Request} r
 * @return        {*}
 * @author       : fourteen
 */
func (a *roleApi) Detail(r *ghttp.Request) {

	params := r.GetMap()
	if err := gconv.Struct(params, &a.IdValidate); err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, "参数错误")
	}
	// 数据校验
	if e := gvalid.CheckStruct(r.Context(), a.IdValidate, &a.IdValidate); e != nil {
		response.Json(r, statusCode.ERROR_PARAMS, e.String(), e.Maps())
	}
	info, err := service.Role.Detail(r.Context(), params)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", info)
}

/**
 * @description  : 添加
 * @param         {*ghttp.Request} r
 * @return        {*}
 * @author       : fourteen
 */
func (a *roleApi) Add(r *ghttp.Request) {

	params := r.GetMap()
	if err := gconv.Struct(params, &a.addValidate); err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, "参数错误")
	}
	// 数据校验
	if e := gvalid.CheckStruct(r.Context(), a.addValidate, &a.addValidate); e != nil {
		response.Json(r, statusCode.ERROR_PARAMS, e.String(), e.Maps())
	}
	info, err := service.Role.Add(r.Context(), params)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", info)
}

/**
 * @description  : 修改
 * @param         {*ghttp.Request} r
 * @return        {*}
 * @author       : fourteen
 */
func (a *roleApi) Edit(r *ghttp.Request) {

	params := r.GetMap()
	if err := gconv.Struct(params, &a.editValidate); err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, "参数错误")
	}
	// 数据校验
	if e := gvalid.CheckStruct(r.Context(), a.editValidate, &a.editValidate); e != nil {
		response.Json(r, statusCode.ERROR_PARAMS, e.String(), e.Maps())
	}
	info, err := service.Role.Edit(r.Context(), params)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", info)
}

/**
 * @description  : 获取角色权限
 * @param         {*ghttp.Request} r
 * @return        {*}
 * @author       : fourteen
 */
func (a *roleApi) GetRoleApis(r *ghttp.Request) {
	roleId := gconv.Int(r.Get("id"))
	list, err := service.Role.RoleApiList(r.Context(), roleId)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", list)
}

/**
 * @description  : 获取角色权限
 * @param         {*ghttp.Request} r
 * @return        {*}
 * @author       : fourteen
 */
func (a *roleApi) SetRoleApis(r *ghttp.Request) {
	params := r.GetMap()

	if err := gconv.Struct(params, &a.roleApiAddValidate); err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, "参数错误")
	}
	// 数据校验
	if e := gvalid.CheckStruct(r.Context(), a.roleApiAddValidate, &a.roleApiAddValidate); e != nil {
		response.Json(r, statusCode.ERROR_PARAMS, e.String(), e.Maps())
	}

	info, err := service.Role.SetRoleApis(r.Context(), params)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", info)
}
