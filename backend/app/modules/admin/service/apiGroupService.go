package service

import (
	"context"
	"nft_object/app/core"
	"nft_object/app/dao"
	"nft_object/app/model"
	"nft_object/library/helper"

	"github.com/gogf/gf/frame/g"
)

// 接口分组管理
var ApiGroup = apiGroupService{}

type apiGroupService struct {
}

// 列表(分页)
func (s *apiGroupService) Items(ctx context.Context, params g.Map) (core.PageReult, error) {
	res, err := dao.ApiGroup.Pagenation(params)
	return res, err
}

// 列表(所有)
func (s *apiGroupService) List(ctx context.Context, params g.Map) ([]*model.ApiGroup, error) {
	reslut, err := dao.ApiGroup.FindAll()
	if err != nil {
		return nil, err
	}
	res, err := dao.ApiGroup.Strcuts(reslut)
	return res, err
}

// 详情
func (s *apiGroupService) Detail(ctx context.Context, data g.Map) (*model.ApiGroup, error) {
	reslut, err := dao.ApiGroup.Where("id", data["id"]).FindOne()
	if err != nil {
		return nil, err
	}
	res, err := dao.ApiGroup.Strcut(reslut)
	return res, err
}

//修改
func (s *apiGroupService) Edit(ctx context.Context, data g.Map) (int64, error) {
	data["modified_user"] = helper.GetRtx(ctx)
	res, err := dao.ApiGroup.Filter().Data(data).Where("id", data["id"]).Update()
	if err != nil {
		return 0, err
	}
	r, err := res.RowsAffected()
	return r, err
}

//修改
func (s *apiGroupService) Enabled(ctx context.Context, id string, enabled int) (int, error) {

	rtx := helper.GetRtx(ctx)
	res, err := dao.ApiGroup.Fields("enabled,modified_user").Where("id", id).Update(g.Map{
		"enabled":       enabled,
		"modified_user": rtx,
	})
	rows, err := core.BaseDao.HandleExecRes(res, err, "操作")
	if err != nil {
		return rows, err
	}
	if enabled == 0 {
		// 禁用分组的时候 禁用所有接口
		res, err = dao.Api.Where("group_id", id).Update(g.Map{
			"status":        0,
			"modified_user": rtx,
		})
		rows, err = core.BaseDao.HandleExecRes(res, err, "操作")
	}
	return rows, err
}

// 添加
func (s *apiGroupService) Add(ctx context.Context, data g.Map) (int64, error) {
	rtx := helper.GetRtx(ctx)
	data["modified_user"] = rtx
	data["create_user"] = rtx
	res, err := dao.ApiGroup.FieldsEx("id").Insert(data)
	if err != nil {
		return 0, err
	}
	r, err := res.RowsAffected()
	return r, err
}
