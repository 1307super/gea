package define

import "gea/app/model"

// ======= api =====
//分页请求参数
type PostApiPageReq struct {
	PostCode      string `p:"postCode"`      //岗位编码
	Status        string `p:"status"`        //状态
	PostName      string `p:"postName"`      //岗位名称
	BeginTime     string `p:"beginTime"`     //开始时间
	EndTime       string `p:"endTime"`       //结束时间
	OrderByColumn string `p:"orderByColumn"` //排序字段
	IsAsc         string `p:"isAsc"`         //排序方式
	PageNum       int    `p:"pageNum"`       //当前页码
	PageSize      int    `p:"pageSize"`      //每页数
}

//新增页面请求参数
type PostApiCreateReq struct {
	PostName string `p:"postName"  v:"required#岗位名称不能为空"`
	PostCode string `p:"postCode"  v:"required#岗位编码不能为空"`
	PostSort int    `p:"postSort"  v:"required#显示顺序不能为空"`
	Status   string `p:"status"    v:"required#状态不能为空"`
	Remark   string `p:"remark"`
}

//修改页面请求参数
type PostApiUpdateReq struct {
	PostId   int64  `p:"postId" v:"required#主键ID不能为空"`
	PostApiCreateReq
}

// API执行删除内容
type PostApiDeleteReq struct {
	Ids string `p:"ids"  v:"required#请选择要删除的数据记录"`
}

// ======= service =========
// 查询列表返回值
type PostServiceList struct {
	List  []model.SysPost `json:"list"`
	Page  int             `json:"page"`
	Size  int             `json:"size"`
	Total int             `json:"total"`
}
