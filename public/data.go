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

type (
	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * 获取AccessToken响应数据
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	AccessTokenResponse struct {
		ErrorCode   int    `form:"errcode" json:"errcode"`
		ErrorMsg    string `form:"errmsg" json:"errmsg"`
		AccessToken string `form:"access_token" json:"access_token"`
		ExpiresIn   int    `form:"expires_in" json:"expires_in"`
	}

	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * 获取微信服务器IP地址响应数据
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	ServerIpResponse struct {
		ErrorCode int      `form:"errcode" json:"errcode"`
		ErrorMsg  string   `form:"errmsg" json:"errmsg"`
		IpList    []string `form:"ip_list" json:"ip_list"`
	}
)
