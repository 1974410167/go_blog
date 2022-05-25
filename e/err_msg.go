package e

const (
	SUCCESS         = 200
	INVALID_PARAMS  = 400
	SERVERERROR     = 500
	ID_IS_NOT_EXIST = 1001
	Create_ERROR    = 1002
	Auth_Error      = 1003
	Auth_Expire     = 1004
	DeCode_Error    = 1005
)

var MsgFlags = map[int]string{
	SUCCESS:         "成功",
	INVALID_PARAMS:  "参数错误",
	SERVERERROR:     "服务器错误",
	ID_IS_NOT_EXIST: "id不存在",
	Create_ERROR:    "创建失败",
	Auth_Error:      "认证失败",
	Auth_Expire:     "认证过期",
	DeCode_Error:    "解码失败",
}
