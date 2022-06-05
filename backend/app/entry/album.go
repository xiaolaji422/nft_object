package entry

//map[albumId:1.00513963e+08
// gId:1.02501776e+08 gName:iBox-暹罗警长 gNum:1004 gStatus:6 isBuy:1 ownerId:1.1417964e+07 ownerName:iBox priceCny:78888 tokenId:124955]

type Detail struct {
	ID        int     `json:"gId"`
	Name      string  `json:"gName"`
	AlbumId   int     `json:"albumId"`
	Price     float64 `json:"priceCny"`
	Number    string  `json:"gNum"`
	Status    int     `json:"gStatus"`
	IsBuy     int     `json:"isBuy"`
	OwnerId   int     `json:"ownerId"`
	OwnerName string  `json:"ownerName"`
	TokenId   string  `json:"tokenId"`
}
