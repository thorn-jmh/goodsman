package feishu

import "goodsman/config"

var (
	TenantTokenManager *CommonAccessTokenManager
	AppTokenManager    *CommonAccessTokenManager

	CommonClient *FeishuClient

	AppID           string
	AppSecret       string
	TokenExpireTime = 110
	Content_Type    = "application/json; charset=utf-8"
)

var (
	getTenantAccessTokenUrl = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
	getAppAccessTokenUrl    = "https://open.feishu.cn/open-apis/auth/v3/app_access_token/internal"
)

func Init() {
	AppID = config.Base.AppID
	AppSecret = config.Base.AppSecret
	TenantTokenManager = DefaultAccessTokenManager("tenant", getTenantAccessTokenUrl)
	AppTokenManager = DefaultAccessTokenManager("app", getAppAccessTokenUrl)
	CommonClient = NewClient()
}
