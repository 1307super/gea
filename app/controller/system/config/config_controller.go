package config

import (
	"gea/app/controller"
	"gea/app/model"
	configModel "gea/app/model/system/config"
	configService "gea/app/service/system/config"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
	controller.BaseController
}

func (c *Controller)Init(r *ghttp.Request) {
	c.Module = "配置管理"
}

//列表分页数据
func (c *Controller)Get(r *ghttp.Request) {
	var req *configModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}
	rows := make([]configModel.Entity, 0)
	result, page, err := configService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	c.SuccTable(r,page.Total,rows)
}

//新增
func (c *Controller) Post (r *ghttp.Request) {
	var req *configModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}

	if configService.CheckConfigKeyUniqueAll(req.ConfigKey) == "1" {
		 c.Err(r,"参数键名已存在")
	}

	rid, err := configService.AddSave(req, r)

	if err != nil || rid <= 0 {
		 c.Err(r,"新增失败")
	}
	 c.Succ(r,rid)
}

//修改
func (c *Controller) Put(r *ghttp.Request) {
	var req *configModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}

	if configService.CheckConfigKeyUnique(req.ConfigKey, req.ConfigId) == "1" {
		 c.Err(r,"参数键名已存在")
	}

	rs, err := configService.EditSave(req, r)

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

	rs := configService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		 c.Succ(r,)
	} else {
		 c.Err(r,"删除失败")
	}
}

// 详情
func (c *Controller) Info(r *ghttp.Request) {
	id := r.GetInt64("id")

	if id <= 0 {
		 c.Err(r,"请选择参数")
	}

	entity, err := configService.SelectRecordById(id)

	if err != nil || entity == nil {
		 c.Err(r,"参数不存在")
	}
	 c.Succ(r,entity)
}



//导出
func (c *Controller) Export(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,model.Buniss_Export)
	var req *configModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}
	url, err := configService.Export(req)

	if err != nil {
		 c.Err(r,"导出失败")
	}
	 c.Succ(r,url)
}

// 根据键获取值
func (c *Controller) GetValueByKey(r *ghttp.Request){
	key := r.GetString("configKey")
	if key == "" {
		c.Err(r,"请输入配置名")
	}
	c.Succ(r,configService.GetValueByKey(key))
	//response.SucessResp(r).SetMsg(configService.GetValueByKey(key)).SetData(configService.GetValueByKey(key)).WriteJsonExit()
}
