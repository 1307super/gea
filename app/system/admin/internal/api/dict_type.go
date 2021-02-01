package api

import (
	"gea/app/system/admin/internal/define"
	"gea/app/system/admin/internal/service"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var DictType = new(dictTypeApi)

type dictTypeApi struct {
	BaseController
}

func(a *dictTypeApi)Init(r *ghttp.Request) {
	a.Module = "字典管理"
}

//列表分页数据
func(a *dictTypeApi) Get(r *ghttp.Request) {
	var req *define.DictTypeApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	result := service.DictType.GetList(req)
	if result != nil {
		a.SuccTable(r, result.Total, result.List)
	}else{
		a.SuccTable(r, 0, nil)
	}
}

//新增页面保存
func(a *dictTypeApi) Post(r *ghttp.Request) {
	var req *define.DictTypeApiCreateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}

	rid, err := service.DictType.Create(r.Context(),req)

	if err != nil {
		a.Err(r,err.Error())
	}
	if rid <= 0  {
		a.Err(r,"保存失败")
	}
	a.Succ(r,rid)
}

//修改页面保存
func(a *dictTypeApi) Put(r *ghttp.Request) {
	var req *define.DictTypeApiUpdateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}

	rs, err := service.DictType.Update(r.Context(),req)

	if err != nil {
		a.Err(r,err.Error())
	}
	if rs <= 0  {
		a.Err(r,"保存失败")
	}
	a.Succ(r,)
}

//删除数据
func(a *dictTypeApi) Delete(r *ghttp.Request) {
	var req *define.DictTypeApiDeleteReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	rs := service.DictType.Delete(req.Ids)
	if rs > 0 {
		a.Succ(r,)
	} else {
		a.Err(r,"删除失败")
	}
}
//数据详情
func(a *dictTypeApi) Info(r *ghttp.Request) {
	dictId := gconv.Int64(r.Get("dictId"))
	if dictId <= 0 {
		a.Err(r,"请选择字典类型")
	}
	dict, _ := service.DictType.Info(dictId)

	if dict == nil {
		a.Err(r,"字典不存在")
	}
	a.Succ(r,dict)
}

func(a *dictTypeApi) Optionselect(r *ghttp.Request) {
	var req *define.DictTypeApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	result, err := service.DictType.GetAll(req)
	if err != nil {
		a.Err(r,err.Error())
	}
	a.Succ(r,result)
}

//导出
func(a *dictTypeApi) Export(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,response.Buniss_Export)
	var req *define.DictTypeApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	url, err := service.DictType.Export(req)
	if err != nil {
		a.Err(r,err.Error())
	}
	a.Succ(r,url)
}