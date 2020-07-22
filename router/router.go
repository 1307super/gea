package router

import (
	_ "gea/app/controller"
	"gea/app/controller/monitor/job"
	"gea/app/controller/monitor/logininfor"
	"gea/app/controller/monitor/online"
	"gea/app/controller/monitor/operlog"
	"gea/app/controller/monitor/server"
	"gea/app/controller/system/config"
	"gea/app/controller/system/dept"
	"gea/app/controller/system/dict_data"
	"gea/app/controller/system/dict_type"
	"gea/app/controller/system/index"
	"gea/app/controller/system/menu"
	"gea/app/controller/system/post"
	"gea/app/controller/system/role"
	"gea/app/controller/system/user"
	"gea/app/controller/tool/gen"
	roleService "gea/app/service/system/role"
	userService "gea/app/service/system/user"
	"gea/hook"
	"gea/library/casbin"
	"gea/middleware/auth"
	"gea/middleware/cors"
	"gea/middleware/jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server("admin")
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(cors.CORS)

		// 测试casbin
		group.GET("/casbin/addUserRole",index.AddP)
		group.GET("/casbin/addRolePerm",index.AddRP)
		group.GET("/casbin/delUserRole",index.DeleteR)
		group.GET("/casbin/delRolePerm",index.DeleteP)
		group.GET("/casbin/check",index.CheckP)


		// 不鉴权
		group.GET("/captchaImage",index.CaptchaImage)
		group.POST("/login",index.CheckLogin)
		group.POST("/logout",index.Logout)
		// 下载
		group.GET("/download",index.Download)
		group.Middleware(jwt.JWT,auth.Auth)
		// 操作日志
		group.Hook("/*",ghttp.HOOK_AFTER_OUTPUT,hook.OperationLog)
		// 用户列表
		userController := new(user.Controller)
		group.GET("/getInfo",userController,"GetInfo")
		group.REST("/system/user",userController)
		group.GET("/system/user/info",  userController,"Info")
		group.PUT("/system/user/resetPwd",userController,"ResetPwdSave")
		group.PUT("/system/user/changeStatus",userController,"ChangeStatus")
		group.GET("/system/user/export", userController,"Export")

		// 用户信息
		group.GET("/system/user/profile", userController,"Profile")
		group.PUT("/system/user/profile", userController,"Update")
		group.PUT("/system/user/profile/updatePwd", userController,"UpdatePassword")
		group.POST("/system/user/profile/avatar", userController,"UpdateAvatar")

		// 角色
		roleController := new(role.Controller)
		group.REST("/system/role",roleController)
		group.GET("/system/role/info",  roleController,"Info")
		group.PUT("/system/role/dataScope",  roleController,"AuthDataScopeSave")
		group.PUT("/system/role/changeStatus", roleController,"ChangeStatus")

		// 菜单
		menuController := new(menu.Controller)
		// 获取路由
		group.GET("/getRouters", menuController, "GetRouter")
		group.REST("/system/menu",menuController)
		group.GET("/system/menu/info",menuController,"Info")
		group.GET("/system/menu/treeselect",  menuController,"MenuTreeData")
		group.GET("/system/menu/roleMenuTreeselect",  menuController,"RoleMenuTreeData")

		// 部门
		deptController := new(dept.Controller)
		group.REST("/system/dept",deptController)
		group.GET("/system/dept/info",deptController,"Info")
		group.GET("/system/dept/treeselect",  deptController,"TreeData") // 用户列表树形列表
		group.GET("/system/dept/roleDeptTreeselect",deptController,"RoleDeptTreeData")

		// 岗位
		postController := new(post.Controller)
		group.REST("/system/post",postController)
		group.GET("/system/post/info",postController,"Info")

		// 字典
		dictTypeController := new(dict_type.Controller)
		group.REST("/system/dict/type",dictTypeController)
		group.GET("/system/dict/type/info",  dictTypeController,"Info")
		group.GET("/system/dict/type/optionselect", dictTypeController,"Optionselect")

		// 字典数据
		dictDataController := new(dict_data.Controller)
		group.REST("/system/dict/data",dictDataController)
		group.GET("/system/dict/datas",  dictDataController,"List")
		group.GET("/system/dict/data/info",dictDataController,"Info")

		// 配置文件
		configController := new(config.Controller)
		group.REST("/system/config",configController)
		group.GET("/system/config/info",configController,"Info")
		group.GET("/system/configKey",  configController,"GetValueByKey")

		// 日志
		operlogController := new(operlog.Controller)
		group.REST("/monitor/operlog",operlogController)
		group.DELETE("/monitor/operlog/clean",  operlogController,"Clean")       // 操作日志

		logininforController := new(logininfor.Controller)
		group.REST("/monitor/logininfor",logininforController)
		group.DELETE("/monitor/logininfor/clean", logininforController,"Clean") // 登录日志

		// 在线用户
		onlineController := new(online.Controller)
		group.REST("/monitor/online",onlineController)

		// 定时任务
		jobController := new(job.Controller)
		group.REST("/monitor/job",jobController)
		group.GET("/monitor/job/info", jobController,"Info")
		group.PUT("/monitor/job/run", jobController,"Start")
		group.PUT("/monitor/job/changeStatus",  jobController,"ChangeStatus")
		// 任务日志
		group.GET("/monitor/jobLog",  jobController,"GetLog")
		group.DELETE("/monitor/jobLog",jobController,"DeleteLog")
		group.DELETE("/monitor/jobLog/clean",jobController,"CleanLog")

		// 服务监控
		group.GET("/monitor/server",  server.Server)

		// 代码生成
		genController := new(gen.Controller)
		group.REST("/tool/gen",genController)
		group.GET("/tool/gen/db/list", genController,"DataList")
		group.GET("/tool/gen/preview",  genController,"Preview")
		group.GET("/tool/gen/info", genController,"Info")
		group.GET("/tool/gen/batchGenCode", genController,"GenCode")
	})
	// 允许的权限
	auth.ALLOW_PERMISSION.Append("/getInfo","/getRouters",
										"/system/dict/datas", "/system/user/profile",
										"/system/user/profile/updatePwd","/system/user/profile/avatar",
										"/system/menu/treeselect","/system/menu/roleMenuTreeselect",
										"/system/dept/treeselect","/system/dept/roleDeptTreeselect",
										"/system/dict/datas","/system/configKey")
	initCasbin()
}

func initCasbin() {
	// 注册
	casbin.Register()
	// 加载用户与角色关系
	userService.LoadUserRole("")
	// 加载角色与资源关系
	roleService.LoadRolePolicy("")
}
