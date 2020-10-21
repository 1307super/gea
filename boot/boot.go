package boot

import (
	"fmt"
	"gea/library/apollo"
	"gea/library/logstash"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"os"
)

var Active string

func init() {
	initConfig()
	apollo.New()
	logstash.New(Active)
}

// 初始化配置
func initConfig() {
	// 获取环境变量
	active := os.Getenv("GF_PROFILES_ACTIVE")
	glog.Infof("从环境变量获取环境：%s",active)
	if active == "" {
		active = gconv.String(g.Cfg().Get("profile.active"))
		glog.Infof("从配置文件获取环境：%s",active)
	}
	switch active {
	case "dev":
		g.Cfg().SetFileName("config-dev.toml")
	case "pro":
		g.Cfg().SetFileName("config-pro.toml")
	case "local":
		g.Cfg().SetFileName("config-local.toml")
	}
	fmt.Println("当前配置环境：",active)
	Active = gconv.String(active)
}