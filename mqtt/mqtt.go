package mqtt

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
	"wss/lib/setting"
)

var Mqttclient mqtt.Client
var MqttToken mqtt.Token

func Init() {
	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	fmt.Println("Mqtt Init....")
	opts := mqtt.NewClientOptions().
		AddBroker("tcp://" + setting.MqttSetting.ServerIp + ":" + setting.MqttSetting.ServerPort).
		SetClientID(setting.MqttSetting.ClientID)
	Mqttclient = mqtt.NewClient(opts)
	// 链接mqtt 服务器
	if MqttToken = Mqttclient.Connect(); MqttToken.Wait() && MqttToken.Error() != nil {
		fmt.Println("Mqtt Connect Error!")
		//panic(MqttToken.Error())
	}
	MqttToken.Wait()
	// 订阅系统
	if MqttToken := Mqttclient.Subscribe(setting.MqttSetting.SysSubscribe, 0, SubscribeCallback); MqttToken.Wait() && MqttToken.Error() != nil {
		fmt.Println("Mqtt Sub System Err : ", MqttToken.Error())
	}
	go loopConntect()
}

// 自动重连机制
func loopConntect() {
	for {
		time.Sleep(time.Second * 10)
		if !Mqttclient.IsConnected() {
			if MqttToken = Mqttclient.Connect(); MqttToken.Wait() && MqttToken.Error() != nil {
				//panic(MqttToken.Error())
				log.Println("Mqtt ReConnect Err")
			} else {
				log.Println("Mqtt ReConnect Ok")
				// 订阅系统
				if MqttToken := Mqttclient.Subscribe("/vhake/hub/pub/face", 0, SubscribeCallback); MqttToken.Wait() && MqttToken.Error() != nil {
					fmt.Println("Mqtt Sub System Err : ", MqttToken.Error())
				}
			}
		}
	}

}
