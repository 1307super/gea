// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"gea/app/model/internal"
)

// SysMenu is the golang structure for table sys_menu.
type SysMenu internal.SysMenu

// Fill with you ideas below.

type SysMenuExtend struct {
	SysMenu
	ParentName string          `json:"parentName"`         // 父菜单名称
	Children   []*SysMenuExtend `json:"children,omitempty"` // 子菜单
}

// 路由
type RouterExtend struct {
	AlwaysShow bool           `json:"alwaysShow,omitempty"` // 总是显示
	Children   []RouterExtend `json:"children,omitempty"`    // 子菜单
	Component  string         `json:"component,omitempty"`             // 组件地址
	Hidden     bool           `json:"hidden"`                // 是否隐藏
	Meta       Meta           `json:"meta"`                  // meta
	Name       string         `json:"name"`                  // 名称
	Path       string         `json:"path"`                  // 地址
	Redirect   string         `json:"redirect,omitempty"`    // 跳转链接
	//SysMenu
	//ParentName string `json:"parentName"` // 父菜单名称
}
type Meta struct {
	Title string `json:"title"` // 标题
	Icon  string `json:"icon"`  // 图标
}

// 角色权限树
type RoleMenuTree struct {
	Children []RoleMenuTree `json:"children"`
	Id       int64          `json:"id"`
	Label    string         `json:"label"`
}