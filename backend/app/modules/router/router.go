package router

// 权限验证路由
import (
	"nft_object/app/modules/api"
	"nft_object/middleware/auth"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// init 统一路由注册.
func init() {
	s := g.Server()

	// 采用驼峰命名方式
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Group("/account_lock_free", func(group *ghttp.RouterGroup) {
			var account_lock_api = api.AccountLockApi()
			// 中间件
			group.GET("/lock", account_lock_api.Lock)
		})
		group.Group("/notice", func(group *ghttp.RouterGroup) {
			var notice_api = api.NoticeApi()
			group.Middleware(auth.MiddlewareAuth)
			// 中间件
			group.GET("/query_notice", notice_api.QueryNotice)
			group.GET("/query_his_notice", notice_api.QueryHistoryNotice)
		})

		group.Group("/album", func(group *ghttp.RouterGroup) {
			var notice_api = api.NFTAlbumApi()
			group.Middleware(auth.MiddlewareAuth)
			// 中间件
			group.GET("/search", notice_api.Search)
			group.GET("/price_list", notice_api.PriceList)
		})
		group.Group("/account", func(group *ghttp.RouterGroup) {
			var account_api = api.AccountApi()
			group.Middleware(auth.MiddlewareAuth)
			// 中间件
			group.GET("/list", account_api.List)
			group.POST("/save", account_api.Save)
		})

		group.Group("/account_lock", func(group *ghttp.RouterGroup) {
			var account_lock_api = api.AccountLockApi()
			group.Middleware(auth.MiddlewareAuth)
			// 中间件
			group.GET("/list", account_lock_api.List)
			group.POST("/save", account_lock_api.Add)
			group.POST("/cancel", account_lock_api.Cancel)
		})
	})

	// 消息推送
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Group("/send_msg", func(group *ghttp.RouterGroup) {
			var send_msg_api = api.SendMsgApi()
			// group.Middleware(auth.MiddlewareAuth)
			// 中间件
			group.GET("/send_msg", send_msg_api.SendMsg)
		})
	})

}
