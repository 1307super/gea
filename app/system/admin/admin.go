package admin

import (
	"gea/app/shared"
	"gea/app/system/admin/internal/api"
	"gea/app/system/admin/internal/define"
	"gea/app/system/admin/internal/service"
	"gea/app/utils"
	"gea/app/utils/convert"
	"gea/library/casbin"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

func Init() {
	g.View().BindFunc("Capitalize", convert.Capitalize)  // 首字母大写
	g.View().BindFunc("CamelString", gstr.CamelCase)     // 转驼峰
	g.View().BindFunc("CamelLower", gstr.CamelLowerCase) // 转首字母小写驼峰
	s := g.Server("admin")
	adminConfig := g.Cfg().GetMap("admin")
	httpConfig, _ := ghttp.ConfigFromMap(adminConfig)
	s.SetConfig(httpConfig)
	prefix := gconv.String(adminConfig["prefix"])

	// 设置静态访问目录
	s.AddStaticPath(prefix+"/static","/public/static")

	shared.GfAdminToken = gtoken.GfToken{
		ServerName:       "admin",
		CacheMode:        g.Cfg().GetInt8("gtoken.CacheMode"),
		Timeout:          g.Cfg().GetInt("gtoken.Timeout"),
		MaxRefresh:       g.Cfg().GetInt("gtoken.MaxRefresh"),
		EncryptKey:       g.Cfg().GetBytes("gtoken.EncryptKey"),
		AuthFailMsg:      g.Cfg().GetString("gtoken.AuthFailMsg"),
		MultiLogin:       g.Cfg().GetBool("gtoken.MultiLogin"),
		LoginPath:        "/login",
		LoginBeforeFunc:  service.User.Login,
		LoginAfterFunc:   service.User.LoginAfter,
		LogoutPath:       "/logout",
		LogoutBeforeFunc: service.User.Logout,
		AuthBeforeFunc:   service.Middleware.Auth,
	}

	// 不鉴权，或仅登录鉴权
	s.Group(prefix, func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware.CORS)
		// 上下文
		group.Middleware(service.Middleware.Ctx)
		// 不鉴权
		group.GET("/captchaImage", api.Login.CaptchaImage)
		// 下载
		group.GET("/download", utils.Download)
		group.Middleware(service.Middleware.LoginAuth)
		// 登录鉴权
		group.GET("/getInfo", api.User.GetUserInfo)
		group.GET("/getRouters", api.Menu.GetRouter)
		group.GET("/system/dict/datas", api.DictData.GetAll)
		group.GET("/system/configKey", api.Config.GetValueByKey)
		// 用户信息
		group.GET("/system/user/profile", api.User.GetProfile)
		group.PUT("/system/user/profile", api.User.UpdateProfile)
		group.PUT("/system/user/profile/updatePwd", api.User.UpdatePassword)
		group.POST("/system/user/profile/avatar", api.User.UpdateAvatar)
		group.GET("/system/menu/treeselect", api.Menu.MenuTreeData)
		group.GET("/system/menu/roleMenuTreeselect", api.Menu.RoleMenuTreeData)
		group.GET("/system/dept/treeselect", api.Dept.TreeData) // 用户列表树形列表
		group.GET("/system/dept/roleDeptTreeselect", api.Dept.RoleDeptTreeData)
	})
	// 需鉴权路由
	s.Group(prefix, func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware.CORS)
		// 上下文
		group.Middleware(service.Middleware.Ctx)
		shared.GfAdminToken.Middleware(group)
		// 菜单鉴权
		group.Middleware(service.Middleware.OperationLog)
		// 用户列表
		group.REST("/system/user", api.User)
		group.GET("/system/user/info", api.User.Info)
		group.PUT("/system/user/resetPwd", api.User.ResetPwdSave)
		group.PUT("/system/user/changeStatus", api.User.ChangeStatus)
		group.GET("/system/user/export", api.User.Export)

		// 角色
		group.REST("/system/role", api.Role)
		group.GET("/system/role/info", api.Role.Info)
		group.GET("/system/role/export", api.Role.Export)
		group.PUT("/system/role/dataScope", api.Role.AuthDataScopeSave)
		group.PUT("/system/role/changeStatus", api.Role.ChangeStatus)

		// 菜单
		// 获取路由
		group.REST("/system/menu", api.Menu)
		group.GET("/system/menu/info", api.Menu.Info)

		// 部门
		group.REST("/system/dept", api.Dept)
		group.GET("/system/dept/info", api.Dept.Info)

		// 岗位
		group.REST("/system/post", api.Post)
		group.GET("/system/post/info", api.Post.Info)
		group.GET("/system/post/export", api.Post.Export)

		// 字典
		group.REST("/system/dict/type", api.DictType)
		group.GET("/system/dict/type/info", api.DictType.Info)
		group.GET("/system/dict/type/optionselect", api.DictType.Optionselect)
		group.GET("/system/dict/type/export", api.DictType.Export)

		// 字典数据
		group.REST("/system/dict/data", api.DictData)
		group.GET("/system/dict/data/info", api.DictData.Info)
		group.GET("/system/dict/data/export", api.DictData.Export)

		// 配置文件
		group.REST("/system/config", api.Config)
		group.GET("/system/config/info", api.Config.Info)
		group.GET("/system/config/export", api.Config.Export)

		// 日志
		group.REST("/monitor/operlog", api.Operlog)
		group.DELETE("/monitor/operlog/clean", api.Operlog.Clean) // 清空操作日志
		group.GET("/monitor/operlog/export", api.Operlog.Export)  // 导入操作日志

		// 登录日志
		group.REST("/monitor/logininfor", api.Logininfor)
		group.DELETE("/monitor/logininfor/clean", api.Logininfor.Clean)
		group.GET("/monitor/logininfor/export", api.Logininfor.Export)

		// 在线用户
		group.REST("/monitor/online", api.Online)

		// 定时任务
		group.REST("/monitor/job", api.Job)
		group.GET("/monitor/job/info", api.Job.Info)
		group.PUT("/monitor/job/run", api.Job.Start)
		group.PUT("/monitor/job/changeStatus", api.Job.ChangeStatus)

		// 任务日志
		group.REST("/monitor/jobLog", api.JobLog)
		group.DELETE("/monitor/jobLog/clean", api.JobLog.Clean)

		// 服务监控
		group.GET("/monitor/server", api.Server.Server)

		// 代码生成
		group.REST("/tool/gen", api.GenTable)
		group.GET("/tool/gen/db/list", api.GenTable.DataList)
		group.GET("/tool/gen/preview", api.GenTable.Preview)
		group.GET("/tool/gen/info", api.GenTable.Info)
		group.GET("/tool/gen/batchGenCode", api.GenTable.GenCode)

	})
	initCasbin()
	initSysDict()
	initTask()
	s.Start()
}


func initCasbin() {
	// 注册
	casbin.Register()
	// 加载用户与角色关系
	service.User.LoadUserRole("")
	// 加载角色与资源关系
	service.Role.LoadRolePolicy("")
}

// 初始化字典
func initSysDict() {
	dictTypeList, _ := service.DictType.GetAll(nil)
	p := &define.DictDataApiSelectPageReq{}
	if len(dictTypeList) > 0 {
		for _, dictType := range dictTypeList {
			p.DictType = dictType.DictType
			service.DictData.GetAll(p)
		}
	}
}
// 初始化定时任务
func initTask()  {
	service.Job.Restart()
}
