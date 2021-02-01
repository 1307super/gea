package api

import (

	"gea/app/system/admin/internal/define"
	"gea/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

var JobLog = new(jobLogApi)

type jobLogApi struct {
	BaseController
}

func (a *jobLogApi) Init(r *ghttp.Request) {
	a.Module = "定时任务日志管理"
}

//列表分页数据
func (a *jobLogApi) Get(r *ghttp.Request) {
	var req *define.JobLogApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	result := service.JobLog.GetList(req)
	if result != nil {
		a.SuccTable(r, result.Total, result.List)
	} else {
		a.SuccTable(r, 0, nil)
	}
}

//删除数据
func (a *jobLogApi) DeleteLog(r *ghttp.Request) {
	var req *define.JobLogApiDeleteReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	rs := service.JobLog.Delete(req.Ids)
	if rs > 0 {
		a.Succ(r, rs)
	} else {
		a.Err(r, "删除失败")
	}
}

//清空记录
func (a *jobLogApi) Clean(r *ghttp.Request) {
	rs := service.JobLog.Clean()
	if rs > 0 {
		a.Succ(r, rs)
	} else {
		a.Err(r, "")
	}
}
