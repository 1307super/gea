package menu

import (
	"gea/app/model"
	menuModel "gea/app/model/system/menu"
	userService "gea/app/service/system/user"
	"gea/app/utils/convert"
	"gea/app/utils/page"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"time"
)

//根据主键查询数据
func SelectRecordById(id int64) (*menuModel.EntityExtend, error) {
	return menuModel.SelectRecordById(id)
}

//根据条件查询数据
func SelectListAll(params *menuModel.SelectPageReq) ([]*menuModel.EntityExtend, error) {
	return menuModel.SelectListAll(params)
}

//根据条件分页查询数据
func SelectListPage(params *menuModel.SelectPageReq) ([]menuModel.Entity, *page.Paging, error) {
	return menuModel.SelectListPage(params)
}

//根据主键删除数据
func DeleteRecordById(id int64) bool {
	result, err := menuModel.Delete("menu_id", id)
	if err == nil {
		affected, _ := result.RowsAffected()
		if affected > 0 {
			return true
		}
	}
	ClearCache()
	return false
}

//添加数据
func AddSave(req *menuModel.AddReq, r *ghttp.Request) (int64, error) {
	var menu menuModel.Entity
	if req.MenuType == "F" {
		req.Component = ""
		req.Path = ""
		req.Icon = ""
	}
	if req.MenuType == "M" {
		req.Component = ""
		req.Perms = ""
	}
	menu.MenuName = req.MenuName
	menu.Visible = req.Visible
	menu.ParentId = req.ParentId
	menu.Remark = ""
	menu.MenuType = req.MenuType
	menu.Path = req.Path
	menu.Perms = req.Perms
	menu.Method = req.Method
	menu.IsFrame = gconv.Int(req.IsFrame)
	menu.Icon = req.Icon
	menu.OrderNum = req.OrderNum
	menu.Component = req.Component
	menu.Status = gconv.Int(req.Status)
	menu.CreateTime = gtime.Now()
	menu.CreateBy = ""

	user, _ := userService.GetProfileApi(r.GetInt64("jwtUid"))

	if user != nil {
		menu.CreateBy = user.LoginName
	}

	var err error
	tx, err := g.DB().Begin()
	if err != nil {
		return 0, err
	}

	result, err := tx.Table("sys_menu").Insert(menu)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil || id <= 0 {
		tx.Rollback()
		return 0, err
	}
	ClearCache()
	return id, tx.Commit()
}

//修改数据
func EditSave(req *menuModel.EditReq, r *ghttp.Request) (int64, error) {

	menu, err := menuModel.FindOne("menu_id=?", req.MenuId)

	if err != nil {
		return 0, err
	}

	if menu == nil {
		return 0, gerror.New("角色不存在")
	}

	if req.MenuType == "F" {
		req.Component = ""
		req.Path = ""
		req.Icon = ""
	}
	if req.MenuType == "M" {
		req.Component = ""
		req.Perms = ""
	}
	menu.MenuName = req.MenuName
	menu.Visible = req.Visible
	menu.ParentId = req.ParentId
	menu.Remark = ""
	menu.MenuType = req.MenuType
	menu.Path = req.Path
	menu.Perms = req.Perms
	menu.Method = req.Method
	menu.IsFrame = gconv.Int(req.IsFrame)
	menu.Icon = req.Icon
	menu.OrderNum = req.OrderNum
	menu.Component = req.Component
	menu.Status = gconv.Int(req.Status)
	menu.CreateTime = gtime.Now()
	menu.CreateBy = ""

	user, _ := userService.GetProfileApi(r.GetInt64("jwtUid"))

	if user == nil {
		menu.UpdateBy = user.LoginName
	}

	tx, err := g.DB().Begin()
	if err != nil {
		return 0, err
	}

	_, err = tx.Table("sys_menu").Update(menu, "menu_id="+gconv.String(menu.MenuId))

	if err != nil {
		tx.Rollback()
		return 0, err
	}
	ClearCache()
	return 1, tx.Commit()
}

