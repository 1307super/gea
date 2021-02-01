package api

import (

	"gea/app/system/admin/internal/define"
	"gea/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

var Dept = new(deptApi)

type deptApi struct {
	BaseController
}


func (a *deptApi)Init(r *ghttp.Request) {
	a.Module = "部门管理"
}

//列表分页数据
func (a *deptApi)Get(r *ghttp.Request) {
	var req *define.DeptApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	result, err := service.Dept.GetAll(req)
	if err != nil {
		a.Err(r,err.Error())
	}
	a.Succ(r,result)
}


//新增页面保存
func (a *deptApi)Post(r *ghttp.Request) {
	var req *define.DeptApiCreateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}

	rid, err := service.Dept.Create(r.Context(),req)

	if err != nil {
		a.Err(r,err.Error())
	}
	if rid <= 0 {
		a.Err(r,"添加失败")
	}
	a.Succ(r,)
}

//修改页面保存
func (a *deptApi)Put(r *ghttp.Request) {
	var req *define.DeptApiUpdateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}


	rs, err := service.Dept.Update(r.Context(),req)

	if err != nil {
		a.Err(r,err.Error())
	}
	if rs <= 0 {
		a.Err(r,"修改失败")
	}
	a.Succ(r,rs)
}

//删除数据
func (a *deptApi) Delete(r *ghttp.Request) {
	id := r.GetInt64("id")
	rs := service.Dept.Delete(id)
	if rs > 0 {
		a.Succ(r,)
	} else {
		a.Err(r,"")
	}
}

// 详情
func (a *deptApi)Info(r *ghttp.Request) {
	id := r.GetInt64("id")
	if id <= 0 {
		a.Err(r,"请选择部门")
	}
	result := service.Dept.Info(id)
	a.Succ(r,result)
}

//加载部门列表树结构的数据
func (a *deptApi)TreeData(r *ghttp.Request) {
	result, err := service.Dept.DeptTree(0, "", "")
	if err != nil {
		a.Err(r,"获取部门列表失败")
	}
	a.Succ(r,service.Dept.BuildDepts(result))
}


//加载角色部门（数据权限）列表树
func (a *deptApi)RoleDeptTreeData(r *ghttp.Request) {
	roleId := r.GetInt64("roleId")
	if roleId <= 0{
		a.Err(r,"请选择角色")
	}

	result, err := service.Dept.RoleDeptTreeData(roleId)
	if err != nil {
		a.Err(r,err.Error())
	}
	a.Succ(r,result)
}
