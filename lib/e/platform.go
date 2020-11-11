package e

import (
	"wss/lib/utils"
	"wss/models"
)

const (
	PLATFORM_ACTION_PROJECT_RECEIVE  = "PROJECT_RECEIVE"  //项目接收
	PLATFORT_ACTION_MATERIAL_RECEIVE = "MATERIAL_RECEIVE" //下料单接收
	PLATFORT_ACTION_SEND             = "SEND"             // 发货
	PLATFORT_ACTION_SEND_RECEIVER    = "SEND_RECEIVER"    // 发货收货
	PLATFORT_ACTION_PR               = "PR"               // 请款申请
)

// 用于Callback 请求数据集合
type HttpCallbackData struct {
	Code        int         `json:"code"`   // 状态码 默认0成功 1失败 其他专用错误码
	Msg         string      `json:"msg"`    // 主体消息
	Action      string      `json:"action"` // 行为 利用行为二次解析Data结构体
	CallbackUrl string      `json:"callback_url"`
	Data        interface{} `json:"data"` // 主体数据根据Action反序列
}

// 项目接收回调结构体
type PlatformProjectReceiveCallback struct {
	Id          string         `json:"project_id"`
	ProjectName string         `json:"project_name"` // 项目名称
	State       int            `json:"state"`        // 0 在建中 1已完成
	CompanyId   int64          `json:"company_id"`   //关联企业
	Company     models.Company `json:"company"`      //关联企业
	BindState   int            `json:"bind_state"`   // 绑定工程账号的状态,手动新建的项目默认已绑定.未绑定,待处理,已绑定
	PlatformKey string         `json:"platform_key"` // 平台key
	PlatformUid string         `json:"platform_uid"` // 平台用户id
	PlatformId  string         `json:"platform_id"`  // 平台用户id
	CreatedAt   utils.Time     `json:"created_at"`
	Status      int            `json:"status"` // 0未同步 1已同步
	SupplierId  string         `json:"supplier_id"`
}

type PlatformPRCreateCallback struct {
}

// 实现回调
func (d *HttpCallbackData) RequestCallback(cb interface{}) error {
	cb_str := utils.JsonEncode(d)
	return utils.HttpPostJson(d.CallbackUrl, cb_str, &cb)
}
