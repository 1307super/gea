package define

import "gea/app/model"

// ========= api ========
//分页请求参数
type RoleApiSelectPageReq struct {
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
type RoleApiCreateReq struct {
	RoleName string `p:"roleName"  v:"required#角色名称不能为空"`
	RoleKey  string `p:"roleKey"  v:"required#权限字符不能为空"`
	RoleSort string `p:"roleSort"  v:"required#显示顺序不能为空"`
	Status   string `p:"status"`
	Remark   string `p:"remark"`
	MenuIds  string `p:"menuIds" v:"required#菜单不能为空"`
}

//修改页面请求参数
type RoleApiEditReq struct {
	RoleId   int64  `p:"roleId" v:"required#主键ID不能为空"`
	RoleApiCreateReq
}

//数据权限保存请求参数
type RoleApiDataScopeReq struct {
	RoleId    int64  `p:"roleId"  v:"required#角色ID不能为空"`
	RoleName  string `p:"roleName"  v:"required#角色名称不能为空"`
	RoleKey   string `p:"roleKey"  v:"required#权限字符串不能为空"`
	DataScope string `p:"dataScope"  v:"required#数据范围不能为空"`
	DeptIds   string `p:"deptIds"`
}

type RoleApiChangeStatus struct {
	RoleId int64 `p:"roleId" v:"required#主键ID不能为空"`
	Status string `p:"status" v:"required#状态不能为空"`
}

// API执行删除内容
type RoleApiDeleteReq struct {
	Ids string `p:"id"  v:"required#请选择要删除的数据记录"`
}


// ========= service ========
// 查询列表返回值
type RoleServiceList struct {
	List  []model.SysRole `json:"list"`
	Page  int             `json:"page"`
	Size  int             `json:"size"`
	Total int             `json:"total"`
}

// 角色与资源
type RoleServicePermissionForRole struct {
	RoleName string `orm:"roleName"` // 菜单ID
	Path     string `orm:"path"`     // 菜单ID
	Method   string `orm:"method"`   // 菜单ID
}

