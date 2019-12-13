package example

import (
	"errors"
	"zz/zz"
)

//获取微信openid
func WxGetOpenid(code string) (zz.WxCode2SessionResponse, error) {
	var err error
	//换取openid
	codeurl := new(zz.Wx).GetCode2SessionUrl(zz.Wx{
		Code2Session: "https://api.weixin.qq.com/sns/jscode2session?appid={APPID}&secret={SECRET}&js_code={JSCODE}&grant_type=authorization_code",
		AppId:        "example:config.Config.Wx.AppId",
		Secret:       "example:config.Config.Wx.Secret",
		JsCode:       code,
	})
	openIdInfo := zz.WxCode2SessionResponse{}
	err = zz.Get(codeurl, map[string]string{}, &openIdInfo)
	if err != nil {
		return openIdInfo, err
	}
	if openIdInfo.ErrCode != 0 || openIdInfo.OpenId == "" {
		return openIdInfo, errors.New("ErrCode not zero or OpenId is empty")
	}
	return openIdInfo, nil
}
