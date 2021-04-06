package api

import (

	"gea/app/system/admin/internal/define"
	"gea/app/system/admin/internal/service"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

var Role = new(roleApi)

type roleApi struct {
	BaseController
}

func (a *roleApi)Init(r *ghttp.Request) {
	a.Module = "角色管理"
}


//列表分页数据
func (a *roleApi)Get(r *ghttp.Request) {
	var req *define.RoleApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	result := service.Role.GetList(r.Context(),req)
	a.SuccTable(r,result.Total,result.List)
}


//新增页面保存
func (a *roleApi)Post(r *ghttp.Request) {
	var req *define.RoleApiCreateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	rid, err := service.Role.Create(r.Context(),req)
	if err != nil {
		a.Err(r,err.Error())
	}
	if rid <= 0 {
		a.Err(r,"添加失败")
	}
	a.Succ(r,rid)
}

//修改
func (a *roleApi)Put(r *ghttp.Request) {
	var req *define.RoleApiEditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	rs, err := service.Role.Update(r.Context(),req)

	if err != nil {
		a.Err(r,err.Error())
	}
	if rs <= 0 {
		a.Err(r,"修改失败")
	}
	a.Succ(r,rs)
}


//删除数据
func (a *roleApi)Delete(r *ghttp.Request) {
	var req *define.RoleApiDeleteReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	rs := service.Role.Delete(req.Ids)
	if rs {
		a.Succ(r,rs)
	} else {
		a.Err(r,"删除失败")
	}
}

//导出
func (a *roleApi) Export(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,response.Buniss_Export)
	var req *define.RoleApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	url, err := service.Role.Export(req)

	if err != nil {
		a.Err(r,err.Error())
	}
	a.Succ(r,url)
}

func (a *roleApi) Info(r *ghttp.Request) {
	id := r.GetInt64("id")

	if id <= 0 {
		a.Err(r,"参数错误")
	}
	role, err := service.Role.Info(id)
	if err != nil || role == nil {
		a.Err(r,"角色不存在")
	}
	a.Succ(r,role)
}

//数据权限保存
func (a *roleApi) AuthDataScopeSave(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,response.Buniss_Authorize)
	var req *define.RoleApiDataScopeReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	rs, err := service.Role.AuthDataScope(r.Context(),req)
	if err != nil || rs <= 0 {
		a.Err(r,"保存数据失败")
	} else {
		a.Succ(r)
	}
}


// 改变角色权限状态
func (a *roleApi) ChangeStatus(r *ghttp.Request){
	var req *define.RoleApiChangeStatus
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	err := service.Role.ChangeStatus(req.RoleId,req.Status)
	if err != nil {
		a.Err(r,err.Error())
	}
	a.Succ(r)
}
