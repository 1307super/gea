package apollo

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang/glog"
	"github.com/shima-park/agollo"
)

// 初始化apollo配置
func New() {
	apolloConfig := g.Cfg().GetMapStrStr("apollo")
	if apolloConfig["Enable"] == "false" {
		glog.Info("使用config.toml配置")
	}else{
		g.Dump(apolloConfig)

		a, err := agollo.New(apolloConfig["Ip"], apolloConfig["AppId"], agollo.PreloadNamespaces(apolloConfig["NamespaceName"]))
		// error handle...
		if err != nil {
			panic("加载配置失败")
		}
		errorCh := a.Start()  // Start后会启动goroutine监听变化，并更新agollo对象内的配置cache
		// 或者忽略错误处理直接 a.Start()
		// 写入配置
		writeConfig(a.GetNameSpace(apolloConfig["NamespaceName"]))
		watchCh := a.Watch()
		for{
			select{
			case  <- errorCh:
				// handle error
			case resp := <-watchCh:
				OnChange(resp)
			}
		}
	}
}

func writeConfig(config agollo.Configurations) {
	for key,value := range config {
		fmt.Println("key : ", key, ", value :", value)
		g.Cfg().Set(gconv.String(key),value)
	}
}

func OnChange(resp *agollo.ApolloResponse) {
	//write your code here
	fmt.Println("========== 获取到配置更新 ============")
	for _, value := range resp.Changes {
		fmt.Println("change key : ", value.Key, ", value :", value.Value)
		g.Cfg().Set(gconv.String(value.Key),value.Value)
	}
}
