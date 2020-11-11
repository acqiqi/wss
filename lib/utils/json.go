package utils

import (
	"encoding/json"
	"errors"
	"fmt"
)

func JsonDecode(body string, obj interface{}) (err error) {
	strByte := []byte(body)
	if err := json.Unmarshal(strByte, &obj); err != nil {
		fmt.Println(err.Error())
		return errors.New("string 格式不正确")
	}
	return nil
}

func JsonDecodeUrls(body string) (obj []string) {
	strByte := []byte(body)
	if err := json.Unmarshal(strByte, &obj); err != nil {
		fmt.Println(err.Error())
		return []string{}
	}
	return obj
}

func JsonEncode(obj interface{}) string {
	b, _ := json.Marshal(obj)
	return string(b)
}
