package api

import (

	"gea/app/model"
	"gea/app/system/admin/internal/define"
	"gea/app/system/admin/internal/service"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)


var User = new(userApi)

type userApi struct {
	BaseController
}

func (a *userApi)Init(r *ghttp.Request) {
	// 设置模块
	a.Module = "用户管理"
}
// RESTful
//用户列表分页数据
func (a *userApi) Get(r *ghttp.Request) {
	var req *define.UserApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	resp := service.User.GetList(r.Context(),req)
	if resp == nil{
		a.SuccTable(r,0,nil)
	}

	a.SuccTable(r,resp.Total,resp.List)
}
//保存新增用户数据
func (a *userApi) Post(r *ghttp.Request) {
	var req *define.UserApiCreateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	uid, err := service.User.Create(r.Context(),req)
	if err != nil || uid <= 0 {
		a.Err(r,err.Error())
	}
	a.Succ(r,uid)
}
//保存修改用户数据
func (a *userApi) Put(r *ghttp.Request) {
	var req *define.UserApiUpdateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}

	uid, err := service.User.Update(r.Context(),req)
	if err != nil || uid <= 0 {
		a.Err(r,err.Error())
	}
	a.Succ(r,uid)
}
//删除数据
func (a *userApi) Delete(r *ghttp.Request) {
	var req *define.UserApiDeleteReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}

	rs := service.User.Delete(req.Ids)
	if rs > 0 {
		a.Succ(r,rs)
	} else {
		a.Err(r,"删除失败")
	}
}


// 详情
func (a *userApi) Info(r *ghttp.Request) {
	uid := r.GetInt64("userId")
	var userInfo model.SysUserInfo
	// 1.获取角色
	roles, err := service.Role.SelectRoleContact()
	if err != nil {
		a.Err(r,err.Error())
	}
	// 2.获取岗位信息
	posts := service.Post.GetAll(nil)
	if uid > 0 {
		// 3.获取用户信息
		user := service.User.GetInfo(uid)
		// 4.获取角色Id
		roleIds,err := service.Role.GetRoleListByUid(uid)
		// 5.获取岗位Id
		postIds, err := service.Post.GetPostListByUid(uid)
		if err != nil {
			a.Err(r,err.Error())
		}
		userInfo.RoleIds = roleIds
		userInfo.UserExtend = user
		userInfo.PostIds = postIds
	}
	userInfo.Roles = roles
	userInfo.Posts = posts
	a.Succ(r,userInfo)
}

// 管理员重置密码保存
func (a *userApi) ResetPwdSave(r *ghttp.Request) {
	var req *define.UserApiResetPwdReq
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	result, err := service.User.ResetPassword(req)
	if err != nil || !result {
		a.Err(r,err.Error())
	} else {
		a.Succ(r)
	}
}


// 修改用户状态
func (a *userApi)ChangeStatus(r *ghttp.Request) {
	var req *define.UserApiChangeStatus
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	err := service.User.ChangeStatus(req.UserId,req.Status)
	if err != nil {
		a.Err(r,err.Error())
	}
	a.Succ(r)
}

//导出
func (a *userApi)Export(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,response.Buniss_Export)
	var req *define.UserApiSelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	url, err := service.User.Export(req)

	if err != nil {
		a.Err(r,err.Error())
	}
	a.Succ(r,url)
}


// ============ 登录用户 ==========

// 用户资料页
func (a *userApi)GetUserInfo(r *ghttp.Request){
	userInfo, err := service.User.GetUserInfo(r.Context())
	if err != nil {
		a.Err(r,err.Error())
	}
	a.Succ(r,userInfo)
}

// 详情
func (a *userApi) GetProfile(r *ghttp.Request) {
	userInfo,err := service.User.GetProfile(r.Context())
	if err != nil || userInfo == nil{
		a.Err(r,"未查询到该用户")
	}
	a.Succ(r,userInfo)
}

// 修改个人信息
func (a *userApi) UpdateProfile(r *ghttp.Request) {
	var req *define.UserApiProfileReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	err := service.User.UpdateProfile(r.Context(),req)
	if err != nil{
		a.Err(r,err.Error())
	}
	a.Succ(r)
}

// 用户修改密码
func (a *userApi) UpdatePassword(r *ghttp.Request) {
	var req *define.UserApiReSetPasswordReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	err := service.User.UpdatePassword(r.Context(),req)
	if err != nil{
		a.Err(r,err.Error())
	}
	a.Succ(r)
}

// 用户修改头像
func (a *userApi) UpdateAvatar(r *ghttp.Request) {
	file := r.GetUploadFile("avatarfile")
	if file == nil {
		a.Err(r,"请选择需要上传的文件")
	}
	avatar, err := service.User.UpdateAvatar(r.Context(),&define.UserApiAvatarUploadReq{Avatarfile: file})
	if err != nil{
		a.Err(r,err.Error())
	}
	a.Succ(r,avatar)
}



