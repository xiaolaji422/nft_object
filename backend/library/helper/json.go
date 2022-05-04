package helper

import (
	"encoding/json"
	"net/url"
	"strings"
)

// JsonToStr json转字符串
func JsonToStr(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

// StrToJson 字符串转json
func StrToJson(s string) interface{} {
	var v interface{}
	err := json.Unmarshal([]byte(s), v)
	if err != nil {
		return nil
	}
	return v
}

// url编码
func UrlEnCode(s string) string {
	baseUrlStr := url.PathEscape(s)
	baseUrlStr = strings.Replace(baseUrlStr, "&", "%26", -1)
	return baseUrlStr
}

// url编码
func UrlDeCode(s string) (string, error) {
	return url.PathUnescape(s)
}

// 校验json
func JsonValidate(s []byte) bool {
	return json.Valid(s)
}
