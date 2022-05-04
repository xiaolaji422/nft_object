package core

import "github.com/gogf/gf/frame/g"

/*
 * @Description  : api的基类
 * @Autor        : fourteen
 * @Date         : 2021-05-12 15:13:40
 */
type QueryParams struct {
	Order   g.Map  `json:"_order"`
	WhereOr g.Map  `json:"_whereOr"`
	Having  g.Map  `json:"_having"`
	Page    string `d:"1" json:"page"`
	Limit   string `d:"50" json:"limit"`
}

// api 验证器基类
type BaseApi struct {
	IdValidate     *idValidate
	EnableValidate *enableValidate
}

//启用\禁用验证器
type enableValidate struct {
	Id     string `v:"required#id不能为空"`
	Status string `v:"required#id不能为空"`
}

// Id验证器
type idValidate struct {
	Id string `v:"required#id不能为空"`
}
