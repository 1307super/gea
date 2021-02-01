package response

import (
	"github.com/gogf/gf/net/ghttp"
)

const ResponseCtxName  = "ResponseCtx"
const ResponseBunissType  = "ResponseBunissType"

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

// 通用api响应
type ApiResp struct {
	c *CommonRes
	r *ghttp.Request
	s int
}

func NewResp(r *ghttp.Request) *ApiResp {
	return &ApiResp{
		c: &CommonRes{
			Code:  -1,
			Btype: Buniss_Other,
			Msg:   "",
		},
		r: r,
	}
}
//返回一个成功的消息体
func SucessResp(r *ghttp.Request) *ApiResp {
	msg := CommonRes{
		Code:  0,
		Btype: Buniss_Other,
		Msg:   "操作成功",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}

//返回一个错误的消息体
func ErrorResp(r *ghttp.Request) *ApiResp {
	msg := CommonRes{
		Code:  500,
		Btype: Buniss_Other,
		Msg:   "操作失败",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}

func NauthorizedResp(r *ghttp.Request) *ApiResp {
	msg := CommonRes{
		Code:  401,
		Btype: Buniss_Other,
		Msg:   "请登录",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}

//返回一个拒绝访问的消息体
func ForbiddenResp(r *ghttp.Request) *ApiResp {
	msg := CommonRes{
		Code:  403,
		Btype: Buniss_Other,
		Msg:   "无操作权限",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}
//返回404的消息体
func NotFoundResp(r *ghttp.Request) *ApiResp {
	msg := CommonRes{
		Code:  404,
		Btype: Buniss_Other,
		Msg:   "NOT FOUND",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}


//返回一个成功的消息体
func (resp *ApiResp)SucessResp(r *ghttp.Request) *ApiResp {
	msg := CommonRes{
		Code:  0,
		Btype: Buniss_Other,
		Msg:   "操作成功",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}

//返回一个错误的消息体
func (resp *ApiResp)ErrorResp(r *ghttp.Request) *ApiResp {
	msg := CommonRes{
		Code:  500,
		Btype: Buniss_Other,
		Msg:   "操作失败",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}

func (resp *ApiResp) NauthorizedResp(r *ghttp.Request) *ApiResp {
	msg := CommonRes{
		Code:  401,
		Btype: Buniss_Other,
		Msg:   "请登录",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}
//返回一个拒绝访问的消息体
func (resp *ApiResp)ForbiddenResp(r *ghttp.Request) *ApiResp {
	msg := CommonRes{
		Code:  403,
		Btype: Buniss_Other,
		Msg:   "无操作权限",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}

func (resp *ApiResp) SetStatus(status int) *ApiResp {
	resp.s = status
	return resp
}
//设置消息体的内容
func (resp *ApiResp) SetMsg(msg string) *ApiResp {
	resp.c.Msg = msg
	return resp
}

//设置消息体的编码
func (resp *ApiResp) SetCode(code int) *ApiResp {
	resp.c.Code = code
	return resp
}

//设置消息体的数据
func (resp *ApiResp) SetData(data interface{}) *ApiResp {
	resp.c.Data = data
	return resp
}

//设置消息体的业务类型
func (resp *ApiResp) SetBtype(btype BunissType) *ApiResp {
	resp.c.Btype = btype
	return resp
}

func (resp *ApiResp) SetModule(module string) *ApiResp{
	resp.c.Module = module
	return resp
}


//输出json到客户端
func (resp *ApiResp) WriteJsonExit() {
	//resp.r.SetCtxVar(ResponseCtxName,resp.c)
	resp.r.Response.WriteJsonExit(resp.c)
}

//输出json到客户端
func (resp *ApiResp) WriteStatusJsonExit() {
	//resp.r.SetCtxVar(ResponseCtxName,resp.c)
	resp.r.Response.WriteStatus(resp.s)
	resp.r.Response.WriteJsonExit(resp.c)
}
