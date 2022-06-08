package core

import (
	"errors"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/util/gconv"
)

type MapI map[string]interface{}

// db 封装类
type DBProxy struct {
	D          *gdb.Model // db  model
	Table      MapI
	PrimaryKey string // 检索主键
}

// 获取数据库主键
func (dbproxy *DBProxy) GetDB() *gdb.Model {
	return dbproxy.D.Clone()
}

//
func (dbproxy *DBProxy) CheckInfo(val interface{}, keys ...string) (gdb.Record, error) {
	var key = "id"
	if len(keys) > 0 {
		key = keys[0]
	}
	res, err := dbproxy.GetDB().Where(key, val).FindOne()
	if err != nil {
		return res, err
	}
	if res.IsEmpty() {
		return res, errors.New("未找到记录：" + gconv.String(val))
	}
	return res, err
}
