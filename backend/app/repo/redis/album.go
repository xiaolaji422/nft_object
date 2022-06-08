package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"nft_object/app/entry"
	"nft_object/library/redis"
	"strings"
)

const (
	refresh_product_task_key = "nft_object:refresh:product:task:list"
	product_list_key         = "nft_object:album:product:{{product_id}}:zset"
)

type IAlbumRedis interface {
	//  detail
	AlbumDetail(ctx context.Context, albumID string, page int, limit int) ([]*entry.Detail, error)
	// start get detail
	SetAlbumList(ctx context.Context, albumID string) error
	// setting lock list to refresh detail list
	SetLockList(ctx context.Context, albumID string) error
}

var AlbumRedisImpl = func() IAlbumRedis {
	return &album{}
}

type album struct {
}

// impl  detail
func (r *album) AlbumDetail(ctx context.Context, albumID string, page int, limit int) ([]*entry.Detail, error) {
	var data = make([]*entry.Detail, 0)
	var product_key = strings.Replace(product_list_key, "{{product_id}}", albumID, 1)
	res, err := redis.Zset().Zrangebyscore(product_key, 0.0, 99999999999, page, limit)
	if err != nil {
		return data, err
	}
	if res.IsEmpty() {
		return data, err
	}
	list := res.Array()
	for _, item := range list {
		if str, ok := item.(string); ok {
			info := r.handleDetail(str)
			data = append(data, info)
		}
	}
	return data, nil
}

func (r *album) handleDetail(str string) *entry.Detail {
	var (
		info = &entry.Detail{}
	)
	if err := json.Unmarshal([]byte(str), info); err != nil {
		fmt.Println("handleDetail  json unmarshal error:"+err.Error(), str)
	}
	return info
}

// impl  set album list  for start get detail
func (r *album) SetAlbumList(ctx context.Context, albumID string) error {
	return nil
}

// impl  set album list  for start get detail
func (r *album) SetLockList(ctx context.Context, albumID string) error {
	return nil
}
