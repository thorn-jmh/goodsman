package model

type Goodsmsg struct {
	Name        string  `bson:"name" json:"name"`               //名称
	Type        string  `bson:"type" json:"type"`               //型号
	Photo       string  `bson:"photo" json:"photo"`             //照片>>>???????
	Cost        float64 `bson:"cost" json:"cost"`               //金额
	Consumables int     `bson:"consumables" json:"consumables"` //是否为耗材,0：不是，1：是
}

type Goods struct {
	Goods_id   string `bson:"goods_id"`
	Goods_auth int    `bson:"goods_auth"` //权限 0：实习，1：正式，2：非外借
	Number     int    `bson:"number"`     //库存数量
	State      int    `bson:"state"`      //物品状态 0：不在库，1：正常，2：报修，3：报损
	Owner      string `bson:"owner"`      //目前所有者(employee_id)
	// Owner_name string   `bson:"owner_name"`//TODO:
	Goods_msg Goodsmsg `bson:"goods_msg"`
}

type Records struct {
	Employee_id string `bson:"employee_id"`
	Goods_id    string `bson:"goods_id"`
	Date        int64  `bson:"date"`
	State       int    `bson:"state"`
	Del_num     int    `bson:"del_number"`
}

type Manager struct {
	Employee_id string  `bson:"employee_id"`
	Name        string  `bson:"name"`
	Auth        int     `bson:"auth"`
	MaxMoney    float64 `bson:"money"`
}
