package user

import (
	"gea/app/controller"
	"gea/app/model"
	userModel "gea/app/model/system/user"
	postService "gea/app/service/system/post"
	roleService "gea/app/service/system/role"
	userService "gea/app/service/system/user"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
	controller.BaseController
}
func (c *Controller)Init(r *ghttp.Request) {
	// 设置模块
	c.Module = "用户管理"
}

// RESTful
//用户列表分页数据
func (c *Controller) Get(r *ghttp.Request) {
	var req *userModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	rows := make([]userModel.UserListEntity, 0)
	result, page, err := userService.SelectRecordList(req)
	if err == nil && len(result) > 0 {
		rows = result
	}
	c.SuccTable(r,page.Total,rows)
}
//保存新增用户数据
func (c *Controller) Post(r *ghttp.Request) {
	var req *userModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}

	//判断登陆名是否已注册
	isHadName := userService.CheckLoginName(req.LoginName)
	if isHadName {
		c.Err(r,"登陆名已经存在")
	}

	//判断手机号码是否已注册
	isHadPhone := userService.CheckPhoneUniqueAll(req.Phonenumber)
	if isHadPhone {
		c.Err(r,"手机号码已经存在")
	}

	//判断邮箱是否已注册
	isHadEmail := userService.CheckEmailUniqueAll(req.Email)
	if isHadEmail {
		c.Err(r,"邮箱已经存在")
	}

	uid, err := userService.AddSave(req, r)

	if err != nil || uid <= 0 {
		c.Err(r,err.Error())
	}
	c.Succ(r,uid)
}
//保存修改用户数据
func (c *Controller) Put(r *ghttp.Request) {
	var req *userModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}

	//判断手机号码是否已注册
	isHadPhone := userService.CheckPhoneUnique(req.UserId, req.Phonenumber)
	if isHadPhone {
		c.Err(r,"手机号码已经存在")
	}
	//判断邮箱是否已注册
	isHadEmail := userService.CheckEmailUnique(req.UserId, req.Email)
	if isHadEmail {
		c.Err(r,"邮箱已经存在")
	}
	uid, err := userService.EditSave(req, r)
	if err != nil || uid <= 0 {
		c.Err(r,err.Error())
	}
	c.Succ(r,uid)
}
//删除数据
func (c *Controller) Delete(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}

	rs := userService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		c.Succ(r,rs)
	} else {
		c.Err(r,"删除失败")
	}
}


// 详情
func (c *Controller) Info(r *ghttp.Request) {
	uid := r.GetInt64("userId")
	var userInfo userModel.UserInfo
	// 1.获取角色
	roles, err := roleService.SelectRoleContactVo()
	if err != nil {
		c.Err(r,err.Error())
	}
	// 2.获取岗位信息
	posts, err := postService.SelectListAll(nil)
	if err != nil {
		c.Err(r,err.Error())
	}
	if uid > 0 {
		// 3.获取用户信息
		user, err := userService.GetProfileApi(uid)
		// 4.获取角色Id
		roleIds,err := roleService.SelectRoleListIdByUserId(uid)
		// 5.获取岗位Id
		postIds, err := postService.SelectPostsIdByUserId(uid)
		if err != nil {
			c.Err(r,err.Error())
		}
		userInfo.RoleIds = roleIds
		userInfo.User = user
		userInfo.PostIds = postIds
	}
	userInfo.Roles = roles
	userInfo.Posts = posts
	c.Succ(r,userInfo)
}

//重置密码保存
func (c *Controller) ResetPwdSave(r *ghttp.Request) {
	var req *userModel.ResetPwdReq
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}

	result, err := userService.ResetPassword(req)

	if err != nil || !result {
		c.Err(r,err.Error())
	} else {
		c.Succ(r)
	}
}

//导出
func (c *Controller)Export(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,model.Buniss_Export)
	var req *userModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	url, err := userService.Export(req)

	if err != nil {
		c.Err(r,err.Error())
	}
	c.Succ(r,url)
}
// 修改用户状态
func (c *Controller)ChangeStatus(r *ghttp.Request) {
	var req *userModel.ChangeStatus
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	err := userService.ChangeStatus(req.UserId,req.Status)
	if err != nil {
		c.Err(r,err.Error())
	}
	c.Succ(r)
}
