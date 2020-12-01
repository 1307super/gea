package role

import (
	roleModel "gea/app/model/system/role"
	"gea/app/model/system/role_dept"
	"gea/app/model/system/role_menu"
	"gea/app/model/system/user_role"
	menuService "gea/app/service/system/menu"
	userService "gea/app/service/system/user"
	"gea/app/utils/convert"
	"gea/app/utils/excel"
	"gea/app/utils/page"
	"gea/library/casbin"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

//根据主键查询数据
func SelectRecordById(id int64) (*roleModel.Entity, error) {
	return roleModel.FindOne("role_id", id)
}

//根据条件查询数据
func SelectRecordAll(params *roleModel.SelectPageReq) ([]roleModel.EntityFlag, error) {
	return roleModel.SelectListAll(params)
}

//根据条件分页查询数据
func SelectRecordPage(params *roleModel.SelectPageReq) ([]roleModel.Entity, *page.Paging, error) {
	return roleModel.SelectListPage(params)
}

//根据主键删除数据
func DeleteRecordById(id int64) bool {
	result, err := roleModel.Delete("role_id", id)
	if err == nil {
		affected, _ := result.RowsAffected()
		if affected > 0 {
			return true
		}
	}

	return false
}

//添加数据
func AddSave(req *roleModel.AddReq, r *ghttp.Request) (int64, error) {

	var role roleModel.Entity
	role.RoleName = req.RoleName
	role.RoleKey = req.RoleKey
	role.Status = req.Status
	role.Remark = req.Remark
	role.CreateTime = gtime.Now()
	role.CreateBy = ""
	role.DelFlag = "0"
	role.DataScope = "1"
	user, err := userService.GetProfileApi(r.GetInt64("jwtUid"))
	if err != nil {
		return 0, err
	}
	if user != nil {
		role.CreateBy = user.LoginName
	}
	tx, err := g.DB().Begin()
	if err != nil {
		return 0, err
	}

	data := gconv.Map(role)
	delete(data, "flag")
	result, err := tx.Table("sys_role").Insert(data)
	if err != nil {
		return 0, err
	}

	rid, err := result.LastInsertId()
	if err != nil || rid <= 0 {
		tx.Rollback()
		return 0, err
	}
	if req.MenuIds != "" {
		menus := convert.ToInt64Array(req.MenuIds, ",")
		if len(menus) > 0 {
			roleMenus := make([]role_menu.Entity, 0)
			for i := range menus {
				if menus[i] > 0 {
					var tmp role_menu.Entity
					tmp.RoleId = rid
					tmp.MenuId = menus[i]
					roleMenus = append(roleMenus, tmp)
				}
			}
			if len(roleMenus) > 0 {
				_, err := tx.Table("sys_role_menu").Insert(roleMenus)
				if err != nil {
					tx.Rollback()
					return 0, err
				}
				// 加载权限
				go ReloadPermissionsForUser(req.RoleKey)
				// 清空缓存
				go menuService.ClearCache()
			}
		}
	}
	err = tx.Commit()
	return rid, err
}

//修改数据
func EditSave(req *roleModel.EditReq, r *ghttp.Request) (int64, error) {

	role, err := roleModel.FindOne("role_id=?", req.RoleId)

	if err != nil {
		return 0, err
	}

	if role == nil {
		return 0, gerror.New("角色不存在")
	}

	role.RoleName = req.RoleName
	role.RoleKey = req.RoleKey
	role.Status = req.Status
	role.Remark = req.Remark
	role.UpdateTime = gtime.Now()
	role.UpdateBy = ""

	user, err := userService.GetProfileApi(r.GetInt64("jwtUid"))

	if err != nil {
		return 0, err
	}
	if user != nil {
		role.CreateBy = user.LoginName
	}

	tx, err := g.DB().Begin()
	if err != nil {
		return 0, err
	}

	_, err = tx.Table("sys_role").Update(role, "role_id="+gconv.String(role.RoleId))
	if req.MenuIds != "" {
		menus := convert.ToInt64Array(req.MenuIds, ",")
		if len(menus) > 0 {
			roleMenus := make([]role_menu.Entity, 0)
			for i := range menus {
				if menus[i] > 0 {
					var tmp role_menu.Entity
					tmp.RoleId = role.RoleId
					tmp.MenuId = menus[i]
					roleMenus = append(roleMenus, tmp)
				}
			}
			if len(roleMenus) > 0 {
				tx.Table("sys_role_menu").Delete("role_id=?", role.RoleId)
				_, err := tx.Table("sys_role_menu").Insert(roleMenus)
				if err != nil {
					tx.Rollback()
					return 0, err
				}
				// 重置权限
				go ReloadPermissionsForUser(req.RoleKey)
				go menuService.ClearCache()
			}
		}
	}
	return 1, tx.Commit()
}

//保存数据权限
func AuthDataScope(req *roleModel.DataScopeReq, r *ghttp.Request) (int64, error) {
	role, err := roleModel.FindOne("role_id=?", req.RoleId)
	if err != nil {
		return 0, err
	}
	if req.DataScope != "" {
		role.DataScope = req.DataScope
	}

	user, _ := userService.GetProfileApi(r.GetInt64("jwtUid"))

	if user != nil {
		role.UpdateBy = user.LoginName
	}
	role.UpdateTime = gtime.Now()

	tx, err := g.DB().Begin()
	if err != nil {
		return 0, err
	}

	_, err = tx.Table("sys_role").Update(role, "role_id="+gconv.String(role.RoleId))

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if req.DeptIds != "" {
		deptids := convert.ToInt64Array(req.DeptIds, ",")
		if len(deptids) > 0 {
			roleDepts := make([]role_dept.Entity, 0)
			for i := range deptids {
				if deptids[i] > 0 {
					var tmp role_dept.Entity
					tmp.RoleId = role.RoleId
					tmp.DeptId = deptids[i]
					roleDepts = append(roleDepts, tmp)
				}
			}
			if len(roleDepts) > 0 {
				tx.Table("sys_role_dept").Delete("role_id=?", role.RoleId)
				_, err := tx.Table("sys_role_dept").Insert(roleDepts)
				if err != nil {
					tx.Rollback()
					return 0, err
				}
			}
		}
	}
	return 1, tx.Commit()

}

//批量删除数据记录
func DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	roles, err := roleModel.FindAll("role_id in(?)", idarr)
	result, err := roleModel.Delete("role_id in (?)", idarr)
	if err != nil {
		return 0
	}
	nums, _ := result.RowsAffected()
	for _, role := range roles {
		go casbin.DeletePermissionsForUser(role.RoleKey)
	}
	// 删除对应的权限
	db := g.DB()
	db.Table("sys_role_menu").Delete("role_id in(?)",idarr)
	db.Table("sys_role_dept").Delete("role_id in(?)",idarr)
	go menuService.ClearCache()
	return nums
}

