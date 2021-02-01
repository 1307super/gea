package service

import (
	"gea/app/dao"
	"gea/app/system/admin/internal/define"
	"gea/app/utils/convert"
	"gea/app/utils/page"
)

var JobLog = &jobLogService{}
type jobLogService struct {}

//根据条件分页查询数据
func (s *jobLogService)GetList(param *define.JobLogApiSelectPageReq) *define.JobLogServiceList {
	m := dao.SysJobLog.As("t")
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
		if param.JobMessage != "" {
			m = m.Where("t.job_message = ?", param.JobMessage)
		}
		if param.Status != "" {
			m = m.Where("t.status = ?", param.Status)
		}
		if param.ExceptionInfo != "" {
			m = m.Where("t.exception_info = ?", param.ExceptionInfo)
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
	result := &define.JobLogServiceList{
		Page:  page.PageNum,
		Size:  page.Pagesize,
		Total: page.Total,
	}
	if err = m.Structs(&result.List); err != nil {
		return nil
	}
	return result
}

//批量删除数据记录
func (s *jobLogService)Delete(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := dao.SysJobLog.Delete("job_log_id in (?)", idarr)
	if err != nil {
		return 0
	}
	nums, _ := result.RowsAffected()
	return nums
}

// 清空
func (s *jobLogService)Clean() int64 {
	result, err := dao.SysJobLog.Where("job_log_id > ?","0").Delete()
	if err != nil {
		return 0
	}
	nums, _ := result.RowsAffected()
	return nums
}