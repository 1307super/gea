package role

import (
	"gea/app/utils/page"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table sys_role.
type EntityFlag struct {
	RoleId     int64       `orm:"role_id,primary" json:"role_id"`     // 角色ID
	RoleName   string      `orm:"role_name"       json:"role_name"`   // 角色名称
	RoleKey    string      `orm:"role_key"        json:"role_key"`    // 角色权限字符串
	RoleSort   int         `orm:"role_sort"       json:"role_sort"`   // 显示顺序
	DataScope  string      `orm:"data_scope"      json:"dataScope"`  // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status     string      `orm:"status"          json:"status"`      // 角色状态（0正常 1停用）
	DelFlag    string      `orm:"del_flag"        json:"del_flag"`    // 删除标志（0代表存在 2代表删除）
	CreateBy   string      `orm:"create_by"       json:"create_by"`   // 创建者
	CreateTime *gtime.Time `orm:"create_time"     json:"create_time"` // 创建时间
	UpdateBy   string      `orm:"update_by"       json:"update_by"`   // 更新者
	UpdateTime *gtime.Time `orm:"update_time"     json:"update_time"` // 更新时间
	Remark     string      `orm:"remark"          json:"remark"`      // 备注
	Flag       bool        `json:"flag"`                              // 标记
}

//数据权限保存请求参数
type DataScopeReq struct {
	RoleId    int64  `p:"roleId"  v:"required#角色ID不能为空"`
	RoleName  string `p:"roleName"  v:"required#角色名称不能为空"`
	RoleKey   string `p:"roleKey"  v:"required#权限字符串不能为空"`
	DataScope string `p:"dataScope"  v:"required#数据范围不能为空"`
	DeptIds   string `p:"deptIds"`
}

//检查角色名称请求参数
type CheckRoleNameReq struct {
	RoleId   int64  `p:"roleId"  v:"required#角色ID不能为空"`
	RoleName string `p:"roleName"  v:"required#角色名称不能为空"`
}

//检查权限字符请求参数
type CheckRoleKeyReq struct {
	RoleId  int64  `p:"roleId"  v:"required#角色ID不能为空"`
	RoleKey string `p:"roleKey"  v:"required#权限字符不能为空"`
}

//检查角色名称请求参数
type CheckRoleNameALLReq struct {
	RoleName string `p:"roleName"  v:"required#角色名称不能为空"`
}

//检查权限字符请求参数
type CheckRoleKeyALLReq struct {
	RoleKey string `p:"roleKey"  v:"required#权限字符不能为空"`
}

//分页请求参数
type SelectPageReq struct {
	RoleName      string `p:"roleName"`      //角色名称
	Status        string `p:"status"`        //状态
	RoleKey       string `p:"roleKey"`       //角色键
	DataScope     string `p:"dataScope"`     //数据范围
	BeginTime     string `p:"beginTime"`     //开始时间
	EndTime       string `p:"endTime"`       //结束时间
	PageNum       int    `p:"pageNum"`       //当前页码
	PageSize      int    `p:"pageSize"`      //每页数
	OrderByColumn string `p:"orderByColumn"` //排序字段
	IsAsc         string `p:"isAsc"`         //排序方式
}

//新增页面请求参数
type AddReq struct {
	RoleName string `p:"roleName"  v:"required#角色名称不能为空"`
	RoleKey  string `p:"roleKey"  v:"required#权限字符不能为空"`
	RoleSort string `p:"roleSort"  v:"required#显示顺序不能为空"`
	Status   string `p:"status"`
	Remark   string `p:"remark"`
	MenuIds  string `p:"menuIds" v:"required#菜单不能为空"`
}

//修改页面请求参数
type EditReq struct {
	RoleId   int64  `p:"roleId" v:"required#主键ID不能为空"`
	RoleName string `p:"roleName"  v:"required#角色名称不能为空"`
	RoleKey  string `p:"roleKey"  v:"required#权限字符不能为空"`
	RoleSort string `p:"roleSort"  v:"required#显示顺序不能为空"`
	Status   string `p:"status"`
	Remark   string `p:"remark"`
	MenuIds  string `p:"menuIds" v:"required#菜单不能为空"`
}

type ChangeStatus struct {
	RoleId int64 `p:"roleId" v:"required#主键ID不能为空"`
	Status string `p:"status" v:"required#状态不能为空"`
}

// 角色与资源
type PermissionForRole struct {
	RoleName string `orm:"roleName"` // 菜单ID
	Path     string `orm:"path"`     // 菜单ID
	Method   string `orm:"method"`   // 菜单ID
}

//根据条件分页查询角色数据
func SelectListPage(param *SelectPageReq) ([]Entity, *page.Paging, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_role r").Where("r.del_flag = '0'")

	if param.RoleName != "" {
		model.Where("r.role_name like ?", "%"+param.RoleName+"%")
	}

	if param.Status != "" {
		model.Where("r.status = ", param.Status)
	}

	if param.RoleKey != "" {
		model.Where("r.role_key like ?", "%"+param.RoleKey+"%")
	}

	if param.DataScope != "" {
		model.Where("r.data_scope = ", param.DataScope)
	}

	if param.BeginTime != "" {
		model.Where("date_format(r.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
	}

	if param.EndTime != "" {
		model.Where("date_format(r.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
	}

	total, err := model.Count()

	if err != nil {
		return nil, nil, gerror.New("读取行数失败")
	}

	page := page.CreatePaging(param.PageNum, param.PageSize, total)

	model.Limit(page.StartNum, page.Pagesize)
	if param.OrderByColumn != "" {
		model.Order(param.OrderByColumn + " " + param.IsAsc)
	}
	var result []Entity

	err = model.Structs(&result)
	return result, page, err
}

// 导出excel
func SelectListExport(param *SelectPageReq) (gdb.Result, error) {
	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_role r").Where("r.del_flag = '0'")

	if param != nil {
		if param.RoleName != "" {
			model.Where("r.role_name like ?", "%"+param.RoleName+"%")
		}

		if param.Status != "" {
			model.Where("r.status = ", param.Status)
		}

		if param.RoleKey != "" {
			model.Where("r.role_key like ?", "%"+param.RoleKey+"%")
		}

		if param.DataScope != "" {
			model.Where("r.data_scope = ", param.DataScope)
		}

		if param.BeginTime != "" {
			model.Where("date_format(r.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(r.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	//角色序号	角色名称	角色权限	角色排序	数据范围	角色状态
	model.Fields("r.role_id,r.role_name,r.role_key,r.role_sort,r.data_scope,r.status")

	result, err := model.All()
	return result, err
}

//获取所有角色数据
func SelectListAll(param *SelectPageReq) ([]EntityFlag, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_role r").Where("r.del_flag = '0'")
	if param != nil {
		if param.RoleName != "" {
			model.Where("r.role_name like ?", "%"+param.RoleName+"%")
		}

		if param.Status != "" {
			model.Where("r.status = ", param.Status)
		}

		if param.RoleKey != "" {
			model.Where("r.role_key like ?", "%"+param.RoleKey+"%")
		}

		if param.DataScope != "" {
			model.Where("r.data_scope = ", param.DataScope)
		}

		if param.BeginTime != "" {
			model.Where("date_format(r.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(r.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []EntityFlag

	err = model.Structs(&result)
	return result, err
}

func SelectRoleContactVo() ([]EntityFlag, error) {
	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}
	model := db.Table("sys_role r")
	model.Fields("distinct r.role_id, r.role_name, r.role_key, r.role_sort, r.data_scope,r.status, r.del_flag, r.create_time, r.remark")
	//model.LeftJoin("sys_user_role ur", "ur.role_id = r.role_id")
	//model.LeftJoin("sys_user u", "u.user_id = ur.user_id")
	//model.LeftJoin("sys_dept d", "u.dept_id = d.dept_id")
	model.Where("r.del_flag = '0'")
	//model.Where("ur.user_id = ?", userId)
	var result []EntityFlag
	err = model.Structs(&result)
	return result, err
}
func SelectRoleListByUserId (uid int64) ([]EntityFlag, error) {
	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}
	model := db.Table("sys_role r")
	model.LeftJoin("sys_user_role ur", "ur.role_id = r.role_id")
	model.LeftJoin("sys_user u", "u.user_id = ur.user_id")
	model.Where("r.del_flag = '0' and u.del_flag = '0'")
	model.Where("ur.user_id = ?", uid)
	//var result g.Array
	var result []EntityFlag
	err = model.Structs(&result)
	return result, err
}

func GetRoleMenuPolicy(roleName ...string) []PermissionForRole {
	var result []PermissionForRole
	db, err := gdb.Instance()
	if err != nil {
		return result
	}
	model := db.Table("sys_menu m")
	model.LeftJoin("sys_role_menu rm", "m.menu_id = rm.menu_id")
	model.LeftJoin("sys_role ro", "rm.role_id = ro.role_id")
	model.Fields("distinct ro.role_key as roleName,m.perms as path,m.method")
	model.Where("m.status = '0' and ro.status = '0' and m.menu_type != 'M'")
	if len(roleName) > 0 && roleName[0] != "" {
		model.Where("ro.role_key = ?", roleName[0])
	}

	_ = model.Structs(&result)
	return result
}