// 导出excel
func Export(param *roleModel.SelectPageReq) (string, error) {
	result, err := roleModel.SelectListExport(param)
	if err != nil {
		return "", err
	}

	head := []string{"用户名", "呢称", "Email", "电话号码", "性别", "部门", "领导", "状态", "删除标记", "创建人", "创建时间", "备注"}
	key := []string{"role_id", "role_name", "role_key", "role_sort", "data_scope", "status"}
	url, err := excel.DownlaodExcel(head, key, result)

	if err != nil {
		return "", err
	}

	return url, nil
}

//查询角色
func SelectRoleContactVo() (g.Array, error) {
	userRoleFlags, err := roleModel.SelectRoleContactVo()
	if err != nil {
		return nil, gerror.New("未查询到用户角色数据")
	}
	var userRoles g.Array
	for _, userRoleFlag := range userRoleFlags {
		userRoles = append(userRoles, userRoleFlag)
	}
	return userRoles, nil
}
func SelectRoleListByUserId(uid int64) ([]roleModel.EntityFlag, error) {
	roleEntityFlags, err := roleModel.SelectRoleListByUserId(uid)
	if err != nil {
		return nil, gerror.New("未查询到用户角色数据数据")
	}

	return roleEntityFlags, nil
}
func SelectRoleListIdByUserId(uid int64) ([]int64, error) {
	roleEntityFlags, err := roleModel.SelectRoleListByUserId(uid)
	if err != nil {
		return nil, gerror.New("未查询到用户角色数据数据")
	}
	var roleIds []int64
	for _, roleEntityFlag := range roleEntityFlags {
		roleIds = append(roleIds, roleEntityFlag.RoleId)
	}
	return roleIds, nil
}

