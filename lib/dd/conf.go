package dd

// 接口地址
const (
	DD_API_AUTOLOGIN_GETSMS = "mobile_auth/get_reg_sms"    //自动登录获取手机验证码
	DD_API_MOBILE_LOGIN     = "mobile_auth/password_login" //自动登录
	DD_API_AUTOLOGIN        = "mobile_auth/auto_reg"       //自动登录

	DD_API_AUTH_RE_SMS      = "mobile_auth/get_rp_sms"  //获取修改密码短信
	DD_API_AUTH_RE_PASSWORD = "mobile_auth/re_password" //修改密码

	DD_API_AUTH_GET_USER_INFO = "api/v1/get_user_info" //获取userinfo

)
