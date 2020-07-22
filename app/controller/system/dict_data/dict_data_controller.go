package dict_data

import (
	"gea/app/controller"
	"gea/app/model"
	dictModel "gea/app/model/system/dict_data"
	dictService "gea/app/service/system/dict_data"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
	controller.BaseController
}

func (c *Controller)Init(r *ghttp.Request) {
	c.Module = "字典数据管理"
}

//列表数据
func (c *Controller) List (r *ghttp.Request) {
	var req *dictModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}
	result, err := dictService.SelectListAll(req)
	if err != nil {
		 c.Err(r,err.Error())
	}
	 c.Succ(r,result)
}

//列表分页数据
func (c *Controller) Get(r *ghttp.Request) {
	var req *dictModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}
	result, page, err := dictService.SelectListByPage(req)

	if err != nil {
		 c.Err(r,err.Error())
	}
	c.SuccTable(r,page.Total,result)
}

//新增页面保存
func (c *Controller) Post(r *ghttp.Request) {
	var req *dictModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}

	rid, err := dictService.AddSave(req, r)

	if err != nil || rid <= 0 {
		 c.Err(r,err.Error())
	}
	 c.Succ(r,rid)
}

func (c *Controller)Info(r *ghttp.Request) {
	id := r.GetInt64("id")

	if id <= 0 {
		 c.Err(r,"字典数据错误")
	}

	entity, err := dictService.SelectRecordById(id)

	if err != nil || entity == nil {
		 c.Err(r,"字典数据不存在")
	}
	 c.Succ(r,entity)
}

//修改页面保存
func (c *Controller) Put (r *ghttp.Request) {
	var req *dictModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}

	rs, err := dictService.EditSave(req, r)

	if err != nil || rs <= 0 {
		 c.Err(r,"修改失败")
	}
	 c.Succ(r,)
}

//删除数据
func (c *Controller) Delete(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}

	rs := dictService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		 c.Succ(r,)
	} else {
		 c.Err(r,"删除失败")
	}
}

//导出
func (c *Controller) Export(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,model.Buniss_Export)
	var req *dictModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}
	url, err := dictService.Export(req)

	if err != nil {
		 c.Err(r,err.Error())
	}
	 c.Succ(r,url)
}
