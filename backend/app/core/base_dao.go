package core

// dao的基类

import (
	"database/sql"
	"errors"
	"math"
	"nft_object/library/helper"
	"nft_object/statusCode"
	"reflect"
	"strings"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"github.com/gogf/gf/util/gconv"
)

// 数据库操作助手类
var BaseDao = baseDao{}

type baseDao struct {
	columns map[string]interface{}
}

// 公共分页返回的信息
type PageReult struct {
	Total       int        `json:"total"`
	CurrentPage int        `json:"current_page"`
	PerPage     int        `json:"per_page"`
	Pages       float64    `json:"pages"`
	Items       gdb.Result `json:"items"`
}

// where处理对象
type WhereItem struct {
	Operator string      `json:"operator"`      //操作符
	Value    interface{} `json:"value"`         // 对应的值
	Where    string      `d:"ADD" json:"where"` // 查询结果
	Empty    bool
}

type Empty int

//  排序规则的底层
type OrderItem struct {
	Property string
	Sort     string
}

// 数据库操作的类
type GdbModel struct {
	gmvc.M
	DB       gdb.DB
	Table    string
	Where    map[string]WhereItem
	Order    []OrderItem
	Fields   []string
	FieldsEx []string
	Columns  map[string]interface{}
}

// 处理子查询

// 公共分页信息
func (d *baseDao) PagenationCommon(gdbModel GdbModel, params g.Map) (pageReult PageReult, err error) {
	// 分页参数
	page, limit := d.PageInfo(params)
	pageReult.CurrentPage = page
	pageReult.PerPage = limit
	// 赋值字段表
	d.columns = gdbModel.Columns
	if g.IsEmpty(gdbModel.Columns) {
		return pageReult, errors.New("开发人员应传入字段列表")
	}

	// // 子查询处理
	dm := gdbModel.M
	// 查询结果排序
	order, err := d.OrderInfo(gdbModel.Order, params)
	if err != nil {
		return pageReult, err
	}

	//  指定查询字段
	fields, err := d.Fields(gdbModel.Fields)
	if err != nil {
		return pageReult, err
	}
	if len(fields) > 0 {
		dm = dm.Order(strings.Join(fields, ","))
	}
	// 排除关键字
	fieldsEx, err := d.FieldsEx(gdbModel.FieldsEx)
	if err != nil {
		return pageReult, err
	}
	if len(fieldsEx) > 0 {
		dm = dm.FieldsEx(strings.Join(fieldsEx, ","))
	}
	where, whereOr, err := d.Where(gdbModel.Where, params)
	if err != nil {
		return pageReult, err
	}
	// 条件查询
	if !g.IsEmpty(where) {
		for k, v := range where {
			dm = dm.Where(k, v)
		}
	}

	if !g.IsEmpty(whereOr) {
		var str []string
		var values []interface{}
		for k, v := range whereOr {
			str = append(str, k)
			values = append(values, v)

		}
		dm = dm.Where(strings.Join(str, " OR "), values...)
		// dm = dm.WhereOr(whereOr)
	}
	// 排序
	if len(order) > 0 {
		dm = dm.Order(strings.Join(order, ","))
	}
	all, err := dm.Limit((page-1)*limit, limit).FindAll()
	if err != nil {
		return pageReult, err
	}

	count, _ := dm.Count()
	pageReult.Total = count
	pageReult.Pages = math.Ceil(float64(count) / float64(limit))
	pageReult.Items = all
	return pageReult, nil
}

//  校验字段是否存在
func (d *baseDao) checkColumns(fileds string) bool {
	if _, ok := d.columns[fileds]; ok {
		return true
	}
	return false
}

// 处理查询分页条件
func (b *baseDao) PageInfo(data map[string]interface{}) (int, int) {
	page := gconv.Int(helper.GetMapValue(data, "page", 1))
	if page <= 0 {
		page = 1
	}
	limit := gconv.Int(helper.GetMapValue(data, "limit", 10))
	if limit <= 0 {
		limit = 10
	}
	delete(data, "page")
	delete(data, "limit")
	// 设置分页最大上限
	if limit > statusCode.MAX_LIMIT {
		limit = statusCode.MAX_LIMIT
	}
	return page, limit
}

