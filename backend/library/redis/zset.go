package redis

import (
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/frame/g"
)

type IZset interface {
	// 升序分页列表
	Zrangebyscore(keys string, min, max float64, page, limit int) (*gvar.Var, error)
	// 降序分页列表
	Zrevrangebyscore(keys string, min, max float64, page, limit int) (*gvar.Var, error)
}

var (
	Zset = func() IZset {
		return &zset{}
	}
)

type zset struct{}

// 升序分页列表
func (x *zset) Zrangebyscore(keys string, min, max float64, page, limit int) (*gvar.Var, error) {
	if page <= 0 {
		page = 1
	}
	start_limit := (page - 1) * limit
	end_limit := page * limit
	return g.Redis().DoVar("zrangebyscore", keys, min, max, "limit", start_limit, end_limit)
}

// 降序分页列表
func (x *zset) Zrevrangebyscore(keys string, min, max float64, page, limit int) (*gvar.Var, error) {
	if page <= 0 {
		page = 1
	}
	start_limit := (page - 1) * limit
	end_limit := page * limit
	return g.Redis().DoVar("zrevrangebyscore", keys, max, min, "limit", start_limit, end_limit)
}
