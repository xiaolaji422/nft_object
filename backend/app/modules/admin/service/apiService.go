package service

/*
 * @desc: 接口管理的service,接口业务实现
 */

import (
	"context"
	"errors"
	"nft_object/app/core"
	"nft_object/app/dao"
	"nft_object/app/model"
	"nft_object/library/helper"

	"github.com/gogf/gf/frame/g"
)

// 接口管理service
var Api = apiService{}

type apiService struct {
}

// 获取角色api集合
func (s *apiService) GetInfoByPath(ctx context.Context, path string) (*model.Api, error) {
	result, err := dao.Api.Where("route", path).FindOne()
	if err != nil {
		return nil, err
	}
	res, err := dao.Api.Strcut(result)
	return res, err
}

// 列表(分页)
func (s *apiService) Items(ctx context.Context, params g.Map) (core.PageReult, error) {
	res, err := dao.Api.Pagenation(params)
	return res, err
}

// 接口列表(所有)
func (s *apiService) List(ctx context.Context, data g.Map) ([]*model.Api, error) {
	result, err := dao.Api.Where(data).FindAll()
	if err != nil {
		return nil, err
	}
	res, err := dao.Api.Strcuts(result)
	return res, err
}

// 接口详情
func (s *apiService) Detail(ctx context.Context, api g.Map) (*model.Api, error) {
	result, err := dao.Api.Where("id", api["id"]).FindOne()
	if err != nil {
		return nil, err
	}
	res, err := dao.Api.Strcut(result)
	return res, err
}

// 修改
func (s *apiService) Edit(ctx context.Context, api g.Map) (int64, error) {
	api["modefied_user"] = helper.GetRtx(ctx)
	res, err := dao.Api.Filter().Data(api).Where("id", api["id"]).Update()
	if err != nil {
		return 0, err
	}
	r, err := res.RowsAffected()
	return r, err
}

// 添加
func (s *apiService) Add(ctx context.Context, api g.Map) (int64, error) {
	rtx := helper.GetRtx(ctx)
	api["modefied_user"] = rtx
	api["create_user"] = rtx
	res, err := dao.Api.FieldsEx("id").Insert(api)
	if err != nil {
		return 0, err
	}
	r, err := res.RowsAffected()
	return r, err
}

// 启用、禁用
func (s *apiService) Enable(ctx context.Context, id int, enabled int) (int64, error) {

	if enabled == 1 {
		// 启用时需查看分组状态
		apiRes, err := dao.Api.Where("id", id).FindOne()
		if err != nil {
			return 0, err
		}
		info, err := dao.Api.Strcut(apiRes)
		if err != nil {
			return 0, err
		}
		if g.IsEmpty(info) {
			return 0, errors.New("接口不存在")
		}
		// 查看父级的状态
		groupRes, err := dao.ApiGroup.Where("id", info.GroupId).FindOne()
		if err != nil {
			return 0, err
		}
		ginfo, err := dao.ApiGroup.Strcut(groupRes)
		if err != nil {
			return 0, err
		}
		if !g.IsEmpty(ginfo) {
			if ginfo.Enabled == 0 {
				return 0, errors.New("该接口分组已禁用")
			}
		}
	}
	res, _ := dao.Api.Fields("enabled,modified_user").Data(g.Map{
		"enabled":       enabled,
		"modified_user": helper.GetRtx(ctx),
	}).Where("id", id).Update()
	r, err := res.RowsAffected()
	return r, err
}
