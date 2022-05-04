package router

import (
	// 引入具体模块的router 定义

	"net/http"
	_ "nft_object/app/modules"
	"nft_object/library/logge"
	"nft_object/library/throttle"
	"nft_object/middleware/auth"
	"nft_object/statusCode"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

func init() {
	s := g.Server()
	// 设置默认路由
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.RedirectTo("/")
	})

	s.BindMiddleware("/*", auth.MiddlewareCORS)
	// 自定义记录访问日志
	s.BindHookHandlerByMap("/*any", map[string]ghttp.HandlerFunc{
		// 服务开始前钩子  ip限频
		ghttp.HookBeforeServe: func(r *ghttp.Request) {
			// 根据ip访问限流
			if r.Method != "OPTIONS" && r.RequestURI != "/favicon.ico" {
				capacity := g.Config().GetInt("common.accessCapacity")
				seconds := g.Config().GetInt("common.accessSeconds")
				if capacity > 0 && seconds > 0 {
					if !throttle.New().Check(r.GetClientIp(), capacity, seconds) {
						r.Response.WriteStatus(http.StatusTooManyRequests, "No Access!!!")
						r.ExitAll()
					}
				}
			}
		},
		// 输出前的钩子  记录日志
		ghttp.HookAfterOutput: func(r *ghttp.Request) {
			response := ""
			if r.Method == "POST" {
				// 最多记录输出信息的600字符
				response = gstr.SubStr(gconv.String(r.Response.Buffer()), 0, 600)
			}
			logge.Write("access_"+gconv.String(r.Response.Status), "info", map[string]interface{}{
				"url":      r.GetUrl(),
				"params":   r.GetRequestMap(),
				"header":   r.Header,
				"response": response,
			}, r.GetClientIp(), r.GetCtxVar(statusCode.SESSION_CACHE_ADMIN_NAME, "admin"), r.Response.Status)
		},
	})
}
