package redis

import (
	"context"
	"fmt"
	"nft_object/app/model"
	"nft_object/library/helper"
	"nft_object/library/redis"
	"time"
)

//  锁单redis
// todo    获取某个锁单信息  // 商品id
// 删除某个锁单信息

type IAccountRedis interface {
	GetInfo(ctx context.Context, id int) (*model.AdminAccount, error)
	DelInfo(ctx context.Context, id int) error
	SetInfo(ctx context.Context, info *model.AdminAccount) error
}

var AccountRedisImpl = func() IAccountRedis {
	return &account_redis{
		key:      "nft_object:refresh:account:info:",
		duration: time.Hour * 2,
	}
}

type account_redis struct {
	key      string
	duration time.Duration
}

// 获取信息
func (r *account_redis) GetInfo(ctx context.Context, id int) (*model.AdminAccount, error) {
	var info = &model.AdminAccount{}
	res, err := redis.GetVar(r.handleRedisKey(ctx, id))
	if err != nil {
		return nil, err
	}
	err = res.Struct(&info)
	return info, err
}

// 设置信息
func (r *account_redis) SetInfo(ctx context.Context, info *model.AdminAccount) error {
	return redis.SetTimeOut(r.handleRedisKey(ctx, info.Id), info, r.duration)
}

// 删除信息
func (r *account_redis) DelInfo(ctx context.Context, id int) error {
	_, err := redis.DelVar(r.handleRedisKey(ctx, id))
	return err
}

func (r *account_redis) handleRedisKey(ctx context.Context, id int) string {
	admin_id := helper.GetAdminId(ctx)
	return r.key + fmt.Sprintf("%d_%d", admin_id, id)
}
