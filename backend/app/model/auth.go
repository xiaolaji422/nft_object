package model

import "github.com/gogf/gf/os/gtime"

// 管理员账户
type Admin struct {
	Id           int         `orm:"id,primary"        json:"id"`            //
	LoginName    string      `orm:"login_name,unique" json:"login_name"`    // 登录账户
	RoleName     string      `orm:"role_name"         json:"role_name"`     // 用户角色
	Password     string      `orm:"password"          json:"password"`      // 密码
	Enabled      int         `orm:"enabled"           json:"enabled"`       // 状态
	Level        int         `orm:"level"             json:"level"`         // 层级
	IsAdmin      int         `orm:"is_admin"          json:"is_admin"`      // 是否是超级管理员
	ModifiedUser string      `orm:"modified_user"     json:"modified_user"` // 配置人
	ModifiedTime *gtime.Time `orm:"modified_time"     json:"modified_time"` // 更新时间
	CreateTime   *gtime.Time `orm:"create_time"       json:"create_time"`   // 创建时间
}

// Role is the golang structure for table t_auth_role.
type Role struct {
	Id           int         `orm:"id,primary"    json:"id"`            //
	Name         string      `orm:"name"          json:"name"`          //
	Enabled      int         `orm:"enabled"       json:"enabled"`       //
	CreateUser   string      `orm:"create_user"   json:"create_user"`   //
	CreateTime   *gtime.Time `orm:"create_time"   json:"create_time"`   //
	ModifiedUser string      `orm:"modified_user" json:"modified_user"` //
	ModifiedTime *gtime.Time `orm:"modified_time" json:"modified_time"` //
}

// Api is the golang structure for table t_auth_api.
type Api struct {
	Id           int         `orm:"id,primary"    json:"id"`            //
	GroupId      int         `orm:"group_id"      json:"group_id"`      //
	Methods      string      `orm:"methods"       json:"methods"`       //
	Name         string      `orm:"name"          json:"name"`          //
	Route        string      `orm:"route"         json:"route"`         // ·
	Enabled      int         `orm:"enabled"       json:"enabled"`       //
	Limit        int         `orm:"limit"         json:"limit"`         //
	CreateUser   string      `orm:"create_user"   json:"create_user"`   //
	CreateTime   *gtime.Time `orm:"create_time"   json:"create_time"`   //
	ModifiedUser string      `orm:"modified_user" json:"modified_user"` //
	ModifiedTime *gtime.Time `orm:"modified_time" json:"modified_time"` //
}

// AdminApi is the golang structure for table t_auth_admin_api.
type AdminApi struct {
	Id           int         `orm:"id,primary"    json:"id"`            //
	AdminId      int         `orm:"admin_id"      json:"admin_id"`      //
	ApiId        int         `orm:"api_id"        json:"api_id"`        //
	ApiName      string      `orm:"api_name"      json:"api_name"`      //
	Enabled      int         `orm:"enabled"       json:"enabled"`       // ״̬
	CreateUser   string      `orm:"create_user"   json:"create_user"`   //
	CreateTime   *gtime.Time `orm:"create_time"   json:"create_time"`   //
	ModifiedUser string      `orm:"modified_user" json:"modified_user"` //
	ModifiedTime *gtime.Time `orm:"modified_time" json:"modified_time"` //
}

// RoleApi is the golang structure for table t_auth_role_api.
type RoleApi struct {
	Id           int         `orm:"id,primary"     json:"id"`            //
	RoleId       int         `orm:"role_id,unique" json:"role_id"`       //
	Apis         string      `orm:"apis"           json:"apis"`          // ·
	Enabled      int         `orm:"enabled"        json:"enabled"`       //
	CreateUser   string      `orm:"create_user"    json:"create_user"`   //
	CreateTime   *gtime.Time `orm:"create_time"    json:"create_time"`   //
	ModifiedUser string      `orm:"modified_user"  json:"modified_user"` //
	ModifiedTime *gtime.Time `orm:"modified_time"  json:"modified_time"` //
}

// AdminRole is the golang structure for table t_auth_admin_role.
type AdminRole struct {
	Id           int         `orm:"id,primary"    json:"id"`            //
	AdminId      int         `orm:"admin_id"      json:"admin_id"`      //
	RoleId       int         `orm:"role_id"       json:"role_id"`       //
	RoleName     string      `orm:"role_name"     json:"role_name"`     //
	Enabled      int         `orm:"enabled"       json:"enabled"`       // ״̬
	CreateUser   string      `orm:"create_user"   json:"create_user"`   //
	CreateTime   *gtime.Time `orm:"create_time"   json:"create_time"`   //
	ModifiedUser string      `orm:"modified_user" json:"modified_user"` //
	ModifiedTime *gtime.Time `orm:"modified_time" json:"modified_time"` //
}

// ApiGroup is the golang structure for table t_auth_api_group.
type ApiGroup struct {
	Id           int         `orm:"id,primary"    json:"id"`            //
	Name         string      `orm:"name"          json:"name"`          //
	Enabled      int         `orm:"enabled"       json:"enabled"`       //
	CreateUser   string      `orm:"create_user"   json:"create_user"`   //
	CreateTime   *gtime.Time `orm:"create_time"   json:"create_time"`   //
	ModifiedUser string      `orm:"modified_user" json:"modified_user"` //
	ModifiedTime *gtime.Time `orm:"modified_time" json:"modified_time"` //
}

// NoticeWarning is the golang structure for table t_notice_warning.
type NoticeWarning struct {
	Id          int         `orm:"id,primary"  json:"id"`          // ID
	Platform    int         `orm:"platform"    json:"platform"`    // 公告平台
	OriginalId  int         `orm:"original_id" json:"original_id"` // 原始的公告id
	Name        string      `orm:"name"        json:"name"`        //
	Description string      `orm:"description" json:"description"` // 公告告警描述
	Content     string      `orm:"content"     json:"content"`     //
	NoticeUrl   string      `orm:"notice_url"  json:"notice_url"`  //
	NoticeTime  *gtime.Time `orm:"notice_time" json:"notice_time"` // 公告创建时间
	CreateTime  *gtime.Time `orm:"create_time" json:"create_time"` // 创建时间
	UpdateTime  *gtime.Time `orm:"update_time" json:"update_time"` // 更新时间
	Enabled     int         `orm:"enabled"     json:"enabled"`     // 状态描述, 1: 正常使用，0: 删除
}