//批量删除数据记录
func DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := menuModel.Delete("menu_id in (?)", idarr)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

//加载所有菜单列表树
func MenuTreeData(userId int64) ([]menuModel.RoleMenuTree, error) {
	var menuList []*menuModel.EntityExtend
	var err error
	if userService.IsAdmin(userId) {
		menuList, err = SelectMenuAll()
	} else {
		menuList, err = SelectMenusListByUserId(userId)
	}
	if err != nil {
		return nil, err
	}
	menuList = getChildPerms(menuList, 0)
	return BuildRoleMenus(menuList), nil
}

//获取管理员菜单树数据
func SelectMenuAll() ([]*menuModel.EntityExtend, error) {
	//从缓存读取
	cache,_ := gcache.Get(model.MENU_TREE_CACHE)
	if cache != nil {
		return cache.([]*menuModel.EntityExtend), nil
	}

	//从数据库中读取
	var result []*menuModel.EntityExtend
	result, err := SelectListAll(nil)

	if err != nil {
		return nil, err
	}
	//存入缓存
	gcache.Set(model.MENU_TREE_CACHE, result, time.Hour)
	return result, nil
}

// 根据用户id获取菜单树
func SelectMenusListByUserId(userId int64) ([]*menuModel.EntityExtend, error) {
	result, err := menuModel.SelectMenusListByUserId(userId)
	if err != nil {
		return nil, err
	}
	//g.Dump(result)
	//存入缓存
	gcache.Set(model.MENU_TREE_CACHE+gconv.String(userId), result, time.Hour)
	return result, nil
}

//获取用户的菜单数据
func SelectMenuNormalByUser(userId int64) ([]*menuModel.EntityExtend, error) {
	if userService.IsAdmin(userId) {
		return SelectMenuNormalAll()
	} else {
		return SelectMenusByUserId(userId)
	}
}

//获取管理员菜单数据
func SelectMenuNormalAll() ([]*menuModel.EntityExtend, error) {
	//从缓存读取
	cache,_ := gcache.Get(model.MENU_CACHE)
	if cache != nil {
		return cache.([]*menuModel.EntityExtend), nil
	}

	//从数据库中读取
	var result []*menuModel.EntityExtend
	result, err := menuModel.SelectMenuNormalAll()

	if err != nil {
		return nil, err
	}

	result = getChildPerms(result, 0)
	//存入缓存
	gcache.Set(model.MENU_CACHE, result, time.Hour)
	return result, nil
}

//根据用户ID读取菜单数据
func SelectMenusByUserId(userId int64, router ...bool) ([]*menuModel.EntityExtend, error) {
	var result []*menuModel.EntityExtend

	//从缓存读取
	cache,_ := gcache.Get(model.MENU_CACHE + gconv.String(userId))

	if cache != nil {
		return cache.([]*menuModel.EntityExtend), nil
	}

	//从数据库中读取
	result, err := menuModel.SelectMenusByUserId(gconv.Int64(userId))
	if err != nil {
		return nil, err
	}
	result = getChildPerms(result, 0)
	//g.Dump(result)
	//存入缓存
	gcache.Set(model.MENU_CACHE+gconv.String(userId), result, time.Hour)
	return result, nil
}

//检查菜单名是否唯一
func CheckMenuNameUniqueAll(menuName string, parentId int64) string {
	menu, err := menuModel.FindOne("menu_name=? and parent_id=?", menuName, parentId)
	if err == nil && menu != nil && menu.MenuId > 0 {
		return "1"
	}
	return "0"
}

//检查菜单名是否唯一
func CheckMenuNameUnique(menuName string, menuId, parentId int64) string {
	menu, err := menuModel.FindOne("menu_name=? and menu_id <> ? and parent_id=?", menuName, menuId, parentId)
	if err == nil && menu != nil && menu.MenuId > 0 {
		return "1"
	}
	return "0"
}

