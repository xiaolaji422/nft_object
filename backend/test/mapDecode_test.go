package test

import (
	"net/url"
	"reflect"
	"testing"
)

// 将map中的数据urlDecode
func mapDecode(data map[string]interface{}) map[string]interface{} {
	for k, v := range data {
		data[k], _ = url.QueryUnescape(v.(string))
	}
	return data
}

func TestMapDecode(t *testing.T) {
	get := mapDecode(map[string]interface{}{"k": "%E4%BD%A0%E5%A5%BD"})
	want := map[string]interface{}{"k": "你好"}
	if !reflect.DeepEqual(get, want) {
		t.Error("not equal")
	}
}
