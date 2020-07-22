package online

import (
	"gea/app/controller"
	"gea/app/model"
	onlineModel "gea/app/model/monitor/online"
	onlineService "gea/app/service/monitor/online"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
	controller.BaseController
}

func (c *Controller) Init(r *ghttp.Request) {
	c.Module = "在线用户管理"
}

//列表分页数据
func (c *Controller) Get(r *ghttp.Request) {
	var req *onlineModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	rows := make([]onlineModel.Entity, 0)
	result, page, err := onlineService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	c.SuccTable(r,page.Total,rows)
}

//用户强退
func (c *Controller) Delete(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,model.Buniss_ForceLogout)
	tokenStr := r.GetString("tokenStr")
	if tokenStr == "" {
		c.Err(r,"参数错误")
	}
	onlineService.ForceLogout(tokenStr)
	c.Succ(r)
}
