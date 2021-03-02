package utils

import (
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
	"os"
)

//下载文件
func Download(r *ghttp.Request) {
	fileName := r.GetQueryString("fileName")
	del := r.GetQueryBool("delete")

	if fileName == "" {
		response.NotFoundResp(r).WriteJsonExit()
		return
	}

	// 创建路径
	curDir, err := os.Getwd()
	if err != nil {
		response.NotFoundResp(r).WriteJsonExit()
		return
	}

	filepath := curDir + "/public/upload/" + fileName
	file, err := os.Open(filepath)

	defer file.Close()

	if err != nil {
		response.NotFoundResp(r).WriteJsonExit()
		return
	}

	r.Response.ServeFileDownload(filepath,fileName)
	if del {
		os.Remove(filepath)
	}

}
