package core

import (
	"context"
	"nft_object/library/runtime"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// 公共Api基类
type CoreApi struct {
	CheckRules CheckRule
}

// 校验规则
type CheckRule map[string][]string

// 校验规则
func (c *CoreApi) CheckParams(r *ghttp.Request, scene ...string) (context.Context, g.Map, error) {
	var (
		params = r.GetMap()
		ctx    = r.Context()
		err    error
		keys   = "default"
		rules  = []string{}
	)
	if len(scene) == 1 {
		keys = scene[0]
	} else {
		keys = runtime.GetFuntName(2)
	}

	if v, ok := c.CheckRules[keys]; ok {
		rules = v
	}
	if len(rules) > 0 {
		err = gvalid.CheckMap(ctx, params, rules)
	}
	return ctx, params, err
}
