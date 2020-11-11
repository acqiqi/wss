package mqtt

//发送消息
func Publish(topic string, qos byte, payload string) {
	Mqttclient.Publish(topic, qos, false, payload)
}

func PublishByte(topic string, qos byte, payload []byte) {
	Mqttclient.Publish(topic, qos, false, payload)
}

//
//func PublishCallback(topic string, data MqttDataModel) (err error) {
//	//处理最后一段 /gateway/ambit/{{imei}}/pub-sub
//	topic_arr := strings.Split(topic, "/")
//	if len(topic_arr) != 5 {
//		return errors.New("topic不正确")
//	}
//	if topic_arr[len(topic_arr)-1] == "sub" {
//		topic_arr[len(topic_arr)-1] = "pub"
//	} else {
//		return errors.New("topic非sub")
//	}
//
//	pub_data := utils.JsonEncode(data)
//	str_topic := strings.Join(topic_arr, "/")
//	log.Println("PubLish Data:" + pub_data)
//	Publish(str_topic, 0x00, pub_data)
//	return nil
//}
