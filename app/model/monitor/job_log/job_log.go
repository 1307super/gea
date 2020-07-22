// ==========================================================================
// GEAGO自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2020-02-18 15:44:13
// 生成路径: app/model/module/log/log.go
// 生成人：yunjie
// ==========================================================================
package job_log

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"gea/app/utils/page"
)

//分页请求参数
type SelectPageReq struct {
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

//根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) ([]Entity, *page.Paging, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_job_log t")

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

		if param.JobMessage != "" {
			model.Where("t.job_message = ?", param.JobMessage)
		}

		if param.Status != "" {
			model.Where("t.status = ?", param.Status)
		}

		if param.ExceptionInfo != "" {
			model.Where("t.exception_info = ?", param.ExceptionInfo)
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

	model := db.Table("sys_job_log t")

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

		if param.JobMessage != "" {
			model.Where("t.job_message = ?", param.JobMessage)
		}

		if param.Status != "" {
			model.Where("t.status = ?", param.Status)
		}

		if param.ExceptionInfo != "" {
			model.Where("t.exception_info = ?", param.ExceptionInfo)
		}
		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	//"任务日志ID","任务名称","任务组名","调用目标字符串","日志信息","执行状态（0正常 1失败）","异常信息","创建时间",
	model.Fields(" t.job_log_id ,t.job_name ,t.job_group ,t.invoke_target ,t.job_message ,t.status ,t.exception_info ,t.create_time")

	result, _ := model.All()
	return result, nil
}

//获取所有数据
func SelectListAll(param *SelectPageReq) ([]Entity, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_job_log t")

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

		if param.JobMessage != "" {
			model.Where("t.job_message = ?", param.JobMessage)
		}

		if param.Status != "" {
			model.Where("t.status = ?", param.Status)
		}

		if param.ExceptionInfo != "" {
			model.Where("t.exception_info = ?", param.ExceptionInfo)
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []Entity
	model.Structs(&result)
	return result, nil
}
