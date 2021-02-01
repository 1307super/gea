package service

import (
	"context"
	"fmt"
	"gea/app/dao"
	"gea/app/model"
	"gea/app/shared"
	"gea/app/system/admin/internal/define"
	"gea/app/utils/convert"
	"gea/app/utils/excel"
	"gea/app/utils/page"
	"gea/app/utils/random"
	"gea/app/utils/token"
	"gea/library/casbin"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"os"
	"strings"
	"time"
)

// 用户管理服务
var User = &userService{
	AvatarUploadPath:      g.Cfg().GetString(`upload.path`) + `/avatar`,
	AvatarUploadUrlPrefix: `/upload/avatar`,
	UserNoPassTimePrefix:  `user_nopass_`,
	UserLockPrefix:        `user_lock_`,
}

type userService struct {
	AvatarUploadPath      string // 头像上传路径
	AvatarUploadUrlPrefix string // 头像上传对应的URL前缀
	UserNoPassTimePrefix  string // 密码未通过前缀
	UserLockPrefix        string // 用户锁定前缀
}

//func init() {
//	// 启动时创建头像存储目录
//	if !gfile.Exists(User.AvatarUploadPath) {
//		gfile.Mkdir(User.AvatarUploadPath)
//	}
//}

// 执行登录
func (s *userService) Login(ctx context.Context, loginReq *define.UserServiceLoginReq) error {
	if s.CheckLock(loginReq.UserName) {
		return gerror.New("账号已锁定，请30分钟后再试")
	}
	//查询用户信息
	user, err := s.GetUserByUserName(loginReq.UserName)
	if err != nil {
		return err
	}
	if user == nil {
		return gerror.New("用户或密码不正确")
	}
	//校验密码
	pwdtoken := user.LoginName + loginReq.Password + user.Salt
	pwdtoken = gmd5.MustEncryptString(pwdtoken)
	if !strings.EqualFold(user.Password, pwdtoken) {
		return gerror.New("密码错误")
	}
	//获取用户token
	jwtToken, err := token.New().SetUserId(gconv.String(user.UserId)).SetLoginName(user.LoginName).CreateToken()
	if err != nil {
		return gerror.New("token生成失败")
	}

	// 登录成功 移除登陆次数记录
	s.RemovePasswordCounts(loginReq.UserName)
	customCtx := shared.Context.Get(ctx)
	var sysUserExtend *model.SysUserExtend
	if err := gconv.Struct(user,&sysUserExtend); err != nil {
		return gerror.New("登录失败")
	}
	customCtx.Token = jwtToken
	shared.Context.SetUser(ctx, sysUserExtend)
	return nil
}

// 根据账号和密码查询用户信息，一般用于账号密码登录。
// 注意password参数传入的是按照相同加密算法加密过后的密码字符串。
func (s *userService) GetUserByUserName(loginname string) (*model.SysUser, error) {
	return dao.SysUser.Where(g.Map{
		dao.SysUser.Columns.LoginName: loginname,
	}).One()
}

func (s *userService) GetUser(ctx context.Context) (*model.SysUserExtend, error) {
	customCtx := shared.Context.Get(ctx)
	if customCtx != nil && customCtx.Uid != "" {
		// 查询用户
		var user *model.SysUserExtend
		sysUser, err := dao.SysUser.Where(g.Map{
			dao.SysUser.Columns.UserId: customCtx.Uid,
		}).One()
		if err != nil {
			return nil, err
		}
		if sysUser == nil {
			return nil, gerror.New("请登录")
		}
		if err := gconv.Struct(sysUser, &user); err != nil {
			return nil, gerror.New("请登录")
		}
		sysDept,err := dao.SysDept.Where(dao.SysDept.Columns.DeptId,sysUser.DeptId).One()
		if err == nil && sysDept != nil{
			if err := gconv.Struct(sysDept, &user.Dept); err != nil {
				return nil, gerror.New("请登录")
			}
		}
		return user, nil
	}
	return nil, gerror.New("上下文错误")
}

func (s *userService) GetUserInfo(ctx context.Context) (*model.SysUserInfo, error) {
	//user, err := s.GetUser(ctx)
	user := shared.Context.Get(ctx).User
	//if err != nil {
	//	return nil, err
	//}
	userInfo := new(model.SysUserInfo)
	userInfo.User = user
	userInfo.Roles = Role.GetRolePermission(ctx)
	userInfo.Permissions = Menu.GetMenuPermission(ctx)
	return userInfo, nil
}

