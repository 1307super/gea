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
	"gea/app/utils/ip"
	"gea/app/utils/page"
	"gea/app/utils/random"
	"gea/app/utils/response"
	"gea/library/casbin"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/mssola/user_agent"
	"os"
	"strings"
	"time"
)

// 用户管理服务
var User = &userService{
	AvatarUploadPath:       `/static/upload/avatar/`,
	AvatarUploadUrlPrefix:  `/static/upload/avatar/`,
	UserNoPassTimePrefix:   `user_nopass_`,
	UserLockPrefix:         `user_lock_`,
	UserCachePrefix:        `user_cache_`,
	UserProfileCachePrefix: `user_profile_cache_`,
}

type userService struct {
	AvatarUploadPath       string // 头像上传路径
	AvatarUploadUrlPrefix  string // 头像上传对应的URL前缀
	UserNoPassTimePrefix   string // 密码未通过前缀
	UserLockPrefix         string // 用户锁定前缀
	UserCachePrefix        string // 用户缓存前缀
	UserProfileCachePrefix string // 用户profile缓存前缀
}

//func init() {
//	// 启动时创建头像存储目录
//	if !gfile.Exists(User.AvatarUploadPath) {
//		gfile.Mkdir(User.AvatarUploadPath)
//	}
//}

// 执行登录
func (s *userService) Login(r *ghttp.Request) (string, interface{}) {
	var (
		req    *define.UserApiLoginReq
		user   *model.SysUser
		status = "0"
		msg    string
	)
	defer func() {
		// 记录登录日志
		if status == "0" {
			errTimes := s.SetPasswordCounts(req.UserName)
			having := 5 - errTimes
			go Logininfor.Create(status, req.UserName, r.GetClientIp(), r.Header.Get("User-Agent"), msg)

			response.ErrorResp(r).SetCode(500).SetMsg("账号或密码不正确,还有" + gconv.String(having) + "次之后账号将锁定").WriteJsonExit()
		}
		if status == "2" {
			response.ErrorResp(r).SetCode(500).SetMsg(msg).WriteJsonExit()
		}
		go Logininfor.Create(status, req.UserName, r.GetClientIp(), r.Header.Get("User-Agent"), msg)
	}()
	//获取参数
	if err := r.Parse(&req); err != nil {
		msg = err.Error()
		r.Exit()
	}

	if s.CheckLock(req.UserName) {
		status = "2"
		msg = "账号已锁定，请30分钟后再试"
		r.Exit()
	}
	//查询用户信息
	user, err := s.GetUserByUserName(req.UserName)
	if err != nil {
		msg = err.Error()
		r.Exit()
	}
	if user == nil {
		msg = "用户或密码不正确"
		r.Exit()
	}
	//校验密码
	pwdtoken := user.LoginName + req.Password + user.Salt
	pwdtoken = gmd5.MustEncryptString(pwdtoken)
	if !strings.EqualFold(user.Password, pwdtoken) {
		msg = "用户或密码不正确"
		r.Exit()
	}
	// 获取用户信息
	customCtx := shared.Context.Get(r.Context())
	customCtx.Uid = user.UserId
	shared.Context.Init(r, customCtx)
	userInfo, err := s.GetProfile(r.Context())
	if err != nil || userInfo == nil {
		msg = "用户不存在"
		r.Exit()
	}
	status = "1"
	msg = "登录成功"
	return user.LoginName, userInfo
}

