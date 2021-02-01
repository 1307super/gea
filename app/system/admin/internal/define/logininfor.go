package define

import "gea/app/model"

// ========== api =============
//查询列表请求参数
type LogininforApiSelectPageReq struct {
	LoginName     string `p:"loginName"`     //登陆名
	Status        string `p:"status"`        //状态
	Ipaddr        string `p:"ipaddr"`        //登录地址
	BeginTime     string `p:"beginTime"`     //数据范围
	EndTime       string `p:"endTime"`       //开始时间
	PageNum       int    `p:"pageNum"`       //当前页码
	PageSize      int    `p:"pageSize"`      //每页数
	OrderByColumn string `p:"orderByColumn"` //排序字段
	IsAsc         string `p:"isAsc"`         //排序方式
}

// API执行删除内容
type LogininforApiDeleteReq struct {
	Ids string `p:"id"  v:"required#请选择要删除的数据记录"`
}

// ======= service =========
// 查询列表返回值
type LogininforServiceList struct {
	List  []model.SysLogininfor `json:"list"`
	Page  int                   `json:"page"`
	Size  int                   `json:"size"`
	Total int                   `json:"total"`
}
