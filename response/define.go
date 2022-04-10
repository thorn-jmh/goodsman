package response

import "net/http"

var (
	SUCCESS           int = 0   //Success
	AUTH_ERROR        int = -1  //用户权限不足
	REST_MONEY_ERROR  int = -2  //用户余额不足
	NO_CERTAIN_GOODS  int = -10 // 不存在该物品
	PARAMS_ERROR      int = -11 //传入参数错误
	DATABASE_ERROR    int = 1   //数据库错误
	NO_BORROW_RECORDS int = 2   //没有借物记录
	STOCK_ERROR       int = 11  //物品供应不足
	FEISHU_ERROR      int = 2   //与飞书交互错误
)

var msgFlags = map[int]string{
	SUCCESS:           "Success",
	AUTH_ERROR:        "用户权限不足",
	REST_MONEY_ERROR:  "用户余额不足",
	NO_CERTAIN_GOODS:  "不存在该物品",
	PARAMS_ERROR:      "传入参数错误",
	DATABASE_ERROR:    "数据库错误",
	NO_BORROW_RECORDS: "用户没有借用该物品的记录",
	STOCK_ERROR:       "物品供应不足",
	FEISHU_ERROR:      "飞书交互错误",
}

var httpStatus = map[int]int{
	SUCCESS:           http.StatusOK,
	AUTH_ERROR:        http.StatusUnauthorized,
	REST_MONEY_ERROR:  http.StatusUnauthorized,
	NO_CERTAIN_GOODS:  http.StatusBadRequest,
	PARAMS_ERROR:      http.StatusBadRequest,
	DATABASE_ERROR:    http.StatusInternalServerError,
	NO_BORROW_RECORDS: http.StatusBadRequest,
	STOCK_ERROR:       http.StatusInternalServerError,
	FEISHU_ERROR:      http.StatusInternalServerError,
}