func SelectMenuListByRoleId(roleId int64) (g.Array, error) {
	return menuModel.SelectMenuIds(roleId)
}

// 获取子级
func getChildPerms(menus []*menuModel.EntityExtend, parentId int64) []*menuModel.EntityExtend {
	var result []*menuModel.EntityExtend
	for _, menu := range menus {
		if menu.ParentId == parentId {
			recursionFn(menus, menu)
			result = append(result, menu)
		}
	}
	return result
}

// 递归
func recursionFn(menus []*menuModel.EntityExtend, entity *menuModel.EntityExtend) {
	childMenus := getChildList(menus, entity)
	entity.Children = childMenus
	for _, childMenu := range childMenus {
		if hasChild(menus, childMenu) {
			for _, menu := range childMenus {
				recursionFn(menus, menu)
			}
		}
	}
}

// 获取子级列表
func getChildList(menus []*menuModel.EntityExtend, entity *menuModel.EntityExtend) []*menuModel.EntityExtend {
	var tlist []*menuModel.EntityExtend
	for _, menu := range menus {
		if menu.ParentId == entity.MenuId {
			tlist = append(tlist, menu)
		}
	}
	return tlist
}

// 判断是是否还有子级
func hasChild(menus []*menuModel.EntityExtend, entity *menuModel.EntityExtend) bool {
	return len(getChildList(menus, entity)) > 0
}

// 构建前端路由所需要的菜单
func BuildMenus(menus []*menuModel.EntityExtend) []menuModel.RouterExtend {
	var routerExtends []menuModel.RouterExtend
	for _, menu := range menus {
		var routerExtend menuModel.RouterExtend
		routerExtend.Hidden = menu.Visible == 1
		routerExtend.Name = gstr.UcFirst(menu.Path)
		routerExtend.Path = getRouterPath(menu)
		if gstr.Equal(menu.Component, "") {
			routerExtend.Component = "Layout"
		} else {
			routerExtend.Component = menu.Component
		}
		routerExtend.Meta.Title = menu.MenuName
		routerExtend.Meta.Icon = menu.Icon
		cMenus := menu.Children
		if len(cMenus) > 0 && gstr.Equal(menu.MenuType, "M") {
			routerExtend.AlwaysShow = true
			routerExtend.Redirect = "noRedirect"
			routerExtend.Children = BuildMenus(cMenus)
		}
		routerExtends = append(routerExtends, routerExtend)
	}
	return routerExtends
}

func getRouterPath(menu *menuModel.EntityExtend) string {
	routerPath := menu.Path
	// 非外链并且是一级目录
	if 0 == menu.ParentId && menu.IsFrame == 1 {
		routerPath = "/" + menu.Path
	}
	return routerPath
}

// 构建前端角色所需要的菜单树
func BuildRoleMenus(menus []*menuModel.EntityExtend) []menuModel.RoleMenuTree {
	var roleMenuTrees []menuModel.RoleMenuTree
	for _, menu := range menus {
		var roleMenuTree menuModel.RoleMenuTree
		roleMenuTree.Id = menu.MenuId
		roleMenuTree.Label = menu.MenuName
		rMenus := menu.Children
		if len(rMenus) > 0 {
			roleMenuTree.Children = BuildRoleMenus(rMenus)
		}
		roleMenuTrees = append(roleMenuTrees, roleMenuTree)
	}
	return roleMenuTrees
}

// 获取菜单数据权限
//func SelectMenuPermsByUserId(uid int64) (*garray.StrArray,*garray.StrArray) {
func SelectMenuPermsByUserId(uid int64) *garray.StrArray {
	return menuModel.SelectMenuPermsByUserId(uid)
}

// 清除缓存
func ClearCache() {
	if cacheKeys, err := gcache.KeyStrings(); err == nil {
		// 删除跟menu有关的缓存
		for _,v := range cacheKeys {
			if gstr.ContainsI(v,"menu") {
				gcache.Remove(v)
			}
		}
	}
	gcache.Remove(model.MENU_CACHE)
}
