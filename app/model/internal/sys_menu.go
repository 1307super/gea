// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// SysMenu is the golang structure for table sys_menu.
type SysMenu struct {
    MenuId     int64       `orm:"menu_id,primary" json:"menu_id"`     // 菜单ID                         
    MenuName   string      `orm:"menu_name"       json:"menu_name"`   // 菜单名称                       
    ParentId   int64       `orm:"parent_id"       json:"parent_id"`   // 父菜单ID                       
    OrderNum   int         `orm:"order_num"       json:"order_num"`   // 显示顺序                       
    Path       string      `orm:"path"            json:"path"`        // 请求地址                       
    Component  string      `orm:"component"       json:"component"`   // 组件路径                       
    IsFrame    uint        `orm:"is_frame"        json:"is_frame"`    // 打开方式（1页签 2新窗口）      
    MenuType   string      `orm:"menu_type"       json:"menu_type"`   // 菜单类型（M目录 C菜单 F按钮）  
    Visible    uint        `orm:"visible"         json:"visible"`     // 菜单状态（0显示 1隐藏）        
    Status     uint        `orm:"status"          json:"status"`      // 菜单状态（0正常 1停用）        
    Perms      string      `orm:"perms"           json:"perms"`       // 权限标识                       
    Icon       string      `orm:"icon"            json:"icon"`        // 菜单图标                       
    CreateBy   string      `orm:"create_by"       json:"create_by"`   // 创建者                         
    CreateTime *gtime.Time `orm:"create_time"     json:"create_time"` // 创建时间                       
    UpdateBy   string      `orm:"update_by"       json:"update_by"`   // 更新者                         
    UpdateTime *gtime.Time `orm:"update_time"     json:"update_time"` // 更新时间                       
    Remark     string      `orm:"remark"          json:"remark"`      // 备注                           
    Url        string      `orm:"url"             json:"url"`         // 接口地址                       
    Method     string      `orm:"method"          json:"method"`      // 请求方法                       
}