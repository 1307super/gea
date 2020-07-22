package user

import (
	deptModel "gea/app/model/system/dept"
	postModel "gea/app/model/system/post"
	roleModel "gea/app/model/system/role"
	userModel "gea/app/model/system/user"
	"gea/app/model/system/user_post"
	"gea/app/model/system/user_role"
	"gea/app/service/permission"
	"gea/app/utils/convert"
	"gea/app/utils/excel"
	"gea/app/utils/page"
	"gea/app/utils/random"
	"gea/library/casbin"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

//根据主键查询用户信息
func SelectRecordById(id int64) (*userModel.Entity, error) {
	return userModel.FindOne("user_id", id)
}

// 根据条件分页查询用户列表
func SelectRecordList(param *userModel.SelectPageReq) ([]userModel.UserListEntity, *page.Paging, error) {
	return userModel.SelectPageList(param)
}

// 导出excel
func Export(param *userModel.SelectPageReq) (string, error) {
	result, err := userModel.SelectExportList(param)
	if err != nil {
		return "", err
	}

	head := []string{"用户名", "呢称", "Email", "电话号码", "性别", "部门", "领导", "状态", "删除标记", "创建人", "创建时间", "备注"}
	key := []string{"login_name", "user_name", "email", "phonenumber", "sex", "dept_name", "leader", "status", "del_flag", "create_by", "create_time", "remark"}
	url, err := excel.DownlaodExcel(head, key, result)

	if err != nil {
		return "", err
	}

	return url, nil
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func SignIn(loginnName, password string) (*userModel.Entity, error) {
	//查询用户信息
	user, err := userModel.FindOne("login_name=?", loginnName)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, gerror.New("用户或密码不正确")
	}

	//校验密码
	token := user.LoginName + password + user.Salt

	token = gmd5.MustEncryptString(token)

	if !strings.EqualFold(user.Password, token) {
		return nil, gerror.New("密码错误")
	}
	return user, nil
}


//新增用户
func AddSave(req *userModel.AddReq, r *ghttp.Request) (int64, error) {
	var user userModel.Entity
	user.LoginName = req.LoginName
	user.UserName = req.UserName
	user.Email = req.Email
	user.Phonenumber = req.Phonenumber
	user.Status = req.Status
	user.Sex = req.Sex
	user.DeptId = req.DeptId
	user.Remark = req.Remark

	//生成密码
	newSalt := random.GenerateSubId(6)
	newToken := req.LoginName + req.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken

	user.CreateTime = gtime.Now()

	createUser,_ := GetProfileApi(r.GetInt64("jwtUid"))

	if createUser != nil {
		user.CreateBy = createUser.LoginName
	}

	user.DelFlag = "0"

	var err error
	tx, err := g.DB().Begin()
	if err != nil {
		return 0, err
	}

	result, err := tx.Table("sys_user").Insert(user)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	uid, err := result.LastInsertId()

	if err != nil || uid <= 0 {
		return 0, err
	}

	//增加岗位数据
	if req.PostIds != "" {
		postIds := convert.ToInt64Array(req.PostIds, ",")
		userPosts := make([]user_post.Entity, 0)
		for i := range postIds {
			if postIds[i] > 0 {
				var userPost user_post.Entity
				userPost.UserId = uid
				userPost.PostId = postIds[i]
				userPosts = append(userPosts, userPost)
			}
		}
		if len(userPosts) > 0 {
			_, err := tx.Table("sys_user_post").Insert(userPosts)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}

	}

	//增加角色数据
	if req.RoleIds != "" {
		roleIds := convert.ToInt64Array(req.RoleIds, ",")
		userRoles := make([]user_role.Entity, 0)
		for i := range roleIds {
			if roleIds[i] > 0 {
				var userRole user_role.Entity
				userRole.UserId = uid
				userRole.RoleId = roleIds[i]
				userRoles = append(userRoles, userRole)
			}
		}
		if len(userRoles) > 0 {
			_, err := tx.Table("sys_user_role").Insert(userRoles)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	}
	go ReloadUserRole(req.LoginName)
	return uid, tx.Commit()
}

//新增用户
func EditSave(req *userModel.EditReq, r *ghttp.Request) (int64, error) {
	user, err := userModel.FindOne("user_id=?", req.UserId)
	if err != nil || user == nil {
		return 0, err
	}

	user.UserName = req.UserName
	user.Email = req.Email
	user.Phonenumber = req.Phonenumber
	user.Status = req.Status
	user.Sex = req.Sex
	user.DeptId = req.DeptId
	user.Remark = req.Remark

	user.UpdateTime = gtime.Now()

	updateUser, _ := GetProfileApi(r.GetInt64("jwtUid"))

	if updateUser != nil {
		user.UpdateBy = updateUser.LoginName
	}

	tx, err := g.DB().Begin()
	if err != nil {
		return 0, err
	}
	_, err = tx.Table("sys_user").Update(user, "user_id="+gconv.String(user.UserId))

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	//增加岗位数据
	if req.PostIds != "" {
		postIds := convert.ToInt64Array(req.PostIds, ",")
		userPosts := make([]user_post.Entity, 0)
		for i := range postIds {
			if postIds[i] > 0 {
				var userPost user_post.Entity
				userPost.UserId = user.UserId
				userPost.PostId = postIds[i]
				userPosts = append(userPosts, userPost)
			}
		}
		if len(userPosts) > 0 {
			tx.Table("sys_user_post").Delete("user_id=?", user.UserId)
			_, err := tx.Table("sys_user_post").Insert(userPosts)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}

	}

	//增加角色数据
	if req.RoleIds != "" {
		roleIds := convert.ToInt64Array(req.RoleIds, ",")
		userRoles := make([]user_role.Entity, 0)
		for i := range roleIds {
			if roleIds[i] > 0 {
				var userRole user_role.Entity
				userRole.UserId = user.UserId
				userRole.RoleId = roleIds[i]
				userRoles = append(userRoles, userRole)
			}
		}
		if len(userRoles) > 0 {
			tx.Table("sys_user_role").Delete("user_id=?", user.UserId)
			_, err := tx.Table("sys_user_role").Insert(userRoles)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	}
	go ReloadUserRole(user.LoginName)
	return 1, tx.Commit()
}

//根据主键删除用户信息
func DeleteRecordById(id int64) bool {
	result, err := userModel.Delete("user_id", id)
	if err == nil {
		affected, _ := result.RowsAffected()
		if affected > 0 {
			return true
		}
	}

	return false
}

//批量删除用户记录
func DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := userModel.Update(g.Map{
		"del_flag": 2,
	},"user_id in(?)",idarr)
	//result, err := userModel.Delete("user_id in (?)", idarr)
	if err != nil {
		return 0
	}
	nums, _ := result.RowsAffected()
	return nums
}

//判断是否是系统管理员
func IsAdmin(userId int64) bool {
	if userId == 1 {
		return true
	} else {
		return false
	}
}

// 检查登陆名是否存在,存在返回true,否则false
func CheckLoginName(loginName string) bool {
	if i, err := userModel.FindCount("login_name", loginName); err != nil {
		return false
	} else {
		return i == 1
	}
}


//更新用户信息详情
func UpdateProfile(profile *userModel.ProfileReq, r *ghttp.Request) error {
	user,_ := GetProfileApi(r.GetInt64("jwtUid"))

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

	_, err := user.Update()
	if err != nil {
		return gerror.New("保存数据失败")
	}
	return nil
}

//更新用户头像
func UpdateAvatar(avatar string, r *ghttp.Request) error {
	user,_ := GetProfileApi(r.GetInt64("jwtUid"))

	if avatar != "" {
		user.Avatar = avatar
	}

	_, err := user.Update()
	if err != nil {
		return gerror.New("保存数据失败")
	}

	return nil
}

//修改用户密码
func UpdatePassword(profile *userModel.PasswordReq, r *ghttp.Request) error {
	user,_ := GetProfileApi(r.GetInt64("jwtUid"))

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
	token := user.LoginName + profile.OldPassword + user.Salt
	token = gmd5.MustEncryptString(token)

	if !strings.EqualFold(token, user.Password) {
		return gerror.New("原密码不正确")
	}

	//新校验密码
	newSalt := random.GenerateSubId(6)
	newToken := user.LoginName + profile.NewPassword + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken
	_, err := user.Update()
	if err != nil {
		return gerror.New("保存数据失败")
	}
	return nil
}

//重置用户密码
func ResetPassword(params *userModel.ResetPwdReq) (bool, error) {

	user, err := userModel.FindOne("user_id=?", params.UserId)

	if err != nil {
		return false, gerror.New("用户不存在")
	}

	//新校验密码
	newSalt := random.GenerateSubId(6)
	newToken := user.LoginName + params.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken

	_, err = user.Update()
	if err != nil {
		return false, gerror.New("保存数据失败")
	}

	return true, nil
}

//检查邮箱是否已使用
func CheckEmailUnique(userId int64, email string) bool {
	return userModel.CheckEmailUnique(userId, email)
}

//检查邮箱是否存在,存在返回true,否则false
func CheckEmailUniqueAll(email string) bool {
	return userModel.CheckEmailUniqueAll(email)
}

//检查手机号是否已使用,存在返回true,否则false
func CheckPhoneUnique(userId int64, phone string) bool {
	return userModel.CheckPhoneUnique(userId, phone)
}

//检查手机号是否已使用 ,存在返回true,否则false
func CheckPhoneUniqueAll(phone string) bool {
	return userModel.CheckPhoneUniqueAll(phone)
}

// 查询已分配用户角色列表
func SelectAllocatedList(roleId int64, loginName, phonenumber string) ([]userModel.Entity, error) {
	return userModel.SelectAllocatedList(roleId, loginName, phonenumber)
}

// 查询未分配用户角色列表
func SelectUnallocatedList(roleId int64, loginName, phonenumber string) ([]userModel.Entity, error) {
	return userModel.SelectUnallocatedList(roleId, loginName, phonenumber)
}

// 获得用户信息详情
func GetProfileApi(uid int64) (*userModel.UserEntityExtend, error) {
	user, err := userModel.SelectUserByUid(uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}


func ChangeStatus(userId int64, status string) error {
	if IsAdmin(userId) {
		return gerror.New("不能停用超级管理员")
	}
	_, err := userModel.Update(g.Map{
		"status": status,
	},"user_id=?",userId)
	if err != nil {
		return err
	}
	return nil
}

func GetProfile(uid int64) (*userModel.UserInfo, error) {
	user,err := GetProfileApi(uid)
	if err != nil {
		return nil,err
	}
	var userInfo = new(userModel.UserInfo)
	// 获取部门
	dept,err := deptModel.SelectDeptById(user.DeptId)
	if err != nil {
		return nil,gerror.New("未获取到部门")
	}
	user.Dept = dept
	// 获取角色
	roles, err := roleModel.SelectRoleListByUserId(uid)
	if err != nil {
		return nil,gerror.New("未获取到角色")
	}
	// 获取岗位
	posts, err := postModel.SelectPostsByUserId(uid)
	if err != nil {
		return nil,gerror.New("未获取到岗位")
	}
	roleNames := garray.New()
	postNames := garray.New()
	for _, role := range roles {
		roleNames.Append(role.RoleName)
	}
	for _, postflag := range posts {
		postNames.Append(postflag.PostName)
	}
	user.Roles = roles
	userInfo.User = user
	userInfo.PostGroup = postNames.Join(",")
	userInfo.RoleGroup = roleNames.Join(",")
	return userInfo,nil
}

func GetUserInfo(uid int64) (*userModel.UserInfo, error) {
	user,err := GetProfileApi(uid)
	if err != nil {
		return nil,err
	}
	userInfo := new(userModel.UserInfo)
	userInfo.User = user
	userInfo.Roles = permission.GetRolePermission(uid)
	userInfo.Permissions = permission.GetMenuPermission(uid)
	return userInfo,nil
}

func ReloadUserRole(userName string) {
	casbin.DeleteRolesForUser(userName)
	LoadUserRole(userName)
}
// 加载所有用户角色
func LoadUserRole(userName string) {
	roleForUsers := userModel.GetUserRolePolicy(userName)
	for _, roleForUser := range roleForUsers{
		casbin.AddRoleForUser(roleForUser.UserName,roleForUser.RoleName)
	}
}