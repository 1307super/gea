package api

func init() {
	//s := g.Server("api")
	//g1 := router.New("api", "/api/v1",cors.CORS)
	//apiRole := new(system.RoleController)
	//g1.GET("/role", "",apiRole.Role)
	//g1.GET("/index","",login.Index)
	//g1.POST("/auth/login","",login.Login)


	////// 不需要鉴权
	//s.Group("/api/v1", func(group *ghttp.RouterGroup) {
	//	group.Middleware(cors.CORS)
	//	group.POST("/login", login.Login)
	//	// 需要鉴权
	//	group.Middleware(jwt.JWT)
	//	group.GET("/index", login.Index)
	//})


}
