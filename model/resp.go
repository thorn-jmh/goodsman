package model

type GetGoodsMsgResp struct {
	Goods_auth int      `json:"goods_auth"`
	Number     int      `json:"number"`
	State      int      `json:"state"`
	Owner      string   `json:"owner"`
	Goods_msg  Goodsmsg `json:"goods_msg"`
}

type UserAuthResp struct {
	Authority int     `json:"authority"`
	Money     float64 `json:"money"`
}
