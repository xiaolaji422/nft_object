package service

import (
	"context"
	"fmt"
	"nft_object/app/core"
	"nft_object/app/dao"
	"nft_object/app/model"
	"nft_object/library/helper"
	"strings"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/pkg/errors"
)

// 角色权限业务处理
var AdminRole = adminRoleService{}

type adminRoleService struct{}

// 角色api集合（id）
func (s *adminRoleService) AdminRoleList(ctx context.Context, adminId int) ([]*gvar.Var, error) {
	res, err := dao.AdminRole.Fields("role_id").Where(g.Map{"admin_id": adminId, "enabled": 1}).Array()
	if err != nil {
		return nil, err
	}
	return res, err
}

//获取角色api详情
func (s *adminRoleService) AdminRoleListInfo(ctx context.Context, adminId int) (gdb.Result, error) {
	res, err := dao.AdminRole.Where(g.Map{"admin_id": adminId}).FindAll()
	if err != nil {
		return nil, err
	}
	return res, err
}

// 获取角色下的账户id
func (s *adminRoleService) GetAdminIdsByRoleId(ctx context.Context, params g.Map) ([]int, error) {
	where := []core.WhereItem{{Where: "role_id"}, {Where: "enabled", Value: 1}}
	res, err := dao.AdminRole.Fields("admin_id").Where(where, params).Array()
	if err != nil {
		return nil, err
	}
	reeData := gconv.SliceInt(res)
	return reeData, err
}

// 批量/添加用户角色
func (s *adminRoleService) AddAdminRole(ctx context.Context, adminIds []string, roleId int) (int, error) {
	// 获取角色的信息
	role, err := Role.Info(ctx, roleId)
	if err != nil {
		return 0, err
	}
	if g.IsEmpty(role) {
		err = errors.New("角色不存在")
		return 0, err
	}
	// 成功数量
	update_datas := make([]g.Map, 0, 20)
	if len(adminIds) > 0 {
		for _, v := range adminIds {
			item := g.Map{
				"role_id":       roleId,
				"role_name":     role.Name,
				"admin_id":      gconv.Int(v),
				"modified_user": helper.GetRtx(ctx),
				"enabled":       1,
			}
			update_datas = append(update_datas, item)
		}
	}
	res, err := dao.AdminRole.Save(update_datas)
	successNum, err := core.BaseDao.HandleExecRes(res, err, "新增权限")
	if err != nil {
		return 0, err
	}
	//   更新用户名
	for _, v := range adminIds {
		_, err := s.updateRoleName(ctx, nil, gconv.Int(v))
		if err != nil {
			fmt.Println("更新用户角色失败：" + err.Error())
			successNum--
		}
	}
	return successNum, err
}

//  删除用户角色
func (s *adminRoleService) DelAdminRole(ctx context.Context, id int) (rows int, err error) {
	authInfo, err := s.CheckPrikey(ctx, id)
	if err != nil {
		return 0, err
	}
	// 同步 数据库
	_, err = dao.AdminRole.Where("id", id).Delete(g.Map{"id": id})

	if err != nil {
		return 0, err
	}

	count, err := s.updateRoleName(ctx, nil, authInfo.AdminId)
	return count, err
}

//  根据Id查找数据
func (s *adminRoleService) CheckPrikey(ctx context.Context, id int) (*model.AdminRole, error) {
	where := g.Map{
		"id": id,
	}
	//查找数据
	sqlRes, err := dao.AdminRole.Where(where).FindOne()
	if err != nil {
		return nil, err
	}
	info, err := dao.AdminRole.Strcut(sqlRes)
	if err != nil {
		return nil, err
	}
	if g.IsEmpty(info) {
		return nil, errors.New("记录不存在或已被删除")
	}
	return info, err
}

//启用/禁用
func (s *adminRoleService) EnableRecord(ctx context.Context, id int, enabled int) (int, error) {
	info, err := s.CheckPrikey(ctx, id)
	if err != nil {
		return 0, err
	}
	res, err := dao.AdminRole.Where("id", id).Update(g.Map{
		"enabled":       enabled,
		"moditied_user": helper.GetRtx(ctx),
	})

	_, err = core.BaseDao.HandleExecRes(res, err, "启用/禁用")
	if err != nil {
		return 0, err
	}
	count, err := s.updateRoleName(ctx, nil, info.AdminId)
	return count, err
}

// 处理添加的相关业务
func (s *adminRoleService) HandleAdd(ctx context.Context, tx *gdb.TX, data g.Map) (rows int, err error) {
	// 首先查找有没有
	count, _ := dao.AdminRole.Where(data).Count()

	if count <= 0 {
		_, err = dao.AdminRole.Filter().TX(tx).Data(data).Save()
		if err != nil {
			return
		}
	} else {
		_, err = dao.AdminRole.TX(tx).Where(data).Update(g.Map{
			"enabled": 1,
		})
		if err != nil {
			return
		}
	}
	_, err = s.updateRoleName(ctx, tx, gconv.Int(data["admin_id"]))
	return 1, err
}

func (s *adminRoleService) updateRoleName(ctx context.Context, tx *gdb.TX, admin_id int) (rows int, err error) {
	// 首先查找有没有
	listWhere := g.Map{
		"admin_id": admin_id,
		"enabled":  1,
	}
	sqlRes, err := dao.AdminRole.Fields("role_name").Where(listWhere).All()
	if err != nil {
		return
	}
	role_names := make([]string, 0, 6)
	//  更新用户名称
	if !g.IsEmpty(sqlRes) {
		for _, v := range sqlRes {
			role_names = append(role_names, v["role_name"].String())
		}
	}
	// 更新用户角色
	updateData := g.Map{
		"role_name":     strings.Join(role_names, ","),
		"modified_user": helper.GetRtx(ctx),
		"is_admin":      1,
	}
	dm := dao.Admin.Where("id", admin_id)
	if tx != nil {
		dm = dm.TX(tx)
	}
	_, err = dm.Data(updateData).Update()
	if err != nil {
		return
	}
	return 1, err
}
