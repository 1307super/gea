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
	"os"
)

var Active string

func init() {
	initConfig()
	initApollo()
	initLogStash()
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
	logstashConfig := g.Cfg().GetMapStrStr("logstash")
	if logstashConfig["Enable"] == "false" {
		return
	}
	glog.SetWriter(&MyLogStashWriter{
		logst:  logstash.New(logstashConfig["Ip"], gconv.Int(logstashConfig["Port"]), gconv.Int(logstashConfig["TimeOut"])),
		logger: glog.New(),
	})
}

type MyLogStashWriter struct {
	logst  *logstash.Logstash
	logger *glog.Logger
}

type LogMessage struct {
	AppName string `json:"appName"`
	Env     string `json:"env"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (w *MyLogStashWriter) Write(p []byte) (n int, err error) {
	message := LogMessage{
		AppName: gconv.String(g.Cfg().GetString("logstash.AppName")),
		Env:     Active,
		Type:    gconv.String(g.Cfg().GetString("logstash.Type")),
		Message: gconv.String(p),
	}
	_, err = w.logst.Connect()
	if err == nil {
		jsonmessage, _ := gjson.Encode(message)
		w.logst.Writeln(gconv.String(jsonmessage))
	}
	return w.logger.Write(p)
}

