package service

import (
	"context"
	"errors"
	"nft_object/app/core"
	"nft_object/app/dao"
	"nft_object/app/model"
	"nft_object/library/auth"
	"nft_object/library/helper"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

// 账户管理业务
var Admin = adminService{}

type adminService struct{}

//  列表(分页)
func (s *adminService) Items(ctx context.Context, params g.Map) (core.PageReult, error) {
	res, err := dao.Admin.Pagenation(params)
	// 对接trpc
	return res, err
}

// 列表(所有)
func (s *adminService) List(ctx context.Context, params g.Map) ([]*model.Admin, error) {
	reslut, err := dao.Admin.FindAll()
	if err != nil {
		return nil, err
	}
	res, err := dao.Admin.Strcuts(reslut)
	return res, err
}

// 获取账户详情
func (s *adminService) Info(ctx context.Context, params g.Map) (*model.Admin, error) {
	record, err := dao.Admin.Where(params).FindOne()
	if err != nil {
		return nil, err
	}
	res, err := dao.Admin.Strcut(record)
	return res, err
}

// 获取信息
func (s *adminService) InfoByLoginName(ctx context.Context, loginName string) (*model.Admin, error) {
	record, err := dao.Admin.Clone().Ctx(ctx).Where("login_name", loginName).FindOne()
	if err != nil {
		return nil, err
	}
	if record.IsEmpty() {
		return nil, errors.New("账号不存在")
	}

	var info = &model.Admin{}

	err = gconv.Struct(record, info)
	return info, err
}

// 获取账户详情(ask)
func (s *adminService) FindByLoginName(login_name string) (*model.Admin, error) {
	record, err := dao.Admin.Where("login_name", login_name).FindOne()
	if err != nil {
		return nil, err
	}
	res, err := dao.Admin.Strcut(record)
	return res, err
}

// 获取人员接口列表
func (s *adminService) AdminApiList(ctx context.Context, adminId int) ([]int, error) {
	// 断言
	adminApiList, err := AdminApi.AbleApiByAdmin(ctx, adminId)
	if err != nil {
		return nil, err
	}
	var res []int
	if len(adminApiList) > 0 {
		for _, v := range adminApiList {
			res = append(res, gconv.Int(v.ApiId))
		}
	}
	return res, err
}

// 获取人员被禁用的接口列表
func (s *adminService) AdminDisableApiList(ctx context.Context, adminId int) ([]int, error) {
	// 断言
	adminApiList, err := AdminApi.DisableApiByAdmin(ctx, adminId)
	if err != nil {
		return nil, err
	}
	var res []int
	if len(adminApiList) > 0 {
		for _, v := range adminApiList {
			res = append(res, gconv.Int(v.ApiId))
		}
	}
	return res, err
}

// 获取人员所有接口接口列表
func (s *adminService) AdminAllApi(ctx context.Context, adminId int) ([]int, error) {
	// 先获取该人的所有的角色
	var roleApi []int
	roleList, err := AdminRole.AdminRoleList(ctx, adminId)
	if err != nil {
		return roleApi, err
	}
	roleApi, err = Role.RoleApiList(ctx, roleList)
	if err != nil {
		return roleApi, err
	}

	adminApi, err := Admin.AdminApiList(ctx, adminId)
	if err != nil {
		return nil, err
	}
	// 定制禁用的api
	disAbleAdminApi, err := Admin.AdminDisableApiList(ctx, adminId)
	if err != nil {
		return nil, err
	}
	// 合并   清洗去重
	roleApi = helper.ClearSliceInt(append(roleApi, adminApi...), disAbleAdminApi)
	return roleApi, err
}

// 注册人员
func (s *adminService) Register(ctx context.Context, login_name, password string) error {
	// 先获取该人的所有的角色
	var admin = model.Admin{
		LoginName: login_name,
		Password:  auth.Md5Encrypt(password),
		Level:     1,
	}
	checkCount, err := dao.Admin.Where("login_name", login_name).Count()
	if err != nil {
		return err
	}
	if checkCount > 0 {
		return errors.New("该账号已注册")
	}

	res, err := dao.Admin.Clone().Ctx(ctx).Data(admin).Fields("login_name,password,level").Insert()
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("注册失败")
	}
	return nil
}
