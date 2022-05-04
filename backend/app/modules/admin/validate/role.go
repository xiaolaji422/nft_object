package validate

// 角色管理验证器

// 添加验证结构体
type RoleAddValidate struct {
	Name string `v:"required#名称不能为空"`
}

// 修改验证结构体
type RoleEditValidate struct {
	RoleAddValidate
	Id string `v:"required#id不能为空"`
}
