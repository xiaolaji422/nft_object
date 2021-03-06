// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// AdminDao is the manager for logic model data accessing and custom defined data operations functions management.
type AdminDao struct {
	gmvc.M              // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      adminColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB       // DB is the raw underlying database management object.
	Table  string       // Table is the underlying table name of the DAO.
}

// AdminColumns defines and stores column names for table t_admin.
type adminColumns struct {
	Id           string //
	LoginName    string // 登录账户
	RoleName     string // 用户角色
	Password     string // 密码
	Enabled      string // 状态
	Level        string // 层级
	IsAdmin      string // 是否是超级管理员
	ModifiedUser string // 配置人
	ModifiedTime string // 更新时间
	CreateTime   string // 创建时间
}

// NewAdminDao creates and returns a new DAO object for table data access.
func NewAdminDao() *AdminDao {
	columns := adminColumns{
		Id:           "id",
		LoginName:    "login_name",
		RoleName:     "role_name",
		Password:     "password",
		Enabled:      "enabled",
		Level:        "level",
		IsAdmin:      "is_admin",
		ModifiedUser: "modified_user",
		ModifiedTime: "modified_time",
		CreateTime:   "create_time",
	}
	return &AdminDao{
		C:     columns,
		M:     g.DB("default").Model("t_admin").Safe(),
		DB:    g.DB("default"),
		Table: "t_admin",
	}
}
