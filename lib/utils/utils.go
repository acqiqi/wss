package utils

import "wss/lib/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}

// 结构体转map
func StructToMap(in_data interface{}) (out_data map[string]interface{}, err error) {
	json_str := JsonEncode(in_data)
	err = JsonDecode(json_str, &out_data)
	return
}
