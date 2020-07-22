package user

import (
	userModel "gea/app/model/system/user"
	userService "gea/app/service/system/user"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"os"
)

//用户资料页面
func (c *Controller) Profile(r *ghttp.Request) {
	uid := r.GetInt64("jwtUid")
	userInfo, err := userService.GetProfile(uid)
	if err != nil || userInfo == nil{
		 c.Err(r,"未查询到该用户")
	}
	 c.Succ(r,userInfo)
}

//用户资料页面
func (c *Controller) GetInfo(r *ghttp.Request) {
	uid := r.GetInt64("jwtUid")
	userInfo, err := userService.GetUserInfo(uid)
	if err != nil {
		 c.Err(r,err.Error())
	}
	 c.Succ(r,userInfo)
}

//修改用户信息
func (c *Controller) Update(r *ghttp.Request) {
	var req *userModel.ProfileReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}
	err := userService.UpdateProfile(req, r)

	if err != nil {
		 c.Err(r,err.Error())
	} else {
		 c.Succ(r,)
	}
}

//修改用户密码
func (c *Controller) UpdatePassword(r *ghttp.Request) {
	var req *userModel.PasswordReq
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}

	err := userService.UpdatePassword(req, r)

	if err != nil {
		 c.Err(r,err.Error())
	} else {
		 c.Succ(r,)
	}
}

//保存头像
func (c *Controller) UpdateAvatar(r *ghttp.Request) {
	user,_ := userService.GetProfileApi(r.GetInt64("jwtUid"))

	curDir, err := os.Getwd()
	if err != nil {
		 c.Err(r,err.Error())
	}

	saveDir := curDir + "/public/upload/avatar/" + gconv.String(user.UserId) +"/"

	files := r.GetUploadFile("avatarfile")
	if files == nil{
		 c.Err(r,"没有获取到上传文件")
	}

	filename, err := files.Save(saveDir,true)
	if err != nil {
		 c.Err(r,err.Error())
	}

	avatar := "/upload/avatar/" + gconv.String(user.UserId) +"/" + filename

	err = userService.UpdateAvatar(avatar, r)

	if err != nil {
		 c.Err(r,err.Error())
	} else {
		 c.Succ(r,avatar)
	}
}
