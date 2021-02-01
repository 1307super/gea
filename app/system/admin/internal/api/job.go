package api

import (
	"gea/app/system/admin/internal/define"
	"gea/app/system/admin/internal/service"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

var Job = new(jobApi)

type jobApi struct {
	BaseController
}

func (a *jobApi) Init(r *ghttp.Request) {
	a.Module = "定时任务管理"
}

//列表分页数据
func (a *jobApi) Get(r *ghttp.Request) {
	var req *define.JobApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	result := service.Job.GetList(req)
	if result != nil {
		a.SuccTable(r, result.Total, result.List)
	} else {
		a.SuccTable(r, 0, nil)
	}
}

//新增页面保存
func (a *jobApi) Post(r *ghttp.Request) {
	var req *define.JobApiCreateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}

	id, err := service.Job.Create(r.Context(),req)

	if err != nil {
		a.Err(r, err.Error())
	}
	if  id <= 0 {
		a.Err(r, "保存失败")
	}
	a.Succ(r, id)
}

//修改页面保存
func (a *jobApi) Put(r *ghttp.Request) {
	var req *define.JobApiUpdateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	rs, err := service.Job.Update(r.Context(),req)
	if err != nil {
		a.Err(r, err.Error())
	}
	if  rs <= 0 {
		a.Err(r, "修改失败")
	}
	a.Succ(r, rs)
}

//删除数据
func (a *jobApi) Delete(r *ghttp.Request) {
	var req *define.JobApiDeleteReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	rs := service.Job.Delete(req.Ids)
	if rs > 0 {
		a.Succ(r, rs)
	} else {
		a.Err(r, "删除失败")
	}
}

func (a *jobApi) Info(r *ghttp.Request) {
	id := r.GetInt64("id")
	if id <= 0 {
		a.Err(r, "参数错误")
	}
	entity, err := service.Job.Info(id)
	if err != nil || entity == nil {
		a.Err(r, "数据不存在")
	}
	a.Succ(r, entity)
}



//导出
func (a *jobApi)Export(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,response.Buniss_Export)
	var req *define.JobApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	url, err := service.Job.Export(req)
	if err != nil {
		a.Err(r,err.Error())
	}
	a.Succ(r,url)
}

//启动 调用一次
func (a *jobApi) Start(r *ghttp.Request) {
	jobId := r.GetInt64("jobId")
	if jobId <= 0 {
		a.Err(r, "参数错误")
	}
	job, err := service.Job.Info(jobId)
	if err != nil || job == nil {
		a.Err(r, "任务不存在")
	}
	job.MisfirePolicy = "0"
	job.CronExpression = "* * * * * *"
	err = service.Job.Start(job)
	if err != nil {
		a.Err(r, err.Error())
	} else {
		a.Succ(r)
	}
}

func (a *jobApi) ChangeStatus(r *ghttp.Request) {
	jobId := r.GetInt64("jobId")
	status := r.GetString("status")
	if jobId <= 0 {
		a.Err(r, "参数错误")
	}
	job, err := service.Job.Info(jobId)
	if job == nil || err != nil {
		a.Err(r, "任务不存在")
	}
	if status == "0" {
		// 启动
		err = service.Job.Start(job)
		if err != nil {
			a.Err(r,"任务启动失败")
		}
	} else {
		// 停止
		err = service.Job.Stop(job)
		if err != nil {
			a.Err(r,"任务停止失败")
		}
	}
	a.Succ(r)
}