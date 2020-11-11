package mqtt

//回调结构体
type MqttDataModel struct {
	Action    string      `json:"action"`     //行为 也就是路由
	Data      interface{} `json:"data"`       //数据内容
	Callback  string      `json:"callback"`   //回调 一般存储url 或者订阅报文
	Code      int         `json:"code"`       //状态码
	Message   string      `json:"message"`    //发送消息以及返回消息
	MessageId string      `json:"message_id"` //消息id 用于对照callback的
}
