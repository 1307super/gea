package menu

import (
	"gea/app/model/system/role_menu"
	"gea/app/utils/page"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/util/gconv"
)

const AdminAuthMenu = "AdminAuthMenu"

// Entity is the golang structure for table sys_menu.
type EntityExtend struct {
	Entity
	ParentName string          `json:"parentName"`         // 父菜单名称
	Children   []*EntityExtend `json:"children,omitempty"` // 子菜单
}

// 路由
type RouterExtend struct {
	AlwaysShow bool           `json:"always_show,omitempty"` // 总是显示
	Children   []RouterExtend `json:"children,omitempty"`    // 子菜单
	Component  string         `json:"component"`             // 组件地址
	Hidden     bool           `json:"hidden"`                // 是否隐藏
	Meta       Meta           `json:"meta"`                  // meta
	Name       string         `json:"name"`                  // 名称
	Path       string         `json:"path"`                  // 地址
	Redirect   string         `json:"redirect,omitempty"`    // 跳转链接
	//Entity
	//ParentName string `json:"parentName"` // 父菜单名称
}
type Meta struct {
	Title string `json:"title"` // 标题
	Icon  string `json:"icon"`  // 图标
}

// 角色权限树
type RoleMenuTree struct {
	Children []RoleMenuTree `json:"children"`
	Id       int64          `json:"id"`
	Label    string         `json:"label"`
}

//检查菜单名称请求参数
type CheckMenuNameReq struct {
	MenuId   int64  `p:"menuId"  v:"required#菜单ID不能为空"`
	ParentId int64  `p:"parentId"  v:"required#父菜单ID不能为空"`
	MenuName string `p:"menuName"  v:"required#菜单名称不能为空"`
}

//检查菜单名称请求参数
type CheckMenuNameALLReq struct {
	ParentId int64  `p:"parentId"  v:"required#父菜单ID不能为空"`
	MenuName string `p:"menuName"  v:"required#菜单名称不能为空"`
}

//分页请求参数
type SelectPageReq struct {
	MenuName  string `p:"menuName"`      //菜单名称
	Status    string `p:"status"`        //状态
	BeginTime string `p:"beginTime"`     //开始时间
	EndTime   string `p:"endTime"`       //结束时间
	PageNum   int    `p:"pageNum"`       //当前页码
	PageSize  int    `p:"pageSize"`      //每页数
	SortName  string `p:"orderByColumn"` //排序字段
	SortOrder string `p:"isAsc"`         //排序方式
}

//新增页面请求参数
type AddReq struct {
	ParentId  int64  `p:"parentId"  v:"required#父节点不能为空"`
	MenuType  string `p:"menuType"  v:"required#菜单类型不能为空"`
	MenuName  string `p:"menuName"  v:"required#菜单名称不能为空"`
	OrderNum  int    `p:"orderNum" v:"required#显示排序不能为空"`
	Path      string `p:"path"`
	Icon      string `p:"icon"`
	IsFrame   string `p:"is_frame"`
	Perms     string `p:"perms"`
	Visible   int    `p:"visible"`
	Status    string `p:"status"`
	Component string `p:"component"`
	Method    string `p:"method"`
}

//修改页面请求参数
type EditReq struct {
	MenuId    int64  `p:"menuId" v:"required#主键ID不能为空"`
	ParentId  int64  `p:"parentId"  v:"required#父节点不能为空"`
	MenuType  string `p:"menuType"  v:"required#菜单类型不能为空"`
	MenuName  string `p:"menuName"  v:"required#菜单名称不能为空"`
	OrderNum  int    `p:"orderNum" v:"required#显示排序不能为空"`
	Path      string `p:"path"`
	Icon      string `p:"icon"`
	IsFrame   string `p:"is_frame"`
	Perms     string `p:"perms"`
	Visible   int    `p:"visible"`
	Status    string `p:"status"`
	Component string `p:"component"`
	Method    string `p:"method"`
}

//根据主键查询数据
func SelectRecordById(id int64) (*EntityExtend, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}
	var result EntityExtend
	model := db.Table("sys_menu t")
	model.Fields("t.menu_id, t.parent_id, t.menu_name, t.order_num, t.path,t.method,t.component, t.is_frame, t.menu_type, t.visible,t.status, t.perms, t.icon, t.remark,(SELECT menu_name FROM sys_menu WHERE menu_id = t.parent_id) parent_name")
	model.Where("menu_id", id)
	err = model.Struct(&result)
	if err != nil {
		return nil, gerror.New("获取数据失败")
	}
	return &result, nil
}

