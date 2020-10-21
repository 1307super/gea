package logstash

import (
	"gea/app/utils/logstash"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

type MyLogStashWriter struct {
	logst  *logstash.Logstash
	logger *glog.Logger
	Active string
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
		Env:     w.Active,
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


// 初始化logstash
func New(active string) {
	logstashConfig := g.Cfg().GetMapStrStr("logstash")
	if logstashConfig["Enable"] == "false" {
		return
	}
	glog.SetWriter(&MyLogStashWriter{
		logst:  logstash.New(logstashConfig["Ip"], gconv.Int(logstashConfig["Port"]), gconv.Int(logstashConfig["TimeOut"])),
		logger: glog.New(),
		Active: active,
	})
}
