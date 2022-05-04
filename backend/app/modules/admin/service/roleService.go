package service

import (
	"context"
	"nft_object/app/core"
	"nft_object/app/dao"
	"nft_object/app/model"
	"nft_object/library/helper"
	"reflect"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

// 角色管理
var Role = roleService{}

type roleService struct{}

// 列表(所有)
func (s *roleService) List(ctx context.Context, params g.Map) ([]*model.Role, error) {
	result, err := dao.Role.FindAll()
	if err != nil {
		return nil, err
	}
	res, err := dao.Role.Strcuts(result)
	return res, err
}

// 详情
func (s *roleService) Info(ctx context.Context, roleId int) (*model.Role, error) {
	result, err := dao.Role.Where("id", roleId).FindOne()
	if err != nil {
		return nil, err
	}
	res, err := dao.Role.Strcut(result)
	return res, err
}

// 列表(分页)
func (s *roleService) Items(ctx context.Context, params g.Map) (core.PageReult, error) {
	res, err := dao.Role.Pagenation(params)
	return res, err
}

// 详情
func (s *roleService) Detail(ctx context.Context, data g.Map) (*model.Role, error) {
	result, err := dao.Role.Where("id", data["id"]).FindOne()
	if err != nil {
		return nil, err
	}
	res, err := dao.Role.Strcut(result)
	return res, err
}

// 修改
func (s *roleService) Edit(ctx context.Context, data g.Map) (int64, error) {
	data["modified_user"] = helper.GetRtx(ctx)
	res, err := dao.Role.Filter().Data(data).Where("id", data["id"]).Update()
	if err != nil {
		return 0, err
	}
	r, err := res.RowsAffected()
	return r, err
}

// 添加
func (s *roleService) Add(ctx context.Context, data g.Map) (int64, error) {
	rtx := helper.GetRtx(ctx)
	data["modified_user"] = rtx
	data["created_user"] = rtx
	res, _ := dao.Role.Filter().Insert(data)
	r, err := res.RowsAffected()
	return r, err
}

//-----------------角色权限管理
// 获取角色接口列表
func (s *roleService) RoleApiList(ctx context.Context, roleIds interface{}) ([]int, error) {
	roleApiDao := dao.RoleApi.Filter()
	apiList := []int{}
	// 断言
	switch reflect.TypeOf(roleIds).Kind() {
	case reflect.Int:
		where := g.Map{
			helper.StrCase("role_id"): roleIds,
		}
		apiResult, _ := roleApiDao.Where(where).FindOne()
		apiRes, _ := dao.RoleApi.Strcut(apiResult)

		if apiRes != nil {
			apiList = getSplitSlice(apiRes.Apis, apiList)
		}
	case reflect.Slice:
		where := g.Map{
			helper.StrCase("role_id") + " in(?)": roleIds,
		}
		apiResult, err := roleApiDao.Where(where).FindAll()
		if err != nil {
			return apiList, err
		}
		apiRes, err := dao.RoleApi.Strcuts(apiResult)
		if err != nil {
			return apiList, err
		}
		if len(apiRes) > 0 {
			for _, v := range apiRes {
				apiList = getSplitSlice(v.Apis, apiList)
			}
		}
	}
	apiList = helper.ArrayUniqueInt(apiList)
	return apiList, nil
}

// 助手函数
func getSplitSlice(apis string, data []int) []int {
	splitRes := strings.Split(apis, ",")
	if len(splitRes) > 0 {
		for _, v := range splitRes {
			data = append(data, gconv.Int(v))
		}
	}
	return data
}

// 设置角色接口列表
func (s *roleService) SetRoleApis(ctx context.Context, data g.Map) (int64, error) {
	// 判断是否存在  存在即修改  不存在就添加
	count, err := dao.RoleApi.Fields("apis").Where(helper.StrCase("role_id"), data["roleId"]).Count()
	if err != nil {
		return 0, err
	}
	rtx := helper.GetRtx(ctx)
	data["modified_user"] = rtx
	// 事务管理
	query := dao.RoleApi.Filter()
	if count > 0 {

		query = query.Where(helper.StrCase("role_id"), data["roleId"])
	} else {
		data["create_user"] = rtx
	}
	res, err := query.Save(data)
	if err != nil {
		return 0, err
	}
	r, err := res.RowsAffected()
	return r, err
}
