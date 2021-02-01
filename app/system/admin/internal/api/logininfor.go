package api

import (
	"gea/app/system/admin/internal/define"
	"gea/app/system/admin/internal/service"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

var Logininfor = new(logininforApi)

type logininforApi struct {
	BaseController
}

func (a *logininforApi) Init(r *ghttp.Request) {
	a.Module = "登录日志管理"
}

//列表分页数据
func (a *logininforApi) Get(r *ghttp.Request) {
	var req *define.LogininforApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	result := service.Logininfor.GetList(req)
	if result != nil {
		a.SuccTable(r, result.Total, result.List)
	} else {
		a.SuccTable(r, 0, nil)
	}
}

//删除数据
func (a *logininforApi) Delete(r *ghttp.Request) {
	var req *define.LogininforApiDeleteReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	rs := service.Logininfor.Delete(req.Ids)
	if rs > 0 {
		a.Succ(r, rs)
	} else {
		a.Err(r, "删除失败")
	}
}

//清空记录
func (a *logininforApi) Clean(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType, response.Buniss_Clean)
	rs := service.Logininfor.Clean()
	if rs > 0 {
		a.Succ(r, rs)
	} else {
		a.Err(r, "清空失败")
	}
}

//导出
func (a *logininforApi) Export(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType, response.Buniss_Export)
	var req *define.LogininforApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	url, err := service.Logininfor.Export(req)
	if err != nil {
		a.Err(r, err.Error())
	} else {
		a.Succ(r, url)
	}
}
