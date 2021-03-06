// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"nft_object/app/dao/internal"
	"nft_object/app/model"

	"github.com/gogf/gf/database/gdb"
)

// adminRoleDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type adminRoleDao struct {
	*internal.AdminRoleDao
}

var (
	// 对外对象
	AdminRole adminRoleDao
)

func init() {
	AdminRole = adminRoleDao{
		internal.NewAdminRoleDao(),
	}
}

// Fill with you ideas below.

// 处理成Struct
func (m *adminRoleDao) Strcut(res gdb.Record) (info *model.AdminRole, err error) {
	if err = res.Struct(&info); err != nil {
		return nil, err
	}
	return
}

// 处理成数组
func (d *adminRoleDao) Strcuts(res gdb.Result) (list []*model.AdminRole, err error) {
	if err = res.Structs(&list); err != nil {
		return nil, err
	}
	return
}
