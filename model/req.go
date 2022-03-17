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
