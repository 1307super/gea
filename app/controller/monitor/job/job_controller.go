package job

import (
	"gea/app/controller"
	"gea/app/model"
	jobModel "gea/app/model/monitor/job"
	jobService "gea/app/service/monitor/job"
	"gea/app/utils/convert"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcron"
)

type Controller struct {
	controller.BaseController
}

func (c *Controller) Init(r *ghttp.Request) {
	c.Module = "定时任务管理"
}

//列表分页数据
func (c *Controller) Get(r *ghttp.Request) {
	var req *jobModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r, err.Error())
	}
	rows := make([]jobModel.Entity, 0)
	result, page, err := jobService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	c.SuccTable(r, page.Total, rows)
}

//新增页面保存
func (c *Controller) Post(r *ghttp.Request) {
	var req *jobModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r, err.Error())
	}

	id, err := jobService.AddSave(req, r)

	if err != nil || id <= 0 {
		c.Err(r, err.Error())
	}
	c.Succ(r, id)
}

func (c *Controller) Info(r *ghttp.Request) {
	id := r.GetInt64("id")
	if id <= 0 {
		c.Err(r, "参数错误")
	}
	entity, err := jobService.SelectRecordById(id)
	if err != nil || entity == nil {
		c.Err(r, "数据不存在")
	}
	c.Succ(r, entity)
}

//修改页面保存
func (c *Controller) Put(r *ghttp.Request) {
	var req *jobModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r, err.Error())
	}
	rs, err := jobService.EditSave(req, r)
	if err != nil || rs <= 0 {
		c.Err(r, err.Error())
	}
	c.Succ(r, rs)
}

//删除数据
func (c *Controller) Delete(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r, err.Error())
	}

	idarr := convert.ToInt64Array(req.Ids, ",")
	list, _ := jobModel.FindAll("job_id in (?)", idarr)
	if list != nil && len(list) > 0 {
		for _, j := range list {
			gcron.Remove(j.JobName)
		}
	}

	rs := jobService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		c.Succ(r, rs)
	} else {
		c.Err(r, "删除失败")
	}
}

//导出
func Export(r *ghttp.Request) {
	var req *jobModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("定时任务管理", req).WriteJsonExit()
	}
	url, err := jobService.Export(req)

	if err != nil {
		response.ErrorResp(r).SetMsg(err.Error()).Log("定时任务管理", req).WriteJsonExit()
	}
	response.SucessResp(r).SetMsg(url).Log("定时任务管理", req).WriteJsonExit()
}

//启动 调用一次
func (c *Controller) Start(r *ghttp.Request) {
	jobId := r.GetInt64("jobId")
	if jobId <= 0 {
		c.Err(r, "参数错误")
	}
	job, _ := jobService.SelectRecordById(jobId)
	if job == nil {
		c.Err(r, "任务不存在")
	}
	job.MisfirePolicy = "0"
	job.CronExpression = "* * * * * *"
	err := jobService.Start(job)
	if err != nil {
		c.Err(r, err.Error())
	} else {
		c.Succ(r)
	}
}

func (c *Controller) ChangeStatus(r *ghttp.Request) {
	jobId := r.GetInt64("jobId")
	status := r.GetString("status")
	if jobId <= 0 {
		c.Err(r, "参数错误")
	}
	job, _ := jobService.SelectRecordById(jobId)
	if job == nil {
		c.Err(r, "任务不存在")
	}
	if status == "0" {
		// 启动
		err := jobService.Start(job)
		if err != nil {
			c.Err(r,"任务启动失败")
		}
	} else {
		// 停止
		err := jobService.Stop(job)
		if err != nil {
			c.Err(r,"任务停止失败")
		}
	}
	c.Succ(r)
}

//停止
//func Stop(r *ghttp.Request) {
//	jobId := r.GetFormInt64("jobId")
//	if jobId <= 0 {
//		response.ErrorResp(r).SetMsg("参数错误").WriteJsonExit()
//	}
//	job, _ := jobService.SelectRecordById(jobId)
//	if job == nil {
//		response.ErrorResp(r).SetMsg("任务不存在").WriteJsonExit()
//	}
//	err := jobService.Stop(job)
//	if err != nil {
//		response.ErrorResp(r).SetMsg(err.Error()).WriteJsonExit()
//	} else {
//		response.SucessResp(r).SetMsg("停止成功").WriteJsonExit()
//	}
//}
