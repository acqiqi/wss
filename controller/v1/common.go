package v1

import (
	"github.com/gin-gonic/gin"
	"wss/lib/e"
	"wss/lib/setting"
	"wss/mqtt"
)

func PushMqttMsg(c *gin.Context) {
	data := struct {
		Msg string `json:"msg"`
		//Users  []receiver_service.UserAdd         `json:"users"`
	}{}
	if err := c.BindJSON(&data); err != nil {
		e.ApiErr(c, err.Error())
		return
	}

	mqtt.Publish(setting.MqttSetting.SysPublish+"mdzz", 0x01, data.Msg)

	e.ApiOk(c, "mdzz", e.GetEmptyStruct())

}
