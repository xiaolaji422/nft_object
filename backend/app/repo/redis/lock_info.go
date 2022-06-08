package redis

import (
	"context"
	"nft_object/app/model"
	"nft_object/library/redis"
	"strconv"
	"time"
)

//  锁单信息redis
type ILockInfoRedis interface {
	GetInfo(ctx context.Context, id int) (*model.AccountLock, error)
	SetInfo(ctx context.Context, info *model.AccountLock) error
	DelInfo(ctx context.Context, id int) error
}

var LockInfoRedisImpl = func() ILockInfoRedis {
	return &lock_info_redis{
		key:      "nft_object:refresh:lock:info:set",
		duration: time.Hour * 2,
	}
}

type lock_info_redis struct {
	key      string
	duration time.Duration
}

// 获取信息
func (r *lock_info_redis) GetInfo(ctx context.Context, id int) (*model.AccountLock, error) {
	var info = &model.AccountLock{}
	res, err := redis.HashGetVar(r.key, strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	err = res.Struct(&info)
	return info, err
}

// 设置信息
func (r *lock_info_redis) SetInfo(ctx context.Context, info *model.AccountLock) error {
	return redis.HashSetVar(r.key, strconv.Itoa(info.Id), info)
}

// 删除信息
func (r *lock_info_redis) DelInfo(ctx context.Context, id int) error {
	return redis.HashDel(r.key, strconv.Itoa(id))
}
