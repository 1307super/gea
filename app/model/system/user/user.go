package user

import (
	deptModel "gea/app/model/system/dept"
	"gea/app/model/system/post"
	"gea/app/model/system/role"
	"gea/app/utils/page"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// Fill with you ideas below.

type UserInfo struct {
	User           *UserEntityExtend `json:"user,omitempty"`
	Permissions    *garray.StrArray  `json:"permissions,omitempty"`
	Roles          g.Array           `json:"roles,omitempty"`
	RoleIds        []int64           `json:"roleIds,omitempty"`
	Posts          []post.EntityFlag `json:"posts,omitempty"`
	PostIds        []int64           `json:"postIds,omitempty"`
	PostGroup      string            `json:"postGroup,omitempty"`
	RoleGroup      string            `json:"roleGroup,omitempty"`
}

type UserEntityExtend struct {
	Entity
	Dept  *deptModel.EntityExtend `json:"dept"`
	Roles []role.EntityFlag       `json:"roles"`
}

//修改用户资料请求参数
type ProfileReq struct {
	UserName    string `p:"userName"  v:"required|length:5,30#请输入用户名称|用户名称长度为:min到:max位"`
	Phonenumber string `p:"phonenumber"  v:"required|phone#请输入手机号码|请输入正确的手机号码"`
	Email       string `p:"email"  v:"required|email#请输入邮箱地址|请输入正确的电子邮箱"`
	Sex         string `p:"sex"  v:"required#请输入用户名称"`
}

//修改密码请求参数
type PasswordReq struct {
	OldPassword string `p:"oldPassword" v:"required|length:5,30#请输入旧密码|旧密码长度为:min到:max位"`
	NewPassword string `p:"newPassword" v:"required|length:5,30#请输入旧密码|旧密码长度为:min到:max位"`
	//Confirm     string `p:"confirm" v:"required|length:5,30#请输入确认密码|确认密码长度为:min到:max位"`
}

//重置密码请求参数
type ResetPwdReq struct {
	UserId   int64  `p:"userId"  v:"required|min:1#请输入用户ID|请输入正确的用户ID"`
	Password string `p:"password" v:"required|length:5,30#请输入密码|密码长度为:min到:max位"`
}

//检查email请求参数
type CheckEmailReq struct {
	UserId int64  `p:"userId"  v:"required|min:1#请输入用户ID|请输入正确的用户ID"`
	Email  string `p:"email"  v:"required|email#请输入邮箱地址|请输入正确的电子邮箱"`
}

//检查email请求参数
type CheckEmailAllReq struct {
	Email string `p:"email"  v:"required|email#请输入邮箱地址|请输入正确的电子邮箱"`
}

//检查phone请求参数
type CheckLoginNameReq struct {
	LoginName string `p:"loginName"  v:"required#请输入登陆名"`
}

//检查phone请求参数
type CheckPhoneReq struct {
	UserId      int64  `p:"userId"  v:"required|min:1#请输入用户ID|请输入正确的用户ID"`
	Phonenumber string `p:"phonenumber"  v:"required|phone#请输入手机号码|请输入正确的手机号码"`
}

//检查phone请求参数
type CheckPhoneAllReq struct {
	Phonenumber string `p:"phonenumber"  v:"required|phone#请输入手机号码|请输入正确的手机号码"`
}

//检查密码请求参数
type CheckPasswordReq struct {
	Password string `p:"password"  v:"required#请输入手机号码"`
}

// 修改状态
type ChangeStatus struct {
	UserId int64  `p:"userId" v:"required#主键ID不能为空"`
	Status string `p:"status" v:"required#状态不能为空"`
}

//查询用户列表请求参数
type SelectPageReq struct {
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
}

//用户列表数据结构
type UserListEntity struct {
	UserId      int64       `json:"user_id"`     // 用户ID
	DeptId      int64       `json:"dept_id"`     // 部门ID
	LoginName   string      `json:"login_name"`  // 登录账号
	UserName    string      `json:"user_name"`   // 用户昵称
	Email       string      `json:"email"`       // 用户邮箱
	Avatar      string      `json:"avatar"`      // 头像路径
	Phonenumber string      `json:"phonenumber"` // 手机号码
	Password    string      `json:"password"`    // 密码
	Sex         string      `json:"sex"`         // 用户性别（0男 1女 2未知）
	Salt        string      `json:"salt"`        // 盐加密
	Status      string      `json:"status"`      // 帐号状态（0正常 1停用）
	DelFlag     string      `json:"del_flag"`    // 删除标志（0代表存在 2代表删除）
	LoginIp     string      `json:"login_ip"`    // 最后登陆IP
	LoginDate   *gtime.Time `json:"login_date"`  // 最后登陆时间
	CreateBy    string      `json:"create_by"`   // 创建者
	CreateTime  *gtime.Time `json:"create_time"` // 创建时间
	Remark      string      `json:"remark"`      // 备注
	DeptName    string      `json:"dept_name"`   // 部门名称
	Leader      string      `json:"leader"`      // 负责人
}

//新增用户资料请求参数
type AddReq struct {
	UserName    string `p:"userName"  v:"required|length:1,30#请输入用户名称|用户名称长度为:min到:max位"`
	Phonenumber string `p:"phonenumber"  v:"required|phone#请输入手机号码|请输入正确的手机号码"`
	Email       string `p:"email"  v:"required|email#请输入邮箱地址|请输入正确的电子邮箱"`
	LoginName   string `p:"loginName"  v:"required#请输入登陆名"`
	Password    string `p:"password"  v:"required|length:5,30#请输入密码|用户密码长度为:min到:max位"`
	DeptId      int64  `p:"deptId" v:"required#请选择部门"`
	Sex         string `p:"sex"  v:"required#请输入用户名称"`
	Status      string `p:"status"`
	RoleIds     string `p:"roleIds"`
	PostIds     string `p:"postIds"`
	Remark      string `p:"remark"`
}

//新增用户资料请求参数
type EditReq struct {
	UserId      int64  `p:"userId" v:"required#用户ID不能为空"`
	UserName    string `p:"userName"  v:"required|length:5,30#请输入用户名称|用户名称长度为:min到:max位"`
	Phonenumber string `p:"phonenumber"  v:"required|phone#请输入手机号码|请输入正确的手机号码"`
	Email       string `p:"email"  v:"required|email#请输入邮箱地址|请输入正确的电子邮箱"`
	DeptId      int64  `p:"deptId" v:"required#请选择部门"`
	Sex         string `p:"sex"  v:"required#请输入用户名称"`
	Status      string `p:"status"`
	RoleIds     string `p:"roleIds"`
	PostIds     string `p:"postIds"`
	Remark      string `p:"remark"`
}


// 角色与用户关联关系
type RoleForUser struct {
	RoleName string
	UserName string
}

// 根据条件分页查询用户列表
func SelectPageList(param *SelectPageReq) ([]UserListEntity, *page.Paging, error) {
	db, err := gdb.Instance()
	if err != nil {
		return nil, nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_user u").LeftJoin("sys_dept d", "u.dept_id = d.dept_id")
	model.Where(" u.del_flag = '0' ")

	if param != nil {
		if param.LoginName != "" {
			model.Where("u.login_name like ?", "%"+param.LoginName+"%")
		}

		if param.Phonenumber != "" {
			model.Where("u.phonenumber like ?", "%"+param.Phonenumber+"%")
		}

		if param.Status != "" {
			model.Where("u.status = ?", param.Status)
		}

		if param.BeginTime != "" {
			model.Where("date_format(u.create_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(u.create_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
		}

		if param.DeptId != 0 {
			model.Where("(u.dept_id = ? OR u.dept_id IN ( SELECT t.dept_id FROM sys_dept t WHERE FIND_IN_SET (?,ancestors) ))", param.DeptId, param.DeptId)
		}
	}

	total, err := model.Count()

	if err != nil {
		return nil, nil, gerror.New("读取行数失败")
	}

	page := page.CreatePaging(param.PageNum, param.PageSize, total)

	model.Fields("u.user_id, u.dept_id, u.login_name, u.user_name, u.email, u.avatar, u.phonenumber, u.password,u.sex, u.salt, u.status, u.del_flag, u.login_ip, u.login_date, u.create_by, u.create_time, u.remark,d.dept_name, d.leader")

	model.Order(param.SortName + " " + param.SortOrder)

	model.Limit(page.StartNum, page.Pagesize)

	var result []UserListEntity

	err = model.Structs(&result)
	return result, page, err
}

// 导出excel
func SelectExportList(param *SelectPageReq) (gdb.Result, error) {
	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_user u").LeftJoin("sys_dept d", "u.dept_id = d.dept_id")

	model.Where(" u.del_flag = '0' ")

	if param != nil {
		if param.LoginName != "" {
			model.Where("u.login_name like ?", "%"+param.LoginName+"%")
		}

		if param.Phonenumber != "" {
			model.Where("u.phonenumber like ?", "%"+param.Phonenumber+"%")
		}

		if param.Status != "" {
			model.Where("u.status = ?", param.Status)
		}

		if param.BeginTime != "" {
			model.Where("date_format(u.create_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(u.create_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
		}

		if param.DeptId != 0 {
			model.Where("(u.dept_id = ? OR u.dept_id IN ( SELECT t.dept_id FROM sys_dept t WHERE FIND_IN_SET (?,ancestors) ))", param.DeptId)
		}
	}

	//用户名  呢称 Email 电话号码 性别 部门 领导  状态 删除标记 创建人 创建时间 备注
	model.Fields("u.login_name, u.user_name, u.email, u.phonenumber, u.sex,d.dept_name, d.leader,  u.status, u.del_flag, u.create_by, u.create_time, u.remark")

	result, err := model.All()
	return result, err
}

// 根据条件分页查询已分配用户角色列表
func SelectAllocatedList(roleId int64, loginName, phonenumber string) ([]Entity, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_user u")
	model.LeftJoin("sys_dept d", "u.dept_id = d.dept_id")
	model.LeftJoin("sys_user_role ur", " u.user_id = ur.user_id")
	model.LeftJoin("sys_role r", "r.role_id = ur.role_id")

	model.Where("u.del_flag =?", 0)
	model.Where("r.role_id = ?", roleId)

	if loginName != "" {
		model.Where("u.login_name like ?", "%"+loginName+"%")
	}

	if phonenumber != "" {
		model.Where("u.phonenumber like ?", "%"+phonenumber+"%")
	}

	var result []Entity
	model.Structs(&result)
	return result, nil
}

// 根据条件分页查询未分配用户角色列表
func SelectUnallocatedList(roleId int64, loginName, phonenumber string) ([]Entity, error) {
	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_user u")
	model.LeftJoin("sys_dept d", "u.dept_id = d.dept_id")
	model.LeftJoin("sys_user_role ur", "u.user_id = ur.user_id")
	model.LeftJoin("sys_role r", "r.role_id = ur.role_id")

	model.Where("u.user_id not in (select u.user_id from sys_user u inner join sys_user_role ur on u.user_id = ur.user_id and ur.role_id = ?)", roleId)

	if loginName != "" {
		model.Where("u.login_name like ?", "%"+loginName+"%")
	}

	if phonenumber != "" {
		model.Where("u.phonenumber like ?", "%"+phonenumber+"%")
	}

	model.Fields("distinct u.user_id, u.dept_id, u.login_name, u.user_name, u.email, u.avatar, u.phonenumber, u.status, u.create_time")

	var result []Entity
	err = model.Structs(&result)
	return result, err
}

//检查邮箱是否已使用
func CheckEmailUnique(userId int64, email string) bool {
	rs, err := FindCount("email=? AND user_id<>?", email, userId)
	if err != nil {
		return false
	}

	if rs > 0 {
		return true
	} else {
		return false
	}
}

//检查邮箱是否存在,存在返回true,否则false
func CheckEmailUniqueAll(email string) bool {
	rs, err := FindCount("email=?", email)
	if err != nil {
		return false
	}

	if rs > 0 {
		return true
	} else {
		return false
	}
}

//检查手机号是否已使用,存在返回true,否则false
func CheckPhoneUnique(userId int64, phone string) bool {
	rs, err := FindCount("phonenumber = ? AND user_id<>?", phone, userId)
	if err != nil {
		return false
	}

	if rs > 0 {
		return true
	} else {
		return false
	}
}

//检查手机号是否已使用 ,存在返回true,否则false
func CheckPhoneUniqueAll(phone string) bool {
	rs, err := FindCount("phonenumber = ?", phone)
	if err != nil {
		return false
	}

	if rs > 0 {
		return true
	} else {
		return false
	}
}

//根据登陆名查询用户信息
func SelectUserByLoginName(loginName string) (*Entity, error) {
	return FindOne("login_name", loginName)
}

//根据登陆名查询用户信息
func SelectUserByUid(uid int64) (*UserEntityExtend, error) {
	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}
	model := db.Table("sys_user")
	model.Where("user_id =?", uid)
	var result *UserEntityExtend
	err = model.Struct(&result)
	return result, err
}

//根据手机号查询用户信息
func SelectUserByPhoneNumber(phonenumber string) (*Entity, error) {
	return FindOne("phonenumber", phonenumber)
}

// 获取用户角色关系
func GetUserRolePolicy(userName ...string) []RoleForUser {
	var roleForUser []RoleForUser
	db, err := gdb.Instance()
	if err != nil {
		return roleForUser
	}
	model := db.Table("sys_role r")
	model.Fields("distinct role_key as roleName,u.login_name as userName")
	model.LeftJoin("sys_user_role ur", "ur.role_id = r.role_id")
	model.LeftJoin("sys_user u", "u.user_id = ur.user_id")
	model.LeftJoin("sys_dept d", "u.dept_id = d.dept_id")
	model.Where("r.del_flag = '0' and u.del_flag = '0'")
	if len(userName) > 0 && userName[0] != "" {
		model.Where("u.login_name = ?",userName[0])
	}
	model.Structs(&roleForUser)
	return roleForUser
}
