package api

import (
	"gea/app/system/admin/internal/define"
	"gea/app/system/admin/internal/service"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

var Online = new(onlineApi)

type onlineApi struct {
	BaseController
}

func (a *onlineApi) Init(r *ghttp.Request) {
	a.Module = "在线用户管理"
}

//列表分页数据
func (a *onlineApi) Get(r *ghttp.Request) {
	var req *define.OnlineApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	result := service.Online.GetList(req)
	if result != nil {
		a.SuccTable(r, result.Total, result.List)
	} else {
		a.SuccTable(r, 0, nil)
	}
}

//用户强退
func (a *onlineApi) Delete(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,response.Buniss_ForceLogout)
	tokenStr := r.GetString("tokenStr")
	if tokenStr == "" {
		a.Err(r,"参数错误")
	}
	service.Online.ForceLogout(tokenStr)
	a.Succ(r)
}


