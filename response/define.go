package response

import "net/http"

var (
	SUCCESS      int = 0   //Success
	AUTH_ERROR   int = -1  //用户权限不足
	PARAMS_ERROR int = -11 //传入参数错误
)

var msgFlags = map[int]string{
	SUCCESS:      "Success",
	AUTH_ERROR:   "用户权限不足",
	PARAMS_ERROR: "传入参数错误",
}

var httpStatus = map[int]int{
	SUCCESS:      http.StatusOK,
	AUTH_ERROR:   http.StatusUnauthorized,
	PARAMS_ERROR: http.StatusBadRequest,
}
