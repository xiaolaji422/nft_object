package service

import (
	"context"
	"errors"
	"nft_object/app/core"
	"nft_object/app/dao"
	"nft_object/app/model"
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
	_, err = s.GetDB().Ctx(ctx).Data(info).Save()
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
