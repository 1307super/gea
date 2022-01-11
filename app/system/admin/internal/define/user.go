package define

import (
	"gea/app/model"
	"github.com/gogf/gf/net/ghttp"
)

// ========== api =============
// api接口登录参数
type UserApiLoginReq struct {
	UserName     string `p:"username"  v:"required|length:5,30#请输入账号|账号长度为:min到:max位"`
	Password     string `p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
	ValidateCode string `p:"validateCode" v:"required|length:4,30#请输入验证码|验证码长度不够"`
	IdKey        string `p:"idkey" v:"required|length:4,30#请输入验证码id|验证码id长度不够"`
}

//查询用户列表请求参数
type UserApiSelectPageReq struct {
	LoginName   string `p:"loginName"`     //登陆名
	Status      string `p:"status"`        //状态
	Phonenumber string `p:"phonenumber"`   //手机号码
	BeginTime   string `p:"beginTime"`     //数据范围
	EndTime     string `p:"endTime"`       //开始时间
	DeptId      int64  `p:"deptId"`        //结束时间
	PageNum     int    `p:"pageNum"`       //当前页码
	PageSize    int    `p:"pageSize"`      //每页数
	SortName    string `p:"orderByColumn"` //排序字段
	SortOrder   string `p:"isAsc"`         //排序方式
	Export int
}

type UserApiCreateBase struct {
	UserName    string `p:"user_name"  v:"required|length:1,30#请输入用户名称|用户名称长度为:min到:max位"`
	Phonenumber string `p:"phonenumber"  v:"required|phone#请输入手机号码|请输入正确的手机号码"`
	Email       string `p:"email"  v:"required|email#请输入邮箱地址|请输入正确的电子邮箱"`
	DeptId      int64  `p:"deptId" v:"required#请选择部门"`
	Sex         string `p:"sex"  v:"required#请输入用户名称"`
	Status      string `p:"status"`
	RoleIds     string `p:"roleIds"`
	PostIds     string `p:"postIds"`
	Remark      string `p:"remark"`
}
//新增用户资料请求参数
type UserApiCreateReq struct {
	UserApiCreateBase
	LoginName   string `p:"login_name"  v:"required|length:5,30#请输入登陆名#请输入登陆名|登陆名长度为:min到:max位"`
	Password    string `p:"password"  v:"required|length:5,30#请输入密码|用户密码长度为:min到:max位"`
}
//修改用户资料请求参数
type UserApiUpdateReq struct {
	UserId      int64  `p:"userId" v:"required#用户ID不能为空"`
	UserApiCreateBase
}

// API执行删除内容
type UserApiDeleteReq struct {
	Ids string `p:"ids"  v:"required#请选择要删除的数据记录"`
}

//重置密码请求参数
type UserApiResetPwdReq struct {
	UserId   int64  `p:"userId"  v:"required|min:1#请输入用户ID|请输入正确的用户ID"`
	Password string `p:"password" v:"required|length:5,30#请输入密码|密码长度为:min到:max位"`
}

//修改用户资料请求参数
type UserApiProfileReq struct {
	UserName    string `p:"user_name"  v:"required|length:5,30#请输入用户名称|用户名称长度为:min到:max位"`
	Phonenumber string `p:"phonenumber"  v:"required|phone#请输入手机号码|请输入正确的手机号码"`
	Email       string `p:"email"  v:"required|email#请输入邮箱地址|请输入正确的电子邮箱"`
	Sex         string `p:"sex"  v:"required#请输入用户名称"`
}

//修改密码请求参数
type UserApiReSetPasswordReq struct {
	OldPassword string `p:"oldPassword" v:"required|length:5,30#请输入旧密码|旧密码长度为:min到:max位"`
	NewPassword string `p:"newPassword" v:"required|length:5,30#请输入旧密码|旧密码长度为:min到:max位"`
	//Confirm     string `p:"confirm" v:"required|length:5,30#请输入确认密码|确认密码长度为:min到:max位"`
}

// 管理员重置密码请求参数
type UserApiAdminResetPwdReq struct {
	UserId   int64  `p:"userId"  v:"required|min:1#请输入用户ID|请输入正确的用户ID"`
	Password string `p:"password" v:"required|length:5,30#请输入密码|密码长度为:min到:max位"`
}

// 头像上传
type UserApiAvatarUploadReq struct {
	Avatarfile       *ghttp.UploadFile `json:"avatarfile"`// 上传文件对象
}

// 修改状态
type UserApiChangeStatus struct {
	UserId int64  `p:"userId" v:"required#主键ID不能为空"`
	Status string `p:"status" v:"required#状态不能为空"`
}

// ======= service =========
// service登录参数
type UserServiceLoginReq struct {
	UserName string `p:"username"  v:"required|length:5,30#请输入账号|账号长度为:min到:max位"`
	Password string `p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
}
// 查询列表返回值
type UserServiceList struct {
	List  []model.UserListItem `json:"list"`
	Page  int                  `json:"page"`
	Size  int                  `json:"size"`
	Total int                  `json:"total"`
}
// 角色与用户关联关系
type UserServiceRoleForUser struct {
	RoleName string
	UserName string
}
