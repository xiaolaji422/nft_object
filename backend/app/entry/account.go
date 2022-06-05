package entry

type AccountData struct {
	Id        int    `orm:"id,primary"        json:"id"`         // ID
	AdminId   int    `orm:"admin_id"    json:"admin_id"`         // 账户id
	AccountId int    `orm:"account_id,unique" json:"account_id"` // 账号id
	AlbumId   string `orm:"album_id,unique"   json:"album_id"`   // 商品id
	AlbumName string `orm:"album_name"        json:"album_name"` // 商品名称
	Min       string `orm:"min"               json:"min"`        // 最低价格
	Max       string `orm:"max"               json:"max"`        // 最高价格
	Enabled   int    `orm:"enabled"           json:"enabled"`    // 状态描述, 1: 正常使用，0: 删除
}
