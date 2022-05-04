package validate

// 接口分组验证器

// 添加验证结构体
type ApiGroupAddValidate struct {
	Name   string `v:"required#名称不能为空"`
	Status int    `v:"required#状态不能为空"`
}

// 修改验证结构体
type ApiGroupEditValidate struct {
	Id int `v:"required#id不能为空"`
}
