// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// SysUserOnline is the golang structure for table sys_user_online.
type SysUserOnline struct {
    Token          string      `orm:"token,primary"    json:"token"`            // 用户会话token                    
    LoginName      string      `orm:"login_name"       json:"login_name"`       // 登录账号                         
    DeptName       string      `orm:"dept_name"        json:"dept_name"`        // 部门名称                         
    Ipaddr         string      `orm:"ipaddr"           json:"ipaddr"`           // 登录IP地址                       
    LoginLocation  string      `orm:"login_location"   json:"login_location"`   // 登录地点                         
    Browser        string      `orm:"browser"          json:"browser"`          // 浏览器类型                       
    Os             string      `orm:"os"               json:"os"`               // 操作系统                         
    Status         string      `orm:"status"           json:"status"`           // 在线状态on_line在线off_line离线  
    StartTimestamp *gtime.Time `orm:"start_timestamp"  json:"start_timestamp"`  // 创建时间                         
    LastAccessTime *gtime.Time `orm:"last_access_time" json:"last_access_time"` // 最后访问时间                     
    ExpireTime     int         `orm:"expire_time"      json:"expire_time"`      // 超时时间，单位为分钟             
}