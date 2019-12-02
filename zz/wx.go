package zz

import (
	"fmt"
	"strings"
)

//wx基本信息
type Wx struct {
	Code2Session string
	AppId        string
	Secret       string
	JsCode       string
}

//code 换取 openid
type WxCode2SessionResponse struct {
	OpenId     string `json:"openid"`      //用户唯一标识
	SessionKey string `json:"session_key"` //会话密钥
	UnionId    string `json:"unionid"`     //用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。
	ErrCode    int    `json:"errcode"`     //错误码
	ErrMsg     string `json:"errmsg"`      //错误信息
}

//encryptedData 解密之后的字符串
type WxEncryptedData struct {
	OpenId    string `json:"openId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
	UnionId   string `json:"unionId"`
	WaterMark struct {
		Appid     string `json:"appid"`
		Simestamp string `json:"simestamp"`
	} `json:"watermark"`
}

func (wx Wx) GetCode2SessionUrl(info Wx) string {
	urlString := strings.Replace(info.Code2Session, "{APPID}", info.AppId, 1)
	urlString = strings.Replace(urlString, "{SECRET}", info.Secret, 1)
	urlString = strings.Replace(urlString, "{JSCODE}", info.JsCode, 1)
	fmt.Printf("qingfeng----%s====", urlString)
	return urlString
}
