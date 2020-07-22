// ==========================================================================
// GEAGO自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2020-02-18 15:44:13
// 生成路径: app/model/module/job/job.go
// 生成人：yunjie
// ==========================================================================
package job

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"gea/app/utils/page"
)

//新增页面请求参数
type AddReq struct {
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
type EditReq struct {
	JobName        string `p:"jobName" `
	JobParams      string `p:"jobParams"` // 任务参数
	JobGroup       string `p:"jobGroup" `
	JobId          int64  `p:"jobId" v:"required#主键ID不能为空"`
	InvokeTarget   string `p:"invokeTarget" `
	CronExpression string `p:"cronExpression" `
	MisfirePolicy  string `p:"misfirePolicy" `
	Concurrent     string `p:"concurrent" `
	Status         string `p:"status" v:"required#状态（0正常 1暂停）不能为空"`
	Remark         string `p:"remark" `
}

//分页请求参数
type SelectPageReq struct {
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

//根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) ([]Entity, *page.Paging, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_job t")

	if param != nil {

		if param.JobName != "" {
			model.Where("t.job_name like ?", "%"+param.JobName+"%")
		}

		if param.JobGroup != "" {
			model.Where("t.job_group = ?", param.JobGroup)
		}

		if param.InvokeTarget != "" {
			model.Where("t.invoke_target = ?", param.InvokeTarget)
		}

		if param.CronExpression != "" {
			model.Where("t.cron_expression = ?", param.CronExpression)
		}

		if param.MisfirePolicy != "" {
			model.Where("t.misfire_policy = ?", param.MisfirePolicy)
		}

		if param.Concurrent != "" {
			model.Where("t.concurrent = ?", param.Concurrent)
		}

		if param.Status != "" {
			model.Where("t.status = ?", param.Status)
		}
		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	total, err := model.Count()

	if err != nil {
		return nil, nil, gerror.New("读取行数失败")
	}

	page := page.CreatePaging(param.PageNum, param.PageSize, total)

	model.Limit(page.StartNum, page.Pagesize)

	if param.OrderByColumn != "" {
		model.Order(param.OrderByColumn + " " + param.IsAsc)
	}

	var result []Entity
	model.Structs(&result)
	return result, page, nil
}

// 导出excel
func SelectListExport(param *SelectPageReq) (gdb.Result, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_job t")

	if param != nil {

		if param.JobName != "" {
			model.Where("t.job_name like ?", "%"+param.JobName+"%")
		}

		if param.JobGroup != "" {
			model.Where("t.job_group = ?", param.JobGroup)
		}

		if param.InvokeTarget != "" {
			model.Where("t.invoke_target = ?", param.InvokeTarget)
		}

		if param.CronExpression != "" {
			model.Where("t.cron_expression = ?", param.CronExpression)
		}

		if param.MisfirePolicy != "" {
			model.Where("t.misfire_policy = ?", param.MisfirePolicy)
		}

		if param.Concurrent != "" {
			model.Where("t.concurrent = ?", param.Concurrent)
		}

		if param.Status != "" {
			model.Where("t.status = ?", param.Status)
		}
		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	//"任务ID","任务名称","任务组名","调用目标字符串","cron执行表达式","计划执行错误策略（1立即执行 2执行一次 3放弃执行）","是否并发执行（0允许 1禁止）","状态（0正常 1暂停）","创建者","创建时间","更新者","更新时间","备注信息",
	model.Fields(" t.job_id ,t.job_name ,t.job_group ,t.invoke_target ,t.cron_expression ,t.misfire_policy ,t.concurrent ,t.status ,t.create_by ,t.create_time ,t.update_by ,t.update_time ,t.remark")

	result, _ := model.All()
	return result, nil
}

//获取所有数据
func SelectListAll(param *SelectPageReq) ([]Entity, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_job t")

	if param != nil {

		if param.JobName != "" {
			model.Where("t.job_name like ?", "%"+param.JobName+"%")
		}

		if param.JobGroup != "" {
			model.Where("t.job_group = ?", param.JobGroup)
		}

		if param.InvokeTarget != "" {
			model.Where("t.invoke_target = ?", param.InvokeTarget)
		}

		if param.CronExpression != "" {
			model.Where("t.cron_expression = ?", param.CronExpression)
		}

		if param.MisfirePolicy != "" {
			model.Where("t.misfire_policy = ?", param.MisfirePolicy)
		}

		if param.Concurrent != "" {
			model.Where("t.concurrent = ?", param.Concurrent)
		}

		if param.Status != "" {
			model.Where("t.status = ?", param.Status)
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []Entity

	err = model.Structs(&result)
	return result, err
}
