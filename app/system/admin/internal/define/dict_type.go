package define

import "gea/app/model"

// ======== api =======
//分页请求参数
type DictTypeApiSelectPageReq struct {
	DictName      string `p:"dictName"`      //字典名称
	DictType      string `p:"dictType"`      //字典类型
	Status        string `p:"status"`        //字典状态
	BeginTime     string `p:"beginTime"`     //开始时间
	EndTime       string `p:"endTime"`       //结束时间
	OrderByColumn string `p:"orderByColumn"` //排序字段
	IsAsc         string `p:"isAsc"`         //排序方式
	PageNum       int    `p:"pageNum"`       //当前页码
	PageSize      int    `p:"pageSize"`      //每页数
}

//新增页面请求参数
type DictTypeApiCreateReq struct {
	DictName string `p:"dictName"  v:"required#字典名称不能为空"`
	DictType string `p:"dictType"  v:"required#字典类型不能为空"`
	Status   string `p:"status"  v:"required#状态不能为空"`
	Remark   string `p:"remark"`
}

//修改页面请求参数
type DictTypeApiUpdateReq struct {
	DictId   int64  `p:"dictId" v:"required#主键ID不能为空"`
	DictTypeApiCreateReq
}

// API执行删除内容
type DictTypeApiDeleteReq struct {
	Ids string `p:"ids"  v:"required#请选择要删除的数据记录"`
}

// ====== service =======
// 查询列表返回值
type DictTypeServiceList struct {
	List  []model.SysDictType `json:"list"`
	Page  int                 `json:"page"`
	Size  int                 `json:"size"`
	Total int                 `json:"total"`
}
