package config

import (
	"github.com/gogf/gf/frame/g"
	"github.com/zouyx/agollo/v3"
)

func Get(key string)interface{} {
	// 判断是否使用apollo
	apollo := g.Cfg().GetMapStrStr("apollo")
	if apollo["Enable"] == "true" {
		return agollo.GetValue(key)
	}
	return g.Cfg().Get(key)
}
func GetString(key string) string{
	// 判断是否使用apollo
	apollo := g.Cfg().GetMapStrStr("apollo")
	if apollo["Enable"] == "true" {
		return agollo.GetValue(key)
	}
	return g.Cfg().GetString(key)
}
