package service

import (
	"context"
	"fmt"
	"gea/app/dao"
	"gea/app/model"
	"gea/app/shared"
	"gea/app/system/admin/internal/define"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"time"
)

const AdminAuthMenu = "AdminAuthMenu"

// 登陆用户的菜单列表缓存前缀
const MENU_CACHE = "menu_cache"

const MENU_TREE_CACHE = "menu_tree_cache"

const (
	YES_FRAME   = 0            // 是否菜单外链 是
	NO_FRAME    = 1            // 是否菜单外链（否）
	TYPE_DIR    = "M"          // 菜单类型（目录）
	TYPE_MENU   = "C"          //  菜单类型（菜单）
	TYPE_BUTTON = "F"          // 菜单类型（按钮）
	LAYOUT      = "Layout"     // Layout组件标识
	PARENT_VIEW = "ParentView" // ParentView组件标识
)

var Menu = &menuService{}

type menuService struct {
}

func (s *menuService) GetAll(param *define.MenuApiSelectPageReq) ([]*model.SysMenuExtend, error) {
	m := dao.SysMenu.As("m")
	if param != nil {
		if param.MenuName != "" {
			m = m.Where("m.menu_name like ?", "%"+param.MenuName+"%")
		}

		if param.Status != "" {
			m = m.Where("m.status = ", param.Status)
		}

		if param.BeginTime != "" {
			m = m.Where("date_format(m.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			m = m.Where("date_format(m.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}
	m = m.Order("m.order_num asc")
	var result []*model.SysMenuExtend
	err := m.Structs(&result)
	return result, err
}

//添加数据
func (s *menuService) Create(ctx context.Context, req *define.MenuApiCreateReq) (int64, error) {
	user := shared.Context.Get(ctx).User

	if s.CheckMenuNameUniqueAll(req.MenuName, req.ParentId) == "1" {
		return 0, gerror.New("菜单名称已存在")
	}

	var menu model.SysMenu
	menu.CreateTime = gtime.Now()
	menu.CreateBy = user.UserExtend.LoginName

	var editReq *define.MenuApiEditReq
	gconv.Struct(req, &editReq)
	return s.save(&menu, editReq)
}

//修改数据
func (s *menuService) Update(ctx context.Context, req *define.MenuApiEditReq) (int64, error) {

	user := shared.Context.Get(ctx).User
	if s.CheckMenuNameUnique(req.MenuName, req.MenuId, req.ParentId) == "1" {
		return 0, gerror.New("菜单名称已存在")
	}
	menu, err := dao.SysMenu.FindOne(dao.SysMenu.Columns.MenuId, req.MenuId)
	if err != nil {
		return 0, err
	}
	if menu == nil {
		return 0, gerror.New("菜单不存在")
	}
	menu.UpdateBy = user.UserExtend.LoginName
	menu.UpdateTime = gtime.Now()
	return s.save(menu, req)
}

func (s *menuService) save(menu *model.SysMenu, req *define.MenuApiEditReq) (int64, error) {
	if req.MenuType == TYPE_BUTTON {
		req.Component = ""
		req.Path = ""
		req.Icon = ""
	}
	if req.MenuType == TYPE_DIR {
		req.Component = ""
		req.Perms = ""
	}
	menu.MenuName = req.MenuName
	menu.Visible = gconv.Uint(req.Visible)
	menu.ParentId = req.ParentId
	menu.Remark = ""
	menu.MenuType = req.MenuType
	menu.Path = req.Path
	menu.Perms = req.Perms
	menu.Method = req.Method
	menu.IsFrame = gconv.Uint(req.IsFrame)
	menu.Icon = req.Icon
	menu.OrderNum = req.OrderNum
	menu.Component = req.Component
	menu.Status = gconv.Uint(req.Status)

	var err error
	tx, err := g.DB().Begin()
	if err != nil {
		return 0, err
	}
	result, err := dao.SysMenu.TX(tx).Data(menu).Save()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if menu.MenuId == 0 {
		// 新增
		id, err := result.LastInsertId()
		if err != nil || id <= 0 {
			tx.Rollback()
			return 0, err
		}
	}
	s.ClearCache()
	return 1, tx.Commit()
}

func (s *menuService) Info(id int64) (*model.SysMenuExtend, error) {
	var result *model.SysMenuExtend
	m := dao.SysMenu.As("t").Fields("t.menu_id, t.parent_id, t.menu_name, t.order_num, t.path,t.method,t.component, t.is_frame, t.menu_type, t.visible,t.status, t.perms, t.icon, t.remark,(SELECT menu_name FROM sys_menu WHERE menu_id = t.parent_id) parent_name")
	m = m.Where(dao.SysMenu.Columns.MenuId, id)
	if err := m.Struct(&result); err != nil {
		return nil, gerror.New("获取数据失败")
	}
	return result, nil
}

func (s *menuService) Delete(id int64) bool {
	result, err := dao.SysMenu.Delete(dao.SysMenu.Columns.MenuId, id)
	if err == nil {
		affected, _ := result.RowsAffected()
		if affected > 0 {
			return true
		}
	}
	s.ClearCache()
	return false
}

//加载所有菜单列表树
func (s *menuService) MenuTreeData(ctx context.Context) ([]model.RoleMenuTree, error) {
	userId := shared.Context.Get(ctx).User.UserExtend.UserId
	var menuList []*model.SysMenuExtend
	var err error
	if IsAdmin(userId) {
		menuList, err = s.SelectMenuAll()
	} else {
		menuList, err = s.GetMenusListByUserId(userId)
	}
	if err != nil {
		return nil, err
	}
	menuList = s.getChildPerms(menuList, 0)
	return s.BuildRoleMenus(menuList), nil
}

//获取管理员菜单树数据
func (s *menuService) SelectMenuAll() ([]*model.SysMenuExtend, error) {
	//从缓存读取
	cache, _ := gcache.Get(MENU_TREE_CACHE)
	if cache != nil {
		return cache.([]*model.SysMenuExtend), nil
	}
	//从数据库中读取
	var result []*model.SysMenuExtend
	result, err := s.GetAll(nil)

	if err != nil {
		return nil, err
	}
	//存入缓存
	gcache.Set(MENU_TREE_CACHE, result, time.Hour)
	return result, nil
}

// 根据用户id获取菜单
func (s *menuService) GetMenusListByUserId(userId int64) ([]*model.SysMenuExtend, error) {
	var result []*model.SysMenuExtend
	m := dao.SysMenu.As("m")
	m = m.LeftJoin("sys_role_menu rm", "m.menu_id = rm.menu_id")
	m = m.LeftJoin("sys_user_role ur", "rm.role_id = ur.role_id")
	m = m.LeftJoin("sys_role ro", "ur.role_id = ro.role_id")
	m = m.Fields("distinct m.menu_id, m.parent_id, m.menu_name, m.path, m.component, m.visible, m.status, ifnull(m.perms,'') as perms, m.is_frame, m.menu_type, m.icon, m.order_num, m.create_time")
	m = m.Where("ur.user_id = ? ", userId).Where(" m.menu_type in ('M', 'C')")
	m = m.Order("m.parent_id, m.order_num")
	if err := m.Structs(&result); err != nil {
		return nil, gerror.New("读取数据失败")
	}
	return result, nil
}

//根据角色ID查询菜单id
func (s *menuService) GetMenuIdsByRoleId(roleId int64) (g.Array, error) {
	var result g.Array
	roleMenus, err := dao.SysRoleMenu.FindAll(g.Map{
		dao.SysRoleMenu.Columns.RoleId: roleId,
	})
	if err != nil {
		return nil, gerror.New("读取数据失败")
	}
	for _, roleMenu := range roleMenus {
		result = append(result, roleMenu.MenuId)
	}
	return result, nil
}

// 获取用户的菜单数据
func (s *menuService) SelectMenuNormalByUser(ctx context.Context) ([]*model.SysMenuExtend, error) {
	userId := shared.Context.Get(ctx).Uid
	if IsAdmin(userId) {
		return s.GettMenuNormalAll()
	} else {
		return s.GetMenusByUserId(userId)
	}
}

// 获取管理员菜单数据
func (s *menuService) GettMenuNormalAll() ([]*model.SysMenuExtend, error) {
	//从缓存读取
	cache, _ := gcache.Get(MENU_CACHE)
	if cache != nil {
		return cache.([]*model.SysMenuExtend), nil
	}

	var result []*model.SysMenuExtend
	m := dao.SysMenu.As("m").Fields("distinct m.menu_id, m.parent_id, m.menu_name, m.path, m.component, m.visible, m.status, ifnull(m.perms,'') as perms, m.is_frame, m.menu_type, m.icon, m.order_num, m.create_time")
	m = m.Where(" m.status = '0'")
	m = m.Where(" m.menu_type in ('M', 'C')")
	m = m.Order("m.parent_id, m.order_num")
	if err := m.Structs(&result); err != nil {
		return nil, gerror.New("读取数据失败")
	} else {
		result = s.getChildPerms(result, 0)
		//存入缓存
		gcache.Set(MENU_CACHE, result, time.Hour)
		return result, nil
	}
}

// 根据用户ID读取菜单数据
func (s *menuService) GetMenusByUserId(userId int64, router ...bool) ([]*model.SysMenuExtend, error) {
	var result []*model.SysMenuExtend

	//从缓存读取
	cache, _ := gcache.Get(MENU_CACHE + gconv.String(userId))

	if cache != nil {
		return cache.([]*model.SysMenuExtend), nil
	}
	//从数据库中读取
	result, err := s.GetMenusListByUserId(gconv.Int64(userId))
	if err != nil {
		return nil, err
	}
	result = s.getChildPerms(result, 0)
	//存入缓存
	gcache.Set(MENU_CACHE+gconv.String(userId), result, time.Hour)
	return result, nil
}

// 获取子级
func (s *menuService) getChildPerms(menus []*model.SysMenuExtend, parentId int64) []*model.SysMenuExtend {
	var result []*model.SysMenuExtend
	for _, menu := range menus {
		if menu.ParentId == parentId {
			s.recursionFn(menus, menu)
			result = append(result, menu)
		}
	}
	return result
}

// 递归
func (s *menuService) recursionFn(menus []*model.SysMenuExtend, entity *model.SysMenuExtend) {
	childMenus := s.getChildList(menus, entity)
	entity.Children = childMenus
	for _, childMenu := range childMenus {
		if s.hasChild(menus, childMenu) {
			for _, menu := range childMenus {
				s.recursionFn(menus, menu)
			}
		}
	}
}

// 获取子级列表
func (s *menuService) getChildList(menus []*model.SysMenuExtend, entity *model.SysMenuExtend) []*model.SysMenuExtend {
	var tlist []*model.SysMenuExtend
	for _, menu := range menus {
		if menu.ParentId == entity.MenuId {
			tlist = append(tlist, menu)
		}
	}
	return tlist
}

// 判断是是否还有子级
func (s *menuService) hasChild(menus []*model.SysMenuExtend, entity *model.SysMenuExtend) bool {
	return len(s.getChildList(menus, entity)) > 0
}

// 构建前端路由所需要的菜单
func (s *menuService) BuildMenus(menus []*model.SysMenuExtend) []model.RouterExtend {
	var routerExtends []model.RouterExtend
	for _, menu := range menus {
		var routerExtend model.RouterExtend
		routerExtend.Hidden = menu.Visible == 1
		routerExtend.Name = s.getRouteName(menu)
		routerExtend.Path = s.getRouterPath(menu)
		routerExtend.Component = s.getComponent(menu)
		routerExtend.Meta.Title = menu.MenuName
		routerExtend.Meta.Icon = menu.Icon
		cMenus := menu.Children
		if len(cMenus) > 0 && gstr.Equal(menu.MenuType, TYPE_DIR) {
			routerExtend.AlwaysShow = true
			routerExtend.Redirect = "noRedirect"
			routerExtend.Children = s.BuildMenus(cMenus)
		} else if s.isMeunFrame(menu) {
			var childrenExtends []model.RouterExtend
			var childrenExtend model.RouterExtend
			childrenExtend.Path = menu.Path
			childrenExtend.Component = menu.Component
			childrenExtend.Name = gstr.UcFirst(menu.Path)
			childrenExtend.Meta.Title = menu.MenuName
			childrenExtend.Meta.Icon = menu.Icon
			childrenExtends = append(childrenExtends, childrenExtend)
			routerExtend.Children = childrenExtends
		}
		routerExtends = append(routerExtends, routerExtend)
	}
	return routerExtends
}

// 获取组件信息
func (s *menuService) getComponent(menu *model.SysMenuExtend) string {
	component := LAYOUT
	if menu.Component != "" && !s.isMeunFrame(menu) {
		component = menu.Component
	} else if menu.Component == "" && s.isParentView(menu) {
		component = PARENT_VIEW
	} else if menu.IsFrame == YES_FRAME {
		component = ""
	}
	return component
}

// 是否为菜单内部跳转
func (s *menuService) isMeunFrame(menu *model.SysMenuExtend) bool {
	return menu.ParentId == 0 && TYPE_MENU == menu.MenuType && menu.IsFrame == NO_FRAME
}

// 是否为parent_view组件
func (s *menuService) isParentView(menu *model.SysMenuExtend) bool {
	return menu.ParentId != 0 && menu.MenuType == TYPE_DIR
}

// 获取路由名称
func (s *menuService) getRouteName(menu *model.SysMenuExtend) string {
	routerName := gstr.UcFirst(menu.Path)
	// 非外链并且是一级目录
	if s.isMeunFrame(menu) {
		routerName = ""
	}
	return routerName
}

// 获取路由地址
func (s *menuService) getRouterPath(menu *model.SysMenuExtend) string {
	routerPath := menu.Path
	// 非外链并且是一级目录
	if 0 == menu.ParentId && menu.MenuType == TYPE_DIR && menu.IsFrame == NO_FRAME {
		routerPath = "/" + menu.Path
	} else if s.isMeunFrame(menu) {
		routerPath = "/"
	}
	return routerPath
}

// 构建前端角色所需要的菜单树
func (s *menuService) BuildRoleMenus(menus []*model.SysMenuExtend) []model.RoleMenuTree {
	var roleMenuTrees []model.RoleMenuTree
	for _, menu := range menus {
		var roleMenuTree model.RoleMenuTree
		roleMenuTree.Id = menu.MenuId
		roleMenuTree.Label = menu.MenuName
		rMenus := menu.Children
		if len(rMenus) > 0 {
			roleMenuTree.Children = s.BuildRoleMenus(rMenus)
		}
		roleMenuTrees = append(roleMenuTrees, roleMenuTree)
	}
	return roleMenuTrees
}

func (s *menuService) GetMenuPermission(ctx context.Context) *garray.StrArray {
	customCtx := shared.Context.Get(ctx)
	perms := garray.NewStrArray()
	if IsAdmin(customCtx.Uid) {
		perms.Append("*:*:*")
	} else {
		// 缓存
		//从缓存读取
		cache, _ := gcache.Get(AdminAuthMenu + gconv.String(customCtx.Uid))
		if cache != nil {
			return cache.(*garray.StrArray)
		}
		var Entity []*model.SysMenu
		if err := dao.SysMenu.As("m").LeftJoin("sys_role_menu rm", "m.menu_id = rm.menu_id").LeftJoin("sys_user_role ur", "rm.role_id = ur.role_id").LeftJoin("sys_role ro", "ur.role_id = ro.role_id").Fields("distinct m.perms,m.method").Where("m.status = '0' and ro.status = '0' and ur.user_id = ?", customCtx.Uid).Structs(&Entity); err == nil {
			if len(Entity) > 0 {
				for _, entity := range Entity {
					if entity.Perms != "" {
						perms.Append(entity.Method + ":" + entity.Perms)
					}
				}
				// 设置缓存
				gcache.Set(AdminAuthMenu+gconv.String(customCtx.Uid), perms, 0)
			}
		}
	}
	return perms
}

//检查菜单名是否唯一
func (s *menuService) CheckMenuNameUniqueAll(menuName string, parentId int64) string {
	menu, err := dao.SysMenu.FindOne(g.Map{
		dao.SysMenu.Columns.MenuName: menuName,
		dao.SysMenu.Columns.ParentId: parentId,
	})
	if err == nil && menu != nil && menu.MenuId > 0 {
		return "1"
	}
	return "0"
}

//检查菜单名是否唯一
func (s *menuService) CheckMenuNameUnique(menuName string, menuId, parentId int64) string {
	menu, err := dao.SysMenu.FindOne(g.Map{
		dao.SysMenu.Columns.MenuName:                     menuName,
		dao.SysMenu.Columns.ParentId:                     parentId,
		fmt.Sprintf("%s <>", dao.SysMenu.Columns.MenuId): menuId,
	})
	if err == nil && menu != nil && menu.MenuId > 0 {
		return "1"
	}
	return "0"
}

// 清除缓存
func (s *menuService) ClearCache() {
	if cacheKeys, err := gcache.KeyStrings(); err == nil {
		// 删除跟menu有关的缓存
		for _, v := range cacheKeys {
			if gstr.ContainsI(v, "menu") {
				gcache.Remove(v)
			}
		}
	}
	gcache.Remove(MENU_CACHE)
}
