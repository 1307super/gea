package define

import (
	"gea/app/model"
	"github.com/gogf/gf/os/gtime"
)

//分页请求参数
type OnlineApiSelectPageReq struct {
	Token          string      `p:"token"`          //用户会话id
	LoginName      string      `p:"loginName"`      //登录账号
	DeptName       string      `p:"deptName"`       //部门名称
	Ipaddr         string      `p:"ipaddr"`         //登录IP地址
	LoginLocation  string      `p:"loginLocation"`  //登录地点
	Browser        string      `p:"browser"`        //浏览器类型
	Os             string      `p:"os"`             //操作系统
	Status         string      `p:"status"`         //在线状态on_line在线off_line离线
	StartTimestamp *gtime.Time `p:"startTimestamp"` //session创建时间
	LastAccessTime *gtime.Time `p:"lastAccessTime"` //session最后访问时间
	ExpireTime     int         `p:"expireTime"`     //超时时间，单位为分钟
	BeginTime      string      `p:"beginTime"`      //开始时间
	EndTime        string      `p:"endTime"`        //结束时间
	PageNum        int         `p:"pageNum"`        //当前页码
	PageSize       int         `p:"pageSize"`       //每页数
	OrderByColumn  string      `p:"orderByColumn"`  //排序字段
	IsAsc          string      `p:"isAsc"`          //排序方式
}

// ======= service =========
// 查询列表返回值
type OnlineServiceList struct {
	List  []model.SysUserOnline `json:"list"`
	Page  int                   `json:"page"`
	Size  int                   `json:"size"`
	Total int                   `json:"total"`
}
