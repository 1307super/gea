package menu

import (
	"gea/app/controller"
	menuModel "gea/app/model/system/menu"
	menuService "gea/app/service/system/menu"
	userService "gea/app/service/system/user"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
	controller.BaseController
}

func (c *Controller)Init(r *ghttp.Request) {
	c.Module = "菜单管理"
}

//列表分页数据
func (c *Controller) Get(r *ghttp.Request) {
	var req *menuModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	result, err := menuService.SelectListAll(req)
	if err != nil {
		c.Err(r,err.Error())
	}
	c.Succ(r,result)
}

//新增页面保存
func (c *Controller)Post(r *ghttp.Request) {
	var req *menuModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}

	if menuService.CheckMenuNameUniqueAll(req.MenuName, req.ParentId) == "1" {
		c.Err(r,"菜单名称已存在")
	}

	id, err := menuService.AddSave(req, r)

	if err != nil || id <= 0 {
		c.Err(r,"添加失败")
	}
	c.Succ(r,id)
}

// 详情
func (c *Controller)Info(r *ghttp.Request) {
	menuId := r.GetInt64("id")
	if menuId < 0 {
		c.Err(r,"请选择菜单")
	}
	menu, err := menuService.SelectRecordById(menuId)
	if err !=nil{
		c.Err(r,err.Error())
	}
	c.Succ(r,menu)
}

//修改页面保存
func (c *Controller)Put(r *ghttp.Request) {
	var req *menuModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}

	if menuService.CheckMenuNameUnique(req.MenuName, req.MenuId, req.ParentId) == "1" {
		c.Err(r,"菜单名称已存在")
	}
	rs, err := menuService.EditSave(req, r)
	if err != nil || rs <= 0 {
		c.Err(r,"修改失败")
	}
	c.Succ(r,rs)
}

//删除数据
func (c *Controller)Delete(r *ghttp.Request) {
	id := r.GetInt64("id")
	rs := menuService.DeleteRecordById(id)
	if rs {
		c.Succ(r)
	} else {
		c.Err(r,"删除失败")
	}
}


//加载所有菜单列表树
func (c *Controller)MenuTreeData(r *ghttp.Request) {
	user,_ := userService.GetProfileApi(r.GetInt64("jwtUid"))
	if user == nil {
		c.Err(r,"登录超时")
	}
	ztrees, err := menuService.MenuTreeData(user.UserId)
	if err != nil {
		c.Err(r,err.Error())
	}
	c.Succ(r,ztrees)
}


//加载角色菜单列表树
func (c *Controller)RoleMenuTreeData(r *ghttp.Request) {
	roleId := r.GetInt64("roleId")
	user, _ := userService.GetProfileApi(r.GetInt64("jwtUid"))
	if user == nil || user.UserId <= 0 {
		c.Err(r,"登录超时")
	}
	// 1 查询出用户所拥有的权限
	menus, err :=menuService.MenuTreeData(user.UserId)
	if err != nil {
		c.Err(r,err.Error())
	}
	// 2 查询出角色拥有的菜单
	menuIds, err := menuService.SelectMenuListByRoleId(roleId)
	if err != nil {
		c.Err(r,err.Error())
	}

	c.Succ(r,g.Map{
		"menus": menus,
		"checkedKeys": menuIds,
	})
}

func (c *Controller)GetRouter(r *ghttp.Request){
	user,_ := userService.GetProfileApi(r.GetInt64("jwtUid"))
	if user == nil || user.UserId <= 0 {
		c.Err(r,"登录超时")
	}

	result, err := menuService.SelectMenuNormalByUser(user.UserId)
	if err != nil {
		c.Err(r,err.Error())
	}
	c.Succ(r,menuService.BuildMenus(result))
}
