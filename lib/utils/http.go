package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const DEFAULT_PLATFORM_KEY = "MATERIAL"

// Http Get Json 请求
func HttpGetJson(url string, cb interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return errors.New(string(b))
	}
	JsonDecode(string(b), &cb)
	return nil
}

// Http Get String 请求
func HttpGetString(url string) (cb string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return "", errors.New(string(b))
	}
	return string(b), nil
}

// Http Post Json 请求
func HttpPostJson(url string, body interface{}, cb interface{}) error {

	requestBody := JsonEncode(body)
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("PlatformKey", DEFAULT_PLATFORM_KEY)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return errors.New(string(b))
	}
	strs := string(b)
	log.Println(strs)
	log.Println(string(b))
	return JsonDecode(string(b), &cb)
}

// Http Post Json 请求 带Header
func HttpPostJsonHeader(url string, body interface{}, headers map[string]string, cb interface{}) error {

	requestBody := JsonEncode(body)
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("PlatformKey", DEFAULT_PLATFORM_KEY)
	for i, v := range headers {
		req.Header.Set(i, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return errors.New(string(b))
	}
	log.Println(string(b))
	JsonDecode(string(b), &cb)
	return nil
}

// Http Post Json 请求 不带回调
func HttpPostJsonNotCallback(url string, body interface{}, platform_key string) error {
	requestBody := JsonEncode(body)
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("PlatformKey", platform_key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return errors.New(string(b))
	}
	log.Println(string(b))
	return nil
}

//保存图片
func HttpGetDownload(url string, filepath string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	_ = ioutil.WriteFile(filepath, body, 0755)
}

func HttpGetIO(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	//defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp.Body, nil
}

// Post获取Bytes
func HttpPostBytes(url string, body interface{}) (cb []byte, err error) {
	requestBody := JsonEncode(body)
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, errors.New(string(b))
	}
	return b, nil
}
