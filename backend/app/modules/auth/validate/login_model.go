package validate

// 权限验证验证器

// code 登录验证
type LoginByCodeParams struct {
	Code string `v:"required#code不能为空"`
}

//
type ApiReqMap struct {
	Code string
}