//根据条件分页查询数据
func SelectListPage(param *SelectPageReq) ([]Entity, *page.Paging, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_menu m")

	if param != nil {
		if param.MenuName != "" {
			model.Where("m.menu_name like ?", "%"+param.MenuName+"%")
		}

		if param.Status != "" {
			model.Where("m.status = ", param.Status)
		}

		if param.BeginTime != "" {
			model.Where("date_format(m.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(m.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	total, err := model.Count()

	if err != nil {
		return nil, nil, gerror.New("读取行数失败")
	}

	page := page.CreatePaging(param.PageNum, param.PageSize, total)

	model.Limit(page.StartNum, page.Pagesize)

	var result []Entity

	err = model.Structs(&result)

	if err != nil {
		return nil, nil, gerror.New("读取数据失败")
	} else {
		return result, page, nil
	}
}

//获取所有数据
func SelectListAll(param *SelectPageReq) ([]*EntityExtend, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_menu m")
	if param != nil {
		if param.MenuName != "" {
			model.Where("m.menu_name like ?", "%"+param.MenuName+"%")
		}

		if param.Status != "" {
			model.Where("m.status = ", param.Status)
		}

		if param.BeginTime != "" {
			model.Where("date_format(m.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(m.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []*EntityExtend

	err = model.Structs(&result)
	return result, err
}

//根据用户ID读取菜单数据
func SelectMenusListByUserId(userId int64) ([]*EntityExtend, error) {
	var result []*EntityExtend

	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_menu m")
	model.LeftJoin("sys_role_menu rm", "m.menu_id = rm.menu_id")
	model.LeftJoin("sys_user_role ur", "rm.role_id = ur.role_id")
	model.LeftJoin("sys_role ro", "ur.role_id = ro.role_id")
	model.Fields("distinct m.menu_id, m.parent_id, m.menu_name, m.path, m.component, m.visible, m.status, ifnull(m.perms,'') as perms, m.is_frame, m.menu_type, m.icon, m.order_num, m.create_time")
	model.Where("ur.user_id = ? ", userId)
	model.Order("m.parent_id, m.order_num")
	model.Structs(&result)

	if err != nil {
		return nil, gerror.New("读取数据失败")
	} else {
		return result, nil
	}
}

// 获取管理员菜单数据
func SelectMenuNormalAll() ([]*EntityExtend, error) {
	var result []*EntityExtend

	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_menu m")
	model.Fields("distinct m.menu_id, m.parent_id, m.menu_name, m.path, m.component, m.visible, m.status, ifnull(m.perms,'') as perms, m.is_frame, m.menu_type, m.icon, m.order_num, m.create_time")
	model.Where(" m.status = '0'")
	model.Where(" m.menu_type in ('M', 'C')")
	model.Order("m.parent_id, m.order_num")
	model.Structs(&result)

	if err != nil {
		return nil, gerror.New("读取数据失败")
	} else {
		return result, nil
	}
}

//根据用户ID读取菜单数据
func SelectMenusByUserId(userId int64) ([]*EntityExtend, error) {
	var result []*EntityExtend

	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_menu m")
	model.LeftJoin("sys_role_menu rm", "m.menu_id = rm.menu_id")
	model.LeftJoin("sys_user_role ur", "rm.role_id = ur.role_id")
	model.LeftJoin("sys_role ro", "ur.role_id = ro.role_id")
	model.LeftJoin("sys_user u", "ur.user_id = u.user_id")
	model.Fields("distinct m.menu_id, m.parent_id, m.menu_name, m.path,m.component, m.visible, m.status, ifnull(m.perms,'') as perms, m.is_frame, m.menu_type, m.icon, m.order_num, m.create_time")
	model.Where("u.user_id = ? and m.menu_type in ('M', 'C') and m.status = '0' and ro.status = '0'", userId)
	model.Order("m.parent_id, m.order_num")
	model.Structs(&result)

	if err != nil {
		return nil, gerror.New("读取数据失败")
	} else {
		return result, nil
	}
}

//根据角色ID查询菜单id
func SelectMenuIds(roleId int64) (g.Array, error) {
	var result g.Array
	roleMenus, err := role_menu.FindAll(g.Map{
		"role_id": roleId,
	})
	if err != nil {
		return nil, gerror.New("读取数据失败")
	}

	for _, roleMenu := range roleMenus {
		result = append(result, roleMenu.MenuId)
	}
	return result, nil
}

//func SelectMenuPermsByUserId(uid int64) (*garray.StrArray,*garray.StrArray) {
func SelectMenuPermsByUserId(uid int64) *garray.StrArray {
	// 缓存
	//从缓存读取
	cache := gcache.Get(AdminAuthMenu + gconv.String(uid))
	if cache != nil {
		return cache.(*garray.StrArray)
	}
	var result []*Entity
	perms := garray.NewStrArray()

	db, err := gdb.Instance()
	if err != nil {
		return perms
	}

	model := db.Table("sys_menu m")
	model.LeftJoin("sys_role_menu rm", "m.menu_id = rm.menu_id")
	model.LeftJoin("sys_user_role ur", "rm.role_id = ur.role_id")
	model.LeftJoin("sys_role ro", "ur.role_id = ro.role_id")
	model.Fields("distinct m.perms,m.method")
	model.Where("m.status = '0' and ro.status = '0' and ur.user_id = ?", uid)
	model.Structs(&result)

	if len(result) <= 0 {
		return perms
	}
	for _, entity := range result {
		if entity.Perms != "" {
			perms.Append(entity.Method + ":" + entity.Perms)
		}
	}
	// 设置缓存
	gcache.Set(AdminAuthMenu+gconv.String(uid), perms, 0)
	return perms
}


