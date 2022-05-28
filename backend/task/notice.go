package task

import (
	"context"
	"fmt"
	"nft_object/app/modules/service"
	"time"

	"github.com/gogf/gf/os/gtimer"
)

func init() {
	interval := time.Second * 3 // 3秒一次
	gtimer.AddSingleton(interval, func() {
		// 获取是否有最新的公告
		msg := service.NewMessage("TASK", "我获取到了新的信息")
		_, err := service.SendMsgImpl().SendAll(context.Background(), []string{}, msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	})
}
