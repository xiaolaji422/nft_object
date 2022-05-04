package helper

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/gogf/gf/frame/g"
)

/**
* @description  :   获取唯一id
* @param         prefix 指定前缀
* @rule   前缀+ 类型+ 年月日时分秒 + 15位随机字符
* 唯一id生成规则：
# 话题|评论 前缀+ 类型+ 年月日时分秒 + 15位随机字符
# TQ20210510161426mlil0oKdfetpiWD
# T 话题 三种（用户提问Q 话题H 音视频V）
# C 评论 三种（用户提问Q 话题H 音视频V）
* @return        {*}
* @author       : fourteen
*/
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// 获取唯一id
// len[0] 标识随机字符串长度
func GetUniqueId(prefix string, len ...int) string {
	res := prefix
	var strLen = 15
	timeStr := time.Now().Format("20060102150405")
	if !g.IsEmpty(len) {
		strLen = len[0]
	}
	randomStr := RandStringBytesMaskImprSrc(strLen)
	res = res + string(timeStr) + randomStr
	return res
}

var src = rand.NewSource(time.Now().UnixNano())

// 获取指定长度的随机字符串
func RandStringBytesMaskImprSrc(n int) string {
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

/**
 * @description  : 获取一个字符串最后两位的hash取值
 * @param         {string} str
 * @param         {int } num  最后几位
 * @return        {*}
 * @author       : fourteen
 */
func GetStringHashInt(valuesStr string, num int) (int64, error) {
	keyStr := valuesStr[len(valuesStr)-num:]
	fmt.Printf("%v,%T", keyStr, keyStr)
	items := []rune(keyStr)
	if len(items) <= 0 {
		return 0, errors.New("字符串不存在")
	}
	total := 0
	for _, v := range items {
		total = total + int(v)
	}
	return int64(total), nil
}
