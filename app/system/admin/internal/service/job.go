package service

import (
	"context"
	"gea/app/dao"
	"gea/app/model"
	"gea/app/shared"
	"gea/app/system/admin/internal/define"
	"gea/app/task"
	"gea/app/utils/convert"
	"gea/app/utils/excel"
	"gea/app/utils/page"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

var Job = &jobService{}

type jobService struct{}


func (s *jobService)Info(id int64) (*model.SysJob,error) {
	return dao.SysJob.FindOne(dao.SysJob.Columns.JobId,id)
}
//根据条件分页查询数据
func (s *jobService)GetList(param *define.JobApiSelectPageReq) *define.JobServiceList {
	m := dao.SysJob.As("t")

	if param != nil {

		if param.JobName != "" {
			m = m.Where("t.job_name like ?", "%"+param.JobName+"%")
		}

		if param.JobGroup != "" {
			m = m.Where("t.job_group = ?", param.JobGroup)
		}

		if param.InvokeTarget != "" {
			m = m.Where("t.invoke_target = ?", param.InvokeTarget)
		}

		if param.CronExpression != "" {
			m = m.Where("t.cron_expression = ?", param.CronExpression)
		}

		if param.MisfirePolicy != "" {
			m = m.Where("t.misfire_policy = ?", param.MisfirePolicy)
		}

		if param.Concurrent != "" {
			m = m.Where("t.concurrent = ?", param.Concurrent)
		}

		if param.Status != "" {
			m = m.Where("t.status = ?", param.Status)
		}
		if param.BeginTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	total, err := m.Count()

	if err != nil {
		return nil
	}

	page := page.CreatePaging(param.PageNum, param.PageSize, total)

	m = m.Limit(page.StartNum, page.Pagesize)

	if param.OrderByColumn != "" {
		m = m.Order(param.OrderByColumn + " " + param.IsAsc)
	}
	result := &define.JobServiceList{
		Page:  page.PageNum,
		Size:  page.Pagesize,
		Total: page.Total,
	}
	if err = m.Structs(&result.List); err != nil {
		return nil
	}
	return result
}

//添加数据
func (s *jobService)Create(ctx context.Context, req *define.JobApiCreateReq) (int64, error) {
	user := shared.Context.Get(ctx).User
	//检查任务名称是否存在
	rs := gcron.Search(req.JobName)
	if rs != nil {
		return 0, gerror.New("任务名称已经存在")
	}
	//可以task目录下是否绑定对应的方法
	//f := task.GetByName(req.JobName)
	//if f == nil {
	//	return 0, gerror.New("当前task目录下没有绑定这个方法")
	//}
	var entity model.SysJob
	entity.JobGroup = req.JobGroup
	entity.CreateTime = gtime.Now()
	entity.CreateBy = user.LoginName
	var editReq *define.JobApiUpdateReq
	gconv.Struct(req,&editReq)
	return s.save(&entity,editReq)
}

//修改数据
func (s *jobService)Update(ctx context.Context,req *define.JobApiUpdateReq) (int64, error) {
	user := shared.Context.Get(ctx).User
	//检查任务名称是否存在
	tmp := gcron.Search(req.JobName)

	if tmp != nil {
		gcron.Remove(req.JobName)
	}

	//可以task目录下是否绑定对应的方法
	//f := task.GetByName(req.JobName)
	//if f.Func == nil {
	//	return 0, gerror.New("当前task目录下没有绑定这个方法")
	//}

	entity, err := dao.SysJob.FindOne(dao.SysJob.Columns.JobId,req.JobId)
	if err != nil {
		return 0, err
	}

	if entity == nil {
		return 0, gerror.New("数据不存在")
	}

	entity.UpdateTime = gtime.Now()
	entity.UpdateBy = user.LoginName
	return s.save(entity,req)
}

func (s *jobService)save(job *model.SysJob,req *define.JobApiUpdateReq)(int64, error) {
	var (
		rs  int64
		err error
	)
	job.JobName = req.JobName
	job.JobParams = req.JobParams
	job.InvokeTarget = req.InvokeTarget
	job.CronExpression = req.CronExpression
	job.MisfirePolicy = req.MisfirePolicy
	job.Concurrent = req.Concurrent
	job.Status = req.Status
	job.Remark = req.Remark

	result, err := dao.SysJob.Data(job).Save()
	if err != nil {
		return 0, err
	}
	if job.JobId == 0 {
		// 新增
		rs, err = result.LastInsertId()
	} else {
		rs, err = result.RowsAffected()
	}
	if err != nil {
		return 0, err
	}
	if req.Status == "0" {
		// 启动
		if err = s.Start(job); err != nil {
			return rs,gerror.New("启动任务失败，您可以手动启动")
		}
	}
	return rs,err
}

//批量删除数据记录
func (s *jobService)Delete(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	list, _ := dao.SysJob.FindAll("job_id in (?)", idarr)
	if list != nil && len(list) > 0 {
		for _, j := range list {
			gcron.Remove(j.JobName)
		}
	}
	result, err := dao.SysJob.Delete("job_id in (?)", idarr)
	if err != nil {
		return 0
	}
	nums, _ := result.RowsAffected()
	return nums
}

// 导出excel
func (s *jobService)Export(param *define.JobApiSelectPageReq) (string, error) {
	m := dao.SysJob.As("t")
	if param != nil {
		if param.JobName != "" {
			m = m.Where("t.job_name like ?", "%"+param.JobName+"%")
		}
		if param.JobGroup != "" {
			m = m.Where("t.job_group = ?", param.JobGroup)
		}

		if param.InvokeTarget != "" {
			m = m.Where("t.invoke_target = ?", param.InvokeTarget)
		}

		if param.CronExpression != "" {
			m = m.Where("t.cron_expression = ?", param.CronExpression)
		}

		if param.MisfirePolicy != "" {
			m = m.Where("t.misfire_policy = ?", param.MisfirePolicy)
		}

		if param.Concurrent != "" {
			m = m.Where("t.concurrent = ?", param.Concurrent)
		}

		if param.Status != "" {
			m = m.Where("t.status = ?", param.Status)
		}
		if param.BeginTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	//"任务ID","任务名称","任务组名","调用目标字符串","cron执行表达式","计划执行错误策略（1立即执行 2执行一次 3放弃执行）","是否并发执行（0允许 1禁止）","状态（0正常 1暂停）","创建者","创建时间","更新者","更新时间","备注信息",
	m = m.Fields(" t.job_id ,t.job_name ,t.job_group ,t.invoke_target ,t.cron_expression ,t.misfire_policy ,t.concurrent ,t.status ,t.create_by ,t.create_time ,t.update_by ,t.update_time ,t.remark")
	result, err := m.M.All()
	if err != nil {
		return "", err
	}
	head := []string{"任务ID", "任务名称", "任务组名", "调用目标字符串", "cron执行表达式", "计划执行错误策略（1立即执行 2执行一次 3放弃执行）", "是否并发执行（0允许 1禁止）", "状态（0正常 1暂停）", "创建者", "创建时间", "更新者", "更新时间", "备注信息"}
	key := []string{"job_id", "job_name", "job_group", "invoke_target", "cron_expression", "misfire_policy", "concurrent", "status", "create_by", "create_time", "update_by", "update_time", "remark"}
	url, err := excel.DownlaodExcel(head, key, result)
	if err != nil {
		return "", err
	}
	return url, nil
}

//启动任务
func (s *jobService)Start(entity *model.SysJob) error {
	//可以task目录下是否绑定对应的方法
	f := task.GetByName(entity.JobName)
	if f.Func == nil {
		return gerror.New("当前task目录下没有绑定这个方法")
	}

	//传参
	paramArr := strings.Split(entity.JobParams, "|")
	f.EditParams(f.JobName, paramArr)
	rs := gcron.Search(entity.JobName)
	if rs == nil {
		var err error
		if entity.MisfirePolicy == "1" {
			// 判断是否并发执行
			if entity.Concurrent == "1"{
				// 允许并发
				_, err = gcron.Add(entity.CronExpression, f.Run, entity.JobName)
			}else{
				// 同时只允许有一个该任务运行
				_, err = gcron.AddSingleton(entity.CronExpression, f.Run, entity.JobName)
			}
			if err != nil {
				return err
			}

		} else {
			_, err = gcron.AddOnce(entity.CronExpression, f.Run, entity.JobName)
			if err != nil {
				return err
			}
		}
	}
	gcron.Start(entity.JobName)
	if entity.MisfirePolicy == "1" {
		entity.Status = "0"
		dao.SysJob.Data(entity).Save()
	}
	return nil
}
//停止任务
func (s *jobService)Stop(job *model.SysJob) error {
	//可以task目录下是否绑定对应的方法
	f := task.GetByName(job.JobName)
	if f == nil {
		return gerror.New("当前task目录下没有绑定这个方法")
	}

	rs := gcron.Search(job.JobName)

	if rs != nil {
		gcron.Stop(job.JobName)
	}
	job.Status = "1"
	dao.SysJob.Data(job).Save()
	return nil
}