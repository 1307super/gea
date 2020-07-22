package router

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"strings"
)

const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	DELETE  = "DELETE"
	CONNECT = "CONNECT"
	TRACE   = "TRACE"
	HOOK   = "HOOK"
)

var GroupList = make([]*routerGroup, 0)

var PermissionMap = make(map[string]string, 0)

//路由信息
type router struct {
	Method       string            //方法名称
	RelativePath string            //url路径
	Permiss      string            //权限字符串
	HandlerFunc  ghttp.HandlerFunc //执行函数
}

//路由组信息
type routerGroup struct {
	ServerName   string              //服务名称
	RelativePath string              //url路径
	Handlers     []ghttp.HandlerFunc //中间件
	Router       []*router           //路由信息
}

//根据url获取权限字符串
func FindPermission(url string) string {
	return PermissionMap[url]
}

//创建一个路由组
func New(serverName, relativePath string, middleware ...ghttp.HandlerFunc) *routerGroup {
	var rg routerGroup
	rg.ServerName = serverName
	rg.Router = make([]*router, 0)
	rg.RelativePath = relativePath
	rg.Handlers = middleware
	GroupList = append(GroupList, &rg)
	return &rg
}

//添加路由信息
func (group *routerGroup) Handle(method, relativePath, permiss string, handler ghttp.HandlerFunc) *routerGroup {
	var r router
	r.Method = method
	r.Permiss = permiss
	r.RelativePath = relativePath
	r.HandlerFunc = handler
	group.Router = append(group.Router, &r)
	if len(permiss) > 0 {
		if strings.EqualFold(relativePath,"/") {
			PermissionMap[method + ":" + group.RelativePath] = permiss
		}else{
			// 解析路由中的匹配规则 : * {}
			routerRule := gfile.Basename(relativePath)
			if strings.Contains(routerRule,":") || strings.Contains(routerRule,"*") || strings.Contains(routerRule,"{") {
				relativePath = strings.ReplaceAll(relativePath,"/"+routerRule,"")
			}
			PermissionMap[method + ":" +group.RelativePath+relativePath] = permiss
		}
	}
	return group
}

//添加路由信息-ANY
func (group *routerGroup) ANY(relativePath, permiss string, handler ghttp.HandlerFunc) *routerGroup {
	group.Handle("ANY", relativePath, permiss, handler)
	return group
}

//添加路由信息-GET
func (group *routerGroup) GET(relativePath, permiss string, handler ghttp.HandlerFunc) *routerGroup {
	group.Handle(GET, relativePath, permiss, handler)
	return group
}

//添加路由信息-POST
func (group *routerGroup) POST(relativePath, permiss string, handler ghttp.HandlerFunc) *routerGroup {
	group.Handle(POST, relativePath, permiss, handler)
	return group
}

//添加路由信息-OPTIONS
func (group *routerGroup) OPTIONS(relativePath, permiss string, handler ghttp.HandlerFunc) *routerGroup {
	group.Handle(OPTIONS, relativePath, permiss, handler)
	return group
}

//添加路由信息-PUT
func (group *routerGroup) PUT(relativePath, permiss string, handler ghttp.HandlerFunc) *routerGroup {
	group.Handle(PUT, relativePath, permiss, handler)
	return group
}

//添加路由信息-PATCH
func (group *routerGroup) PATCH(relativePath, permiss string, handler ghttp.HandlerFunc) *routerGroup {
	group.Handle(PATCH, relativePath, permiss, handler)
	return group
}

//添加路由信息-HEAD
func (group *routerGroup) HEAD(relativePath, permiss string, handler ghttp.HandlerFunc) *routerGroup {
	group.Handle(HEAD, relativePath, permiss, handler)
	return group
}

//添加路由信息-DELETE
func (group *routerGroup) DELETE(relativePath, permiss string, handler ghttp.HandlerFunc) *routerGroup {
	group.Handle(DELETE, relativePath, permiss, handler)
	return group
}

//添加路由信息-CONNECT
func (group *routerGroup) CONNECT(relativePath, permiss string, handler ghttp.HandlerFunc) *routerGroup {
	group.Handle(CONNECT, relativePath, permiss, handler)
	return group
}

//添加路由信息-TRACE
func (group *routerGroup) TRACE(relativePath, permiss string, handler ghttp.HandlerFunc) *routerGroup {
	group.Handle(TRACE, relativePath, permiss, handler)
	return group
}

func (group *routerGroup) HOOK(relativePath, permiss string, handler ghttp.HandlerFunc) *routerGroup {
	group.Handle(HOOK, relativePath,permiss,handler)
	return group
}
