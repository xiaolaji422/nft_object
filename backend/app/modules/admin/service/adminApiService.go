package service

import (
	"context"
	"nft_object/app/dao"
	"nft_object/app/model"
	"nft_object/library/helper"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

// 账户接口管理
var AdminApi = adminApiService{}

type adminApiService struct{}

// 获取用户api集合
func (s *adminApiService) AdminApiListByAdmin(ctx context.Context, adminId int) ([]*model.AdminApi, error) {
	where := g.Map{
		helper.StrCase("admin_id"): adminId,
	}
	reslut, err := dao.AdminApi.Where(where).FindAll()
	if err != nil {
		return nil, err
	}
	res, err := dao.AdminApi.Strcuts(reslut)
	return res, err
}

// 获取定制启用的用户接口
func (s *adminApiService) AbleApiByAdmin(ctx context.Context, adminId int) ([]*model.AdminApi, error) {
	where := g.Map{
		helper.StrCase("admin_id"): adminId,
		helper.StrCase("enabled"):  1,
	}
	reslut, err := dao.AdminApi.Where(where).FindAll()
	if err != nil {
		return nil, err
	}
	res, err := dao.AdminApi.Strcuts(reslut)
	return res, err
}

// 获取用户的禁用api
func (s *adminApiService) DisableApiByAdmin(ctx context.Context, adminId int) ([]*model.AdminApi, error) {
	where := g.Map{
		helper.StrCase("admin_id"): adminId,
		helper.StrCase("enabled"):  0,
	}
	reslut, err := dao.AdminApi.Where(where).FindAll()
	if err != nil {
		return nil, err
	}
	res, err := dao.AdminApi.Strcuts(reslut)
	return res, err
}

//  启用/禁用
func (s *adminApiService) EnableRecord(ctx context.Context, id int, enabled int) (int, error) {
	res, err := dao.AdminApi.Where("id", id).Update(g.Map{
		"enabled": enabled,
	})
	if err != nil {
		return 0, err
	}
	count, _ := res.RowsAffected()
	return gconv.Int(count), nil
}

// 批量/添加用户角色
func (s *adminApiService) AddAdminApi(ctx context.Context, adminId int, apiIds []string, enabled int) (err error) {
	// 获取角色的信息
	apiInfo, err := Api.List(ctx, g.Map{"id in (?)": apiIds})
	if err != nil {
		return err
	}
	apiInfoMap := make(map[int]string)
	for _, v := range apiInfo {
		apiInfoMap[v.Id] = v.Name
	}
	// 事务管理
	tx, err := g.DB().Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	if len(apiIds) > 0 {
		for _, v := range apiIds {
			apiId := gconv.Int(v)
			apiName := ""
			if _, ok := apiInfoMap[apiId]; ok {
				apiName = apiInfoMap[apiId]
			}
			_, err = handleAdd(ctx, tx, g.Map{
				"api_id":   apiId,
				"api_name": apiName,
				"admin_id": adminId,
				"enabled":  enabled,
			})
			if err != nil {
				return
			}

		}
	}
	return err
}

// 批量/添加用户角色
func (s *adminApiService) AddAdminApis(ctx context.Context, adminIds []string,
	apiIds []string, status int) (successNum int, errNum int, err error) {
	// 事务管理

	for _, v := range adminIds {
		err = s.AddAdminApi(ctx, gconv.Int(v), apiIds, status)
		if err != nil {
			errNum++

		} else {
			successNum++
		}
	}
	// 获取角色的信息
	return
}

// 处理添加的事情
func handleAdd(ctx context.Context, tx *gdb.TX, data g.Map) (int, error) {
	// 首先查找有没有
	where := g.Map{
		helper.StrCase("admin_id"): data["admin_id"],
		helper.StrCase("api_id"):   data["api_id"],
	}
	count, _ := dao.AdminApi.TX(tx).Where(where).Count()
	if count <= 0 {
		_, err := dao.AdminApi.Filter().Data(data).Save()
		if err != nil {
			return 0, err
		}
		_, err = dao.Admin.TX(tx).Where("id", data["admin_id"]).Update(g.Map{
			"is_admin": 1,
		})
		if err != nil {
			return 0, err
		}

	} else {
		_, err := dao.AdminApi.Filter().Where(where).Data(data).Update()
		if err != nil {
			return 0, err
		}
	}
	return 1, nil
}
