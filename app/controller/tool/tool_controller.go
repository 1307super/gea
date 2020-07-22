package tool

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gproc"
	"os"
)

const (
	swaggoRepoPath = "github.com/swaggo/swag/cmd/swag"
)

//swagger文档
func Swagger(r *ghttp.Request) {
	a := r.GetQueryString("a")
	if a == "r" {
		//重新生成文档
		curDir, err := os.Getwd()
		if err != nil {
			r.Response.WriteTpl("error/error.html", g.Map{
				"desc": "参数错误",
			})
			return
		}

		genPath := curDir + "/public/swagger"
		err = generateSwaggerFiles(genPath, false)
		if err != nil {
			r.Response.WriteTpl("error/error.html", g.Map{
				"desc": "参数错误",
			})
			return
		}
	}
	//r.Response.RedirectTo("/swagger/index.html")
	r.Response.RedirectTo("/swagger-vue/doc.html")
}

//自动生成文档
func generateSwaggerFiles(output string, pack bool) error {
	// Temporary storing swagger files directory.
	tempOutputPath := gfile.Join(gfile.TempDir(), "swagger")
	if gfile.Exists(tempOutputPath) {
		gfile.Remove(tempOutputPath)
	}
	gfile.Mkdir(tempOutputPath)
	// Check and install swag tool.
	swag := gproc.SearchBinary("swag")
	if swag == "" {
		err := gproc.ShellRun(fmt.Sprintf(`go get -u %s`, swaggoRepoPath))
		if err != nil {
			return err
		}
	}
	command := fmt.Sprintf(`swag init -o %s`, tempOutputPath)
	result, err := gproc.ShellExec(command)
	if err != nil {
		return gerror.New(result + err.Error())
	}
	if !gfile.Exists(gfile.Join(tempOutputPath, "swagger.json")) {
		return gerror.New("make swagger files failed")
	}
	if !gfile.Exists(output) {
		gfile.Mkdir(output)
	}
	if err = gfile.CopyFile(
		gfile.Join(tempOutputPath, "swagger.json"),
		gfile.Join(output, "swagger.json"),
	); err != nil {
		return err
	}

	// Auto pack.
	if pack && gfile.Exists("swagger") {
		packCmd := fmt.Sprintf(`gf pack %s boot/data-swagger.go -n boot`, "swagger")
		if err := gproc.ShellRun(packCmd); err != nil {
			return err
		}
	}
	return nil
}

// swagger文档ui
func SwaggerUi(r *ghttp.Request) {
	res := g.Map{
		"deepLinking": true,
		"displayOperationId": false,
		"defaultModelsExpandDepth": 1,
		"defaultModelExpandDepth": 1,
		"defaultModelRendering": "example",
		"displayRequestDuration": false,
		"docExpansion": "none",
		"filter": false,
		"operationsSorter": "alpha",
		"showExtensions": false,
		"tagsSorter": "alpha",
		"validatorUrl": "",
		"apisSorter": "alpha",
		"jsonEditor": false,
		"showRequestHeaders": false,
		"supportedSubmitMethods": g.Array{"get", "put", "post", "delete", "options", "head", "patch", "trace"},
	}
	r.Response.WriteJsonExit(res)
}

// swagger资源接口
func SwaggerResources(r *ghttp.Request) {
	res := g.Array{
		//g.Map{"name":"默认接口","url":"/tool/swagger_doc","swaggerVersion":"2.0","location":"/tool/swagger_doc"},
		g.Map{"name":"默认接口","url":"/swagger/swagger.json","swaggerVersion":"2.0","location":"/swagger/swagger.json"},
	}
	r.Response.WriteJsonExit(res)
}
// swagger文档接口
func SwaggerDoc(r *ghttp.Request) {
	curDir, err := os.Getwd()
	if err != nil {
		r.Response.WriteTpl("error/error.html", g.Map{
			"desc": "参数错误",
		})
		return
	}
	genPath := curDir + "/public/swagger/swagger.json"
	fmt.Println(genPath)
	res,err := gjson.Load(genPath)
	if err != nil {
		fmt.Println("我错误了",err.Error())
	}
	r.Response.WriteJsonExit(res.Export())
}