// 登录之后
func (s *userService) LoginAfter(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		response.ErrorResp(r).SetCode(500).SetMsg(respData.Msg).WriteJsonExit()
	}
	token := respData.GetString("token")
	r.Header.Set("Authorization", "Bearer "+token)

	tokenInfo := shared.GfAdminToken.GetTokenData(r)
	if !tokenInfo.Success() {
		response.ErrorResp(r).SetCode(500).SetMsg("登录过期，请重新登录").WriteJsonExit()
	}
	var userInfo *model.SysUserInfo
	if err := gjson.DecodeTo(tokenInfo.GetString("data"), &userInfo); err == nil {
		// 登录成功 移除登陆次数记录
		s.RemovePasswordCounts(userInfo.UserExtend.LoginName)
		userInfo.UserExtend.LoginIp = r.GetClientIp()
		userInfo.UserExtend.LoginDate = gtime.Now()
		dao.SysUser.Data(g.Map{
			dao.SysUser.Columns.LoginIp:   r.GetClientIp(),
			dao.SysUser.Columns.LoginDate: gtime.Now(),
		}).Where(dao.SysUser.Columns.UserId, userInfo.UserExtend.UserId).Update()

		// 下线另一个用户
		if !g.Cfg().GetBool("gtoken.MultiLogin") {
			Online.OffLine(userInfo.UserExtend.LoginName)
		}

		// 记录在线状态
		userAgent := r.Header.Get("User-Agent")
		ua := user_agent.New(userAgent)
		browser, _ := ua.Browser()
		var userOnline model.SysUserOnline
		userOnline.Token = token
		userOnline.LoginName = userInfo.UserExtend.LoginName
		userOnline.Browser = browser
		userOnline.Os = ua.OS()
		userOnline.Ipaddr = r.GetClientIp()
		userOnline.ExpireTime = g.Cfg().GetInt("gtoken.MaxRefresh")
		userOnline.StartTimestamp = gtime.Now()
		userOnline.LastAccessTime = gtime.Now()
		userOnline.Status = "on_line"
		userOnline.LoginLocation = ip.GetCityByIp(r.GetClientIp())
		dao.SysUserOnline.Save(userOnline)
		response.ErrorResp(r).SetCode(0).SetMsg("登录成功").SetData(g.Map{
			"token": token,
		}).WriteJsonExit()
	} else {
		response.ErrorResp(r).SetCode(500).SetMsg("登录过期，请重新登录").WriteJsonExit()
	}
}

func (s *userService) Logout(r *ghttp.Request) bool {
	tokenInfo := shared.GfAdminToken.GetTokenData(r)
	if !tokenInfo.Success() {
		return true
	}
	var userInfo *model.SysUser
	if err := gjson.DecodeTo(tokenInfo.GetString("data"), &userInfo); err != nil {
		return true
	}
	s.RemoveUserCache(userInfo.UserId)
	// 清除在线状态
	dao.SysUserOnline.Data(g.Map{
		dao.SysUserOnline.Columns.Status: "off_line",
	}).Where(dao.SysUserOnline.Columns.Token, tokenInfo.GetString("token")).Update()
	return true
}

// 根据账号和密码查询用户信息，一般用于账号密码登录。
// 注意password参数传入的是按照相同加密算法加密过后的密码字符串。
func (s *userService) GetUserByUserName(loginname string) (*model.SysUser, error) {
	return dao.SysUser.Where(g.Map{
		dao.SysUser.Columns.LoginName: loginname,
	}).One()
}

func (s *userService) GetUser(ctx context.Context) (*model.SysUser, error) {
	customCtx := shared.Context.Get(ctx)
	if customCtx != nil && customCtx.Uid != 0 {
		// 查询用户
		var user *model.SysUser
		cache := s.GetUserCache(s.UserCachePrefix, customCtx.Uid)
		if !cache.IsEmpty() {
			if err := cache.Struct(&user); err == nil {
				return user, nil
			}
		}
		err := dao.SysUser.Where(g.Map{
			dao.SysUser.Columns.UserId: customCtx.Uid,
		}).Struct(&user)
		if err != nil {
			return nil, err
		}
		s.SetUserCache(s.UserCachePrefix, customCtx.Uid, user)
		return user, nil
	}
	return nil, gerror.New("上下文错误")
}

func (s *userService) GetUserInfo(ctx context.Context) (*model.SysUserInfo, error) {
	return s.GetProfile(ctx)
}

