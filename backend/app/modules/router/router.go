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
