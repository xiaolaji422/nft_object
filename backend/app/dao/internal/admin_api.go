// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// AdminApiDao is the manager for logic model data accessing and custom defined data operations functions management.
type AdminApiDao struct {
	gmvc.M                 // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      adminApiColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB          // DB is the raw underlying database management object.
	Table  string          // Table is the underlying table name of the DAO.
}

// AdminApiColumns defines and stores column names for table t_auth_admin_api.
type adminApiColumns struct {
	Id           string //
	AdminId      string //
	ApiId        string //
	ApiName      string //
	Enabled      string // ״̬
	CreateUser   string //
	CreateTime   string //
	ModifiedUser string //
	ModifiedTime string //
}

// NewAdminApiDao creates and returns a new DAO object for table data access.
func NewAdminApiDao() *AdminApiDao {
	columns := adminApiColumns{
		Id:           "id",
		AdminId:      "admin_id",
		ApiId:        "api_id",
		ApiName:      "api_name",
		Enabled:      "enabled",
		CreateUser:   "create_user",
		CreateTime:   "create_time",
		ModifiedUser: "modified_user",
		ModifiedTime: "modified_time",
	}
	return &AdminApiDao{
		C:     columns,
		M:     g.DB("default").Model("t_auth_admin_api").Safe(),
		DB:    g.DB("default"),
		Table: "t_auth_admin_api",
	}
}
