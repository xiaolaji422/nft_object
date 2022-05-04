// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"database/sql"
	"nft_object/app/core"
	"nft_object/app/dao/internal"
	"nft_object/app/model"
	"nft_object/library/helper"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

// adminDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type adminDao struct {
	*internal.AdminDao
}

var (
	// 对外对象
	Admin adminDao
)

func init() {
	Admin = adminDao{
		internal.NewAdminDao(),
	}
}

// Fill with you ideas below.
// 处理成Struct
func (d *adminDao) Strcut(res gdb.Record) (info *model.Admin, err error) {
	if err = res.Struct(&info); err != nil {
		return nil, err
	}
	return
}

// 处理成数组
func (d *adminDao) Strcuts(res gdb.Result) (list []*model.Admin, err error) {
	if err = res.Structs(&list); err != nil {
		return nil, err
	}
	return
}

// 分页查询
func (d *adminDao) Pagenation(params map[string]interface{}) (core.PageReult, error) {
	// 分页信息
	var admin = gconv.Map(model.Admin{})

	if _, ok := admin["id"]; ok {

	}

	res, err := core.BaseDao.PagenationCommon(core.GdbModel{
		M:     d.M,
		DB:    d.DB,
		Table: d.Table,
		Order: []core.OrderItem{
			{
				Property: "id",
				Sort:     "desc",
			},
		},
		Columns: gconv.Map(model.Admin{}),
		Where: map[string]core.WhereItem{
			"admin_id": {Operator: "="},
		},
	}, params)
	if err != nil {
		return res, err
	}
	return res, nil
}

// 用户详情
func (d *adminDao) GetInfoByLoginName(loginName string) *model.Admin {
	where := g.Map{
		helper.StrCase("login_name"): loginName,
	}
	res, err := d.M.Where(where).One()
	if err != nil {
		return nil
	}
	var entity *model.Admin
	if err = res.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return entity
	}
	return entity
}
