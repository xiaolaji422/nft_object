package redis

import (
	"context"
	"nft_object/library/redis"
	"time"
)

//  锁单redis
// todo    获取某个锁单信息  // 商品id
// 删除某个锁单信息
type ILockRedis interface {
	SetLockList(ctx context.Context, id int) error
}

var LockRedisImpl = func() ILockRedis {
	return &lock_redis{
		key: "nft_object:refresh:lock:list",
	}
}

type lock_redis struct {
	key      string
	duration time.Duration
}

// 设置信息
func (r *lock_redis) SetLockList(ctx context.Context, id int) error {
	return redis.ListRPush(r.key, id)
}
