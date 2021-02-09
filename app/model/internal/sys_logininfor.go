// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// SysLogininfor is the golang structure for table sys_logininfor.
type SysLogininfor struct {
    InfoId        int64       `orm:"info_id,primary" json:"info_id"`        // 访问ID                   
    LoginName     string      `orm:"login_name"      json:"login_name"`     // 登录账号                 
    Ipaddr        string      `orm:"ipaddr"          json:"ipaddr"`         // 登录IP地址               
    LoginLocation string      `orm:"login_location"  json:"login_location"` // 登录地点                 
    Browser       string      `orm:"browser"         json:"browser"`        // 浏览器类型               
    Os            string      `orm:"os"              json:"os"`             // 操作系统                 
    Status        string      `orm:"status"          json:"status"`         // 登录状态（0成功 1失败）  
    Msg           string      `orm:"msg"             json:"msg"`            // 提示消息                 
    LoginTime     *gtime.Time `orm:"login_time"      json:"login_time"`     // 访问时间                 
}