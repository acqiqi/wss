package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_AUTH_CHECK_PASSWORD:      "账号或密码错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",

	ERROR_COMPANY_HEADER: "公司参数有误",
	ERROR_COMPANY_NOT:    "公司数据有误",

	API_NOT_AUTH_CODE:               "超时",
	API_USER_STOP_CODE:              "当前用户停用",
	API_USER_DISABLED_CODE:          "禁用",
	API_NOT_AUTH_REG_CODE:           "用户没有手机绑定注册",
	API_NOT_AUTH_BIND_CODE:          "沒有綁定用戶信息", //沒有綁定用戶信息 == 沒註冊
	API_USER_VALIDATE_IMEI_ERR_CODE: "imei不对",   //imei不对
	API_PLATFORM_ERR:                "平台不存在",    //平台不存在
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
