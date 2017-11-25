package public

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
	AccessTokenUri PublicUriType = iota
	ServerIpUri
)

type (
	PublicUriType int

	IPublic interface {
		SetUri(uriType PublicUriType, uri string)

		GetAccessToken() (*AccessTokenResponse, error)
		GetServerIp(accessKey string) (*ServerIpResponse, error)
	}

	Public struct {
		ClientId     string //app id
		ClientSecret string //app secret

		accessTokenUri string //请求AccessToken地址
		serverIpUri    string //请求ServerIp地址
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 初始化WeChat公众号
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewPublic(clientId, clientSecret string) IMini {
	public := new(Public)
	public.ClientId = clientId
	public.ClientSecret = clientSecret

	public.accessTokenUri = "https://api.weixin.qq.com/cgi-bin/token"
	public.serverIpUri = "https://api.weixin.qq.com/cgi-bin/getcallbackip"

	return public
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置Uri
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Public) SetUri(uriType PublicUriType, uri string) {
	switch uriType {
	case AccessTokenUri:
		s.accessTokenUri = uri
	case ServerIpUri:
		s.serverIpUri = uri
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
* 获取session_key
* {
*     "errcode": 40013,
*     "errmsg": "invalid appid"
*     "access_token": "ACCESS_TOKEN",
*     "expires_in": 7200
* }
* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Public) GetAccessToken() (*AccessTokenResponse, error) {
	var response *AccessTokenResponse

	params := map[string]interface{}{
		"appid":      s.ClientId,
		"secret":     s.ClientSecret,
		"grant_type": "client_credential",
	}

	queryString := glib.ToQueryString(params)

	//获取api响应数据
	data, err := glib.HttpGet(s.accessTokenUri, queryString)
	if err == nil {
		//解析json数据
		glib.FromJson(data, &response)
	}

	return response, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
* 获取微信服务器IP地址
* {
*     "errcode": 40013,
*     "errmsg": "invalid appid"
*     "ip_list": [
*         "127.0.0.1",
*         "127.0.0.2",
*         "101.226.103.0/25"
*     ]
* }
* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Public) GetServerIp(accessToken string) (*ServerIpResponse, error) {
	var response *ServerIpResponse

	params := map[string]interface{}{
		"access_token": accessToken,
	}

	queryString := glib.ToQueryString(params)

	//获取api响应数据
	data, err := glib.HttpGet(s.serverIpUri, queryString)
	if err == nil {
		//解析json数据
		glib.FromJson(data, &response)
	}

	return response, err
}
