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

/*
 api分组管理
*/
var ApiGroup = apiGroupApi{}

type apiGroupApi struct {
	core.BaseApi
	addValidate  *validate.ApiGroupAddValidate
	editValidate *validate.ApiGroupEditValidate
}

// 列表(所有)
func (a *apiGroupApi) GroupAll(r *ghttp.Request) {
	list, err := service.ApiGroup.List(r.Context(), r.GetMap())
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", list)
}

// 列表(分页)
func (a *apiGroupApi) Items(r *ghttp.Request) {
	params := r.GetMap()
	list, err := service.ApiGroup.Items(r.Context(), params)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", list)
}

// 详情
func (a *apiGroupApi) Detail(r *ghttp.Request) {

	params := r.GetMap()
	if err := gconv.Struct(params, &a.IdValidate); err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, "参数错误")
	}
	// 数据校验
	if e := gvalid.CheckStruct(r.Context(), a.IdValidate, &a.IdValidate); e != nil {
		response.Json(r, statusCode.ERROR_PARAMS, e.String(), e.Maps())
	}
	info, err := service.ApiGroup.Detail(r.Context(), params)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", info)
}

// 添加
func (a *apiGroupApi) Add(r *ghttp.Request) {

	params := r.GetMap()
	if err := gconv.Struct(params, &a.addValidate); err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, "参数错误")
	}
	// 数据校验
	if e := gvalid.CheckStruct(r.Context(), a.addValidate, &a.addValidate); e != nil {
		response.Json(r, statusCode.ERROR_PARAMS, e.String(), e.Maps())
	}
	rows, err := service.ApiGroup.Add(r.Context(), params)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", rows)
}

// 修改
func (a *apiGroupApi) Edit(r *ghttp.Request) {

	params := r.GetMap()

	if err := gconv.Struct(params, &a.editValidate); err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, "参数错误")
	}
	// 数据校验
	if e := gvalid.CheckStruct(r.Context(), a.editValidate, &a.editValidate); e != nil {
		response.Json(r, statusCode.ERROR_PARAMS, e.String(), e.Maps())
	}

	rows, err := service.ApiGroup.Edit(r.Context(), params)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", rows)
}

// 修改
func (a *apiGroupApi) Enable(r *ghttp.Request) {

	params := r.GetMap()
	rules := []string{
		"id@required#分组id不能为空",
		"enabled@required#修改状态不能为空",
	}
	if e := gvalid.CheckMap(r.Context(), params, rules); e != nil {
		response.Json(r, statusCode.ERROR_PARAMS, e.String(), e.Maps())
	}
	info, err := service.ApiGroup.Enabled(r.Context(), gconv.String(params["id"]), gconv.Int(params["enabled"]))
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", info)
}
