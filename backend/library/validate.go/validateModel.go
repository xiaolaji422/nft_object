package validate

// 修改状态
type EnableValidate struct {
	Id     string `v:"required#id不能为空"`
	Status string `v:"required#id不能为空"`
}

//	Id不能为空信息
type IdValidate struct {
	Id string `v:"required#id不能为空"`
}
