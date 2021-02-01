package define

// ======== api ==========
//分页请求参数
type MenuApiSelectPageReq struct {
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
type MenuApiCreateReq struct {
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
type MenuApiEditReq struct {
	MenuId    int64  `p:"menuId" v:"required#主键ID不能为空"`
	MenuApiCreateReq
}

// API执行删除内容
type MenuApiDeleteReq struct {
	Id int64 `p:"id"  v:"required#请选择要删除的数据记录"`
}