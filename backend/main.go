package main

import (
	_ "nft_object/boot"
	_ "nft_object/router"

	"github.com/gogf/gf/frame/g"
)

/**
* 	@title      	sy服务API
* 	@version     	1.0
* 	@description 	GoFrame`基础开发框架示例服务API接口文档。
*	@schemes     	http
**/
func main() {
	g.Server().Run()
}
