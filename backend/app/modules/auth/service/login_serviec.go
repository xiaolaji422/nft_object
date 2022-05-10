package service

import (
	"context"
	"errors"
	adminModel "nft_object/app/model"
	"nft_object/app/modules/admin/service"
	"nft_object/statusCode"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Login = loginService{}

type loginService struct{}

//  账号密码
func (s *loginService) Login(ctx context.Context, login_name, password string) (map[string]interface{}, error) {
	// 此处应该根据code去获取用户身份信息
	// 假设获取到了
	admin, err := service.Admin.InfoByLoginName(ctx, login_name)
	if err != nil {
		return nil, err
	}
	if admin == nil || admin.Enabled != 1 {
		return nil, errors.New("账户不存在")
	}

	if admin.Password != password {
		return nil, errors.New("密码错误")
	}

	// todo  判断用户当前状态
	apis, err := HandleLoginBate(ctx, admin)
	var adminInfo = gconv.Map(admin)
	adminInfo["apis"] = apis
	return adminInfo, err
}

//登出
func (s *loginService) LoginOut(r *ghttp.Request) error {
	err := r.Session.Clear()
	if err != nil {
		return err
	}
	return nil
}

//用户信息
func (s *loginService) UserInfo(r *ghttp.Request) g.Map {

	mapData := gconv.Map(r.Session.Get(statusCode.SESSION_ADMIN_INFO))
	// 指定返回信息
	res := g.Map{
		"login_name": mapData["login_name"],
	}
	return res
}

/**
* @description  : 根据code 获取到用户信息

* @todo 		 : 请求code接口 换取用户信息
* @todo 		 : 验证用户是否存在  状态是否允许登录
 */
func getUserByCode(code string) (map[string]interface{}, error) {
	return nil, nil
}

/**
 * @description  : 用户登录涉及到的操作
 * @todo 		 : 登录日志
 */
func HandleLogin(admin *adminModel.Admin, r *ghttp.Request) (map[string]interface{}, error) {
	resMap := gconv.Map(admin)

	// 获取用户的角色列表
	apiList, err := service.Admin.AdminAllApi(r.Context(), admin.Id)
	if err != nil {
		return resMap, err
	}
	// 合并两个
	// 拿到当前用户的api list  设置进session
	resMap["apiList"] = apiList
	// 设置进入session
	r.Session.Set(statusCode.SESSION_ADMIN_INFO, resMap)
	return resMap, nil
}

/**
 * @description  : 用户登录涉及到的操作
 * @todo 		 : 登录日志
 */
func HandleLoginBate(ctx context.Context, admin *adminModel.Admin) ([]int, error) {
	// // 获取用户的角色列表
	apiList, err := service.Admin.AdminAllApi(ctx, admin.Id)
	return apiList, err
}
