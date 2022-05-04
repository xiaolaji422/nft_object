// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"nft_object/app/dao/internal"
	"nft_object/app/model"

	"github.com/gogf/gf/database/gdb"
)

// roleApiDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type roleApiDao struct {
	*internal.RoleApiDao
}

var (
	// 对外对象
	RoleApi roleApiDao
)

func init() {
	RoleApi = roleApiDao{
		internal.NewRoleApiDao(),
	}
}

// 处理成Struct
func (m *roleApiDao) Strcut(res gdb.Record) (info *model.RoleApi, err error) {
	if err = res.Struct(&info); err != nil {
		return nil, err
	}
	return
}

// 处理成数组
func (d *roleApiDao) Strcuts(res gdb.Result) (list []*model.RoleApi, err error) {
	if err = res.Structs(&list); err != nil {
		return nil, err
	}
	return
}
