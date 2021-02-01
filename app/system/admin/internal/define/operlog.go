package define

import "gea/app/model"

// ======== api ==========
//查询列表请求参数
type OperlogApiSelectPageReq struct {
	Title         string `p:"title"`         //系统模块
	OperName      string `p:"operName"`      //操作人员
	BusinessType  string `p:"businessType"`  //操作类型
	Status        string `p:"status"`        //操作类型
	BeginTime     string `p:"beginTime"`     //数据范围
	EndTime       string `p:"endTime"`       //开始时间
	PageNum       int    `p:"pageNum"`       //当前页码
	PageSize      int    `p:"pageSize"`      //每页数
	OrderByColumn string `p:"orderByColumn"` //排序字段
	IsAsc         string `p:"isAsc"`         //排序方式
}

//通用的删除请求
type OperlogApiDeleteReq struct {
	Ids string `p:"ids"  v:"required#请选择要删除的数据记录"`
}

// ========= service ========
// 查询列表返回值
type OperlogServiceList struct {
	List  []model.SysOperLog `json:"list"`
	Page  int                `json:"page"`
	Size  int                `json:"size"`
	Total int                `json:"total"`
}
