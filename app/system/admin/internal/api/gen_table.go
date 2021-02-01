package api

import (
	"gea/app/system/admin/internal/define"
	"gea/app/system/admin/internal/service"
	"gea/app/utils/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var GenTable = new(genTableApi)
type genTableApi struct {
	BaseController
}

func (a *genTableApi) Init(r *ghttp.Request) {
	a.Module = "代码生成管理"
}


func (a *genTableApi) Get (r *ghttp.Request) {
	var req *define.GenTableApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	result := service.GenTable.GetList(req)
	if result != nil {
		a.SuccTable(r, result.Total, result.List)
	} else {
		a.SuccTable(r, 0, nil)
	}
}
//导入表结构（保存）
func (a *genTableApi) Post(r *ghttp.Request) {
	tables := r.GetQueryString("tables")
	if tables == "" {
		a.Err(r,"参数错误")
	}
	if err := service.GenTable.Create(r.Context(),tables);err !=nil{
		a.Err(r,err.Error())
	}
	a.Succ(r)
}
//修改数据保存
func (a *genTableApi)Put(r *ghttp.Request) {
	var req *define.GenTableApiUpdateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	_, err := service.GenTable.Update(r.Context(),req)
	if err != nil {
		a.Err(r,err.Error())
	}
	a.Succ(r)
}
//删除数据
func (a *genTableApi) Delete(r *ghttp.Request) {
	var req *define.GenTableApiDeleteReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}

	rs := service.GenTable.Delete(req.Ids)

	if rs > 0 {
		a.Succ(r)
	} else {
		a.Err(r,"删除失败")
	}
}


func (a *genTableApi) Info(r *ghttp.Request) {
	id := r.GetInt64("id")
	if id <= 0 {
		a.Err(r,"参数错误")
	}
	entity, err := service.GenTable.Info(id)
	if err != nil {
		a.Err(r,"参数不存在")
	}
	if entity != nil {
		a.Succ(r,g.Map{
			"info":        entity,
			"rows": entity.Columns,
		})
	}
	a.Err(r,"参数不存在")

}

//预览代码
func (a *genTableApi) Preview(r *ghttp.Request) {
	tableId := r.GetInt64("tableId")
	if tableId <= 0 {
		a.Err(r,"参数错误")
	}

	result := service.GenTable.Preview(r,tableId)
	if result != nil{
		a.Succ(r,result)
	}
	a.Err(r,"预览失败")

}

//生成代码
func (a *genTableApi) GenCode(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,response.Buniss_Gen)
	tableId := r.GetQueryString("tableId")
	if tableId == "" {
		a.Err(r,"参数错误")
	}

	if err := service.GenTable.GenCode(r,tableId); err != nil{
		a.Succ(r,err.Error())
	}
	a.Succ(r)
}

//查询数据库列表
func (a *genTableApi) DataList(r *ghttp.Request) {
	var req *define.GenTableApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	result := service.GenTable.GetTables(req)
	if result != nil {
		a.SuccTable(r, result.Total, result.List)
	} else {
		a.SuccTable(r, 0, nil)
	}
}

