package boot

import (
	"flag"
	"nft_object/library/dbDriver"
	"nft_object/middleware/cfg"
	"time"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gsession"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

// 用于应用初始化。
func init() {
	gtime.SetTimeZone("Asia/Shanghai")

	c := g.Config()
	s := g.Server()
	//解析命令行参数
	env := flag.String("env", "pro", "环境变量")
	portStr := flag.String("port", "", "端口")
	flag.Parse()

	cfg.Do(*env)

	// if c.GetString("redis.session") != "" {
	// 	// redis sessions
	// 	s.SetConfigWithMap(g.Map{
	// 		"SessionMaxAge":  time.Second * time.Duration(c.GetInt("redis.lifetime", 0)),
	// 		"SessionStorage": gsession.NewStorageRedis(g.Redis()),
	// 	})
	// }

	logPath := c.GetString("server.LogPath", "./logs")
	s.SetLogPath(logPath)
	s.SetAccessLogEnabled(c.GetBool("server.AccessLogEnabled"))
	s.SetErrorLogEnabled(c.GetBool("server.ErrorLogEnabled"))
	s.SetErrorStack(c.GetBool("server.ErrorStack"))
	s.SetDumpRouterMap(c.GetBool("server.DumpRouterMap"))

	// s.SetServerRoot("/public")
	// // 设置静态目录
	// s.AddStaticPath("/public", "/public")
	// 静态文件服务总开关：是否开启/关闭静态文件服务
	// s.SetFileServerEnabled(true)
	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	// s.SetRewrite("/favicon.ico", "/public/favicon.ico")

	// 开启平滑重启
	s.EnableAdmin()
	// 读取敏感词的词库到json
	// var dirtyFilterMap = gcfg.Instance("dirtyFilterMap").ToMap()
	// g.Config().Set("dirtyFilterMapPointer", dirtyFilterMap)
	// 默认读取配置address
	address := c.GetString("server.address", ":8222")
	s.SetAddr(address)

	// 如果有port参数，则设置，否则默认配置
	if *portStr != "" {
		port := gconv.Int(*portStr)
		s.SetPort(port)
	}

	// 注册mysql driver
	gdb.Register("mysql", &dbDriver.MysqlDriver{})

	// 注册session的
	s.SetConfigWithMap(g.Map{
		"SessionMaxAge":  time.Minute * 600,
		"SessionStorage": gsession.NewStorageRedis(g.Redis()),
		// "dirtyFilterMap": &dirtyFilterMap,
	})

}
