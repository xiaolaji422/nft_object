package api

/*
 * @desc: 接口管理的controller,处理接口请求，进行数据校验和请求转发
 */

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

// 接口管理
var Api = apiApi{}

type apiApi struct {
	core.BaseApi
	addValidate  *validate.ApiAddValidate
	editValidate *validate.ApiEditValidate
}

// 列表(所有)
func (a *apiApi) ApiAll(r *ghttp.Request) {
	params := r.GetMap()
	params["enabled"] = 1
	list, err := service.Api.List(r.Context(), params)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", list)
}

// 列表(分页)
func (a *apiApi) Items(r *ghttp.Request) {

	list, err := service.Api.Items(r.Context(), r.GetMap())
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", list)
}

// 详情
func (a *apiApi) Detail(r *ghttp.Request) {
	params := r.GetMap()
	if err := gconv.Struct(params, &a.IdValidate); err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, "参数错误")
	}
	// 数据校验
	ctx := r.Context()
	if e := gvalid.CheckStruct(ctx, a.IdValidate, &a.IdValidate); e != nil {
		response.Json(r, statusCode.ERROR_PARAMS, e.String(), e.Maps())
	}
	info, err := service.Api.Detail(r.Context(), params)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", info)
}

// 添加
func (a *apiApi) Add(r *ghttp.Request) {

	params := r.GetMap()
	if err := gconv.Struct(params, &a.addValidate); err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, "参数错误")
	}
	// 数据校验
	ctx := r.Context()
	if e := gvalid.CheckStruct(ctx, a.addValidate, &a.addValidate); e != nil {
		response.Json(r, statusCode.ERROR_PARAMS, e.String(), e.Maps())
	}
	rows, err := service.Api.Add(r.Context(), params)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", rows)
}

// 修改
func (a *apiApi) Edit(r *ghttp.Request) {

	params := r.GetMap()

	if err := gconv.Struct(params, &a.editValidate); err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, "参数错误")
	}
	// 数据校验
	ctx := r.Context()
	if e := gvalid.CheckStruct(ctx, a.editValidate, &a.editValidate); e != nil {
		response.Json(r, statusCode.ERROR_PARAMS, e.String(), e.Maps())
	}
	rows, err := service.Api.Edit(r.Context(), params)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", rows)
}

// 启用禁用
func (a *apiApi) Enable(r *ghttp.Request) {

	params := r.GetMap()
	rules := []string{
		"id@required#产品ID不能为空",
		"enabled@required#状态值不能为空",
	}
	ctx := r.Context()
	if e := gvalid.CheckMap(ctx, params, rules); e != nil {
		response.Json(r, statusCode.ERROR_PARAMS, e.String(), e.Maps())
	}

	rows, err := service.Api.Enable(r.Context(), gconv.Int(params["id"]), gconv.Int(params["enabled"]))
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", rows)
}
