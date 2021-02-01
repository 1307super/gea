package api

import (
	"gea/app/system/admin/internal/define"
	"gea/app/system/admin/internal/service"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

var Config = new(configApi)

type configApi struct {
	BaseController
}

func (a *configApi) Init(r *ghttp.Request) {
	a.Module = "配置管理"
}

//列表分页数据
func (a *configApi) Get(r *ghttp.Request) {
	var req *define.ConfigApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	result := service.Config.GetList(req)
	if result != nil {
		a.SuccTable(r, result.Total, result.List)
	} else {
		a.SuccTable(r, 0, nil)
	}
}

//新增
func (a *configApi) Post(r *ghttp.Request) {
	var req *define.ConfigApiCreateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	rid, err := service.Config.Create(r.Context(), req)
	if err != nil {
		a.Err(r, err.Error())
	}
	if rid <= 0 {
		a.Err(r, "新增失败")
	}
	a.Succ(r, rid)
}

//修改
func (a *configApi) Put(r *ghttp.Request) {
	var req *define.ConfigApiUpdateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	rs, err := service.Config.Update(r.Context(), req)
	if err != nil {
		a.Err(r, err.Error())
	}
	if rs <= 0 {
		a.Err(r, "修改失败")
	}
	a.Succ(r)
}

//删除数据
func (a *configApi) Delete(r *ghttp.Request) {
	var req *define.ConfigApiDeleteReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	rs := service.Config.Delete(req.Ids)
	if rs > 0 {
		a.Succ(r)
	} else {
		a.Err(r, "删除失败")
	}
}

// 详情
func (a *configApi) Info(r *ghttp.Request) {
	id := r.GetInt64("id")
	if id <= 0 {
		a.Err(r, "请选择参数")
	}
	entity, err := service.Config.Info(id)
	if err != nil || entity == nil {
		a.Err(r, "配置不存在")
	}
	a.Succ(r, entity)
}

//导出
func (a *configApi) Export(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType, response.Buniss_Export)
	var req *define.ConfigApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	url, err := service.Config.Export(req)
	if err != nil {
		a.Err(r, "导出失败")
	}
	a.Succ(r, url)
}

// 根据键获取值
func (a *configApi) GetValueByKey(r *ghttp.Request) {
	key := r.GetString("configKey")
	if key == "" {
		a.Err(r, "请输入配置名")
	}
	a.Succ(r, service.Config.GetValueByKey(key))
}
