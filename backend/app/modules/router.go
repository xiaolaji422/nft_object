package modules

// 此处主要是引入具体模块的路由  主动注册

import (
	// 引入admin模块路由
	_ "nft_object/app/modules/admin"
	// 引入权限认证模块路由
	_ "nft_object/app/modules/auth"

	// 引入系统模块路由
	_ "nft_object/app/modules/system"
)
