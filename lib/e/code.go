package e

const (
	SUCCESS        = 0
	ERROR          = 1
	INVALID_PARAMS = 400

	ERROR_AUTH_CHECK_PASSWORD      = 20000
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	ERROR_COMPANY_HEADER = 30001
	ERROR_COMPANY_NOT    = 30002

	// 主要是用在ucenter
	API_NOT_AUTH_CODE               = 42001 //超时
	API_USER_STOP_CODE              = 42002 //当前用户停用
	API_USER_DISABLED_CODE          = 42003 //禁用
	API_NOT_AUTH_REG_CODE           = 42008 //用户没有手机绑定注册
	API_NOT_AUTH_BIND_CODE          = 42009 //沒有綁定用戶信息 == 沒註冊
	API_USER_VALIDATE_IMEI_ERR_CODE = 42004 //imei不对
	API_PLATFORM_ERR                = 43001 //平台不存在
)
