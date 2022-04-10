package model

type BorrowGoodsRequest struct {
	EmployeeId string `form:"employee_id" json:"employee_id" binding:"required"`
	GoodsId    string `form:"goods_id" json:"goods_id" binding:"required"`
	GoodsNum   int    `form:"goods_num" json:"goods_num" binding:"required"`
}

type Goodsstate struct {
	NewState int `bson:"state" json:"State"`   // 新状态
	DelNum   int `bson:"number" json:"Number"` // 变化数量
}

type ChangeGoodsStateRequest struct {
	EmployeeId string     `form:"employee_id" json:"employee_id" binding:"required"`
	GoodsId    string     `form:"goods_id" json:"goods_id" binding:"required"`
	GoodsState Goodsstate `form:"goods_state" json:"goods_state" binding:"required"`
}

type AddNewGoodsRequest struct {
	EmployeeId string   `form:"employee_id" json:"employee_id" binding:"required"`
	Number     int      `form:"number" json:"number"`
	GoodsAuth  int      `form:"goods_authority" json:"goods_authority"`
	State      int      `form:"state" json:"state"`
	Owner      string   `form:"owner" json:"owner" default:"Null"`
	GoodsMsg   Goodsmsg `form:"goods_msg" json:"goods_msg" binding:"required"`
}

type ReturnAllGoodsRequest struct {
	EmployeeId string `form:"employee_id" json:"employee_id" binding:"required"`
	GoodsId    string `form:"goods_id" json:"goods_id" binding:"required"`
}

type ReturnGoodsRequest struct {
	EmployeeId string `form:"employee_id" json:"employee_id" binding:"required"`
	GoodsId    string `form:"goods_id" json:"goods_id" binding:"required"`
	DelNum     int    `bson:"number" json:"number"` // 变化数量
}

type GetUserIDRequest struct {
	Code string `json:"code" binding:"required"`
}

type ModifyManagerRequest struct {
	EmployeeId string `json:"employee_id" binding:"required"`
	Name       string `json:"name"`
	Token      string `json:"token"`
}

type ChangeManagerStateRequest struct {
	SuperEid   string `json:"super_admin" binding:"required"`
	ManagerEid string `json:"employee_id" binding:"required"`
	NewAuth    int    `json:"new_auth" binding:"required"`
}

type ChangeManagerMaxMoneyRequest struct {
	SuperEid    string `json:"super_admin" binding:"required"`
	ManagerEid  string `json:"employee_id" binding:"required"`
	NewMaxMoney int    `json:"new_max_money" binding:"required"`
}
