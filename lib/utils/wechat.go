package utils

import (
	"errors"
	"fmt"
	"log"
)

type WechatUtils struct {
	// 微信 appid
	Appid string `json:"appid"`
	// 微信 secret
	Secret string `json:"secret"`
	// 默认微信服务器url
	ApiUrl string `json:"api_url"`
	// 小程序接口对应版本 developer为开发版；trial为体验版；formal为正式版；默认为正式版
	MiniprogramType string `json:"miniprogram_type"`

	AccessToken string `json:"access_token"`

	SendMessageData struct {
		// 发送用户的openid
		Touser string `json:"touser"`
		// 订阅消息模板id
		TemplateId string `json:"template_id"`
		// 页面路径
		Page string `json:"page"`
		// 数据
		Data map[string]interface{} `json:"data"`
	}
	SmallQrcodeData struct {
		Path  string `json:"path"`  //页面路径
		Page  string `json:"page"`  //页面 生产二维码时候用这个
		Width int    `json:"width"` //默认430
		Scene string `json:"scene"` //获取参数
	}
}

type WechatValueStruct struct {
	Value string `json:"value"`
}

// 初始化
func (this *WechatUtils) Init(appid string, secret string) (err error) {
	if appid != "" {
		this.Appid = appid
	}
	if secret != "" {
		this.Secret = secret
	}
	this.ApiUrl = "https://api.weixin.qq.com/cgi-bin/"
	this.MiniprogramType = "formal"
	return nil
}

//第一步 获取access token
func (this *WechatUtils) GetAccessToken() error {

	url := this.ApiUrl + fmt.Sprintf("token?grant_type=client_credential&appid=%s&secret=%s",
		this.Appid,
		this.Secret)
	cb := make(map[string]string)
	if err := HttpGetJson(url, &cb); err != nil {
		log.Println(err.Error())
	}
	log.Println(cb)
	if cb["access_token"] == "" {
		return errors.New("获取AccessToken失败")
	} else {
		this.AccessToken = cb["access_token"]
		//SetCacheString("VK_WECHAT_ACCESSTOKEN", this.AccessToken, time.Second*720)
		return nil
	}

}

// 获取微信小程序二维码
func (this *WechatUtils) GetSmallQrcode() (cb []byte, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s",
		this.AccessToken)

	data, err := HttpPostBytes(url, this.SmallQrcodeData)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *WechatUtils) SendMessage() error {
	url := this.ApiUrl + "message/subscribe/send?access_token=" + this.AccessToken

	cb := make(map[string]string)
	if err := HttpPostJson(url, this.SendMessageData, &cb); err != nil {
		return err
	}

	log.Println("cbcbcb")
	log.Println(cb)
	return nil
}
