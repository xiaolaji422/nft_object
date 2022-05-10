package core

import (
	"github.com/gogf/gf/database/gdb"
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
