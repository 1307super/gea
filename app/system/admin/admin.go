package admin

import (
	"gea/app/system/admin/internal/api"
	"gea/app/system/admin/internal/service"
	"gea/app/utils"
	"gea/app/utils/convert"
	"gea/library/casbin"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

func Init() {
	g.View().BindFunc("Capitalize",convert.Capitalize) // 首字母大写
	g.View().BindFunc("CamelString",gstr.CamelCase)    // 转驼峰
	g.View().BindFunc("CamelLower",gstr.CamelLowerCase)    // 转首字母小写驼峰
	s := g.Server("admin")
	adminConfig, _ := ghttp.ConfigFromMap(g.Cfg().GetMap("admin"))
	s.SetConfig(adminConfig)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware.CORS)
		// 操作日志
		//group.Hook("/*",ghttp.HookBeforeOutput,hook.OperationLog)

		// 上下文
		group.Middleware(service.Middleware.Ctx)
		// 不鉴权
		group.GET("/captchaImage",api.Login.CaptchaImage)
		group.POST("/login",api.Login.CheckLogin)
		//group.POST("/logout",api.Login.Logout)
		// 下载
		group.GET("/download",utils.Download)

		// 鉴权
		group.Middleware(service.Middleware.Jwt,service.Middleware.Auth,service.Middleware.OperationLog)
		// 用户列表
		group.GET("/getInfo",api.User.GetUserInfo)
		group.REST("/system/user",api.User)
		group.GET("/system/user/info",  api.User.Info)
		group.PUT("/system/user/resetPwd",api.User.ResetPwdSave)
		group.PUT("/system/user/changeStatus",api.User.ChangeStatus)
		group.GET("/system/user/export", api.User.Export)

		// 用户信息
		group.GET("/system/user/profile", api.User.GetProfile)
		group.PUT("/system/user/profile", api.User.UpdateProfile)
		group.PUT("/system/user/profile/updatePwd", api.User.UpdatePassword)
		group.POST("/system/user/profile/avatar", api.User.UpdateAvatar)


		// 角色
		group.REST("/system/role",api.Role)
		group.GET("/system/role/info",  api.Role.Info)
		group.GET("/system/role/export",  api.Role.Export)
		group.PUT("/system/role/dataScope",  api.Role.AuthDataScopeSave)
		group.PUT("/system/role/changeStatus", api.Role.ChangeStatus)

		// 菜单
		// 获取路由
		group.GET("/getRouters", api.Menu.GetRouter)
		group.REST("/system/menu",api.Menu)
		group.GET("/system/menu/info",api.Menu.Info)
		group.GET("/system/menu/treeselect",  api.Menu.MenuTreeData)
		group.GET("/system/menu/roleMenuTreeselect",  api.Menu.RoleMenuTreeData)

		// 部门
		group.REST("/system/dept",api.Dept)
		group.GET("/system/dept/info",api.Dept.Info)
		group.GET("/system/dept/treeselect",  api.Dept.TreeData) // 用户列表树形列表
		group.GET("/system/dept/roleDeptTreeselect",api.Dept.RoleDeptTreeData)

		// 部门
		//deptController := new(dept.Controller)
		//group.REST("/system/dept",deptController)
		//group.GET("/system/dept/info",deptController,"Info")
		//group.GET("/system/dept/treeselect",  deptController,"TreeData") // 用户列表树形列表
		//group.GET("/system/dept/roleDeptTreeselect",deptController,"RoleDeptTreeData")

		// 岗位
		group.REST("/system/post",api.Post)
		group.GET("/system/post/info",api.Post.Info)
		group.GET("/system/post/export",api.Post.Export)

		//postController := new(post.Controller)
		//group.REST("/system/post",postController)
		//group.GET("/system/post/info",postController,"Info")

		// 字典
		group.REST("/system/dict/type",api.DictType)
		group.GET("/system/dict/type/info",  api.DictType.Info)
		group.GET("/system/dict/type/optionselect", api.DictType.Optionselect)
		group.GET("/system/dict/type/export", api.DictType.Export)

		//dictTypeController := new(dict_type.Controller)
		//group.REST("/system/dict/type",dictTypeController)
		//group.GET("/system/dict/type/info",  dictTypeController,"Info")
		//group.GET("/system/dict/type/optionselect", dictTypeController,"Optionselect")

		// 字典数据
		group.REST("/system/dict/data",api.DictData)
		group.GET("/system/dict/datas",  api.DictData.GetAll)
		group.GET("/system/dict/data/info",api.DictData.Info)
		group.GET("/system/dict/data/export",api.DictData.Export)

		//dictDataController := new(dict_data.Controller)
		//group.REST("/system/dict/data",dictDataController)
		//group.GET("/system/dict/datas",  dictDataController,"List")
		//group.GET("/system/dict/data/info",dictDataController,"Info")

		// 配置文件
		group.REST("/system/config",api.Config)
		group.GET("/system/config/info",api.Config.Info)
		group.GET("/system/configKey",  api.Config.GetValueByKey)
		group.GET("/system/config/export",  api.Config.Export)
		//configController := new(config.Controller)
		//group.REST("/system/config",configController)
		//group.GET("/system/config/info",configController,"Info")
		//group.GET("/system/configKey",  configController,"GetValueByKey")

		// 日志
		group.REST("/monitor/operlog",api.Operlog)
		group.DELETE("/monitor/operlog/clean",  api.Operlog.Clean)       // 清空操作日志
		group.DELETE("/monitor/operlog/export",  api.Operlog.Export)       // 导入操作日志

		//operlogController := new(operlog.Controller)
		//group.REST("/monitor/operlog",operlogController)
		//group.DELETE("/monitor/operlog/clean",  operlogController,"Clean")       // 操作日志

		// 登录日志
		group.REST("/monitor/logininfor",api.Logininfor)
		group.DELETE("/monitor/logininfor/clean", api.Logininfor.Clean)
		group.DELETE("/monitor/logininfor/export", api.Logininfor.Export)
		//logininforController := new(logininfor.Controller)
		//group.REST("/monitor/logininfor",logininforController)
		//group.DELETE("/monitor/logininfor/clean", logininforController,"Clean") // 登录日志

		// 在线用户
		group.REST("/monitor/online",api.Online)
		//onlineController := new(online.Controller)
		//group.REST("/monitor/online",onlineController)

		// 定时任务
		group.REST("/monitor/job",api.Job)
		group.GET("/monitor/job/info", api.Job.Info)
		group.PUT("/monitor/job/run", api.Job.Start)
		group.PUT("/monitor/job/changeStatus",  api.Job.ChangeStatus)

		//jobController := new(job.Controller)
		//group.REST("/monitor/job",jobController)
		//group.GET("/monitor/job/info", jobController,"Info")
		//group.PUT("/monitor/job/run", jobController,"Start")
		//group.PUT("/monitor/job/changeStatus",  jobController,"ChangeStatus")
		// 任务日志
		group.REST("/monitor/jobLog",  api.JobLog)
		//group.DELETE("/monitor/jobLog",jobController,"DeleteLog")
		group.DELETE("/monitor/jobLog/clean",api.JobLog.Clean)

		//group.GET("/monitor/jobLog",  jobController,"GetLog")
		//group.DELETE("/monitor/jobLog",jobController,"DeleteLog")
		//group.DELETE("/monitor/jobLog/clean",jobController,"CleanLog")

		// 服务监控
		group.GET("/monitor/server", api.Server.Server)

		// 代码生成
		group.REST("/tool/gen",api.GenTable)
		group.GET("/tool/gen/db/list", api.GenTable.DataList)
		group.GET("/tool/gen/preview",  api.GenTable.Preview)
		group.GET("/tool/gen/info", api.GenTable.Info)
		group.GET("/tool/gen/batchGenCode", api.GenTable.GenCode)
		//
		//genController := new(gen.Controller)
		//group.REST("/tool/gen",genController)
		//group.GET("/tool/gen/db/list", genController,"DataList")
		//group.GET("/tool/gen/preview",  genController,"Preview")
		//group.GET("/tool/gen/info", genController,"Info")
		//group.GET("/tool/gen/batchGenCode", genController,"GenCode")

	})
	// 允许的权限

	service.ALLOW_PERMISSION.Append("/getInfo","/getRouters",
		"/system/dict/datas", "/system/user/profile",
		"/system/user/profile/updatePwd","/system/user/profile/avatar",
		"/system/menu/treeselect","/system/menu/roleMenuTreeselect",
		"/system/dept/treeselect","/system/dept/roleDeptTreeselect",
		"/system/dict/datas","/system/configKey")

	initCasbin()
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