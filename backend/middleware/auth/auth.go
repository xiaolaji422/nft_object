package auth

import (
	"errors"
	"net/http"
	"nft_object/app/modules/admin/service"
	"nft_object/library/response"
	"nft_object/statusCode"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

// MiddlewareAuth 权限处理
func MiddlewareAuth(r *ghttp.Request) {
	if r.Method == "OPTIONS" {
		r.Middleware.Next()
	}
	// 判断是否登录
	response.CheckAuthSession(r)
	// 判断用户信息
	adminInfo := r.GetCtxVar(statusCode.SESSION_ADMIN_INFO).MapDeep()
	// 超级管理员
	if _, ok := adminInfo["is_admin"]; !ok || gconv.Int(adminInfo["is_admin"]) != 1 {

		// 非超管验证账户信息
		if code, err := checkApiAuth(r); err != nil {
			response.Json(r, code, err.Error())
		}
	}
	r.Middleware.Next()
}

// MiddlewareCORS 跨域处理
func MiddlewareCORS(r *ghttp.Request) {
	// 处理跨域
	options := r.Response.DefaultCORSOptions()
	allowDomain := g.Config().GetArray("common.allowDomain", []string{"localhost:3002"})
	options.AllowDomain = gconv.SliceStr(allowDomain)
	options.AllowHeaders += ",x-requested-with,x-csrf-token, session_id"
	if !r.Response.CORSAllowedOrigin(options) {
		r.Response.WriteStatus(http.StatusForbidden)
		return
	}
	// 请求通过之后 设置上下文
	r.Response.CORS(options)
	r.Middleware.Next()
}

/**
 * @description  :  验证是否有接口权限
 * @param         {*ghttp.Request} r
 * @return        {*}
 * @author       : fourteen
 */
func checkApiAuth(r *ghttp.Request) (int, error) {
	var (
		apiPath = "/" + strings.Trim(r.URL.Path, "/")
	)
	//接口白名单
	whiteApi := gconv.SliceStr(g.Config().GetArray("auth.whiteApi", []string{}))
	if gstr.InArray(whiteApi, apiPath) {
		return 0, nil
	}

	// 获取path 所在的接口的信息
	apiInfo, err := service.Api.GetInfoByPath(r.Context(), apiPath)
	if err != nil {
		return statusCode.FORBIDDEN, err
	}
	// 接口不存在
	if g.IsEmpty(apiInfo) {
		return statusCode.FORBIDDEN, errors.New("接口不存在")
	}

	// 查看接口的状态
	if apiInfo.Enabled != 1 {
		return statusCode.FORBIDDEN, errors.New("接口状态不支持访问")
	}
	// 接口鉴权
	adminInfo := r.GetCtxVar(statusCode.SESSION_ADMIN_INFO).MapDeep()
	authApiList := []string{}
	if _, ok := adminInfo["apis"]; ok {
		authApiList = gconv.SliceStr(adminInfo["apis"])
	}

	apiAuth := gstr.InArray(authApiList, gconv.String(apiInfo.Id))

	if !apiAuth {
		return statusCode.FORBIDDEN_SN, errors.New("用户无权限")
	}
	// todo 接口限频
	limitRes := routeLimit(gconv.Int(adminInfo["id"]), apiInfo.Id, apiInfo.Limit) //apiInfo.Limit)

	return statusCode.FORBIDDEN_BUSY, limitRes
}

//
func routeLimit(adminId int, apiId int, limit int) error {
	if limit <= 0 {
		return nil
	}
	// 限频的key
	cacheKey := gconv.String(adminId) + "api" + gconv.String(apiId) + "T" + gconv.String(gtime.Now().Timestamp())

	// 获取缓存
	cacheNum, _ := g.Redis().DoVar("GET", cacheKey)

	setNum := 1
	if cacheNum != nil {
		setNum = gconv.Int(gconv.String(cacheNum)) + 1
	}
	if setNum > limit {
		return errors.New("访问频率太高啦")
	}
	// 批量处理
	c := g.Redis().Conn()
	defer c.Close()
	c.Send("SET", cacheKey, gconv.Int(setNum))
	c.Send("EXPIRE", cacheKey, time.Second)
	c.Flush()
	c.Receive()
	// 当前时间戳
	return nil
}
