package api

import (
	"nft_object/app/core"
	"nft_object/app/entry"
	"nft_object/app/modules/service"
	"nft_object/library/response"
	"nft_object/statusCode"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 自动锁单
var AccountLockApi = func() *account_lock {
	return &account_lock{
		proxy: service.AccountLockImpl(),
		core: core.CoreApi{
			CheckRules: core.CheckRule{
				// Id        int    `orm:"id,primary"        json:"id"`         // ID
				// AdminId   int    `orm:"admin_id"    json:"admin_id"`         // 账户id
				// AccountId int    `orm:"account_id,unique" json:"account_id"` // 账号id
				// AlbumId   string `orm:"album_id,unique"   json:"album_id"`   // 商品id
				// AlbumName string `orm:"album_name"        json:"album_name"` // 商品名称
				// Min       string `orm:"min"               json:"min"`        // 最低价格
				// Max       string `orm:"max"               json:"max"`        // 最高价格

				"Add": {
					"account_id@required#账户不能为空",
					"album_id@required#商品不能为空",
					"min@required#最低价格不能为空",
					"max@required#最高价格不能为空",
				},
				"Cancel": {
					"id@required#ID不能为空",
				},
				"Lock": {
					"id@required#ID不能为空",
				},
			},
		},
	}
}

// 自动锁单api类
type account_lock struct {
	core  core.CoreApi
	proxy service.IAccountLock
}

// 获取最新公告
func (a *account_lock) Add(r *ghttp.Request) {
	ctx, params, err := a.core.CheckParams(r)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	// 校验内容
	var info = entry.AccountData{}
	err = gconv.Struct(params, &info)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	err = a.proxy.Add(ctx, &info)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok")
}

// 锁单回调
func (a *account_lock) Lock(r *ghttp.Request) {
	ctx, params, err := a.core.CheckParams(r)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}

	err = a.proxy.Lock(ctx, r.GetInt("id"), params)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok")
}

func (a *account_lock) Cancel(r *ghttp.Request) {
	ctx, _, err := a.core.CheckParams(r)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	err = a.proxy.Cacnel(ctx, r.GetInt("id"))
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok")
}

// 获取商品价格列表
func (a *account_lock) List(r *ghttp.Request) {
	ctx, _, err := a.core.CheckParams(r)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}

	res, err := a.proxy.List(ctx)
	if err != nil {
		response.Json(r, statusCode.ERROR_PARAMS, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "ok", res)
}
