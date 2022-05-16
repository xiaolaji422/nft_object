package service

// 业务层代码
import (
	"context"
	"fmt"
	"nft_object/app/core"
	"nft_object/app/dao"
	"nft_object/app/repo/redis"
	"nft_object/library/helper"
	"nft_object/library/pagenation"
	"time"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/util/gconv"
)

// 公告的接口

// 公告业务对外提供的服务
var NoticeImpl = func() INotice {
	return &notice{
		DBProxy: &core.DBProxy{
			D: dao.NoticeWarning.M,
		},
	}
}

const layout = ""
const NEW_NOTICE_ID = "nft:new_notice:"     // 人的redis
const NEW_NOTICE_SYS = "nft:new_notice:sys" // 系统redis
const NEW_NOTICE_DURATION = 5               // 过期时间

type INotice interface {
	//  获取最新公告
	GetNewNotice(ctx context.Context) (gdb.Result, bool, error)
	// 历史公告
	GetHistoryNotice(ctx context.Context, params core.MapI) (interface{}, error)
}

// 公告业务类
type notice struct {
	*core.DBProxy
}

//  获取新的公告 五分钟以内的
func (s *notice) GetNewNotice(ctx context.Context) (gdb.Result, bool, error) {
	// 从redis取最新的id
	var (
		isWarning  = false
		login_name = helper.GetRtx(ctx)
	)
	last_time, err := s.handleNoticeNewer(login_name)
	if err != nil {
		return nil, isWarning, err
	}

	// 查找五分钟以内的
	where := core.MapI{
		"notice_time >=": time.Now().Add(-time.Minute * 5).Format("2006-01-02 15:04:05"),
	}
	res, err := s.GetDB().Ctx(ctx).Where(where).Limit(5).OrderBy("notice_time desc").FindAll()
	if res.IsEmpty() {
		res, err = s.GetDB().Ctx(ctx).Limit(1).OrderBy("notice_time desc").FindAll()
	}
	if err != nil {
		return nil, isWarning, err
	}
	var newer_time = time.Time{}
	if len(res) > 0 {
		for _, v := range res {
			if notice_time, ok := v["notice_time"]; ok {
				time_v, err := time.Parse("2006-01-02 15:04:05", notice_time.String())
				if err != nil {
					return res, isWarning, err
				}
				if time_v.After(newer_time) {
					newer_time = time_v
				}
			}
		}
	}

	// 需要响铃
	if newer_time.After(last_time) {
		isWarning = true
	}
	err = redis.NoticeRedisImpl().SetAdminNewerTime(login_name, newer_time)
	return res, isWarning, err
}

//  历史公告
func (s *notice) GetHistoryNotice(ctx context.Context, params core.MapI) (interface{}, error) {
	var (
		tp    = pagenation.GetPagenation(s.D.Ctx(ctx))
		page  = gconv.Int(helper.GetMapValue(params, "page", 1))
		limit = gconv.Int(helper.GetMapValue(params, "limit", 50))
		// admin_name = helper.GetRtx(ctx)
	)

	// last_id, err := s.handleNoticeNewer(admin_name)
	// if err != nil {
	// 	return nil, err
	// }
	// 分页信息
	tp = tp.PageInfo(page, limit)
	tp.OrderInfo(pagenation.OrderItem{
		Property: "notice_time",
		Sort:     "desc",
	})
	tp.Rules(
		&pagenation.Rule{Field: "notice_time", Operator: pagenation.OP_GT},
	)

	where := core.MapI{
		"notice_time": time.Now().Add(-time.Minute * 5).Format("2006-01-02 15:04:05"),
	}
	res, err := tp.Search(where)
	if err != nil {
		return res, err
	}
	return res, err
	// var newer_time = time.Time{}
	// if len(res.Items) > 0 {
	// 	for _, v := range res.Items {
	// 		if notice_time, ok := v["notice_time"]; ok {

	// 			time_v, err := time.Parse("2006-01-02 15:04:05", notice_time.String())
	// 			if err != nil {
	// 				return res, err
	// 			}
	// 			if time_v.After(newer_time) {
	// 				newer_time = time_v
	// 			}

	// 		}
	// 	}
	// }

	// redis.NoticeRedisImpl().SetSysNewerTime(newer_time)

	// return res, err
}

// 获取最新id
func (s *notice) handleNoticeNewer(login_name string) (time.Time, error) {
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		if r := recover(); r != nil {
			fmt.Println("handleNewNotice error:", r)
		}
	}()
	last_time, err := redis.NoticeRedisImpl().GetAdminNewerTime(login_name)
	return last_time, err
}
