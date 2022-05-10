// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"nft_object/app/core"
	"nft_object/app/dao/internal"
	"nft_object/app/model"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/util/gconv"
)

// apiDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type apiDao struct {
	*internal.ApiDao
}

var (
	// 对外对象
	Api apiDao
)

func init() {
	Api = apiDao{
		internal.NewApiDao(),
	}
}

// 处理成Struct
func (m *apiDao) Strcut(res gdb.Record) (info *model.Api, err error) {
	if err = res.Struct(&info); err != nil {
		return nil, err
	}
	return
}

// 处理成数组
func (d *apiDao) Strcuts(res gdb.Result) (list []*model.Api, err error) {
	if err = res.Structs(&list); err != nil {
		return nil, err
	}
	return
}

// 分页查询
func (d *apiDao) Pagenation(params map[string]interface{}) (core.PageReult, error) {
	// 分页信息
	res, err := core.BaseDao.PagenationCommon(core.GdbModel{
		M:     d.M,
		DB:    d.DB,
		Table: d.Table,
		Where: map[string]core.WhereItem{
			"id":       {Operator: "="},
			"group_id": {Operator: "="},
		},
		Fields:  []string{"id", "route", "group_id"},
		Columns: gconv.Map(model.Api{}),
		Order: []core.OrderItem{
			{Property: "group_id", Sort: "asc"},
			{Property: "id", Sort: "asc"},
		},
	}, params)
	if err != nil {
		return res, err
	}
	return res, nil
}