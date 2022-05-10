package pagenation

import (
	"math"
	"nft_object/app/core"
	"nft_object/library/helper"
	"strings"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/util/gconv"
)

// dao的基类
type EmptyValue struct{}

// 数据库操作助手类
const (
	PAGE_PAGE  = "page"
	PAGE_LIMIT = "limit"

	MAX_LIMIT = 200 // 分页最大行数

	//  查询规则
	OP_EQ      = "="
	OP_NEQ     = "!="
	OP_LTE     = ">="
	OP_GTE     = "<="
	OP_GT      = "<"
	OP_LT      = ">"
	OP_BETWEEN = "between"
	OP_LIKE    = "like"
	OP_LIKEALL = "likeall"
	OP_IN      = "in"
	OP_NULL    = "null"
	OP_NOTNULL = "notnull"
	OP_EMPTY   = "empty"
)

// 数据库操作的类
type Pagenation struct {
	M        *gdb.Model
	params   core.MapI
	columns  core.MapI
	limit    int              // 行数
	page     int              // 分页
	rules    map[string]*Rule // 查询规则  eq 等于  lt >
	fieldsEx []string         // 指定不查出字段 eg: content
	fields   []string         // 指定查出字段 eg：id
	empty    []string         // 允许为空的字段	eg："name"
	order    []string         // 排序字段	eg："create_tiem desc"
	where    core.MapI        // 查询字段	eg：""
	whereOr  core.MapI        // or查询字段	eg：""
}

// 公共分页返回的信息
var GetPagenation = func(M *gdb.Model) *Pagenation {
	return &Pagenation{M: M}
}

type PageReult struct {
	Total       int        `json:"total"`
	CurrentPage int        `json:"current_page"`
	PerPage     int        `json:"per_page"`
	Pages       float64    `json:"pages"`
	Items       gdb.Result `json:"items"`
}

// 处理子查询

// 公共分页信息
func (p *Pagenation) Search(params core.MapI) (pageReult PageReult, err error) {
	// 特殊处理
	p.params = params
	dm := p.M
	// 一顿操作
	p = p.PageInfo()
	p = p.OrderInfo()
	p, err = p.Where(params)
	if err != nil {
		return pageReult, err
	}

	pageReult.CurrentPage = p.page
	pageReult.PerPage = p.limit
	// 条件查询
	if len(p.where) > 0 {
		dm = dm.Where(p.where)
	}

	if len(p.whereOr) > 0 {
		var str []string
		var values []interface{}
		for k, v := range p.whereOr {
			str = append(str, k)
			values = append(values, v)
		}
		dm = dm.Where(strings.Join(str, " OR "), values...)
	}
	// 先查出数量   fields 之后会报错
	count, _ := dm.Count()
	// 查出关键字
	if len(p.fields) > 0 {
		dm = dm.Fields(strings.Join(p.fields, ","))
	}

	// 排除关键字
	if len(p.fieldsEx) > 0 {
		dm = dm.FieldsEx(strings.Join(p.fieldsEx, ","))
	}

	// 排序
	if len(p.order) > 0 {
		dm = dm.Order(strings.Join(p.order, ","))
	}
	all, err := dm.Limit((p.page-1)*p.limit, p.limit).FindAll()
	if err != nil {
		return pageReult, err
	}
	pageReult.Total = count
	pageReult.Pages = math.Ceil(float64(count) / float64(p.limit))
	pageReult.Items = all
	return pageReult, nil
}

// 处理查询分页条件
func (p *Pagenation) PageInfo(limit ...int) *Pagenation {
	if len(limit) > 0 {
		if limit[0] >= 1 {
			p.page = limit[0]
		}
	}
	if len(limit) > 1 {
		if limit[1] >= 1 {
			p.limit = limit[1]
		}
	}

	if p.page > 0 && p.limit > 0 {
		// 删除相关的参数
		delete(p.params, PAGE_PAGE)
		delete(p.params, PAGE_LIMIT)
		return p
	}

	if p.page == 0 {
		p.page = 1
	}

	if p.limit == 0 {
		p.limit = 10
	}

	if v, ok := p.params[PAGE_PAGE]; ok {
		p.page = gconv.Int(v)
		if p.page <= 0 {
			p.page = 1
		}
		delete(p.params, PAGE_PAGE)
	}

	if v, ok := p.params[PAGE_LIMIT]; ok {
		p.limit = gconv.Int(v)
		if p.page <= 0 {
			p.limit = 10
		}
		// 设置分页最大上限
		if p.limit > MAX_LIMIT {
			p.limit = MAX_LIMIT
		}
		delete(p.params, PAGE_LIMIT)
	}
	return p
}

// 处理查询的排序条件
type OrderItem struct {
	Property string
	Sort     string
}

// 处理查询的排序条件
func (p *Pagenation) OrderInfo(order ...OrderItem) *Pagenation {
	var orderMap []string

	//  如果传入参数中存在_order  MOrder置空
	if _, ok := p.params["_order_property"]; ok {
		proStr := gconv.String(p.params["_order_property"])
		//  重置为只有一个的order
		var sort_str = gconv.String(helper.GetMapValue(p.params, "_order_sort", "ASC"))

		delete(p.params, "_order_property")
		delete(p.params, "_order_sort")
		if strings.ToTitle(sort_str) != "DESC" {
			sort_str = "ASC"
		} else {
			sort_str = "DESC"
		}

		if p.checkColumns(proStr) {
			order = []OrderItem{
				{
					Property: proStr,
					Sort:     sort_str,
				},
			}
		} else {
			return p
		}

	}
	if len(order) == 0 {
		return p
	}
	//  组装排序规则
	for _, v := range order {
		if strings.ToTitle(v.Sort) != "DESC" {
			v.Sort = "ASC"
		}
		if p.checkColumns(v.Property) {
			sort_str := v.Property + " " + v.Sort
			orderMap = append(orderMap, sort_str)
		}
	}
	if len(orderMap) > 0 {
		p.order = append(p.order, orderMap...)
	}

	return p
}

// 处理查询指定字段
func (p *Pagenation) Fields(fields ...string) *Pagenation {
	// 置为空
	p.fields = make([]string, 0)
	if len(fields) > 0 {
		p.fields = append(p.fields, fields...)
	}
	return p
}

// 处理查询指定排除字段
func (p *Pagenation) FieldsEx(fieldsex ...string) *Pagenation {
	// 置为空
	p.fieldsEx = make([]string, 0)
	if len(fieldsex) > 0 {
		p.fieldsEx = append(p.fieldsEx, fieldsex...)
	}
	return p
}

// 设置 可以为空的字段
func (p *Pagenation) Empty(fields ...string) *Pagenation {
	// 置为空
	p.empty = make([]string, 0)
	p.empty = append(p.empty, fields...)
	return p
}
