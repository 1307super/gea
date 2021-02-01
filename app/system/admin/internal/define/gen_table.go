package define

import "gea/app/model"

// ============ api ===============
//分页请求参数
type GenTableApiSelectPageReq struct {
	TableName    string `p:"tableName"`    //表名称
	TableComment string `p:"tableComment"` //表描述
	BeginTime    string `p:"beginTime"`    //开始时间
	EndTime      string `p:"endTime"`      //结束时间
	PageNum      int    `p:"pageNum"`      //当前页码
	PageSize     int    `p:"pageSize"`     //每页数
}

//修改页面请求参数
type GenTableApiUpdateReq struct {
	TableId        int64  `p:"tableId" v:"required#主键ID不能为空"`
	TableName      string `p:"tableName"  v:"required#表名称不能为空"`
	TableComment   string `p:"tableComment"  v:"required#表描述不能为空"`
	ClassName      string `p:"className" v:"required#实体类名称不能为空"`
	FunctionAuthor string `p:"functionAuthor"  v:"required#作者不能为空"`
	TplCategory    string `p:"tplCategory"`
	PackageName    string `p:"packageName" v:"required#生成包路径不能为空"`
	ModuleName     string `p:"moduleName" v:"required#生成模块名不能为空"`
	BusinessName   string `p:"businessName" v:"required#生成业务名不能为空"`
	FunctionName   string `p:"functionName" v:"required#生成功能名不能为空"`
	Remark         string `p:"remark"`
	Params         string `p:"params"`
	Columns        string `p:"columns"`
}

//通用的删除请求
type GenTableApiDeleteReq struct {
	Ids string `p:"ids"  v:"required#请选择要删除的数据记录"`
}

// ======= service =========
// 查询列表返回值
type GenTableServiceList struct {
	List  []model.GenTable `json:"list"`
	Page  int              `json:"page"`
	Size  int              `json:"size"`
	Total int              `json:"total"`
}

type GenTableServicePreview struct {
	ListKey       string
	AppJsKey      string
	TreeKey       string
	SqlKey        string
	EntityKey     string
	ModelKey      string
	ExtendKey     string
	ServiceKey    string
	RouterKey     string
	ControllerKey string
	VoKey         string
	DtoKey        string
}

type GenTableServiceGenFile struct {
	TemplatePath   string
	GenTableEntity *model.GenTableExtend
	GenFileName    string
}
