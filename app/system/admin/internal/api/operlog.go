package api

import (
	"gea/app/system/admin/internal/define"
	"gea/app/system/admin/internal/service"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

var Operlog = new(operlogApi)

type operlogApi struct {
	BaseController
}

func (a *operlogApi) Init(r *ghttp.Request) {
	a.Module = "操作日志管理"
}

//列表分页数据
func (a *operlogApi) Get(r *ghttp.Request) {
	var req *define.OperlogApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	result := service.OperLog.GetList(req)
	if result != nil {
		a.SuccTable(r, result.Total, result.List)
	} else {
		a.SuccTable(r, 0, nil)
	}
}

//删除数据
func (a *operlogApi) Delete(r *ghttp.Request) {
	var req *define.OperlogApiDeleteReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	rs := service.OperLog.Delete(req.Ids)
	if rs > 0 {
		a.Succ(r, rs)
	} else {
		a.Err(r, "删除失败")
	}
}

//清空记录
func (a *operlogApi) Clean(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType, response.Buniss_Clean)
	rs := service.OperLog.Clean()
	if rs > 0 {
		a.Succ(r, rs)
	} else {
		a.Err(r, "清空失败")
	}
}

//导出
func (a *operlogApi) Export(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType, response.Buniss_Export)
	var req *define.OperlogApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	url, err := service.OperLog.Export(req)
	if err != nil {
		a.Err(r, err.Error())
	} else {
		a.Succ(r, url)
	}
}
