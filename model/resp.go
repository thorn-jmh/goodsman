package model

type GetGoodsMsgResp struct {
	Goods_auth int      `json:"goods_authority"`
	Number     int      `json:"number"`
	State      int      `json:"state"`
	Owner      string   `json:"owner"`
	Goods_msg  Goodsmsg `json:"goods_msg"`
}

type UserAuthResp struct {
	Authority int     `json:"authority"`
	Money     float64 `json:"money"`
}

type AddNewGoodsResp struct {
	GoodsId string `json:"goods_uuid"`
}

type GetUserIDResp struct {
	Employee_id   string `json:"employee_id"`
	Access_token  string `json:"access_token"`
	Expires_in    int64  `json:"expires_in"`
	Refresh_token string `json:"refresh_token"`
}
