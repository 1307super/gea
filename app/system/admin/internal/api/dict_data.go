package api

import (
	"gea/app/system/admin/internal/define"
	"gea/app/system/admin/internal/service"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var DictData = new(dictDataApi)

type dictDataApi struct {
	BaseController
}

func (a *dictDataApi) Init(r *ghttp.Request) {
	a.Module = "字典数据管理"
}

//列表数据
func (a *dictDataApi) GetAll (r *ghttp.Request) {
	var req *define.DictDataApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	result, err := service.DictData.GetAll(req)
	if err != nil {
		a.Err(r,err.Error())
	}
	a.Succ(r,result)
}

//列表分页数据
func (a *dictDataApi) Get(r *ghttp.Request) {
	var req *define.DictDataApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	result := service.DictData.GetList(req)
	if result != nil {
		a.SuccTable(r, result.Total, result.List)
	} else {
		a.SuccTable(r, 0, nil)
	}
}


//新增页面保存
func (a *dictDataApi) Post(r *ghttp.Request) {
	var req *define.DictDataApiCreateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	rid, err := service.DictData.Create(r.Context(), req)
	if err != nil {
		a.Err(r, err.Error())
	}
	if rid <= 0 {
		a.Err(r, "保存失败")
	}
	a.Succ(r, rid)
}

//修改页面保存
func (a *dictDataApi) Put(r *ghttp.Request) {
	var req *define.DictDataApiUpdateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	rs, err := service.DictData.Update(r.Context(), req)
	if err != nil {
		a.Err(r, err.Error())
	}
	if rs <= 0 {
		a.Err(r, "保存失败")
	}
	a.Succ(r)
}

//删除数据
func (a *dictDataApi) Delete(r *ghttp.Request) {
	var req *define.DictDataApiDeleteReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	rs := service.DictData.Delete(req.Ids)
	if rs > 0 {
		a.Succ(r)
	} else {
		a.Err(r, "删除失败")
	}
}

//数据详情
func (a *dictDataApi) Info(r *ghttp.Request) {
	dictId := gconv.Int64(r.Get("id"))
	if dictId <= 0 {
		a.Err(r, "请选择字典类型")
	}
	dict, _ := service.DictData.Info(dictId)

	if dict == nil {
		a.Err(r, "字典不存在")
	}
	a.Succ(r, dict)
}


//导出
func (a *dictDataApi) Export(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType, response.Buniss_Export)
	var req *define.DictDataApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	url, err := service.DictData.Export(req)
	if err != nil {
		a.Err(r, err.Error())
	}
	a.Succ(r, url)
}
