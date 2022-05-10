package runtime

import (
	"runtime"
	"strings"
)

func GetFuntName(skip ...int) string {
	var i = 2
	if len(skip) > 0 {
		i = skip[0]
	}
	var str = "GetFuntName"
	if pc, _, _, ok := runtime.Caller(i); ok {

		pc_name := runtime.FuncForPC(pc).Name()
		keys_list := strings.Split(pc_name, ".")
		str = keys_list[len(keys_list)-1]
	}
	return str
}
