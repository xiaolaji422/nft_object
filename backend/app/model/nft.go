package model

import "github.com/gogf/gf/os/gtime"

// NftAlbum is the golang structure for table t_nft_album.
type NftAlbum struct {
	Id          int         `orm:"id,primary"   json:"id"`           // ID
	Platform    int         `orm:"platform"     json:"platform"`     // 平台
	OriginalId  int         `orm:"original_id"  json:"original_id"`  // 原始的nft的id
	Name        string      `orm:"name"         json:"name"`         // 专辑名称
	CoverPic    string      `orm:"cover_pic"    json:"cover_pic"`    // 封面
	Description string      `orm:"description"  json:"description"`  // 描述
	OnSale      int         `orm:"on_sale"      json:"on_sale"`      // 是否在售
	OnSaleTime  *gtime.Time `orm:"on_sale_time" json:"on_sale_time"` // 发售时间
	SellLimit   int         `orm:"sell_limit"   json:"sell_limit"`   // 发售总量
	CreateTime  *gtime.Time `orm:"create_time"  json:"create_time"`  // 创建时间
	UpdateTime  *gtime.Time `orm:"update_time"  json:"update_time"`  // 更新时间
	Enabled     int         `orm:"enabled"      json:"enabled"`      // 状态描述, 1: 正常使用，0: 删除
}

// NoticeWarning is the golang structure for table t_notice_warning.
type NoticeWarning struct {
	Id          int         `orm:"id,primary"  json:"id"`          // ID
	Platform    int         `orm:"platform"    json:"platform"`    // 公告平台
	OriginalId  int         `orm:"original_id" json:"original_id"` // 原始的公告id
	Name        string      `orm:"name"        json:"name"`        //
	Description string      `orm:"description" json:"description"` // 公告告警描述
	Content     string      `orm:"content"     json:"content"`     //
	NoticeUrl   string      `orm:"notice_url"  json:"notice_url"`  //
	NoticeTime  *gtime.Time `orm:"notice_time" json:"notice_time"` // 公告创建时间
	CreateTime  *gtime.Time `orm:"create_time" json:"create_time"` // 创建时间
	UpdateTime  *gtime.Time `orm:"update_time" json:"update_time"` // 更新时间
	Enabled     int         `orm:"enabled"     json:"enabled"`     // 状态描述, 1: 正常使用，0: 删除
}
