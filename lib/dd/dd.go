package dd

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"wss/lib/e"
	"wss/lib/setting"
	"wss/lib/utils"
)

type UCUtils struct {
	PlatformKey    string `json:"platform_key"`
	PlatformSecret string `json:"platform_secret"`
	Token          string `json:"token"`
}

var UC_UTILS_BASE_URL = ""

const (
	UC_MY_PLATFORM_KEY   = "material"
	UC_API_GET_TOKEN     = "platform_auth/get_access_token" //获取token
	UC_API_MOBILE_CHECK  = "platform/v1/mobile_check"       //检测手机号是否注册
	UC_API_MOBILE_REG    = "platform/v1/mobile_reg"         //手机号注册
	UC_API_GET_USER_INFO = "api/v1/get_user_info"           //用户信息

)

func (d *UCUtils) Setup(platform_key, platform_secret string) {
	d.PlatformKey = platform_key
	d.PlatformSecret = platform_secret
}

func (d *UCUtils) GetToken() (string, error) {
	data := new(TokenData)

	if err := httpPostJson(setting.PlatformSetting.ApiBaseUrl+UC_API_GET_TOKEN, map[string]string{
		"platform_key":    setting.PlatformSetting.PlatformKey,
		"platform_secret": setting.PlatformSetting.PlatformSecret,
	}, map[string]string{}, &data); err != nil {
		return "", err
	}
	if data.Code == 0 {
		d.Token = data.Data.Token
		log.Println(d.Token)
		return data.Data.Token, nil
	} else {
		return "", errors.New("获取token失败")
	}
}

// 查询是否注册
func (d *UCUtils) GetUserInfoOrMobile(mobile string) (UserInfoData, error) {
	data := new(UserInfoData)

	if err := httpPostJson(setting.PlatformSetting.ApiBaseUrl+UC_API_MOBILE_CHECK, map[string]string{
		"mobile": mobile,
	}, map[string]string{
		"Authentication": d.Token,
	}, &data); err != nil {
		return UserInfoData{}, err
	}
	log.Println(data)
	if data.Code == 0 {

		return *data, nil
	} else {
		return UserInfoData{}, errors.New(data.Msg)
	}
}

// 注册
func (d *UCUtils) UserMobileReg(mobile string) (UserInfoData, error) {
	data := new(UserInfoData)

	if err := httpPostJson(setting.PlatformSetting.ApiBaseUrl+UC_API_MOBILE_REG, map[string]string{
		"mobile": mobile,
	}, map[string]string{
		"Authentication": d.Token,
	}, &data); err != nil {
		return UserInfoData{}, err
	}
	log.Println(data)
	if data.Code == 0 {
		return *data, nil
	} else {
		return UserInfoData{}, errors.New(data.Msg)
	}
}

// 获取用户信息
func (d *UCUtils) UserGetInfo(user_token string) (UserInfoData, error) {
	data := new(UserInfoData)

	if err := httpPostJson(setting.PlatformSetting.ApiBaseUrl+UC_API_GET_USER_INFO, e.GetEmptyStruct(), map[string]string{
		"Authorization": user_token,
		"PlatformKey":   setting.PlatformSetting.PlatformKey,
	}, &data); err != nil {
		return UserInfoData{}, err
	}
	log.Println(data)
	if data.Code == 0 {
		return *data, nil
	} else {
		return *data, errors.New(data.Msg)
	}
}

// Http Post Json 请求
func httpPostJson(url string, body interface{}, header map[string]string, cb interface{}) error {

	requestBody := utils.JsonEncode(body)
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	for i, v := range header {
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
	utils.JsonDecode(string(b), &cb)
	return nil
}
