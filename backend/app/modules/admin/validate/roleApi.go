package validate

// 角色权限验证器

// 添加验证结构体
type RoleApiAddValidate struct {
	RoleId string `v:"required#角色id不能为空"`
	Apis   string `v:"required#权限列表不能为空"`
}

// 修改验证结构体
type RoleApiEditValidate struct {
	RoleApiAddValidate
	Id string `v:"required#id不能为空"`
}
