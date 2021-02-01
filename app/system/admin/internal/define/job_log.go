package define

import "gea/app/model"

// ============ api =============
//分页请求参数
type JobLogApiSelectPageReq struct {
	JobLogId      int64  `p:"jobLogId"`      //任务日志ID
	JobName       string `p:"jobName"`       //任务名称
	JobGroup      string `p:"jobGroup"`      //任务组名
	InvokeTarget  string `p:"invokeTarget"`  //调用目标字符串
	JobMessage    string `p:"jobMessage"`    //日志信息
	Status        string `p:"status"`        //执行状态（0正常 1失败）
	ExceptionInfo string `p:"exceptionInfo"` //异常信息
	BeginTime     string `p:"beginTime"`     //开始时间
	EndTime       string `p:"endTime"`       //结束时间
	PageNum       int    `p:"pageNum"`       //当前页码
	PageSize      int    `p:"pageSize"`      //每页数
}
//通用的删除请求
type JobLogApiDeleteReq struct {
	Ids string `p:"ids"  v:"required#请选择要删除的数据记录"`
}

// ============ service =============
// 查询列表返回值
type JobLogServiceList struct {
	List  []model.SysJobLog `json:"list"`
	Page  int            `json:"page"`
	Size  int            `json:"size"`
	Total int            `json:"total"`
}