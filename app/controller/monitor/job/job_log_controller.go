package job

import (
	"gea/app/model"
	jobLogModel "gea/app/model/monitor/job_log"
	jobLogService "gea/app/service/monitor/job_log"
	"github.com/gogf/gf/net/ghttp"
)

//列表分页数据
func (c *Controller) GetLog (r *ghttp.Request) {
	var req *jobLogModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	rows := make([]jobLogModel.Entity, 0)
	result, page, err := jobLogService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	c.SuccTable(r,page.Total,rows)
}

//删除数据
func (c *Controller) DeleteLog(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	rs := jobLogService.DeleteRecordByIds(req.Ids)
	if rs > 0 {
		c.Succ(r,rs)
	} else {
		c.Err(r,"删除失败")
	}
}

//清空记录
func (c *Controller)CleanLog(r *ghttp.Request) {
	rs := jobLogService.DeleteRecordAll()
	if rs > 0 {
		c.Succ(r,rs)
	} else {
		c.Err(r,"")
	}
}


