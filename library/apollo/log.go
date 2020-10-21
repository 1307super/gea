package apollo

import "github.com/gogf/gf/os/glog"

//DefaultLogger 默认日志实现
type DefaultLogger struct {
}

//Debugf debug 格式化
func (d *DefaultLogger) Debugf(format string, params ...interface{}) {
	glog.Debugf(format, params)
}

//Infof 打印info
func (d *DefaultLogger) Infof(format string, params ...interface{}) {
	glog.Infof(format, params)
}

//Warnf warn格式化
func (d *DefaultLogger) Warnf(format string, params ...interface{}) {
	glog.Warningf(format, params)
}

//Errorf error格式化
func (d *DefaultLogger) Errorf(format string, params ...interface{}) {
	glog.Errorf(format, params)
}

//Debug 打印debug
func (d *DefaultLogger) Debug(v ...interface{}) {
	glog.Debug(v)
}

//Info 打印Info
func (d *DefaultLogger) Info(v ...interface{}) {
	glog.Info(v)
}

//Warn 打印Warn
func (d *DefaultLogger) Warn(v ...interface{}) {
	glog.Warning(v)
}

//Error 打印Error
func (d *DefaultLogger) Error(v ...interface{}) {
	glog.Error(v)
}
