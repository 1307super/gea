package shared

import (
	"gea/app/model"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
)

var GfAdminToken gtoken.GfToken

var GfToken gtoken.GfToken

func GetAdminUser(r *ghttp.Request) (userInfo *model.SysUserInfo) {
	tokenInfo := GfAdminToken.GetTokenData(r)
	if !tokenInfo.Success() {
		return nil
	}
	if err := gjson.DecodeTo(tokenInfo.GetString("data"), &userInfo); err == nil {
		return
	}
	return nil
}