package auth

import (
	userService "gea/app/service/system/user"
	"gea/app/utils/response"
	"gea/library/casbin"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/net/ghttp"
)

const ALL_PERMISSION = "*:*:*"
const SUPER_ADMIN = "admin"

var ALLOW_PERMISSION garray.StrArray

// 鉴权中间件，只有登录成功之后才能通过
func Auth(r *ghttp.Request) {
	//根据url判断是否有权限
	url := r.Request.URL.Path
	//获取用户信息
	user, err := userService.GetUserInfo(r.GetInt64("jwtUid"))
	if err != nil || user == nil {
		response.ErrorResp(r).SetStatus(401).SetMsg("用户不存在").WriteJsonExit()
	}

	// 如果是不是超管 校验权限
	if !hasPermissions(r,user.Permissions) {
		// casbin校验
		// 判断判断
		isOk, err := casbin.Enforce(user.User.LoginName, url, r.Method)
		if err != nil || !isOk {
			response.ErrorResp(r).SetStatus(401).SetMsg("您没有权限").WriteJsonExit()
		}
	}

	r.Middleware.Next()
}

func hasPermissions(r *ghttp.Request,userPermissions *garray.StrArray) bool {
	return userPermissions.Contains(ALL_PERMISSION) || ALLOW_PERMISSION.Contains(r.Request.URL.Path)
}
