package mqtt

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
	"log"
)

var SubscribeList []string

// SubscribeHandler
//var SubscribeHandler mqtt.MessageHandler

// sub消息回调
func SubscribeCallback(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("收到消息")
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	fmt.Println("MsgId--->", msg.MessageID())

	data := msg.Payload()
	str := string(data)
	log.Println(str)

	//Publish("/device/ambit/test/pub", 2, "false")

	//encodeString := base64.StdEncoding.EncodeToString(data)
	//cache.Cache.Set(msg.Topic(), encodeString, time.Second*10)
}

// 新增订阅
func AddSubscribe(subscride string) error {
	//查询订阅是否重复
	for _, v := range SubscribeList {
		if subscride == v {
			return errors.New("重复订阅")
		}
	}

	if Mqttclient.Subscribe(subscride, 2, SubscribeCallback); MqttToken.Wait() && MqttToken.Error() != nil {
		fmt.Println(MqttToken.Error())
		return MqttToken.Error()
	}
	SubscribeList = append(SubscribeList, subscride)
	fmt.Println("addSub----->list:", SubscribeList)
	return nil
}

func RemoveSubscride(subscride string) error {
	if Mqttclient.Subscribe(subscride, 2, SubscribeCallback); MqttToken.Wait() && MqttToken.Error() != nil {
		fmt.Println(MqttToken.Error())
		return MqttToken.Error()
	}
	// 删除订阅List
	for i, v := range SubscribeList {
		if subscride == v {
			SubscribeList = append(SubscribeList[:i], SubscribeList[i+1:]...)
		}
		fmt.Println("unSub----->list:", SubscribeList)
	}
	return nil
}
