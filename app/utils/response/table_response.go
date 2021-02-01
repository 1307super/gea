package response

import (
	"github.com/gogf/gf/net/ghttp"
)

// 通用api响应
type TableResp struct {
	c *TableDataInfo
	r *ghttp.Request
}

//返回一个成功的消息体
func BuildTable(r *ghttp.Request, total int, rows interface{}) *TableResp {
	msg := TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: total,
		Rows:  rows,
	}
	a := TableResp{
		c: &msg,
		r: r,
	}
	return &a
}

// 总数
func (resp *TableResp) SetTotal(total int) *TableResp {
	resp.c.Total = total
	return resp
}
// 数据
func (resp *TableResp) SetRows(rows interface{}) *TableResp {
	resp.c.Rows = rows
	return resp
}

//返回一个成功的消息体
func (resp *TableResp)BuildTable(r *ghttp.Request, total int, rows interface{}) *TableResp {
	msg := TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: total,
		Rows:  rows,
	}
	return &TableResp{c: &msg, r: r}
}

//输出json到客户端
func (resp *TableResp) WriteJsonExit() {
	resp.r.Response.WriteJsonExit(resp.c)
}