/**
 *	处理查询排序条件
 *	morder 模型写死的查询条件
 *	data  前端传入的排序字段
 */
func (d *baseDao) OrderInfo(MOrder []OrderItem, data map[string]interface{}) ([]string, error) {
	var orderMap []string
	//  如果传入参数中存在_order  MOrder置空
	if _, ok := data["_order_property"]; ok {
		proStr := gconv.String(data["_order_property"])
		//  重置为只有一个的order
		var sort_str = gconv.String(helper.GetMapValue(data, "_order_sort", "ASC"))

		delete(data, "_order_property")
		delete(data, "_order_sort")
		if strings.ToTitle(sort_str) != "DESC" {
			sort_str = "ASC"
		} else {
			sort_str = "DESC"
		}

		if d.checkColumns(proStr) {
			MOrder = []OrderItem{
				{
					Property: proStr,
					Sort:     sort_str,
				},
			}
		} else {
			return orderMap, errors.New("sort property is not exit:" + proStr)
		}

	}
	if g.IsEmpty(MOrder) {
		return orderMap, nil
	}
	//  组装排序规则
	for _, v := range MOrder {
		if strings.ToTitle(v.Sort) != "DESC" {
			v.Sort = "ASC"
		}
		if d.checkColumns(v.Property) {
			sort_str := v.Property + " " + v.Sort
			orderMap = append(orderMap, sort_str)
		}
	}
	return orderMap, nil
}

// 处理查询指定字段
func (d *baseDao) Fields(fields []string) ([]string, error) {
	var (
		filter []string
		// 默认升序
	)
	if g.IsEmpty(fields) {
		return filter, nil
	}

	for _, v := range fields {
		if d.checkColumns(v) {
			filter = append(filter, v)
		} else {
			return filter, errors.New("fields is not suport:" + v)
		}
	}
	return filter, nil
}

// 处理查询指定排除字段
func (d *baseDao) FieldsEx(fieldsEx []string) ([]string, error) {
	var (
		filter []string
		// 默认升序
	)
	if g.IsEmpty(fieldsEx) {
		return filter, nil
	}

	for _, v := range fieldsEx {
		if d.checkColumns(v) {
			filter = append(filter, v)
		}
	}
	// d.dm = d.dm.FieldsEx(strings.Join(filter, ","))
	return filter, nil
}

// 处理排序条件具体函数
// func (d *baseDao) handleOrderMap(data map[string]interface{}) []map[string]interface{} {
// 	// 组装数据  组装目标  []map[string]interface  [{property:name,sort:desc,weight:1}]

// 	var orderSlice []map[string]interface{}

// 	for k, v := range data {
// 		itemMap := map[string]interface{}{
// 			"property": k,
// 			"sort":     "asc",
// 			"weight":   0,
// 		}
// 		switch reflect.TypeOf(v).Kind() {
// 		case reflect.String:
// 			if strings.ToTitle(gconv.String(v)) == "DESC" || strings.ToTitle(gconv.String(v)) == "ASC" {
// 				itemMap["sort"] = strings.ToTitle(gconv.String(v))
// 			}
// 		case reflect.Map:
// 			customSort := helper.GetMapValue(gconv.Map(v), "sort", "ASC")
// 			customWeight := helper.GetMapValue(gconv.Map(v), "weight", 1)
// 			itemMap["sort"] = strings.ToTitle(gconv.String(customSort))
// 			itemMap["weight"] = gconv.Int(customWeight)
// 		}
// 		orderSlice = append(orderSlice, itemMap)
// 	}
// 	// 排序一下 根据权重排序
// 	res := helper.SortMap(orderSlice, "weight")

// 	return res
// }

