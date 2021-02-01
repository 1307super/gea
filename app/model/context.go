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
	Token string       // jwttoken
	Uid   string       // 用户id
	User  *SysUserExtend // 上下文用户信息
	Data  g.Map        // 自定KV变量，业务模块根据需要设置，不固定
}

// 请求上下文中的用户信息
//type ContextUser struct {
//	SysUserExtend
//	//Id        int64  // 用户ID
//	//LoginName string // 用户账号
//	//Nickname  string // 用户名称
//	//Avatar    string // 用户头像
//	//DeptId    int64  // 部门id
//	//DeptName  string // 部门名称
//}
