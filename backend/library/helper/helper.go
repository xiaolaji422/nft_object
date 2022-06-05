package helper

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"nft_object/statusCode"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
)

/**
 * @description  :
 * @param         {*} string
 * @return        {*} string
 * @author       : fourteen
 */
func StrCase(str string) string {
	return gstr.SnakeCase(str)
}

// 获取上下文中的RTX
func GetRtx(ctx context.Context) string {

	RTX := ctx.Value(statusCode.SESSION_CACHE_ADMIN_NAME)
	if g.IsEmpty(RTX) {
		return "admin"
	}
	return gconv.String(RTX)
}

func GetAdminId(ctx context.Context) int {
	RTX := ctx.Value(statusCode.SESSION_CACHE_ADMIN_ID)
	if g.IsEmpty(RTX) {
		return 0
	}
	return gconv.Int(RTX)
}

// 获取字符串首字母
func FirstLetter(s string) string {
	_, size := utf8.DecodeRuneInString(s)
	return s[:size]
}

func StrPad(input string, padLength int, padString string, padType string) string {
	var output string

	inputLength := len(input)
	padStringLength := len(padString)

	if inputLength >= padLength {
		return input[0:padLength]
	}

	repeat := math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength))

	switch padType {
	case "RIGHT":
		output = input + strings.Repeat(padString, int(repeat))
		output = output[:padLength]
	case "LEFT":
		output = strings.Repeat(padString, int(repeat)) + input
		output = output[len(output)-padLength:]
	case "BOTH":
		length := (float64(padLength - inputLength)) / float64(2)
		repeat = math.Ceil(length / float64(padStringLength))
		output = strings.Repeat(padString, int(repeat))[:int(math.Floor(float64(length)))] + input +
			strings.Repeat(padString, int(repeat))[:int(math.Ceil(float64(length)))]
	}

	return output
}

// 字符串去重，如：test;test1;test; 去重：test;test1;
func StrUnique(str string, sep ...string) string {
	if str == "" {
		return str
	}
	if len(sep) == 0 {
		// 设置默认分割符
		sep = append(sep, ";")
	}
	strs := strings.Split(str, sep[0])
	result := make([]string, 0, len(strs))
	// 空struct不占内存空间
	temp := map[string]struct{}{}
	for _, item := range strs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	str = strings.Join(result, sep[0])
	return str
}

// 多个map合并
func MergeMap(data map[string]interface{}, d ...map[string]interface{}) map[string]interface{} {
	if len(d) > 0 {
		for _, val := range d {
			for k, v := range val {
				data[k] = v
			}
		}
	}
	return data
}

// 过虑html标签，支持指定保留标签
func StripTags(html string, allowed_tags ...string) string {

	if html == "" {
		return html
	}
	// 允许保留标签变量
	allow_match := []string{}

	// 处理允许保留的标签
	if len(allowed_tags) > 0 {

		allowed_pattern := ""
		allowed_tags_list := strings.Split(allowed_tags[0], ",")
		re, _ := regexp.Compile(`<|>`)
		for _, v := range allowed_tags_list {
			if v == "" {
				continue
			}
			if allowed_pattern != "" {
				allowed_pattern += "|"
			}
			v = re.ReplaceAllString(v, "")
			allowed_pattern += "(?m)<" + v + ".*?>|</" + v + ">"
		}
		var hrefRegexp = regexp.MustCompile(allowed_pattern)
		allow_match = hrefRegexp.FindAllString(html, -1)
	}

	var tagsRegexp = regexp.MustCompile("(?m)<.*?>")
	match := tagsRegexp.FindAllString(html, -1)

	// 循环所有的标签，如指定保留标签，则保留，反之替换空字符
	for _, del := range match {
		if !Contain(del, allow_match) {
			re, _ := regexp.Compile(del)
			html = re.ReplaceAllString(html, "")
		}
	}
	return html
}

// 指定位数随机数
func RandomNumLength(length int) string {
	result := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		num := rand.Intn(10)
		result = result + strconv.Itoa(num)
	}
	return result
}

// 获取 min - max 的一个随机数
func GetRangeRandNum(max int, min int) int {
	return rand.Intn(max-min) + min
}

// GetMapChildren 获取多层级map的子层级数据
// pattern string 如：a.b.c
func GetMapChildren(data interface{}, pattern string) interface{} {
	keys := gstr.Split(pattern, ".")
	for _, k := range keys {
		if _, ok := data.(map[string]interface{})[k]; ok {
			data = data.(map[string]interface{})[k]
		}
	}
	return data
}

