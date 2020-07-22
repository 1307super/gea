package operlog

import (
	"gea/app/controller"
	"gea/app/model"
	operlogModel "gea/app/model/monitor/oper_log"
	operlogService "gea/app/service/monitor/operlog"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
	controller.BaseController
}

func (c *Controller) Init(r *ghttp.Request) {
	c.Module = "操作日志管理"
}


//用户列表分页数据
func (c *Controller)Get(r *ghttp.Request) {
	var req *operlogModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}

	rows := make([]operlogModel.Entity, 0)

	result, page, err := operlogService.SelectPageList(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	c.SuccTable(r,page.Total,rows)
}

//删除数据
func (c *Controller) Delete(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	rs := operlogService.DeleteRecordByIds(req.Ids)
	if rs > 0 {
		c.Succ(r,rs)
	} else {
		c.Err(r,"删除失败")
	}
}

//清空记录
func (c *Controller) Clean(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,model.Buniss_Clean)
	rs := operlogService.DeleteRecordAll()
	if rs > 0 {
		c.Succ(r,rs)
	} else {
		c.Err(r,"清空失败")
	}
}
//导出
func (c *Controller) Export(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,model.Buniss_Export)
	var req *operlogModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	url, err := operlogService.Export(req)

	if err != nil {
		c.Err(r,err.Error())
	} else {
		c.Succ(r,url)
	}
}
