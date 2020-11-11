package mqtt

const (
	ApiSuccess = 0
	ApiError   = 1
	ApiNotAuth = 45000 //未登录
)

var statusText = map[int]string{
	ApiSuccess: "操作成功",
	ApiError:   "操作失败",
	ApiNotAuth: "未登录或登录超时",
}

//获取code对应错误字段
func GetApiMsg(status int) (msg string) {
	return statusText[status]
}
