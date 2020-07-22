package role

import (
	"gea/app/controller"
	"gea/app/model"
	roleModel "gea/app/model/system/role"
	roleService "gea/app/service/system/role"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
	controller.BaseController
}

func (c *Controller)Init(r *ghttp.Request) {
	c.Module = "角色管理"
}


//列表分页数据
func (c *Controller)Get(r *ghttp.Request) {
	var req *roleModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	rows := make([]roleModel.Entity, 0)
	result, page, err := roleService.SelectRecordPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	c.SuccTable(r,page.Total,rows)
}


//新增页面保存
func (c *Controller)Post(r *ghttp.Request) {
	var req *roleModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}

	if roleService.CheckRoleNameUniqueAll(req.RoleName) == "1" {
		c.Err(r,"角色名称已存在")
	}

	if roleService.CheckRoleKeyUniqueAll(req.RoleKey) == "1" {
		c.Err(r,"角色权限已存在")
	}

	rid, err := roleService.AddSave(req, r)
	if err != nil || rid <= 0 {
		c.Err(r,"添加失败")
	}
	c.Succ(r,rid)
}

//修改
func (c *Controller)Put(r *ghttp.Request) {
	var req *roleModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}

	if roleService.CheckRoleNameUnique(req.RoleName, req.RoleId) == "1" {
		c.Err(r,"角色名称已存在")
	}

	if roleService.CheckRoleKeyUnique(req.RoleKey, req.RoleId) == "1" {
		c.Err(r,"角色权限已存在")
	}
	rs, err := roleService.EditSave(req, r)

	if err != nil || rs <= 0 {
		c.Err(r,"修改失败")
	}
	c.Succ(r,rs)
}


//删除数据
func (c *Controller)Delete(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	rs := roleService.DeleteRecordByIds(req.Ids)
	if rs > 0 {
		c.Succ(r,rs)
	} else {
		c.Err(r,"删除失败")
	}
}

//导出
func (c *Controller) Export(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,model.Buniss_Export)
	var req *roleModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	url, err := roleService.Export(req)

	if err != nil {
		c.Err(r,err.Error())
	}
	c.Succ(r,url)
}

func (c *Controller) Info(r *ghttp.Request) {
	id := r.GetInt64("id")

	if id <= 0 {
		c.Err(r,"参数错误")
	}
	role, err := roleService.SelectRecordById(id)
	if err != nil || role == nil {
		c.Err(r,"角色不存在")
	}
	c.Succ(r,role)
}

//数据权限保存
func (c *Controller) AuthDataScopeSave(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,model.Buniss_Authorize)
	var req *roleModel.DataScopeReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	if !roleService.CheckRoleAllowed(req.RoleId) {
		c.Err(r,"不允许操作超级管理员角色")
	}

	rs, err := roleService.AuthDataScope(req, r)
	if err != nil || rs <= 0 {
		c.Err(r,"保存数据失败")
	} else {
		c.Succ(r)
	}
}



// 改变角色权限状态
func (c *Controller) ChangeStatus(r *ghttp.Request){
	var req *roleModel.ChangeStatus
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	err := roleService.ChangeStatus(req.RoleId,req.Status)
	if err != nil {
		c.Err(r,err.Error())
	}
	c.Succ(r)
}
