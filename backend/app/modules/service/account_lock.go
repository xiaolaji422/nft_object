package service

import (
	"context"
	"errors"
	"nft_object/app/core"
	"nft_object/app/dao"
	"nft_object/app/entry"
	"nft_object/app/model"
	"nft_object/app/repo/redis"
	"nft_object/library/helper"
	"nft_object/library/str"

	"github.com/gogf/gf/util/gconv"
)

type IAccountLock interface {
	Add(ctx context.Context, info *entry.AccountData) error
	List(ctx context.Context) ([]*model.AccountLock, error)
	Lock(ctx context.Context, id int, params core.MapI) error
	Cacnel(ctx context.Context, id int) error
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
func (s *account_lock) Add(ctx context.Context, info *entry.AccountData) error {
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
	sqlRes, err := s.GetDB().Ctx(ctx).Data(info).FieldsEx("update_time,create_time,id").Insert()
	if err != nil {
		return err
	}
	id, err := sqlRes.LastInsertId()
	if err != nil {
		return err
	}
	if id > 0 {
		// 删除redis详情，设置锁单列表
		err = s.updateRedis(ctx, int(id))
	}
	return err
}

// 锁单
func (s *account_lock) Lock(ctx context.Context, id int, params core.MapI) error {
	// 保存
	var (
		adminId    = helper.GetAdminId(ctx)
		updateData = core.MapI{
			"id":       id,
			"admin_id": adminId,
			"enabled":  1,
		}
		updateWhere = core.MapI{
			"id":       id,
			"admin_id": adminId,
			"enabled":  1,
		}
	)
	info, err := s.CheckInfo(id)

	if err != nil {
		return err
	}

	if info["enabled"].Int() != 1 {
		return errors.New("单据状态异常，不支持此操作")
	}

	// 更新数据库
	_, err = s.GetDB().Ctx(ctx).Where(updateWhere).Data(updateData).Update()
	if err != nil {
		return err
	}
	// 发送消息
	msg := NewMessage("API", str.String(info))
	err = SendMsgImpl().SendMsg(ctx, info["admin_id"].String(), msg)
	if err != nil {
		return err
	}

	if id > 0 {
		// 删除redis详情，设置锁单列表
		err = redis.LockInfoRedisImpl().DelInfo(ctx, int(id))
	}

	return err
}

func (s *account_lock) Cacnel(ctx context.Context, id int) error {
	// 保存
	var (
		adminId    = helper.GetAdminId(ctx)
		updateData = core.MapI{
			"id":       id,
			"admin_id": adminId,
			"enabled":  0,
		}
		updateWhere = core.MapI{
			"id":       id,
			"admin_id": adminId,
			"enabled":  1,
		}
	)
	info, err := s.CheckInfo(id)

	if err != nil {
		return err
	}

	if info["enabled"].Int() != 1 {
		return errors.New("单据状态异常，不支持此操作")
	}

	if info["admin_id"].Int() != adminId {
		return errors.New("操作异常(code:400003)")
	}

	sqlRes, err := s.GetDB().Ctx(ctx).Where(updateWhere).Data(updateData).Update()

	if id, err := sqlRes.LastInsertId(); err != nil {
		return err
	} else {
		if id > 0 {
			// 删除redis详情，设置锁单列表
			redis.LockInfoRedisImpl().DelInfo(ctx, int(id))
		}
	}
	return err
}

func (s *account_lock) updateRedis(ctx context.Context, id int) error {
	res, err := s.CheckInfo(id)
	if err != nil {
		return err
	} else {
		var info = &model.AccountLock{}
		if err := res.Struct(&info); err != nil {
			return err
		}
		err = redis.LockInfoRedisImpl().SetInfo(ctx, info)
		if err != nil {
			return err
		}
		err = redis.LockRedisImpl().SetLockList(ctx, id)
	}
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
