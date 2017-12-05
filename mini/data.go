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

	UserInfoResponse struct {
		Code          string `form:"code" json:"code"`
		RawData       string `form:"rawData" json:"rawData"`
		EncryptedData string `form:"encryptedData" json:"encryptedData"`
		Signature     string `form:"signature" json:"signature"`
		Iv            string `form:"iv" json:"iv"`
	}

	UserInfoData struct {
		UnionId   string     `form:"unionId" json:"unionId"`
		OpenId    string     `form:"openId" json:"openId"`
		Nickname  string     `form:"nickName" json:"nickName"`
		Gender    int        `form:"gender" json:"gender"` //性别 0：未知、1：男、2：女
		AvatarUrl string     `form:"avatarUrl" json:"avatarUrl"`
		City      string     `form:"city" json:"city"`
		Province  string     `form:"province" json:"province"`
		Country   string     `form:"country" json:"country"`
		Watermark *Watermark `form:"watermark" json:"watermark"`
	}

	ImageResponse struct {
		ErrorCode int    `form:"errcode" json:"errcode"`
		ErrorMsg  string `form:"errmsg" json:"errmsg"`
		Data      string `form:"data" json:"data"`
	}

	UserInfo struct {
		SessionKey string     `form:"session_key" json:"session_key"`
		UnionId    string     `form:"union_id" json:"union_id"`
		OpenId     string     `form:"open_id" json:"open_id"`
		Nickname   string     `form:"nickname" json:"nickname"`
		Gender     string     `form:"gender" json:"gender"`
		Avatar     string     `form:"avatar" json:"avatar"`
		City       string     `form:"city" json:"city"`
		Province   string     `form:"province" json:"province"`
		Country    string     `form:"country" json:"country"`
		Watermark  *Watermark `form:"watermark" json:"watermark"`
	}

	Watermark struct {
		AppId     string `form:"appid" json:"appid"`
		Timestamp uint64 `form:"timestamp" json:"timestamp"`
	}
)