// IsNumeric 判断是否数字
func IsNumeric(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
	case float32, float64, complex64, complex128:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		return gstr.IsNumeric(str)
	}
	return false
}

// SplitAtCommas 逗号分割字符串，双引号内的逗号忽略
func SplitAtCommas(s string) []string {
	res := []string{}
	var beg int
	var inString bool

	for i := 0; i < len(s); i++ {
		if s[i] == ',' && !inString {
			res = append(res, s[beg:i])
			beg = i + 1
		} else if s[i] == '"' {
			if !inString {
				inString = true
			} else if i > 0 && s[i-1] != '\\' {
				inString = false
			}
		}
	}
	return append(res, s[beg:])
}

// EncryptId 加密id，通过生成5位随机数来组合解密
func EncryptId(id string) string {
	// 生成随机数
	rstr := strconv.Itoa(grand.N(10000, 99999))
	return hex.EncodeToString([]byte(base64.StdEncoding.EncodeToString([]byte(id+"@"+rstr)))) + rstr
}

// DecryptId 解密id
func DecryptId(id string) string {
	p := len(id) - 5
	if p <= 0 {
		return ""
	}
	// 去掉后面随机数，再解密
	str := id[:p]
	b, err := hex.DecodeString(str)
	if err != nil {
		return ""
	}
	b, err = base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		return ""
	}
	str = string(b)
	// 如果解密串包涵随机数，则返回正常解密id
	if strings.Contains(str, "@"+id[p:]) {
		str = strings.ReplaceAll(str, "@"+id[p:], "")
		return str
	}
	return ""
}

// ExecGrepCommand 执行管道命令，如：ps -ef | grep http | wc -l
func ExecGrepCommand(strCommand string) (*bytes.Buffer, error) {

	fmt.Println("command: ", strCommand)
	commands := strings.Split(strCommand, "|")
	var stdout *bytes.Buffer
	var err error
	for _, command := range commands {
		if command == "" {
			continue
		}
		stdout, err = ExecCommand(command, stdout)
		if err != nil {
			fmt.Println("execute command error: ", err.Error())
			return stdout, err
		}
	}
	return stdout, err
}

// ExecCommand 执行命令，如：ps -ef
func ExecCommand(strCommand string, prevStdout ...*bytes.Buffer) (*bytes.Buffer, error) {
	params := strings.Split(strCommand, " ")
	name := ""
	args := []string{}
	for _, v := range params {
		if v != "" {
			if name == "" {
				name = v
			} else {
				args = append(args, v)
			}
		}
	}
	cmd := exec.Command(name, args...)
	if len(prevStdout) > 0 && prevStdout[0] != nil {
		// 管道依赖输入
		cmd.Stdin = prevStdout[0]
	}

	stdout, cmdErr := SetCommandStd(cmd)
	err := cmd.Run()
	if err != nil {
		err = errors.New(err.Error() + cmdErr.String())
	}
	return stdout, err
}

// SetCommandStd 设置命令输出
func SetCommandStd(cmd *exec.Cmd) (stdout, stderr *bytes.Buffer) {
	stdout = &bytes.Buffer{}
	stderr = &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	return
}

// ReplaceUrlParams 过滤/替换url参数
func ReplaceUrlParams(urlAddr string, pattern []string) string {
	for _, str := range pattern {
		urlAddr = regexp.MustCompile("(&?)"+str+"=[^&]+").ReplaceAllString(urlAddr, "")
	}
	urlAddr = regexp.MustCompile("\\?&").ReplaceAllString(urlAddr, "?")
	return urlAddr
}

//清洗掉某个slice中的某些不要的
func ClearSliceInt(data []int, disable []int) []int {
	result := []int{}
	if len(data) <= 0 || len(disable) < 0 {
		return data
	}
	apiMap := make(map[int]int)
	for _, v := range data {
		apiMap[v] = 1
	}
	for _, v := range disable {
		apiMap[v] = 0
	}
	for k, v := range apiMap {
		if v > 0 {
			result = append(result, k)
		}
	}
	return result
}

// 清洗掉某个slice中的某些不要的
func ClearSliceString(data []string, disable []string) []string {
	result := []string{}
	if len(data) <= 0 || len(disable) < 0 {
		return data
	}
	apiMap := make(map[string]int)
	for _, v := range data {
		apiMap[v] = 1
	}
	for _, v := range disable {
		apiMap[v] = 0
	}
	for k, v := range apiMap {
		if v > 0 {
			result = append(result, k)
		}
	}
	return result
}
