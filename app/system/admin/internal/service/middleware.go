package service

import (
	"context"
	"fmt"
	"gea/app/model"
	"gea/app/shared"
	"gea/app/utils/response"
	"gea/library/casbin"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

const ALL_PERMISSION = "*:*:*"


const DATA_SCOPE_ALL = "1"            // 全部数据权限
const DATA_SCOPE_CUSTOM = "2"         // 自定数据权限
const DATA_SCOPE_DEPT = "3"           // 部门数据权限
const DATA_SCOPE_DEPT_AND_CHILD = "4" // 部门及以下数据权限
const DATA_SCOPE_SELF = "5"           // 仅本人数据权限

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
func (s *middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// 登录鉴权
func (s *middlewareService) LoginAuth(r *ghttp.Request)  {
	// 从缓存中获取
	user := shared.GetAdminUser(r)
	if user == nil {
		response.NauthorizedResp(r).WriteJsonExit()
	}
	shared.Context.Get(r.Context()).Uid = user.UserExtend.UserId
	shared.Context.SetUser(r.Context(), user)
	r.Middleware.Next()
}

func (s *middlewareService) Auth(r *ghttp.Request) bool {
	// 从缓存中获取
	user := shared.GetAdminUser(r)
	if user == nil {
		response.NauthorizedResp(r).WriteJsonExit()
	}
	prefix := g.Cfg().GetString("admin.prefix")
	//根据url判断是否有权限
	url := r.Request.URL.Path
	url = gstr.Replace(url, prefix, "")
	// 如果是不是超管 校验权限
	if !hasPermissions(user.Permissions) {
		// casbin校验
		// 判断判断
		isOk, err := casbin.Enforce(user.UserExtend.LoginName, url, r.Method)
		if err != nil || !isOk {
			response.ErrorResp(r).SetStatus(401).SetMsg("您没有权限").WriteJsonExit()
		}
	}
	shared.Context.Get(r.Context()).Uid = user.UserExtend.UserId
	shared.Context.SetUser(r.Context(), user)
	return true
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
	resp := new(response.CommonRes)
	gconv.Struct(r.Response.BufferString(), resp)
	resp.Btype = response.BunissType(r.GetCtxVar(response.ResponseBunissType).Int())
	// 存操作日志
	var paramJson string
	param := r.GetMap()
	paramByte, err := gjson.Encode(param)
	if err == nil {
		paramJson = string(paramByte)
	}
	go OperLog.Create(r, resp.Module, paramJson, resp)
}

// 数据权限过滤
func DataScopeFilter(ctx context.Context,deptAlias,userAlias string) string {
	if deptAlias == ""{
		deptAlias = "d"
	}
	// 获取用户角色
	customCtx := shared.Context.Get(ctx)
	var sqlArr garray.StrArray
	roles := customCtx.User.UserExtend.Roles
	for _, role := range roles {
		dataScope := role.DataScope
		if DATA_SCOPE_ALL == dataScope {
			sqlArr.Append("")
		} else if DATA_SCOPE_CUSTOM == dataScope {
			sqlArr.Append(fmt.Sprintf(" or `%s`.dept_id IN ( SELECT dept_id FROM sys_role_dept WHERE role_id = %d ) ",deptAlias,role.RoleId))
		} else if DATA_SCOPE_DEPT == dataScope {
			sqlArr.Append(fmt.Sprintf(" or `%s`.dept_id = %d ",deptAlias,customCtx.User.UserExtend.DeptId))
		} else if DATA_SCOPE_DEPT_AND_CHILD == dataScope {
			sqlArr.Append(fmt.Sprintf(
				" or `%s`.dept_id IN ( SELECT dept_id FROM sys_dept WHERE dept_id = %d or find_in_set( %d , ancestors ) ) ",
				deptAlias,customCtx.User.UserExtend.DeptId,customCtx.User.UserExtend.DeptId,
			))
		} else if DATA_SCOPE_SELF == dataScope {
			if userAlias == ""{
				sqlArr.Append(" or 1 = 0 ")
			}else{
				sqlArr.Append(fmt.Sprintf(" or `%s`.user_id = %d",userAlias,customCtx.User.UserExtend.UserId))
			}

		}
	}
	if !sqlArr.IsEmpty() {
		return gstr.SubStr(sqlArr.Join(""),4)
	}
	return ""
}
