package service

// 业务层代码
import (
	"context"
	"nft_object/app/core"
	"nft_object/app/dao"
	"nft_object/library/pagenation"
)

// 公告的接口

// 公告业务对外提供的服务
var NFTAlbumImpl = func() INFTAlbum {
	return &album{
		DBProxy: &core.DBProxy{
			D: dao.NftAlbum.M,
		},
	}
}

// 过期时间

type INFTAlbum interface {
	// 搜索商品
	Search(ctx context.Context, platform int, keyword string) (pagenation.PageReult, error)
}

// 公告业务类
type album struct {
	*core.DBProxy
}

//  获取新的公告 五分钟以内的
func (s *album) Search(ctx context.Context, platform int, keyword string) (pagenation.PageReult, error) {
	var (
		tp    = pagenation.GetPagenation(s.D.Ctx(ctx))
		page  = 1
		limit = 20
		// admin_name = helper.GetRtx(ctx)
	)
	tp = tp.PageInfo(page, limit)
	tp.OrderInfo(pagenation.OrderItem{
		Property: "update_time",
		Sort:     "desc",
	})
	tp.Rules(
		&pagenation.Rule{Field: "name", Operator: pagenation.OP_LIKEALL},
	)
	platform = 1
	where := core.MapI{
		"platform": platform,
	}
	if len(keyword) > 0 {
		where["name"] = keyword
	}
	res, err := tp.Search(where)
	if err != nil {
		return res, err
	}
	return res, err
}
