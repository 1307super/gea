package define

import (
	"gea/app/model"
	"github.com/gogf/gf/frame/g"
)

// ======== api ==========
//分页请求参数
type DeptApiSelectPageReq struct {
	ParentId  int64  `p:"parentId"`      //父部门ID
	DeptName  string `p:"deptName"`      //部门名称
	Status    string `p:"status"`        //状态
	BeginTime string `p:"beginTime"`     //开始时间
	EndTime   string `p:"endTime"`       //结束时间
	PageNum   int    `p:"pageNum"`       //当前页码
	PageSize  int    `p:"pageSize"`      //每页数
	SortName  string `p:"orderByColumn"` //排序字段
	SortOrder string `p:"isAsc"`         //排序方式
}

//新增页面请求参数
type DeptApiCreateReq struct {
	ParentId int64  `p:"parentId"  v:"required#父节点不能为空"`
	DeptName string `p:"deptName"  v:"required#部门名称不能为空"`
	OrderNum int    `p:"orderNum" v:"required#显示排序不能为空"`
	Leader   string `p:"leader"`
	Phone    string `p:"phone"`
	Email    string `p:"email"`
	Status   string `p:"status"`
}

//修改页面请求参数
type DeptApiUpdateReq struct {
	DeptId int64 `p:"deptId" v:"required#主键ID不能为空"`
	DeptApiCreateReq
}

// ========== service ============
// 查询列表返回值
type DeptServiceList struct {
	List  []model.SysDeptExtend `json:"list"`
	Page  int                   `json:"page"`
	Size  int                   `json:"size"`
	Total int                   `json:"total"`
}

type DeptServiceRoleTreeData struct {
	Depts       []model.DeptExtend `json:"depts"`
	CheckedKeys g.Array            `json:"checkedKeys"`
}
