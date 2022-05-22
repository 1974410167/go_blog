package e

const (
	SUCCESS         = 200
	INVALID_PARAMS  = 400
	SERVERERROR     = 500
	ID_IS_NOT_EXIST = 1001
)

var MsgFlags = map[int]string{
	SUCCESS:         "成功",
	INVALID_PARAMS:  "参数错误",
	SERVERERROR:     "服务器错误",
	ID_IS_NOT_EXIST: "id不存在",
}