// 获取当前登录用户信息
func (s *userService) GetProfile(ctx context.Context) (*model.SysUserInfo, error) {
	// 从缓存中获取
	user, err := s.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	userInfo := new(model.SysUserInfo)
	cache := s.GetUserCache(s.UserProfileCachePrefix, user.UserId)
	if !cache.IsEmpty() {
		if err = cache.Struct(&userInfo); err == nil {
			return userInfo, nil
		}
	}
	// 获取部门
	var (
		userExtend *model.SysUserExtend
		roleResult []model.SysRoleFlag
		postResult []model.SysPostFlag
	)
	if err := gconv.Struct(user, &userExtend); err != nil {
		return nil, gerror.New("请登录")
	}
	sysDept, err := dao.SysDept.Where(dao.SysDept.Columns.DeptId, user.DeptId).One()
	if err == nil && sysDept != nil {
		if err := gconv.Struct(sysDept, &userExtend.Dept); err != nil {
			return nil, gerror.New("请登录")
		}
	}
	// 获取角色
	sysRoleDao := dao.SysRole.As("r").
		LeftJoin("sys_user_role ur", "ur.role_id = r.role_id").
		LeftJoin("sys_user u", "u.user_id = ur.user_id").
		Where("r.del_flag = '0' and u.del_flag = '0'").
		Where("ur.user_id = ?", user.UserId)
	if err := sysRoleDao.Structs(&roleResult); err != nil {
		return nil, gerror.New("未获取到角色")
	}
	// 获取岗位
	sysPostDao := dao.SysPost.As("p").
		LeftJoin("sys_user_post up", "p.post_id = up.post_id").
		LeftJoin("sys_user u", "u.user_id = up.user_id").
		Where("u.user_id = ?", user.UserId).
		Fields("p.post_id, p.post_name, p.post_code")
	if err := sysPostDao.Structs(&postResult); err != nil {
		return nil, gerror.New("未获取到岗位")
	}
	roleNames := garray.New()
	postNames := garray.New()
	for _, role := range roleResult {
		roleNames.Append(role.RoleName)
	}
	for _, postflag := range postResult {
		postNames.Append(postflag.PostName)
	}
	userExtend.Roles = roleResult
	userInfo.UserExtend = userExtend
	userInfo.PostGroup = postNames.Join(",")
	userInfo.RoleGroup = roleNames.Join(",")
	userInfo.Roles = Role.GetRolePermission(ctx)
	userInfo.Permissions = Menu.GetMenuPermission(ctx)
	s.SetUserCache(s.UserProfileCachePrefix, user.UserId, userInfo)
	return userInfo, nil
}

