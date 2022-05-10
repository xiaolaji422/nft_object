package service

// 业务层代码
import (
	"context"
	"fmt"
	"nft_object/app/core"
	"nft_object/app/dao"
	"nft_object/library/helper"
	"nft_object/library/pagenation"
	"time"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
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

const NEW_NOTICE_ID = "nft:new_notice:"     // 人的redis
const NEW_NOTICE_SYS = "nft:new_notice:sys" // 系统redis
const NEW_NOTICE_DURATION = 5               // 过期时间

type INotice interface {
	//  获取最新公告
	GetNewNotice(ctx context.Context) (gdb.Result, error)
	// 历史公告
	GetHistoryNotice(ctx context.Context, params core.MapI) (interface{}, error)
}

// 公告业务类
type notice struct {
	*core.DBProxy
}

//  获取新的公告
func (s *notice) GetNewNotice(ctx context.Context) (gdb.Result, error) {
	// 从redis取最新的id
	last_id, err := s.handleNoticeNewer(helper.GetRtx(ctx))
	if err != nil {
		return nil, err
	}

	where := core.MapI{
		"id > ": last_id,
	}
	res, err := s.GetDB().Where(where).Limit(5).OrderBy("create_time desc").FindAll()
	if err != nil {
		return nil, err
	}
	return res, err
}

//  历史公告
func (s *notice) GetHistoryNotice(ctx context.Context, params core.MapI) (interface{}, error) {
	var (
		tp         = pagenation.GetPagenation(s.D)
		page       = gconv.Int(helper.GetMapValue(params, "page", 1))
		limit      = gconv.Int(helper.GetMapValue(params, "limit", 50))
		admin_name = helper.GetRtx(ctx)
	)

	last_id, err := s.handleNoticeNewer(admin_name)
	if err != nil {
		return nil, err
	}
	// 分页信息
	tp = tp.PageInfo(page, limit)
	tp.Rules(
		&pagenation.Rule{Field: "id", Operator: pagenation.OP_GT},
	)
	where := core.MapI{
		"id": last_id,
	}

	return tp.Search(where)
}

// 获取最新id
func (s *notice) handleNoticeNewer(login_name string) (int, error) {
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		if r := recover(); r != nil {
			fmt.Println("handleNewNotice error:", r)
		}
	}()
	var (
		admin_notice_key = NEW_NOTICE_ID + login_name
		admin_new        = 0 // 此人的最新浏览值
		sys_new          = 0 //系统最新浏览值
	)
	// 当前人的最新
	admin_notice, err := g.Redis().Do("GET", admin_notice_key)
	if err != nil {
		return 0, err
	}
	if !g.IsEmpty(admin_notice) {
		admin_new = gconv.Int(admin_notice)
	}
	// 系统最新
	sys_new_notice, err := g.Redis().Do("GET", NEW_NOTICE_SYS)
	if err != nil {
		return 0, err
	}
	if !g.IsEmpty(sys_new_notice) {
		admin_new = gconv.Int(sys_new_notice)
	}
	// 当前用户最新小于系统最新
	if admin_new < sys_new {
		g.Redis().DoWithTimeout(time.Millisecond*5, "SET", admin_notice_key, sys_new)
	}
	g.Redis().DoWithTimeout(time.Millisecond*5, "SET", admin_notice_key, sys_new)
	return admin_new, nil
}
