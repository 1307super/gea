package dept

import "C"
import (
	"gea/app/controller"
	deptModel "gea/app/model/system/dept"
	deptService "gea/app/service/system/dept"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)


type Controller struct {
	controller.BaseController
}
func (c *Controller)Init(r *ghttp.Request) {
	c.Module = "部门管理"
}

//列表分页数据
func (c *Controller)Get(r *ghttp.Request) {
	var req *deptModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}
	result, err := deptService.SelectListAll(req)
	if err != nil {
		 c.Err(r,err.Error())
	}
	 c.Succ(r,result)
}


//新增页面保存
func (c *Controller)Post(r *ghttp.Request) {
	var req *deptModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}

	if deptService.CheckDeptNameUniqueAll(req.DeptName, req.ParentId) == "1" {
		 c.Err(r,"部门名称已存在")
	}

	rid, err := deptService.AddSave(req, r)

	if err != nil || rid <= 0 {
		 c.Err(r,err.Error())
	}
	 c.Succ(r,)
}

//修改页面保存
func (c *Controller)Put(r *ghttp.Request) {
	var req *deptModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}

	if deptService.CheckDeptNameUnique(req.DeptName, req.DeptId, req.ParentId) == "1" {
		 c.Err(r,"部门名称已存在")
	}

	rs, err := deptService.EditSave(req, r)

	if err != nil || rs <= 0 {
		 c.Err(r,err.Error())
	}
	 c.Succ(r,rs)
}

//删除数据
func (c *Controller) Delete(r *ghttp.Request) {
	id := r.GetInt64("id")

	rs := deptService.DeleteDeptById(id)

	if rs > 0 {
		 c.Succ(r,)
	} else {
		 c.Err(r,"")
	}
}

// 详情
func (c *Controller)Info(r *ghttp.Request) {
	id := r.GetInt64("id")
	if id <= 0 {
		 c.Err(r,"请选择部门")
	}
	result := deptService.SelectDeptById(id)
	 c.Succ(r,result)
}

//加载部门列表树结构的数据
func (c *Controller)TreeData(r *ghttp.Request) {
	result, err := deptService.SelectDeptTree(0, "", "")
	if err != nil {
		 c.Err(r,"获取部门列表失败")
	}
	 c.Succ(r,deptService.BuildDepts(result))
}


//加载角色部门（数据权限）列表树
func (c *Controller)RoleDeptTreeData(r *ghttp.Request) {
	roleId := r.GetInt64("roleId")
	// 1 先查出所有部门
	result, err := deptService.SelectDeptTree(0, "", "")
	if err != nil {
		 c.Err(r,err.Error())
	}
	// 2 查出权限
	deptIds, err := deptService.SelectRoleDeptIds(roleId)
	if err != nil {
		 c.Err(r,err.Error())
	}
	 c.Succ(r,g.Map{
		"depts": deptService.BuildDepts(result),
		"checkedKeys": deptIds,
	})
}