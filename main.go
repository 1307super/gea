package main

import (
	_ "gea/boot"
	_ "gea/router"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/swagger"

	//"github.com/gogf/gf/net/ghttp"
)

// @title GEAGO 自动生成API文档
// @version 1.0
// @description 生成文档请在调试模式下进行<a href="/tool/swagger?a=r">重新生成文档</a>

// @host localhost:8082
// @BasePath /
func main() {
	serverSwitch := g.Cfg().GetBool("status.admin")
	apiSwitch := g.Cfg().GetBool("status.api")
	if serverSwitch {
		admin := g.Server("admin")
		adminConfig, _ := ghttp.ConfigFromMap(g.Cfg().GetMap("admin"))
		admin.SetConfig(adminConfig)
		admin.Plugin(&swagger.Swagger{})
		admin.Start()
	}
	if apiSwitch {
		api := g.Server("api")
		apiConfig, _ := ghttp.ConfigFromMap(g.Cfg().GetMap("api"))
		api.SetConfig(apiConfig)
		api.Start()
	}
	g.Wait()
}