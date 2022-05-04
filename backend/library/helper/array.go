package helper

import (
	"math/rand"
	"reflect"
	"sort"
	"time"

	"github.com/gogf/gf/util/gconv"
)

// InArray 判断value是否在array中
func InArray(a interface{}, v interface{}) bool {
	switch reflect.TypeOf(a).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(a)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(v, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

// Contain 判断searchkey是否在target中，target支持的类型arrary,slice,map
func Contain(searchkey interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == searchkey {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(searchkey)).IsValid() {
			return true
		}
	}
	return false
}

// SortMap 对slice map排序
func SortMap(data []map[string]interface{}, index string) []map[string]interface{} {
	flag := 0
	sort.Slice(data, func(i, j int) bool {
		if flag == 0 {
			t := reflect.ValueOf(data[i][index]).Kind()
			if t == reflect.String {
				flag = 1
			}
		}
		if flag == 1 {
			return gconv.String(data[i][index]) < gconv.String(data[j][index])
		}
		return gconv.Int(data[i][index]) < gconv.Int(data[j][index])
	})
	return data
}

// Sort 对slice map或者slice struct 排序（统一转成slice map去排序）
func Sort(data []interface{}, index string, sortType ...string) []interface{} {
	if len(sortType) == 0 {
		sortType = []string{"asc"}
	}
	flag := 0
	sort.Slice(data, func(i, j int) bool {
		if flag == 0 {
			t := reflect.ValueOf(gconv.Map(data[i])[index]).Kind()
			if t == reflect.String {
				flag = 1
			}
		}
		if flag == 1 {
			if sortType[0] == "desc" {
				return gconv.String(gconv.Map(data[i])[index]) > gconv.String(gconv.Map(data[j])[index])
			}
			return gconv.String(gconv.Map(data[i])[index]) < gconv.String(gconv.Map(data[j])[index])
		}
		if sortType[0] == "desc" {
			return gconv.Int(gconv.Map(data[i])[index]) > gconv.Int(gconv.Map(data[j])[index])
		}
		return gconv.Int(gconv.Map(data[i])[index]) < gconv.Int(gconv.Map(data[j])[index])
	})
	return data
}

// RandSlice 随机取一个slice值
func RandSlice(dataAny interface{}) interface{} {
	data := gconv.SliceAny(dataAny)
	l := len(data)
	if l == 0 {
		return dataAny
	}
	// n := rand.Int() % len(data)
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	n := rand.Intn(l)
	return data[n]
}

// ArrayUnique slice去重
func ArrayUnique(data []string) []string {
	result := make([]string, 0, len(data))
	temp := map[string]struct{}{}
	for _, item := range data {
		if len(item) <= 0 {
			continue
		}
		if _, ok := temp[item]; !ok && item != "" {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// int数组去重
func ArrayUniqueInt(data []int) []int {
	result := []int{}
	temp := map[string]int{}
	for _, item := range data {
		if _, ok := temp[gconv.String(item)]; !ok && item != 0 {
			temp[gconv.String(item)] = 1
			result = append(result, item)
		}
	}
	return result
}

// 获取map中的key
func GetMapValue(data map[string]interface{}, property string, def interface{}) interface{} {
	if _, ok := data[property]; ok {
		return data[property]
	}
	return def
}

// 获取结果集中某一列的值
func ArrayColumn(data []map[string]interface{}, column string) []interface{} {
	newData := make([]interface{}, 10)
	for _, v := range data {
		for tmpK, tmpV := range v {
			if tmpK == column {
				newData = append(newData, tmpV)
			}
		}
	}
	return newData
}

// 查找slice中是个否存在
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
