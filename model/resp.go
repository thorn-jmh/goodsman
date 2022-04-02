package model

type GetGoodsMsgResp struct {
	Goods_auth int      `json:"goods_authority"`
	Number     int      `json:"number"`
	State      int      `json:"state"`
	Owner      string   `json:"owner"`
	Goods_msg  Goodsmsg `json:"goods_msg"`
}

type UserAuthResp struct {
	Name      string  `json:"name"`
	Authority int     `json:"authority"`
	Money     float64 `json:"money"`
}

type AddNewGoodsResp struct {
	GoodsId string `json:"goods_uuid"`
}

type GetUserIDResp struct {
	EmployeeID   string `json:"employee_id"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
