package define

import "gea/app/model"

// Fill with you ideas below.
// ========== api =============
//分页请求参数
type ConfigApiSelectPageReq struct {
	ConfigName string `p:"configName"` //参数名称
	ConfigKey  string `p:"configKey"`  //参数键名
	ConfigType string `p:"configType"` //状态
	BeginTime  string `p:"beginTime"`  //开始时间
	EndTime    string `p:"endTime"`    //结束时间
	PageNum    int    `p:"pageNum"`    //当前页码
	PageSize   int    `p:"pageSize"`   //每页数
}

//新增页面请求参数
type ConfigApiCreateReq struct {
	ConfigName  string `p:"configName"  v:"required#参数名称不能为空"`
	ConfigKey   string `p:"configKey"  v:"required#参数键名不能为空"`
	ConfigValue string `p:"configValue"  v:"required#参数键值不能为空"`
	ConfigType  string `p:"configType"    v:"required#系统内置不能为空"`
	Remark      string `p:"remark"`
}

//修改页面请求参数
type ConfigApiUpdateReq struct {
	ConfigId    int64  `p:"configId" v:"required#主键ID不能为空"`
	ConfigName  string `p:"configName"  v:"required#参数名称不能为空"`
	ConfigKey   string `p:"configKey"  v:"required#参数键名不能为空"`
	ConfigValue string `p:"configValue"  v:"required#参数键值不能为空"`
	ConfigType  string `p:"configType"    v:"required#系统内置不能为空"`
	Remark      string `p:"remark"`
}

//通用的删除请求
type ConfigApiDeleteReq struct {
	Ids string `p:"ids"  v:"required#请选择要删除的数据记录"`
}

// ======= service =========
// 查询列表返回值
type ConfigServiceList struct {
	List  []model.SysConfig `json:"list"`
	Page  int               `json:"page"`
	Size  int               `json:"size"`
	Total int               `json:"total"`
}