// 组装查询条件
func (d *baseDao) Where(whereMap map[string]WhereItem, data g.Map) (whereAdd g.Map, whereOr g.Map, err error) {
	var (
		searchMap = make(map[string]*WhereItem)
	)
	whereAdd = make(map[string]interface{})
	whereOr = make(map[string]interface{})
	if g.IsEmpty(data) {
		return
	}

	// 排除非法字段
	for k, v := range data {
		if d.checkColumns(k) {
			if item, ok := whereMap[k]; ok {
				if g.IsEmpty(v) && !item.Empty {
					continue
				}
				item.Value = v
				searchMap[k] = &item
			} else {

				//  使用外界穿参
				item := d.handleWhere(v)
				searchMap[k] = item
			}
		} else {
			err = errors.New("property is not suport:" + k)
			return
		}
	}
	if g.IsEmpty(searchMap) {
		return
	}
	for k, v := range searchMap {
		// 转化为下划线的搜索
		k = helper.StrCase(k)
		key := ""
		// 跳过循环
		if g.IsEmpty(v.Value) {
			continue
		}
		operator := "="
		if !g.IsEmpty(v.Operator) {
			operator = v.Operator
		}
		switch operator {
		case "between":
			key = k + " " + operator + " ? and ?"
		case "is":
			if reflect.TypeOf(v.Value).Kind() != reflect.String {
				continue
			}
			isNUll := strings.ToUpper(gconv.String(v.Value))
			if isNUll == "NULL" {
				key = k + " " + operator + " NULL"
				v.Value = nil
			} else if isNUll == "NOT NULL" {
				key = k + " " + operator + " NOT NULL"
				v.Value = nil
			} else {
				continue
			}
		case "in":
			key = k + " " + operator + " (?)"
		case "notin":
			key = k + " NOT IN " + " (?)"
		case "like":
			key = k + " " + operator + " ?"
		case "likeall":
			operator = "like"
			key = k + " " + operator + " ?"

		default:
			key = k + " " + operator + " ?"
		}
		if v.Where == "OR" {
			whereOr[key] = v.Value
		} else {
			whereAdd[key] = v.Value
		}
	}
	return
}

// 处理where条件
func (d *baseDao) handleWhere(v interface{}) (item *WhereItem) {
	item = &WhereItem{
		Value:    nil,
		Operator: "=",
		Where:    "AND",
	}
	// v   结构： 1. "小米"  2. {value:"小米",operator:"like"}
	switch reflect.TypeOf(v).Kind() {
	case reflect.String:
		item.Value = v.(string)
	case reflect.Map:
		var itemEn WhereItem
		gconv.Struct(gconv.Map(v), &itemEn) // 转换
		if g.IsEmpty(itemEn.Value) {
			return item
		}
		itemEn.Operator = strings.ToLower(itemEn.Operator)
		itemEn.Where = strings.ToUpper(itemEn.Where)
		itemEn.Value = d.handleWhereValue(itemEn.Operator, itemEn.Value)
		item = &itemEn
	default:
		item.Value = v
	}
	return
}

// 处理where条件的value
func (d *baseDao) handleWhereValue(operator string, value interface{}) interface{} {
	switch operator {
	case "like":
		if reflect.TypeOf(value).Kind() != reflect.String {
			break
		}
		// 此处只使用右匹配
		str := gconv.String(value)
		str = strings.ReplaceAll(str, "%", "\\%")
		return str + "%"
	case "likeall":
		if reflect.TypeOf(value).Kind() != reflect.String {
			break
		}
		// 此处只使用右匹配
		//  特殊字符串替换
		str := gconv.String(value)
		str = strings.ReplaceAll(str, "%", "\\%")
		return "%" + str + "%"
	case "in":
		if reflect.TypeOf(value).Kind() == reflect.Slice {
			return gconv.SliceAny(value)
		} else if reflect.TypeOf(value).Kind() != reflect.String {
			break
		}
		return strings.Split(gconv.String(value), ",")
	case "notin":
		if reflect.TypeOf(value).Kind() == reflect.Slice {
			return gconv.SliceAny(value)
		} else if reflect.TypeOf(value).Kind() != reflect.String {
			break
		}
		return strings.Split(gconv.String(value), ",")
	case "between":
		if reflect.TypeOf(value).Kind() != reflect.String {
			break
		}
		return strings.Split(gconv.String(value), ",")
	default:
		return value
	}
	return value
}

// 解析添加、修改的结果
func (d *baseDao) HandleExecRes(result sql.Result, err error, desc string) (int, error) {
	if err != nil {
		return 0, errors.New("数据库操作失败:" + err.Error())
	}
	r, err := result.RowsAffected()
	if err != nil {
		return 0, errors.New("数据库操作解析失败")
	}
	return int(r), err
}