// 获取当前登录用户信息
func (s *userService) GetProfile(ctx context.Context) (*model.SysUserInfo, error){
	user, err := s.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	userInfo := new(model.SysUserInfo)
	// 获取部门
	var (
		deptResult *model.SysDeptExtend
		roleResult []model.SysRoleFlag
		postResult []model.SysPostFlag
	)
	if err := dao.SysDept.Fields("d.dept_id, d.parent_id, d.ancestors, d.dept_name, d.order_num, d.leader, d.phone, d.email, d.status,(select dept_name from sys_dept where dept_id = d.parent_id) parent_name").Where(dao.SysDept.Columns.DeptId,user.DeptId).Struct(&deptResult);err != nil {
		return nil,gerror.New("未获取到部门")
	}
	user.Dept = deptResult
	// 获取角色
	sysRoleDao := dao.SysRole.As("r")
	sysRoleDao.LeftJoin("sys_user_role ur", "ur.role_id = r.role_id")
	sysRoleDao.LeftJoin("sys_user u", "u.user_id = ur.user_id")
	sysRoleDao.Where("r.del_flag = '0' and u.del_flag = '0'")
	sysRoleDao.Where("ur.user_id = ?", user.UserId)
	if err := sysRoleDao.Structs(&roleResult); err != nil {
		return nil,gerror.New("未获取到角色")
	}
	// 获取岗位
	sysPostDao := dao.SysPost.As("p")
	sysPostDao.LeftJoin("sys_user_post up", "p.post_id = up.post_id")
	sysPostDao.LeftJoin("sys_user u", "u.user_id = up.user_id")
	sysPostDao.Where("u.user_id = ?", user.UserId)
	sysPostDao.Fields("p.post_id, p.post_name, p.post_code")
	if err := sysPostDao.Structs(&postResult); err != nil {
		return nil,gerror.New("未获取到岗位")
	}
	roleNames := garray.New()
	postNames := garray.New()
	for _, role := range roleResult {
		roleNames.Append(role.RoleName)
	}
	for _, postflag := range postResult {
		postNames.Append(postflag.PostName)
	}
	user.Roles = roleResult
	userInfo.User = user
	userInfo.PostGroup = postNames.Join(",")
	userInfo.RoleGroup = roleNames.Join(",")
	return userInfo,nil
}


