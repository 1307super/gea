package apollo

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang/glog"
	"github.com/zouyx/agollo/v4"
	"github.com/zouyx/agollo/v4/env/config"
	"github.com/zouyx/agollo/v4/storage"
	"sync"
)

// 初始化apollo配置
func New() {
	apolloConfig := g.Cfg().GetMapStrStr("apollo")
	if apolloConfig["Enable"] == "false" {
		glog.Info("使用config.toml配置")
	}else{
		g.Dump(apolloConfig)
		c := &config.AppConfig{
			AppID:            apolloConfig["AppId"],
			Cluster:          apolloConfig["Cluster"],
			NamespaceName:    apolloConfig["NamespaceName"],
			IP:               apolloConfig["Ip"],
			NextTryConnTime:  gconv.Int64(apolloConfig["NextTryConnTime"]),
			IsBackupConfig:   gconv.Bool(apolloConfig["IsBackupConfig"]),
		}
		agollo.SetLogger(&DefaultLogger{})
		client,err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
			return c,nil
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("初始化Apollo配置成功")
		c2 := &CustomChangeListener{}
		// 开始监听配置
		client.AddChangeListener(c2)
		writeConfig(c.NamespaceName,client)
	}
}

func writeConfig(namespace string,client *agollo.Client) {
	cache := client.GetConfigCache(namespace)
	cache.Range(func(key, value interface{}) bool {
		fmt.Println("key : ", key, ", value :", value)
		g.Cfg().Set(gconv.String(key),value)
		return true
	})
}

type CustomChangeListener struct {
	wg sync.WaitGroup
}

func (c *CustomChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	//write your code here
	fmt.Println("========== 获取到配置更新 ============")
	for key, value := range changeEvent.Changes {
		fmt.Println("change key : ", key, ", value :", value)
		g.Cfg().Set(gconv.String(key),value.NewValue)
	}
	//fmt.Println(changeEvent.Namespace)
}

func (c *CustomChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	//write your code here
}
