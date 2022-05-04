package service

import (
	"nft_object/app/dao"
	"nft_object/app/model"
)

// 角色权限
var RoleApi = roleApiService{}

type roleApiService struct{}

// 获取角色api集合
func (s *roleApiService) RoleApiList() ([]*model.RoleApi, error) {
	result, err := dao.RoleApi.FindAll()
	if err != nil {
		return nil, err
	}
	res, err := dao.RoleApi.Strcuts(result)
	return res, err
}
