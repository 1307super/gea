package index

import (
	"gea/app/utils/response"
	"gea/app/utils/token"
	"github.com/gogf/gf/net/ghttp"
	"io/ioutil"
	"os"
)



//下载文件
func Download(r *ghttp.Request) {
	fileName := r.GetQueryString("fileName")
	delete := r.GetQueryBool("delete")

	if fileName == "" {
		response.NotFoundResp(r).WriteJsonExit()
		//response.ErrorTpl(r).WriteTpl(g.Map{
		//	"desc": "参数错误",
		//})
		return
	}

	// 创建路径
	curDir, err := os.Getwd()
	if err != nil {
		response.NotFoundResp(r).WriteJsonExit()
		//response.ErrorTpl(r).WriteTpl(g.Map{
		//	"desc": "获取目录失败",
		//})
		return
	}

	filepath := curDir + "/public/upload/" + fileName
	file, err := os.Open(filepath)

	defer file.Close()

	if err != nil {
		response.NotFoundResp(r).WriteJsonExit()
		//response.ErrorTpl(r).WriteTpl(g.Map{
		//	"desc": "参数错误",
		//})
		return
	}

	b, _ := ioutil.ReadAll(file)

	r.Response.Header().Add("Content-Disposition", "attachment")
	r.Response.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	r.Response.Write(b)

	if delete {
		os.Remove(filepath)
	}

}

//注销
func Logout(r *ghttp.Request) {
	token.RemoveCache(token.CacheKey + r.GetString("jwtLoginName"))
	response.SucessResp(r).SetMsg("ok").WriteJsonExit()
}