func (s *userService) GetList(param *define.UserApiSelectPageReq) *define.UserServiceList {
	m := dao.SysUser.As("u").LeftJoin("sys_dept d", "u.dept_id = d.dept_id").Where(fmt.Sprintf("u.%s",dao.SysUser.Columns.DelFlag),"0")
	if param != nil {
		if param.LoginName != "" {
			m = m.Where("u.login_name like ?", "%"+param.LoginName+"%")
		}

		if param.Phonenumber != "" {
			m = m.Where("u.phonenumber like ?", "%"+param.Phonenumber+"%")
		}

		if param.Status != "" {
			m = m.Where("u.status = ?", param.Status)
		}

		if param.BeginTime != "" {
			m = m.Where("date_format(u.create_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}

		if param.EndTime != "" {
			m = m.Where("date_format(u.create_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
		}

		if param.DeptId != 0 {
			m = m.Where("(u.dept_id = ? OR u.dept_id IN ( SELECT t.dept_id FROM sys_dept t WHERE FIND_IN_SET (?,ancestors) ))", param.DeptId, param.DeptId)
		}
	}
	total, err := m.Count()
	if err != nil {
		return nil
	}
	m = m.Fields("u.user_id, u.dept_id, u.login_name, u.user_name, u.email, u.avatar, u.phonenumber, u.password,u.sex, u.salt, u.status, u.del_flag, u.login_ip, u.login_date, u.create_by, u.create_time, u.remark,d.dept_name, d.leader")
	page := page.CreatePaging(param.PageNum, param.PageSize, total)
	m = m.Order(param.SortName + " " + param.SortOrder)
	m = m.Limit(page.StartNum, page.Pagesize)
	result := &define.UserServiceList{
		Page:  page.PageNum,
		Size:  page.Pagesize,
		Total: page.Total,
	}
	if err = m.Structs(&result.List); err != nil {
		return nil
	}
	return result
}

func (s *userService) SelectExportList(param *define.UserApiSelectPageReq) (gdb.Result,error) {
	m := dao.SysUser.As("u").Fields(model.UserListItem{})
	if param != nil {
		if param.LoginName != "" {
			m = m.Where("u.login_name like ?", "%"+param.LoginName+"%")
		}

		if param.Phonenumber != "" {
			m = m.Where("u.phonenumber like ?", "%"+param.Phonenumber+"%")
		}

		if param.Status != "" {
			m = m.Where("u.status = ?", param.Status)
		}

		if param.BeginTime != "" {
			m = m.Where("date_format(u.create_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}

		if param.EndTime != "" {
			m = m.Where("date_format(u.create_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
		}

		if param.DeptId != 0 {
			m = m.Where("(u.dept_id = ? OR u.dept_id IN ( SELECT t.dept_id FROM sys_dept t WHERE FIND_IN_SET (?,ancestors) ))", param.DeptId, param.DeptId)
		}
	}
	result, err := m.M.All()
	return result,err
}

// 创建
func (s *userService) Create(ctx context.Context, req *define.UserApiCreateReq) (int64, error) {
	// 校验
	if s.CheckLoginName(req.LoginName) {
		return 0, gerror.New("登陆名已经存在")
	}
	if s.CheckPhoneUniqueAll(req.Phonenumber) {
		return 0, gerror.New("手机号码已经存在")
	}
	if s.CheckEmailUniqueAll(req.Email) {
		return 0, gerror.New("邮箱已经存在")
	}

	// 获取管理员信息
	adminUser := shared.Context.Get(ctx).User
	var user model.SysUser
	user.LoginName = req.LoginName
	//生成密码
	newSalt := random.GenerateSubId(6)
	newToken := req.LoginName + req.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)
	user.Salt = newSalt
	user.Password = newToken
	user.CreateTime = gtime.Now()
	user.CreateBy = adminUser.LoginName
	user.DelFlag = "0"
	var editReq *define.UserApiUpdateReq
	gconv.Struct(req,&editReq)
	return s.save(&user,editReq)
}

func (s *userService) Update(ctx context.Context, req *define.UserApiUpdateReq) (int64, error) {

	if s.CheckPhoneUnique(req.UserId, req.Phonenumber) {
		return 0, gerror.New("手机号码已经存在")
	}
	if s.CheckEmailUnique(req.UserId, req.Email) {
		return 0, gerror.New("邮箱已经存在")
	}
	adminUser := shared.Context.Get(ctx).User
	user, err := dao.SysUser.FindOne(dao.SysUser.Columns.UserId, req.UserId)
	if err != nil || user == nil {
		return 0, err
	}

	user.UpdateBy = adminUser.LoginName
	user.UpdateTime = gtime.Now()
	return s.save(user,req)
}

func (s *userService) save(user *model.SysUser, req *define.UserApiUpdateReq)(int64,error) {
	user.UserName = req.UserName
	user.Email = req.Email
	user.Phonenumber = req.Phonenumber
	user.Status = req.Status
	user.Sex = req.Sex
	user.DeptId = req.DeptId
	user.Remark = req.Remark

	tx, err := g.DB().Begin()
	if err != nil {
		return 0, err
	}
	result, err := dao.SysUser.TX(tx).Data(user).Save()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if user.UserId == 0 {
		// 新增
		userId,err := result.LastInsertId()
		if err != nil {
			return 0, gerror.New("新增失败")
		}
		user.UserId = userId
	}

	//增加岗位数据
	if req.PostIds != "" {
		postIds := convert.ToInt64Array(req.PostIds, ",")
		userPosts := make([]model.SysUserPost, 0)
		for i := range postIds {
			if postIds[i] > 0 {
				var userPost model.SysUserPost
				userPost.UserId = user.UserId
				userPost.PostId = postIds[i]
				userPosts = append(userPosts, userPost)
			}
		}
		if len(userPosts) > 0 {
			dao.SysUserPost.TX(tx).Delete(dao.SysUserPost.Columns.UserId, user.UserId)
			_, err := dao.SysUserPost.TX(tx).Data(userPosts).Insert()
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}

	}

	//增加角色数据
	if req.RoleIds != "" {
		roleIds := convert.ToInt64Array(req.RoleIds, ",")
		userRoles := make([]model.SysUserRole, 0)
		for i := range roleIds {
			if roleIds[i] > 0 {
				var userRole model.SysUserRole
				userRole.UserId = user.UserId
				userRole.RoleId = roleIds[i]
				userRoles = append(userRoles, userRole)
			}
		}
		if len(userRoles) > 0 {
			dao.SysUserRole.TX(tx).Delete(dao.SysUserRole.Columns.UserId, user.UserId)
			_, err := dao.SysUserRole.TX(tx).Data(userRoles).Insert()
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	}
	go s.ReloadUserRole(user.LoginName)
	return 1, tx.Commit()
}

func (s *userService) Delete(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := dao.SysUser.Where(fmt.Sprintf("%s in(?)", dao.SysUser.Columns.UserId), idarr).Data(g.Map{
		dao.SysUser.Columns.DelFlag: "2",
	}).Update()
	if err != nil {
		return 0
	}
	nums, _ := result.RowsAffected()
	return nums
}

// 获取用户信息
func (s *userService) GetInfo(uid int64) *model.SysUserExtend {
	var user *model.SysUserExtend
	dao.SysUser.Where(dao.SysUser.Columns.UserId, uid).Struct(&user)
	return user
}

//重置用户密码
func (s *userService) ResetPassword(param *define.UserApiResetPwdReq) (bool, error) {
	user, err := dao.SysUser.FindOne(dao.SysUser.Columns.UserId, param.UserId)
	if err != nil {
		return false, gerror.New("用户不存在")
	}
	//新校验密码
	newSalt := random.GenerateSubId(6)
	newToken := user.LoginName + param.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken
	_, err = dao.SysUser.Data(user).Save()
	if err != nil {
		return false, gerror.New("修改失败")
	}
	return true, nil
}


//更新用户信息详情
func (s *userService)UpdateProfile(ctx context.Context, profile *define.UserApiProfileReq) error {
	user, err := s.GetUser(ctx)
	if err != nil {
		return gerror.New("请登录")
	}

	if profile.UserName != "" {
		user.UserName = profile.UserName
	}

	if profile.Email != "" {
		user.Email = profile.Email
	}

	if profile.Phonenumber != "" {
		user.Phonenumber = profile.Phonenumber
	}

	if profile.Sex != "" {
		user.Sex = profile.Sex
	}

	_, err = dao.SysUser.Data(user).Save()
	if err != nil {
		return gerror.New("修改信息失败")
	}
	return nil
}


//更新用户头像
func (s *userService)UpdateAvatar(ctx context.Context,r *define.UserApiAvatarUploadReq) error {
	user, err := s.GetUser(ctx)
	if err != nil {
		return gerror.New("请登录")
	}
	curDir, err := os.Getwd()
	if err != nil {
		return gerror.New("获取路径失败")
	}
	saveDir := curDir + "/public/upload/avatar/" + gconv.String(user.UserId) +"/"

	filename, err := r.Avatarfile.Save(saveDir,true)
	if err != nil {
		return gerror.New(err.Error())
	}
	avatar := "/upload/avatar/" + gconv.String(user.UserId) +"/" + filename
	if avatar != "" {
		user.Avatar = avatar
	}
	_, err = dao.SysUser.Data(user).Save()
	if err != nil {
		return gerror.New("保存数据失败")
	}
	return nil
}

//修改用户密码
func (s *userService)UpdatePassword(ctx context.Context,profile *define.UserApiReSetPasswordReq) error {
	user, err := s.GetUser(ctx)
	if err != nil {
		return gerror.New("请登录")
	}
	if strings.EqualFold(profile.OldPassword, "") {
		return gerror.New("旧密码不能为空")
	}

	if strings.EqualFold(profile.NewPassword, "") {
		return gerror.New("新密码不能为空")
	}

	if strings.EqualFold(profile.NewPassword, profile.OldPassword) {
		return gerror.New("新旧密码不能相同")
	}

	//校验密码
	tokenPwd := user.LoginName + profile.OldPassword + user.Salt
	tokenPwd = gmd5.MustEncryptString(tokenPwd)

	if !strings.EqualFold(tokenPwd, user.Password) {
		return gerror.New("原密码不正确")
	}
	//新校验密码
	newSalt := random.GenerateSubId(6)
	newToken := user.LoginName + profile.NewPassword + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken
	_, err = dao.SysUser.Data(user).Save()
	if err != nil {
		return gerror.New("保存数据失败")
	}
	return nil
}



func (s *userService) ChangeStatus(userId int64, status string) error {
	if IsAdmin(userId) {
		return gerror.New("不能停用超级管理员")
	}
	_, err := dao.SysUser.Where(dao.SysUser.Columns.UserId,userId).Data(g.Map{
		"status": status,
	}).Update()
	if err != nil {
		return err
	}
	return nil
}

// 导出excel
func (s *userService)Export(param *define.UserApiSelectPageReq) (string, error) {
	userList,err := s.SelectExportList(param)
	if err != nil {
		return "",err
	}
	head := []string{"用户名", "呢称", "Email", "电话号码", "性别", "部门", "领导", "状态", "删除标记", "创建人", "创建时间", "备注"}
	key := []string{"login_name", "user_name", "email", "phonenumber", "sex", "dept_name", "leader", "status", "del_flag", "create_by", "create_time", "remark"}
	url, err := excel.DownlaodExcel(head, key, userList)

	if err != nil {
		return "", err
	}

	return url, nil
}

// 检查登陆名是否存在,存在返回true,否则false
func (s *userService) CheckLoginName(loginName string) bool {
	if i, err := dao.SysUser.FindCount(g.Map{
		dao.SysUser.Columns.LoginName: loginName,
	}); err != nil {
		return false
	} else {
		return i == 1
	}
}

//检查邮箱是否已使用
func (s *userService) CheckEmailUnique(userId int64, email string) bool {
	if i, err := dao.SysUser.FindCount(g.Map{
		dao.SysUser.Columns.Email:                        email,
		fmt.Sprintf("%s <>", dao.SysUser.Columns.UserId): userId,
	}); err != nil {
		return false
	} else {
		return i > 1
	}

}

//检查邮箱是否存在,存在返回true,否则false
func (s *userService) CheckEmailUniqueAll(email string) bool {
	if i, err := dao.SysUser.FindCount(g.Map{
		dao.SysUser.Columns.Email: email,
	}); err != nil {
		return false
	} else {
		return i > 1
	}
}

//检查手机号是否已使用,存在返回true,否则false
func (s *userService) CheckPhoneUnique(userId int64, phone string) bool {
	if i, err := dao.SysUser.FindCount(g.Map{
		dao.SysUser.Columns.Phonenumber:                  phone,
		fmt.Sprintf("%s <>", dao.SysUser.Columns.UserId): userId,
	}); err != nil {
		return false
	} else {
		return i > 1
	}
}

//检查手机号是否已使用 ,存在返回true,否则false
func (s *userService) CheckPhoneUniqueAll(phone string) bool {
	if i, err := dao.SysUser.FindCount(g.Map{
		dao.SysUser.Columns.Phonenumber: phone,
	}); err != nil {
		return false
	} else {
		return i > 1
	}
}

//记录密码尝试次数
func (s *userService) SetPasswordCounts(loginName string) int {
	curTimes := 0
	curTimeObj, _ := gcache.Get(s.UserNoPassTimePrefix + loginName)
	if curTimeObj != nil {
		curTimes = gconv.Int(curTimeObj)
	}
	curTimes = curTimes + 1
	gcache.Set(s.UserNoPassTimePrefix+loginName, curTimes, 1*time.Minute)

	if curTimes >= 5 {
		s.Lock(loginName)
	}
	return curTimes
}

//移除密码错误次数
func (s *userService) RemovePasswordCounts(loginName string) {
	gcache.Remove(s.UserNoPassTimePrefix + loginName)
}

//锁定账号
func (s *userService) Lock(loginName string) {
	gcache.Set(s.UserLockPrefix+loginName, true, 30*time.Minute)
}

// 校验是否锁定
func (s *userService) CheckLock(loginName string) bool {
	result := false
	rs, _ := gcache.Get(s.UserLockPrefix + loginName)
	if rs != nil {
		result = true
	}
	return result
}

func (s *userService) ReloadUserRole(userName string) {
	casbin.DeleteRolesForUser(userName)
	s.LoadUserRole(userName)
}

// 加载所有用户角色
func (s *userService) LoadUserRole(userName string) {
	roleForUsers := s.GetUserRolePolicy(userName)
	for _, roleForUser := range roleForUsers {
		casbin.AddRoleForUser(roleForUser.UserName, roleForUser.RoleName)
	}
}

// 获取用户角色关系
func (s *userService)GetUserRolePolicy(userName ...string) []define.UserServiceRoleForUser {
	var roleForUser []define.UserServiceRoleForUser
	m := dao.SysRole.As("r")
	m = m.Fields("distinct role_key as roleName,u.login_name as userName")
	m = m.LeftJoin("sys_user_role ur", "ur.role_id = r.role_id")
	m = m.LeftJoin("sys_user u", "u.user_id = ur.user_id")
	m = m.LeftJoin("sys_dept d", "u.dept_id = d.dept_id")
	m = m.Where("r.del_flag = '0' and u.del_flag = '0'")
	if len(userName) > 0 && userName[0] != "" {
		m = m.Where("u.login_name = ?",userName[0])
	}
	if err := m.Structs(&roleForUser); err != nil {
		return nil
	}
	return roleForUser
}
