package api

import (

	"gea/app/system/admin/internal/define"
	"gea/app/system/admin/internal/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Menu = new(menuApi)

type menuApi struct {
	BaseController
}

func (a *menuApi)Init(r *ghttp.Request) {
	a.Module = "菜单管理"
}

//列表分页数据
func (a *menuApi) Get(r *ghttp.Request) {
	var req *define.MenuApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	result, err :=  service.Menu.GetAll(req)
	if err != nil {
		a.Err(r,err.Error())
	}
	a.Succ(r,result)
}

//新增页面保存
func (a *menuApi)Post(r *ghttp.Request) {
	var req *define.MenuApiCreateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	id, err := service.Menu.Create(r.Context(),req)
	if err != nil || id <= 0 {
		a.Err(r,"添加失败")
	}
	a.Succ(r,id)
}

// 详情
func (a *menuApi)Info(r *ghttp.Request) {
	menuId := r.GetInt64("id")
	if menuId < 0 {
		a.Err(r,"请选择菜单")
	}
	menu, err := service.Menu.Info(menuId)
	if err !=nil{
		a.Err(r,err.Error())
	}
	a.Succ(r,menu)
}

//修改页面保存
func (a *menuApi)Put(r *ghttp.Request) {
	var req *define.MenuApiEditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	rs, err := service.Menu.Update(r.Context(),req)
	if err != nil || rs <= 0 {
		a.Err(r,"修改失败")
	}
	a.Succ(r,rs)
}

//删除数据
func (a *menuApi)Delete(r *ghttp.Request) {
	var req *define.MenuApiDeleteReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	rs := service.Menu.Delete(req.Id)
	if rs {
		a.Succ(r)
	} else {
		a.Err(r,"删除失败")
	}
}


//加载所有菜单列表树
func (a *menuApi)MenuTreeData(r *ghttp.Request) {
	ztrees, err := service.Menu.MenuTreeData(r.Context())
	if err != nil {
		a.Err(r,err.Error())
	}
	a.Succ(r,ztrees)
}


//加载角色菜单列表树
func (a *menuApi)RoleMenuTreeData(r *ghttp.Request) {
	roleId := r.GetInt64("roleId")
	// 1 查询出用户所拥有的权限
	menus, err := service.Menu.MenuTreeData(r.Context())
	if err != nil {
		a.Err(r,err.Error())
	}
	// 2 查询出角色拥有的菜单
	menuIds, err := service.Menu.GetMenuIdsByRoleId(roleId)
	if err != nil {
		a.Err(r,err.Error())
	}

	a.Succ(r,g.Map{
		"menus": menus,
		"checkedKeys": menuIds,
	})
}

func (a *menuApi)GetRouter(r *ghttp.Request){
	result, err := service.Menu.SelectMenuNormalByUser(r.Context())
	if err != nil {
		a.Err(r,err.Error())
	}
	a.Succ(r,service.Menu.BuildMenus(result))
}
