package logge

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

// 日志统一管理（入库到redis，再从redis消费）

// Write 写日志
func Write(filepath, level string, data interface{}, datas ...interface{}) {

	logpath := g.Config().GetString("server.LogPath", "./logs")
	glog.SetPath(logpath)

	fileName := "{Ymd}." + filepath + ".log"
	if level == "error" {
		// 如果是错误日志另写到error目录
		fileName = level + "/" + fileName
		if err := os.MkdirAll(logpath+"/"+level, os.ModePerm); err != nil {
			fmt.Print("文件创建失败")
		}
	}
	msg := " | " + gconv.String(data)

	systemInfo := map[string]interface{}{}
	// 获取执行错误的文件名和行
	_, file, line, ok := runtime.Caller(1)
	if ok {
		systemInfo["file"] = file
		systemInfo["line"] = line
		msg += " | " + gconv.String(systemInfo)
	}

	if len(datas) > 0 {
		for i := 0; i < len(datas); i++ {
			msg += " | " + gconv.String(datas[i])
		}
	}

	glog.File(fileName).Line(false).Stdout(false).Println(msg)
}

/**
 * @description  :  写入sql日志
 * @param         {*} filepath
 * @param         {string} level
 * @param         {interface{}} data
 * @param         {...interface{}} datas
 * @return        {*}
 * @author       : fourteen
 */
func WriteSql(filepath, level string, data interface{}, datas ...interface{}) {

	logpath := g.Config().GetString("server.LogPath", "./logs")
	glog.SetPath(logpath)

	fileName := "{Ymd}." + filepath + ".log"
	msg := "【" + level + "】 | " + gconv.String(data)

	if len(datas) > 0 {
		for i := 0; i < len(datas); i++ {
			msg += " | " + gconv.String(datas[i])
		}
	}
	glog.File(fileName).Line(false).Stdout(false).Println(msg)
}

// WriteError 写error级别日志
func WriteError(filepath string, data interface{}, datas ...interface{}) {
	Write(filepath, "error", data, datas...)
}

// WriteInfo 写info级别日志
func WriteInfo(filepath string, data interface{}, datas ...interface{}) {
	Write(filepath, "info", data, datas...)
}
