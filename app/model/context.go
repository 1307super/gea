package model

import (
	"github.com/gogf/gf/frame/g"
)

const (
	// 上下文变量存储键名，前后端系统共享
	ContextKey = "ContextKey"
)

// 请求上下文结构
type Context struct {
	Uid   int64       // 用户id
	User  *SysUserInfo // 上下文用户信息
	TokenUser *SysUser
	Data  g.Map        // 自定KV变量，业务模块根据需要设置，不固定
}