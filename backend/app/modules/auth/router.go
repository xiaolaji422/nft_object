package auth

// 权限验证路由
import (
	"nft_object/app/modules/auth/api"
	"nft_object/middleware/auth"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// init 统一路由注册.
func init() {
	s := g.Server()
	// 采用驼峰命名方式
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Group("/auth", func(group *ghttp.RouterGroup) {
			// 中间件
			group.GET("/userInfo", api.Login.UserInfo)
			group.POST("/register", api.Login.Register)
			group.POST("/loginOut", api.Login.LoginOut)
			group.POST("/login", api.Login.Login)
		})
		group.Group("/authUser", func(group *ghttp.RouterGroup) {
			// 中间件
			group.Middleware(auth.MiddlewareAuth)
			group.GET("/userInfo", api.Login.UserInfo) //获取登录信息
		})
	})
}
