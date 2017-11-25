package mini

import (
	"fmt"
)

import (
	"github.com/sanxia/glib"
)

/* ================================================================================
 * Oauth
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊
 * ================================================================================ */

const (
	SessionKeyUri MiniUriType = iota
	MiniCodeUri
	QrCodeUri
)

type (
	MiniUriType int

	IMini interface {
		SetUri(uriType MiniUriType, uri string)

		GetUserInfo(code, encryptedData, iv string) (*UserInfo, error)
		GetMiniCodeImage(accessToken, path string, args ...int) (*ImageResponse, error)
		GetQrCodeImage(accessToken, path string, args ...int) (*ImageResponse, error)

		GetSessionKey(code string) (*SessionKeyResponse, error)
		Decrypt(sessionKey, sourceData, iv string) (string, error)
		IsSignature(sessionKey, sourceData, signature string) bool
	}

	Mini struct {
		ClientId     string //app id
		ClientSecret string //app secret

		sessionKeyUri string //请求session_key地址
		miniCodeUri   string //请求小程序码图片
		qrCodeUri     string //请求qr二维码图片
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 初始化WeChat小程序
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewMini(clientId, clientSecret string) IMini {
	mini := new(Mini)
	mini.ClientId = clientId
	mini.ClientSecret = clientSecret

	mini.sessionKeyUri = "https://api.weixin.qq.com/sns/jscode2session"
	mini.miniCodeUri = "https://api.weixin.qq.com/wxa/getwxacode"
	mini.qrCodeUri = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode"

	return mini
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置Uri
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Mini) SetUri(uriType MiniUriType, uri string) {
	switch uriType {
	case SessionKeyUri:
		s.sessionKeyUri = uri
	case MiniCodeUri:
		s.miniCodeUri = uri
	case QrCodeUri:
		s.qrCodeUri = uri
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取微信小程序用户信息
 * {
 * "userInfo:        object
 * "rawData:         string
 * "signature:       string
 * "encryptedData":  string
 * "iv":             string
 * }
 *
 * userInfo:
 * {
 *    nickName:      string
 *    avatarUrl:     string
 *    gender:        string （值为1时是男性，值为2时是女性，值为0时是未知）
 *    country:       string （用户所在国家）
 *    province:      string （用户所在省份）
 *    city:          string （用户所在城市）
 *    language:      string （用户的语言，简体中文为zh_CN）
 * }
 *
 * encryptedData解密后的json结构:
 * {
 *     "openId": "OPENID",
 *     "nickName": "NICKNAME",
 *     "gender": GENDER,
 *     "city": "CITY",
 *     "province": "PROVINCE",
 *     "country": "COUNTRY",
 *     "avatarUrl": "AVATARURL",
 *     "unionId": "UNIONID",
 *     "watermark":{
 *         "appid":"APPID",
 *         "timestamp":TIMESTAMP
 *     }
 * }
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Mini) GetUserInfo(code, encryptedData, iv string) (*UserInfo, error) {
	var userInfo *UserInfo

	//获取会话key
	sessionKeyResponse, err := s.GetSessionKey(code)
	if err != nil {
		return nil, err
	}

	//解密数据
	if data, err := s.Decrypt(sessionKeyResponse.SessionKey, encryptedData, iv); err != nil {
		return nil, err
	} else {
		//解析json数据
		glib.FromJson(data, &userInfo)
	}

	if userInfo != nil {
		//附加上session_key
		userInfo.SessionKey = sessionKeyResponse.SessionKey

		if userInfo.Gender == "0" {
			userInfo.Gender = "secret"
		} else if userInfo.Gender == "1" {
			userInfo.Gender = "male"
		} else if userInfo.Gender == "2" {
			userInfo.Gender = "female"
		}
	}

	return userInfo, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取小程序的小程序码图片
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Mini) GetMiniCodeImage(accessToken, path string, args ...int) (*ImageResponse, error) {
	var imageResponse *ImageResponse

	width := 300
	if len(args) > 0 {
		width = args[0]
	}

	params := map[string]interface{}{
		"path":  path,
		"width": width,
	}

	queryString := glib.ToQueryString(params)

	url := fmt.Sprintf("%s?access_token=%s", s.miniCodeUri, accessToken)

	//获取api响应数据
	data, err := glib.HttpPost(url, queryString)
	if err == nil {
		//解析json数据
		if err := glib.FromJson(data, &imageResponse); err != nil {
			imageResponse = &ImageResponse{
				Data: data,
			}
		}
	}

	return imageResponse, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取小程序的Qr二维码图片
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Mini) GetQrCodeImage(accessToken, path string, args ...int) (*ImageResponse, error) {
	var imageResponse *ImageResponse

	width := 300
	if len(args) > 0 {
		width = args[0]
	}

	params := map[string]interface{}{
		"path":  path,
		"width": width,
	}

	queryString := glib.ToQueryString(params)

	url := fmt.Sprintf("%s?access_token=%s", s.qrCodeUri, accessToken)

	//获取api响应数据
	data, err := glib.HttpPost(url, queryString)
	if err == nil {
		//解析json数据
		if err := glib.FromJson(data, &imageResponse); err != nil {
			imageResponse = &ImageResponse{
				Data: data,
			}
		}
	}

	return imageResponse, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取session_key
 * {
 *     "openid": "OPENID",
 *     "session_key": "SESSIONKEY",
 *     "unionid": "UNIONID"
 * }
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Mini) GetSessionKey(code string) (*SessionKeyResponse, error) {
	var response *SessionKeyResponse

	params := map[string]interface{}{
		"appid":      s.ClientId,
		"secret":     s.ClientSecret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}

	queryString := glib.ToQueryString(params)

	//获取api响应数据
	data, err := glib.HttpGet(s.sessionKeyUri, queryString)
	if err == nil {
		//解析json数据
		glib.FromJson(data, &response)
	}

	return response, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 解密微信小程序数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Mini) Decrypt(sessionKey, sourceData, iv string) (string, error) {
	var data string
	var err error

	sessionKey, err = glib.FromBase64(sessionKey)
	if err != nil {
		return "", err
	}

	sourceData, err = glib.FromBase64(sourceData)
	if err != nil {
		return "", err
	}

	iv, err = glib.FromBase64(iv)
	if err != nil {
		return "", err
	}

	if dataBytes, dataErr := glib.AesDecrypt(
		[]byte(sourceData), []byte(sessionKey), []byte(iv)); dataErr == nil {
		data = string(dataBytes)
	} else {
		err = dataErr
	}
	return data, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断签名数据是否有效
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Mini) IsSignature(sessionKey, sourceData, signature string) bool {
	isOk := false
	data := sourceData + sessionKey
	if signatureResult := glib.Sha1(data); signatureResult == signature {
		isOk = true
	}

	return isOk
}
