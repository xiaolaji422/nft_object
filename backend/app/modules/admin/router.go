package admin

//admin模块路由管理

import (
	"nft_object/app/modules/admin/api"
	"nft_object/middleware/auth"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// init 统一路由注册.
func init() {
	s := g.Server()
	// 采用驼峰命名方式
	s.SetNameToUriType(ghttp.URI_TYPE_CAMEL)
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		// 中间件
		group.Middleware(auth.MiddlewareAuth)
		group.Group("/role", func(group *ghttp.RouterGroup) {
			// 角色路由
			_api := api.Role
			group.GET("/getRoleApis", _api.GetRoleApis)  // 角色权限列表
			group.POST("/setRoleApis", _api.SetRoleApis) // 设置角色权限
			group.GET("/items", _api.Items)              // 获取列表/分页
			group.GET("/all", _api.RoleAll)              // 获取列表/不分页
			group.POST("/add", _api.Add)                 // 添加
			group.GET("/info", _api.Detail)              // 详情
			group.POST("/update", _api.Edit)             // 修改
		})
		group.Group("/admin", func(group *ghttp.RouterGroup) {
			// 账户路由
			_api := api.Admin
			group.GET("/items", _api.Items)            // 列表(分页)
			group.GET("/info", _api.Detail)            // 详情
			group.GET("/roles", _api.Roles)            // 用户角s色
			group.GET("/apis", _api.AdminApis)         // 用户授权接口
			group.GET("/apiArray", _api.AdminApiArray) // 用户定制权限
		})
		group.Group("/api", func(group *ghttp.RouterGroup) {
			// 接口路由
			_api := api.Api
			group.GET("/all", _api.ApiAll)     // 列表(所有)
			group.GET("/items", _api.Items)    // 列表/分页
			group.GET("/info", _api.Detail)    // 详情
			group.POST("/add", _api.Add)       // 添加
			group.POST("/edit", _api.Edit)     // 修改
			group.POST("/enable", _api.Enable) // 启用、禁用
		})
		group.Group("/apiGroup", func(group *ghttp.RouterGroup) {
			// 接口分组路由
			_api := api.ApiGroup
			group.GET("/all", _api.GroupAll)    // 列表(所有)
			group.GET("/items", _api.Items)     // 列表/分页
			group.POST("/add", _api.Add)        // 添加
			group.GET("/info", _api.Detail)     // 详情
			group.POST("/edit", _api.Edit)      // 修改
			group.POST("/enabled", _api.Enable) // 启用禁用
		})
		group.Group("/adminRole", func(group *ghttp.RouterGroup) {
			// 接口分组路由
			_api := api.AdminRole
			group.POST("/enable", _api.EnableRecord) // 启用/禁用
			group.POST("/addRole", _api.AddRole)     // 添加用户角色
			group.POST("/delRole", _api.DelRole)     // 删除用户角色
		})
		group.Group("/adminApi", func(group *ghttp.RouterGroup) {
			// 接口分组路由
			_api := api.AdminApi
			group.POST("/enable", _api.EnableRecord) // 启用/禁用
			group.POST("/addApis", _api.AddApis)     // 添加用户接口

		})
	})
}
