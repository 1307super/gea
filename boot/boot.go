package boot

import (
	"fmt"
	"gea/app/utils/logstash"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"github.com/zouyx/agollo/v3"
	"github.com/zouyx/agollo/v3/env/config"
)

func init() {
	initConfig()
	//initApollo()
	//initLogStash()
}


// 初始化配置
func initConfig() {
	active := g.Cfg().Get("profile.active")
	switch active {
	case "dev":
		g.Cfg().SetFileName("config-dev.toml")
	case "pro":
		g.Cfg().SetFileName("config-pro.toml")
	case "local":
		g.Cfg().SetFileName("config-local.toml")
	}
	fmt.Println("当前配置环境：",active)
}

// 初始化apollo配置
func initApollo() {
	apolloConfig := g.Cfg().GetMapStrStr("apollo")
	if apolloConfig["Enable"] == "false" {
		glog.Info("使用config.toml配置")
	}else{
		readyConfig := &config.AppConfig{
			AppID:            apolloConfig["AppId"],
			Cluster:          apolloConfig["Cluster"],
			NamespaceName:    apolloConfig["NamespaceName"],
			IP:               apolloConfig["Ip"],
			NextTryConnTime:  gconv.Int64(apolloConfig["NextTryConnTime"]),
			IsBackupConfig:   gconv.Bool(apolloConfig["IsBackupConfig"]),
		}
		agollo.InitCustomConfig(func() (appConfig *config.AppConfig, err error) {
			return readyConfig,nil
		})
		err := agollo.Start()
		if err != nil {
			panic(err)
		}
		glog.Info("初始化Apollo配置成功")
	}
}
// 初始化logstash
func initLogStash() {
	glog.SetWriter(&MyLogStashWriter{
		logst:  logstash.New("10.86.13.239", 4560, 5),
		logger: glog.New(),
	})
}

type MyLogStashWriter struct {
	logst  *logstash.Logstash
	logger *glog.Logger
}

type LogMessage struct {
	AppName string `json:"appName"`
	Message string `json:"message"`
}

func (w *MyLogStashWriter) Write(p []byte) (n int, err error) {
	message := LogMessage{
		AppName: gconv.String(g.Cfg().GetString("logger.AppName")),
		Message: gconv.String(p),
	}
	_, err = w.logst.Connect()
	if err == nil {
		jsonmessage, _ := gjson.Encode(message)
		w.logst.Writeln(gconv.String(jsonmessage))
	}
	return w.logger.Write(p)
}