//批量选择用户授权
func InsertAuthUsers(roleId int64, userIds string) int64 {
	idarr := convert.ToInt64Array(userIds, ",")
	var roleUserList []user_role.Entity
	for _, str := range idarr {
		var tmp user_role.Entity
		tmp.UserId = str
		tmp.RoleId = roleId
		roleUserList = append(roleUserList, tmp)
	}
	rs, err := user_role.Insert(roleUserList)
	if err != nil {
		return 0
	}
	nums, err := rs.RowsAffected()
	return nums
}

//取消授权用户角色
func DeleteUserRoleInfo(userId, roleId int64) int64 {
	result, err := user_role.Delete("user_id=? and role_id=?", userId, roleId)
	if err != nil {
		return 0
	}
	nums, _ := result.RowsAffected()
	return nums
}

//批量取消授权用户角色
func DeleteUserRoleInfos(roleId int64, ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := user_role.Delete("role_id=? and user_id in (?)", roleId, idarr)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

//检查角色名是否唯一
func CheckRoleNameUniqueAll(roleName string) string {
	role, err := roleModel.FindOne("role_name=?", roleName)
	if err == nil && role != nil && role.RoleId > 0 {
		return "1"
	}
	return "0"
}

//检查角色键是否唯一
func CheckRoleKeyUniqueAll(roleKey string) string {
	role, err := roleModel.FindOne("role_key=?", roleKey)
	if err == nil && role != nil && role.RoleId > 0 {
		return "1"
	}
	return "0"
}

//检查角色名是否唯一
func CheckRoleNameUnique(roleName string, roleId int64) string {
	role, err := roleModel.FindOne("role_name=? and role_id <> ?", roleName, roleId)
	if err == nil && role != nil && role.RoleId > 0 {
		return "1"
	}
	return "0"
}

//检查角色键是否唯一
func CheckRoleKeyUnique(roleKey string, roleId int64) string {
	role, err := roleModel.FindOne("role_key=? and role_id <> ?", roleKey, roleId)
	if err == nil && role != nil && role.RoleId > 0 {
		return "1"
	}
	return "0"
}

//判断是否是管理员
func IsAdmin(id int64) bool {
	if id == 1 {
		return true
	} else {
		return false
	}
}

//校验角色是否允许操作
func CheckRoleAllowed(id int64) bool {
	if IsAdmin(id) {
		return false
	} else {
		return true
	}
}

func ChangeStatus(roleId int64, status string) error {
	if IsAdmin(roleId) {
		return gerror.New("不能停用超级管理员")
	}
	_, err := roleModel.Update(g.Map{
		"status": status,
	}, "role_id=?", roleId)
	if err != nil {
		return err
	}
	return nil
}

// 重置角色权限
func ReloadPermissionsForUser(roleName string) {
	casbin.DeletePermissionsForUser(roleName)
	LoadRolePolicy(roleName)
}

// 加载所有角色权限
func LoadRolePolicy(roleName string) {
	permissionForRoles := roleModel.GetRoleMenuPolicy(roleName)
	for _, permissionForRole := range permissionForRoles {
		casbin.AddPermissionForUser(permissionForRole.RoleName, permissionForRole.Path, permissionForRole.Method)
	}
}
