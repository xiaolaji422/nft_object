package service

import (
	"context"
	"errors"
	"nft_object/app/core"
	"nft_object/app/dao"
	"nft_object/app/model"
	"nft_object/app/repo/redis"
	"nft_object/library/helper"
)

type IAccount interface {
	Save(ctx context.Context, account, params string) error
	List(ctx context.Context) ([]*model.AdminAccount, error)
}

var AccountImpl = func() IAccount {
	return &account{
		DBProxy: &core.DBProxy{
			D: dao.AdminAccount.M,
		},
	}
}

type account struct {
	*core.DBProxy
}

func (s *account) Save(ctx context.Context, account, params string) error {
	// 保存
	var info = model.AdminAccount{
		Account: account,
		Info:    params,
		Enabled: 1,
		AdminId: helper.GetAdminId(ctx),
	}

	cnt, err := s.GetDB().Ctx(ctx).Where("account != ?", account).Count()

	if err != nil {
		return err
	}
	if cnt > 5 {
		return errors.New("最多持有6个账号")
	}
	sqlRes, err := s.GetDB().Ctx(ctx).Data(info).Save()
	if id, err := sqlRes.LastInsertId(); err != nil {
		return err
	} else {
		if id > 0 {
			// 删除redis详情，设置锁单列表
			s.updateRedis(ctx, int(id))
		}
	}
	return err
}

// 返回列表
func (s *account) List(ctx context.Context) ([]*model.AdminAccount, error) {
	var (
		admin_id = helper.GetAdminId(ctx)
		data     = make([]*model.AdminAccount, 0)
	)
	res, err := s.GetDB().Ctx(ctx).Where("admin_id", admin_id).Limit(10).FindAll()
	if err != nil {
		return nil, err
	}
	err = res.Structs(&data)
	return data, err
}

func (s *account) updateRedis(ctx context.Context, id int) error {
	res, err := s.CheckInfo(id)
	if err != nil {
		return err
	} else {
		var info = &model.AdminAccount{}
		if err := res.Struct(&info); err != nil {
			return err
		}
		err = redis.AccountRedisImpl().SetInfo(ctx, info)
	}
	return err
}
