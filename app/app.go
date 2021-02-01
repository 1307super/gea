package app

import (
	"gea/app/system/admin"
	"github.com/gogf/gf/frame/g"
)

func Run() {
	// 业务初始化
	admin.Init()
	g.Wait()
}
