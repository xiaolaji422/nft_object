// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// NoticeWarningDao is the manager for logic model data accessing and custom defined data operations functions management.
type NoticeWarningDao struct {
	gmvc.M                      // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      noticeWarningColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB               // DB is the raw underlying database management object.
	Table  string               // Table is the underlying table name of the DAO.
}

// NoticeWarningColumns defines and stores column names for table t_notice_warning.
type noticeWarningColumns struct {
	Id          string // ID
	Platform    string // 公告平台
	OriginalId  string // 原始的公告id
	Name        string //
	Description string // 公告告警描述
	Content     string //
	NoticeUrl   string //
	NoticeTime  string // 公告创建时间
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
	Enabled     string // 状态描述, 1: 正常使用，0: 删除
}

// NewNoticeWarningDao creates and returns a new DAO object for table data access.
func NewNoticeWarningDao() *NoticeWarningDao {
	columns := noticeWarningColumns{
		Id:          "id",
		Platform:    "platform",
		OriginalId:  "original_id",
		Name:        "name",
		Description: "description",
		Content:     "content",
		NoticeUrl:   "notice_url",
		NoticeTime:  "notice_time",
		CreateTime:  "create_time",
		UpdateTime:  "update_time",
		Enabled:     "enabled",
	}
	return &NoticeWarningDao{
		C:     columns,
		M:     g.DB("default").Model("t_notice_warning").Safe(),
		DB:    g.DB("default"),
		Table: "t_notice_warning",
	}
}
