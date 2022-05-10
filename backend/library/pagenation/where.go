package pagenation

import (
	"errors"
	"nft_object/app/core"
	"nft_object/library/helper"
	"reflect"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

// 排除关键字
var keyword = []string{"page", "limit", "_order_property", "_order_sort"}

// 组装查询条件
func (p *Pagenation) Where(data core.MapI) (*Pagenation, error) {
	var (
		reqMap    = p.params
		searchMap = make([]*whereItem, 0)
		whereAdd  = make(map[string]interface{})
		whereOr   = make(map[string]interface{})
	)
	if len(data) <= 0 {
		return p, nil
	}
	// 排除非法字段
	for k, v := range reqMap {
		// 排除关键字
		if index := helper.FindStr(keyword, k); index > 0 {
			continue
		}

		// 空值排除
		if g.IsEmpty(v) && !p.checkEmpty(k) {
			continue
		}
		// 校验搜索的值是否合理
		val, err := p.checkWhereValue(v)
		if err != nil {
			return p, err
		}

		// 规则
		var rule *Rule
		if rul, ok := p.rules[k]; ok {
			rule = rul
		} else {
			rule = &Rule{
				Field:    k,
				Operator: OP_EQ,
			}
		}
		// 处理结果  处理为
		searchV := rule.getWhereItems(val)
		if err != nil {
			return p, err
		}
		searchMap = append(searchMap, searchV...)
	}

	if g.IsEmpty(searchMap) {
		return p, nil
	}
	// 组装查询map  结构：{name:	{value："123",operator:"like",where:"add"},age:	{value："12",operator:">"} }
	for _, v := range searchMap {
		// 转化为下划线的搜索
		// k = helper.StrCase(k)
		key := ""

		switch v.Operator {
		case OP_BETWEEN:

			key = v.Fields + " " + v.Operator + " ? and ?"
			// 结果解析
			if reflect.TypeOf(v.Value).Kind() == reflect.String {
				v.Value = strings.Split(gconv.String(v.Value), ",")
			}

		case "is":
			if reflect.TypeOf(v.Value).Kind() != reflect.String {
				continue
			}
			isNUll := strings.ToUpper(gconv.String(v.Value))
			if isNUll == "NULL" {
				key = v.Fields + " " + v.Operator + " NULL"
				v.Value = nil
			} else if isNUll == "NOT NULL" {
				key = v.Fields + " " + v.Operator + " NOT NULL"
				v.Value = nil
			} else {
				continue
			}
		case OP_IN:
			key = v.Fields + " " + v.Operator + " (?)"
			if reflect.TypeOf(v.Value).Kind() == reflect.String {
				v.Value = strings.Split(gconv.String(v.Value), ",")
			}
		case "notin":
			key = v.Fields + " NOT IN " + " (?)"
			if reflect.TypeOf(v).Kind() == reflect.String {
				v.Value = strings.Split(gconv.String(v.Value), ",")
			}
		case OP_LIKE:

		case OP_LIKEALL:
			v.Operator = "like"
			key = v.Fields + " " + v.Operator + " ?"
		default:
			key = v.Fields + " " + v.Operator + " ?"
		}
		value := v.Value
		if v.Where == "OR" {
			whereOr[key] = value
		} else {
			whereAdd[key] = value
		}
	}
	p.where = whereAdd
	p.whereOr = whereOr
	return p, nil
}

// 设置规则
func (p *Pagenation) Rules(rules ...*Rule) *Pagenation {
	if p.rules == nil {
		p.rules = make(map[string]*Rule)
	}
	if len(rules) > 0 {
		for _, v := range rules {
			p.rules[v.Field] = v
		}
	}
	return p
}

// where处理对象
type whereItem struct {
	Operator string      //操作符
	Value    interface{} // 对应的值
	Where    string      `d:"ADD"` // 查询结果
	Fields   string
}

type Rule struct {
	Field    string
	Operator string
	WhereOr  bool     // 是否是where or
	DBFields []string // 映射到数据库的字段
}

// 通过rule 组装whereItems
func (r *Rule) getWhereItems(v interface{}) []*whereItem {
	var res = make([]*whereItem, 0)
	// 多字段
	dbfileds := []string{r.Field}
	if len(r.DBFields) > 0 {
		dbfileds = r.DBFields
	}
	//	操作符

	for _, field := range dbfileds {
		item := &whereItem{
			Value:    v,
			Operator: r.Operator,
			Where:    "AND",
			Fields:   field,
		}
		// or 查询
		if r.WhereOr {
			item.Where = "OR"
		}
		res = append(res, item)
	}
	return res

}

// 校验字段是否存在
func (p *Pagenation) checkColumns(fileds ...string) bool {
	// map不存在  默认不校验
	if len(p.columns) == 0 {
		return true
	}

	if len(fileds) == 0 {
		return true
	}

	for _, v := range fileds {
		if _, ok := p.columns[v]; !ok {
			return false
		}
	}

	return true
}

// 校验字段是否允许为空
func (p *Pagenation) checkEmpty(fileds string) bool {
	return helper.FindStr(p.empty, fileds) >= 0
}

/**
 * 校验用户输入的查询是否是string  int
 */
func (p *Pagenation) checkWhereValue(i interface{}) (res interface{}, err error) {
	i_rty := reflect.TypeOf(i)

	switch i_rty.Kind() {
	case reflect.Func, reflect.Map, reflect.Struct, reflect.Array, reflect.Bool, reflect.Chan, reflect.Interface:
		return res, errors.New("search error: value except value string,int but got :" + i_rty.Kind().String())
	case reflect.Ptr:
		return p.checkWhereValue(reflect.ValueOf(i).Elem().Interface())
	}

	return i, err
}
