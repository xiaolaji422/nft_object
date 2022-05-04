package validate

//接口管理验证器

// 添加验证结构体
type ApiAddValidate struct {
	Name    string `v:"required#名称不能为空"`
	Route   string `v:"required#路由不能为空"`
	Enabled string `v:"required#状态不能为空"`
	GroupId string `v:"required#分组不能为空"`
	Limit   string `v:"min:0|max:20 # 限频最小为0|限频每秒最大值是20"`
	Methods string `v:"required#分组不能为空"`
}

// 修改验证结构体
type ApiEditValidate struct {
	Id string `v:"required#id不能为空"`
	ApiAddValidate
}