func (s *userService) GetList(ctx context.Context, param *define.UserApiSelectPageReq) *define.UserServiceList {
	m := dao.SysUser.As("u").
		LeftJoin("sys_dept d", "u.dept_id = d.dept_id").
		Where(fmt.Sprintf("u.%s", dao.SysUser.Columns.DelFlag), "0")
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
	// 获取资源权限
	dataScope := DataScopeFilter(ctx, "d", "u")
	if dataScope != "" {
		m = m.Where(dataScope)
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

func (s *userService) SelectExportList(param *define.UserApiSelectPageReq) (gdb.Result, error) {
	m := dao.SysUser.As("u").LeftJoin("sys_dept d", "u.dept_id = d.dept_id").Fields("u.login_name, u.user_name, u.email, u.phonenumber, u.sex,d.dept_name, d.leader,  u.status, u.del_flag, u.create_by, u.create_time, u.remark")
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
	return result, err
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
	adminUser.UserExtend.LoginName = req.LoginName
	//生成密码
	newSalt := random.GenerateSubId(6)
	newToken := req.LoginName + req.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)
	user.Salt = newSalt
	user.Password = newToken
	user.CreateTime = gtime.Now()
	user.CreateBy = adminUser.UserExtend.LoginName
	user.DelFlag = "0"
	var editReq *define.UserApiUpdateReq
	gconv.Struct(req, &editReq)
	return s.save(&user, editReq)
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

	user.UpdateBy = adminUser.UserExtend.LoginName
	user.UpdateTime = gtime.Now()
	return s.save(user, req)
}

func (s *userService) save(user *model.SysUser, req *define.UserApiUpdateReq) (int64, error) {
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
		userId, err := result.LastInsertId()
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
	defer func() {
		// 重置权限
		go s.ReloadUserRole(user.LoginName)
		// 重置缓存
		s.RemoveUserCache(user.UserId)
		// 重置菜单
		gcache.Remove(MENU_CACHE + gconv.String(user.UserId))
	}()

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
	s.RemoveUserCache(user.UserId)
	return true, nil
}

//更新用户信息详情
func (s *userService) UpdateProfile(ctx context.Context, profile *define.UserApiProfileReq) error {
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

	s.RemoveUserCache(user.UserId)
	return nil
}

//更新用户头像
func (s *userService) UpdateAvatar(ctx context.Context, r *define.UserApiAvatarUploadReq) (string, error) {
	user, err := s.GetUser(ctx)
	if err != nil {
		return "",gerror.New("请登录")
	}
	curDir, err := os.Getwd()
	if err != nil {
		return "",gerror.New("获取路径失败")
	}
	saveDir := curDir + "/public/" + s.AvatarUploadPath + gconv.String(user.UserId) + "/"
	filename, err := r.Avatarfile.Save(saveDir, true)
	if err != nil {
		return "",gerror.New(err.Error())
	}
	avatar := s.AvatarUploadPath + gconv.String(user.UserId) + "/" + filename
	if avatar != "" {
		user.Avatar = avatar
	}
	_, err = dao.SysUser.Data(user).Save()
	if err != nil {
		return "",gerror.New("保存数据失败")
	}
	s.RemoveUserCache(user.UserId)
	return avatar,nil
}

//修改用户密码
func (s *userService) UpdatePassword(ctx context.Context, profile *define.UserApiReSetPasswordReq) error {
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
	s.RemoveUserCache(user.UserId)
	return nil
}

func (s *userService) ChangeStatus(userId int64, status string) error {
	if IsAdmin(userId) {
		return gerror.New("不能停用超级管理员")
	}
	_, err := dao.SysUser.Where(dao.SysUser.Columns.UserId, userId).Data(g.Map{
		"status": status,
	}).Update()
	if err != nil {
		return err
	}
	s.RemoveUserCache(userId)
	return nil
}

// 导出excel
func (s *userService) Export(param *define.UserApiSelectPageReq) (string, error) {
	userList, err := s.SelectExportList(param)
	if err != nil {
		return "", err
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
func (s *userService) GetUserRolePolicy(userName ...string) []define.UserServiceRoleForUser {
	var roleForUser []define.UserServiceRoleForUser
	m := dao.SysRole.As("r")
	m = m.Fields("distinct role_key as roleName,u.login_name as userName")
	m = m.LeftJoin("sys_user_role ur", "ur.role_id = r.role_id")
	m = m.LeftJoin("sys_user u", "u.user_id = ur.user_id")
	m = m.LeftJoin("sys_dept d", "u.dept_id = d.dept_id")
	m = m.Where("r.del_flag = '0' and u.del_flag = '0'")
	if len(userName) > 0 && userName[0] != "" {
		m = m.Where("u.login_name = ?", userName[0])
	}
	if err := m.Structs(&roleForUser); err != nil {
		return nil
	}
	return roleForUser
}

// 获取缓存用户信息
func (s *userService) GetUserCache(prefix string, userId int64) *g.Var {
	user, err := gcache.GetVar(fmt.Sprintf(prefix+"%d", userId))
	if err != nil {
		return nil
	}
	return user
}

// 设置用户信息缓存
func (s *userService) SetUserCache(prefix string, userId int64, user interface{}) {
	gcache.Set(fmt.Sprintf(prefix+"%d", userId), user, time.Hour*24)
}

// 删除用户信息缓存
func (s *userService) RemoveUserCache(userId int64) {
	gcache.Remove(fmt.Sprintf(s.UserCachePrefix+"%d", userId))
	gcache.Remove(fmt.Sprintf(s.UserProfileCachePrefix+"%d", userId))
}
