package system

import (
	"nft_object/app/modules/system/api"
	"nft_object/middleware/auth"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// init 统一路由注册.
func init() {
	s := g.Server()
	// 采用驼峰命名方式
	s.SetNameToUriType(ghttp.URI_TYPE_CAMEL)
	s.Group("/system", func(group *ghttp.RouterGroup) {
		group.Middleware(auth.MiddlewareAuth)
		group.GET("/getVersion", api.System.GetVersion)
		group.POST("/uploadImg", api.Upload.UpImg)
		group.POST("/uploadImgs", api.Upload.UpImgs)
	})
}
