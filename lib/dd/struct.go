package dd

type UcenterUsers struct {
	Id             int64   `json:"id"`
	Username       string  `json:"username"`         // 账号
	Password       string  `json:"password"`         // 密码
	Mobile         string  `json:"mobile"`           // 手机号
	Nickname       string  `json:"nickname"`         // 昵称
	Email          string  `json:"email"`            // 邮箱
	Avatar         string  `json:"avatar"`           // 头像
	Gender         string  `json:"gender"`           // 性别
	Status         int     `json:"status"`           // 状态 0停用 1启用
	RoleType       int     `json:"role_type"`        // 0无权限 1小区合伙人 2区县级合伙人 3城市合伙人 4省合伙人
	Score          int     `json:"score"`            // 积分
	Money          float64 `json:"money"`            // 余额
	OkMoney        float64 `json:"ok_money"`         // 可提现余额
	NoMoney        float64 `json:"no_money"`         // 不可提现金额
	LastLoginIp    string  `json:"last_login_ip"`    // 最后一次登录ip
	LastLoginTime  int     `json:"last_login_time"`  // 最后一次登录时间戳
	LastLongitude  float64 `json:"last_longitude"`   // 最后一次经度
	LastLatitude   float64 `json:"last_latitude"`    // 最后一次维度
	ShareOne       int     `json:"share_one"`        // 一级分享
	ShareTwo       int     `json:"share_two"`        // 二级分享
	StOne          int     `json:"st_one"`           // 一级师徒
	StTwo          int     `json:"st_two"`           // 二级师徒
	UserKey        string  `json:"user_key"`         // 用户注册唯一key
	WechatUnionid  string  `json:"wechat_unionid"`   // 微信相关unionid
	RegType        int     `json:"reg_type"`         // 注册类型 0手机号验证码 1账号
	RegSource      string  `json:"reg_source"`       // 注册来源 例如 手机 微信 小程序
	RegPlatformKey string  `json:"reg_platform_key"` // 从哪个平台注册的
	IsVip          int     `json:"is_vip"`           // 是否是vip
	BindUserinfo   int     `json:"bind_userinfo"`    // 是否绑定用户信息
	VipEndTime     string  `json:"vip_end_time"`     // vip到期时间
	IsPayPassword  int     `json:"is_pay_password"`  // 是否填写支付密码
	PayPassword    string  `json:"pay_password"`     // 支付密码
	OldManagerCuid int     `json:"old_manager_cuid"`
	IsWork         int     `json:"is_work"`         // 是否工人0否 1认证 2通过 -1拒绝
	IsAuth         int     `json:"is_auth"`         // 是否实名认证 0 否 1审核 2通过 -1拒绝
	IdcardId       string  `json:"idcard_id"`       // 身份证号
	IdcardBom      string  `json:"idcard_bom"`      // 身份证背面
	IdcardTop      string  `json:"idcard_top"`      // 身份证正面
	IdcardRen      string  `json:"idcard_ren"`      // 手持证件
	IdcardAddress  string  `json:"idcard_address"`  // 身份证地址
	IdcardNation   string  `json:"idcard_nation"`   // 身份民族
	IdcardGender   string  `json:"idcard_gender"`   // 性别
	IdcardProvince string  `json:"idcard_province"` // 省
	IdcardCity     string  `json:"idcard_city"`     // 市
	IdcardArea     string  `json:"idcard_area"`     // 区
}

type CallbackData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 登录返回Token 结构
type TokenData struct {
	*CallbackData
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

// 用户Userinfo 结构
type UserInfoData struct {
	*CallbackData
	Data struct {
		Token    string       `json:"token"`
		UserInfo UcenterUsers `json:"user_info"`
	} `json:"data"`
}
