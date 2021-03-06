// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// ApiDao is the manager for logic model data accessing and custom defined data operations functions management.
type ApiDao struct {
	gmvc.M            // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      apiColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB     // DB is the raw underlying database management object.
	Table  string     // Table is the underlying table name of the DAO.
}

// ApiColumns defines and stores column names for table t_auth_api.
type apiColumns struct {
	Id           string //
	GroupId      string //
	Methods      string //
	Name         string //
	Route        string // ·
	Enabled      string //
	Limit        string //
	CreateUser   string //
	CreateTime   string //
	ModifiedUser string //
	ModifiedTime string //
}

// NewApiDao creates and returns a new DAO object for table data access.
func NewApiDao() *ApiDao {
	columns := apiColumns{
		Id:           "id",
		GroupId:      "group_id",
		Methods:      "methods",
		Name:         "name",
		Route:        "route",
		Enabled:      "enabled",
		Limit:        "limit",
		CreateUser:   "create_user",
		CreateTime:   "create_time",
		ModifiedUser: "modified_user",
		ModifiedTime: "modified_time",
	}
	return &ApiDao{
		C:     columns,
		M:     g.DB("default").Model("t_auth_api").Safe(),
		DB:    g.DB("default"),
		Table: "t_auth_api",
	}
}
