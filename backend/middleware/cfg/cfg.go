package cfg

import (
	"os"
	"path/filepath"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/genv"
)

// Do 处理配置信息
func Do(env string) {

	c := g.Config()
	cfgDir, err := c.GetFilePath()
	if err != nil {
		cfgDir = "./config"
	}
	// 根据env参数判断配置文件是否存在
	cfg_path := filepath.Dir(cfgDir) + "/" + env + "/config.toml"

	// 根据环境变量读取配置文件
	_, err = os.Stat(cfg_path)
	if err != nil && os.IsNotExist(err) {
		c.SetFileName("config.toml")
		// 默认读取的是现网配置
		env = "pro"
	} else {
		c.SetFileName(env + "/config.toml")
	}

	// 设置项目环境变量
	genv.Set("env", env)
}
