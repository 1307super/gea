package api

import (
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

type BaseController struct {
	Module string
}

func (c *BaseController) Succ(r *ghttp.Request, data ...interface{}) {
	resp := response.SucessResp(r).SetCode(0).SetMsg("操作成功").SetModule(c.Module)
	if len(data) > 0 {
		resp.SetData(data[0])
	}
	resp.WriteJsonExit()
}
func (c *BaseController) Err(r *ghttp.Request, msg string, data ...interface{}) {
	if msg == "" {
		msg = "操作失败"
	}
	resp := response.ErrorResp(r).SetModule(c.Module)
	resp.SetCode(500).SetMsg(msg)
	if len(data) > 0 {
		resp.SetData(data[0])
	}
	resp.WriteJsonExit()
}
func (c *BaseController) SuccTable(r *ghttp.Request, total int, data interface{}) {
	response.BuildTable(r,total,data).WriteJsonExit()
}


