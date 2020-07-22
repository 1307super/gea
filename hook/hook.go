package hook

import (
	"gea/app/model"
	operlogService "gea/app/service/monitor/operlog"
	userService "gea/app/service/system/user"
	"gea/app/utils/response"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
)


func OperationLog(r *ghttp.Request) {
	// get请求不记录
	if r.Method == "GET" {
		return
	}

	//if !gstr.Contains(r.RequestURI,"operlog") {
		userInfo, err := userService.GetProfileApi(r.GetInt64("jwtUid"))
		if err != nil {
			return
		}
		if userInfo == nil {
			return
		}
		resp := new(model.CommonRes)
		r.GetCtxVar(response.ResponseCtxName).StructDeep(resp)
		resp.Btype = model.BunissType(r.GetCtxVar(response.ResponseBunissType).Int())
		// 存操作日志
		var paramJson string
		param := r.GetMap()
		paramByte, err := gjson.Encode(param)
		if err == nil {
			paramJson = string(paramByte)
		}
		go operlogService.Add(r,resp.Module,paramJson,resp)
	//}

}
