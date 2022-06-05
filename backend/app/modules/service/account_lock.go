package service

import (
	"context"
	"errors"
	"nft_object/app/core"
	"nft_object/app/dao"
	"nft_object/app/entry"
	"nft_object/app/model"
	"nft_object/library/helper"

	"github.com/gogf/gf/util/gconv"
)

type IAccountLock interface {
	Save(ctx context.Context, info *entry.AccountData) error
	List(ctx context.Context) ([]*model.AccountLock, error)
}

var AccountLockImpl = func() IAccountLock {
	return &account_lock{
		DBProxy: &core.DBProxy{
			D: dao.AccountLock.M,
		},
	}
}

type account_lock struct {
	*core.DBProxy
}

// 保存自动锁单
func (s *account_lock) Save(ctx context.Context, info *entry.AccountData) error {
	// 保存
	info.AdminId = helper.GetAdminId(ctx)
	info.Enabled = 1
	if gconv.Float64(info.Max) < gconv.Float64(info.Min) {
		return errors.New("最高价格应该大于等于最低价格")
	}
	if gconv.Float64(info.Min) <= 0 {
		return errors.New("请输入大于0的最低价格")
	}
	check_where := core.MapI{
		"admin_id": info.AdminId,
		"enabled":  1,
	}
	cnt, err := s.GetDB().Ctx(ctx).Where(check_where).Count()

	if err != nil {
		return err
	}
	if cnt > 4 {
		return errors.New("最多建立4个自动锁单订单")
	}
	_, err = s.GetDB().Ctx(ctx).Data(info).Save()
	// 删除redis
	return err
}

// 返回自动锁单列表
func (s *account_lock) List(ctx context.Context) ([]*model.AccountLock, error) {
	var (
		data         = make([]*model.AccountLock, 0)
		search_where = core.MapI{
			"admin_id": helper.GetAdminId(ctx),
			"enabled":  1,
		}
	)

	res, err := s.GetDB().Ctx(ctx).Where(search_where).Limit(10).FindAll()
	if err != nil {
		return nil, err
	}
	err = res.Structs(&data)
	return data, err
}
