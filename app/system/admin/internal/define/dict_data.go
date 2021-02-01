package define

import "gea/app/model"

// ======== api =======
// Fill with you ideas below.
//分页请求参数
type DictDataApiSelectPageReq struct {
	DictType  string `p:"dictType"`  //字典名称
	DictLabel string `p:"dictLabel"` //字典标签
	Status    string `p:"status"`    //状态
	BeginTime string `p:"beginTime"` //开始时间
	EndTime   string `p:"endTime"`   //结束时间
	PageNum   int    `p:"pageNum"`   //当前页码
	PageSize  int    `p:"pageSize"`  //每页数
}

//新增页面请求参数
type DictDataApiCreateReq struct {
	DictLabel string `p:"dictLabel"  v:"required#字典标签不能为空"`
	DictValue string `p:"dictValue"  v:"required#字典键值不能为空"`
	DictType  string `p:"dictType"  v:"required#字典类型不能为空"`
	DictSort  int    `p:"dictSort"  v:"required#字典排序不能为空"`
	Status    string `p:"status"    v:"required#状态不能为空"`
	Remark    string `p:"remark"`
}

//修改页面请求参数
type DictDataApiUpdateReq struct {
	DictCode  int64  `p:"dictCode" v:"required#主键ID不能为空"`
	DictDataApiCreateReq
}

// API执行删除内容
type DictDataApiDeleteReq struct {
	Ids string `p:"ids"  v:"required#请选择要删除的数据记录"`
}

// ====== service =======
// 查询列表返回值
type DictDataServiceList struct {
	List  []model.SysDictData `json:"list"`
	Page  int                 `json:"page"`
	Size  int                 `json:"size"`
	Total int                 `json:"total"`
}
