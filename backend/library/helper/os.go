package helper

import (
	"errors"
	"fmt"
	"net"
	"reflect"
	"strconv"

	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/grand"
)

// GetLocalIP 获取本机ip地址
func GetLocalIP() (ips []string, err error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("get ip interfaces error:", err)
		return
	}

	for _, i := range ifaces {
		addrs, errRet := i.Addrs()
		if errRet != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				if ip.IsGlobalUnicast() {
					ips = append(ips, ip.String())
				}
			}
		}
	}
	return
}

// CreateId 生成一个24位的id
// 第一个参数是节点0~99
// 第二个参数是生成位数，大于等于20
func CreateId(node ...int) string {
	n := 0
	if len(node) > 0 {
		n = node[0]
	}
	str := gtime.Now().Format("ymdHis") + strconv.Itoa(n)

	clen := len(str)
	max := 24
	if len(node) > 1 && node[1] >= 20 {
		max = node[1]
	}
	maxRand, _ := strconv.Atoi(gstr.Repeat("9", max-clen))
	mixRand, _ := strconv.Atoi("1" + gstr.Repeat("0", max-clen-1))
	randNum := grand.N(mixRand, maxRand)

	str += strconv.Itoa(randNum)
	return str
}

// Call 调用方法
func Call(method interface{}, params ...interface{}) ([]reflect.Value, error) {

	if reflect.TypeOf(method).Kind() != reflect.Func {
		return nil, errors.New("the name of input not func!")
	}

	f := reflect.ValueOf(method)
	if len(params) != f.Type().NumIn() {
		return nil, errors.New("the number of input params not match!")
	}
	mp := []reflect.Value{}
	for _, v := range params {
		mp = append(mp, reflect.ValueOf(v))
	}
	return f.Call(mp), nil
}
