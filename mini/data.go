package mini

/* ================================================================================
 * Oauth
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊
 * ================================================================================ */

type (
	SessionKeyResponse struct {
		ErrorCode  int    `form:"errcode" json:"errcode"`
		ErrorMsg   string `form:"errmsg" json:"errmsg"`
		OpenId     string `form:"openid" json:"openid"`
		UnionId    string `form:"unionid" json:"unionid"`
		SessionKey string `form:"session_key" json:"session_key"`
	}

	ImageResponse struct {
		ErrorCode int    `form:"errcode" json:"errcode"`
		ErrorMsg  string `form:"errmsg" json:"errmsg"`
		Data      string `form:"data" json:"data"`
	}

	UserInfoResponse struct {
		Code          string `form:"code" json:"code"`
		RawData       string `form:"rawData" json:"rawData"`
		EncryptedData string `form:"encryptedData" json:"encryptedData"`
		Signature     string `form:"signature" json:"signature"`
		Iv            string `form:"iv" json:"iv"`
	}

	UserInfo struct {
		OpenId     string    `form:"openId" json:"openId"`
		UnionId    string    `form:"unionId" json:"unionId"`
		Nickname   string    `form:"nickName" json:"nickName"`
		Gender     string    `form:"gender" json:"gender"` //性别 0：未知、1：男、2：女
		Avatar     string    `form:"avatarUrl" json:"avatarUrl"`
		City       string    `form:"city" json:"city"`
		Province   string    `form:"province" json:"province"`
		Country    string    `form:"country" json:"country"`
		SessionKey string    `form:"session_key" json:"session_key"`
		Watermark  Watermark `form:"watermark" json:"watermark"`
	}

	Watermark struct {
		AppId     string `form:"appid" json:"appid"`
		Timestamp uint64 `form:"timestamp" json:"timestamp"`
	}
)
