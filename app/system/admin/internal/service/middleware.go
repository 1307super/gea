package service

import (
	"gea/app/model"
	"gea/app/shared"
	"gea/app/utils/response"
	"gea/app/utils/token"
	"gea/library/casbin"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

const ALL_PERMISSION = "*:*:*"

var ALLOW_PERMISSION garray.StrArray

var Middleware = &middlewareService{}

type middlewareService struct{}

// 自定义上下文对象
func (s *middlewareService) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{}
	shared.Context.Init(r, customCtx)
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// 允许接口跨域请求
func (s *middlewareService)CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *middlewareService) Jwt(r *ghttp.Request) {
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

	//判断是否有新的token生成
	if t.NewToken != "" {
		r.Response.Header().Set("nt", t.NewToken)
	}

	// 将当前请求的uid信息保存到请求的上下文c上
	customCtx := &model.Context{
		Token: tokenStr,
		Uid:   t.Claim.UserId,
		Data:  make(g.Map),
	}
	shared.Context.Init(r, customCtx)
	if userEntity, err := User.GetUser(r.Context()); err == nil {
		if userEntity == nil {
			response.NauthorizedResp(r).WriteJsonExit()
		} else {
			shared.Context.SetUser(r.Context(),userEntity)
		}
	} else {
		response.NauthorizedResp(r).WriteJsonExit()
	}
	r.Middleware.Next()
}

func (s *middlewareService) Auth(r *ghttp.Request) {
	//根据url判断是否有权限
	url := r.Request.URL.Path
	//获取用户信息
	user, err := User.GetUserInfo(r.Context())
	if err != nil || user == nil {
		response.ErrorResp(r).SetStatus(401).SetMsg("用户不存在").WriteJsonExit()
	}
	// 如果是不是超管 校验权限
	if !hasPermissions(user.Permissions) {
		// casbin校验
		// 判断判断
		isOk, err := casbin.Enforce(user.User.LoginName, url, r.Method)
		if err != nil || !isOk {
			response.ErrorResp(r).SetStatus(401).SetMsg("您没有权限").WriteJsonExit()
		}
	}
	r.Middleware.Next()
}

func hasPermissions(userPermissions *garray.StrArray) bool {
	return userPermissions.Contains(ALL_PERMISSION)
}


// ====== 操作日志 ==========
func (s *middlewareService) OperationLog(r *ghttp.Request) {
	r.Middleware.Next()
	// get请求不记录
	if r.Method == "GET" {
		return
	}
	//userInfo, err := userService.GetProfileApi(r.GetInt64("jwtUid"))
	//if err != nil {
	//	return
	//}
	//if userInfo == nil {
	//	return
	//}
	resp := new(response.CommonRes)
	//r.GetCtxVar(response.ResponseCtxName).Struct(resp)
	//fmt.Println(r.Response.BufferString())
	gconv.Struct(r.Response.BufferString(),resp)
	resp.Btype = response.BunissType(r.GetCtxVar(response.ResponseBunissType).Int())
	// 存操作日志
	var paramJson string
	param := r.GetMap()
	paramByte, err := gjson.Encode(param)
	if err == nil {
		paramJson = string(paramByte)
	}
	go OperLog.Create(r,resp.Module,paramJson,resp)
	//}
}
