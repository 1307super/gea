package define

import "gea/app/model"

// ========= api ============
//分页请求参数
type JobApiSelectPageReq struct {
	JobId          int64  `p:"jobId"`          //任务ID
	JobName        string `p:"jobName"`        //任务名称
	JobGroup       string `p:"jobGroup"`       //任务组名
	InvokeTarget   string `p:"invokeTarget"`   //调用目标字符串
	CronExpression string `p:"cronExpression"` //cron执行表达式
	MisfirePolicy  string `p:"misfirePolicy"`  //计划执行错误策略（1立即执行 2执行一次 3放弃执行）
	Concurrent     string `p:"concurrent"`     //是否并发执行（0允许 1禁止）
	Status         string `p:"status"`         //状态（0正常 1暂停）
	BeginTime      string `p:"beginTime"`      //开始时间
	EndTime        string `p:"endTime"`        //结束时间
	PageNum        int    `p:"pageNum"`        //当前页码
	PageSize       int    `p:"pageSize"`       //每页数
	OrderByColumn  string `p:"orderByColumn"`  //排序字段
	IsAsc          string `p:"isAsc"`          //排序方式
}

//新增页面请求参数
type JobApiCreateReq struct {
	JobName        string `p:"jobName" `
	JobParams      string `p:"jobParams"` // 任务参数
	JobGroup       string `p:"jobGroup" `
	InvokeTarget   string `p:"invokeTarget" `
	CronExpression string `p:"cronExpression" `
	MisfirePolicy  string `p:"misfirePolicy" `
	Concurrent     string `p:"concurrent" `
	Status         string `p:"status" v:"required#状态（0正常 1暂停）不能为空"`
	Remark         string `p:"remark" `
}

//修改页面请求参数
type JobApiUpdateReq struct {
	JobId int64 `p:"jobId" v:"required#主键ID不能为空"`
	JobApiCreateReq
}


//通用的删除请求
type JobApiDeleteReq struct {
	Ids string `p:"ids"  v:"required#请选择要删除的数据记录"`
}

// ========= service ============
// 查询列表返回值
type JobServiceList struct {
	List  []model.SysJob `json:"list"`
	Page  int            `json:"page"`
	Size  int            `json:"size"`
	Total int            `json:"total"`
}
