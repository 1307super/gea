package jwt

import (
	"gea/app/utils/response"
	"gea/app/utils/token"
	"github.com/gogf/gf/net/ghttp"
	"strings"
)

// 鉴权中间件，只有登录成功之后才能通过
func JWT(r *ghttp.Request) {
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	// 这里的具体实现方式要依据你的实际业务情况决定
	authHeader := r.Request.Header.Get("Authorization")
	// 按空格分割
	tokenStr := ""
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) > 1 {
			tokenStr = parts[1]
		}
	} else {
		tokenStr = r.Request.Header.Get("token")
	}

	if tokenStr == "" {
		tokenStr = r.GetQueryString("token")
	}

	if tokenStr == "" {
		response.NauthorizedResp(r).WriteJsonExit()
		return
	}

	// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
	t, err := token.VerifyAuthToken(tokenStr)
	if err != nil {
		response.NauthorizedResp(r).WriteJsonExit()
		return
	}

	// 将当前请求的uid信息保存到请求的上下文c上
	r.SetParam("jwtUid", t.Claim.UserId)
	r.SetParam("jwtLoginName", t.Claim.LoginName)
	//判断是否有新的token生成
	if t.NewToken != "" {
		r.Response.Header().Set("nt", t.NewToken)
	}
	r.Middleware.Next()
}
