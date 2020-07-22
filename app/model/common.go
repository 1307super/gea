package model

//SESSION前缀
const (
	USER_SESSION_MARK = "user_info"
)

// 登陆用户的菜单列表缓存前缀
const MENU_CACHE = "menu_cache"

const MENU_TREE_CACHE = "menu_tree_cache"

type BunissType int

//业务类型
const (
	Buniss_Other       BunissType = 0 //0其它
	Buniss_Add         BunissType = 1 //1新增
	Buniss_Edit        BunissType = 2 //2修改
	Buniss_Del         BunissType = 3 //3删除
	Buniss_Authorize   BunissType = 4 //4授权
	Buniss_Export      BunissType = 5 //5导出
	Buniss_Import      BunissType = 6 //6导入
	Buniss_ForceLogout BunissType = 7 //7强退
	Buniss_Gen         BunissType = 8 //8生成代码
	Buniss_Clean       BunissType = 9 //9清空数据
)

//响应结果
const (
	SUCCESS      = 0   // 成功
	ERROR        = 500 //错误
	UNAUTHORIZED = 403 //无权限
	FAIL         = -1  //失败
)

// 通用api响应
type CommonRes struct {
	Code   int         `json:"code"`   //响应编码 0 成功 500 错误 403 无权限  -1  失败
	Msg    string      `json:"msg"`    //消息
	Data   interface{} `json:"data"`   //数据内容
	Btype  BunissType  `json:"otype"`  //业务类型
	Module string      `json:"module"` // 模块
}

// 验证码响应
type CaptchaRes struct {
	Code  int         `json:"code"`  //响应编码 0 成功 500 错误 403 无权限
	Msg   string      `json:"msg"`   //消息
	Data  interface{} `json:"data"`  //数据内容
	IdKey string      `json:"idkey"` //验证码ID
}

//通用分页表格响应
type TableDataInfo struct {
	Total int         `json:"total"` //总数
	Rows  interface{} `json:"rows"`  //数据
	Code  int         `json:"code"`  //响应编码 0 成功 500 错误 403 无权限
	Msg   string      `json:"msg"`   //消息
}

//通用的树形结构
type Ztree struct {
	Id      int64  `json:"id"`      //节点ID
	Pid     int64  `json:"pId"`     //节点父ID
	Name    string `json:"name"`    //节点名称
	Title   string `json:"title"`   //节点标题
	Checked bool   `json:"checked"` //是否勾选
	Open    bool   `json:"open"`    //是否展开
	Nocheck bool   `json:"nocheck"` //是否能勾选
}

//通用的删除请求
type RemoveReq struct {
	Ids string `p:"ids"  v:"required#请选择要删除的数据记录"`
}

//通用详情请求
type DetailReq struct {
	Id int64 `json:"id"` //主键ID
}

//通用修改请求
type EditReq struct {
	Id int64 `json:"id"` //主键ID
}
