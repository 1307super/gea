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
	"gea/library/casbin"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var Role = &roleService{}

type roleService struct {}

func IsAdmin(uid int64) bool {
	if uid == 1 {
		return true
	}
	return false
}

//根据条件分页查询角色数据
func (s *roleService)GetList(param *define.RoleApiSelectPageReq) *define.RoleServiceList {

	m := dao.SysRole.As("r").Where(fmt.Sprintf("r.%s",dao.SysRole.Columns.DelFlag),"0")

	if param.RoleName != "" {
		m = m.Where("r.role_name like ?", "%"+param.RoleName+"%")
	}

	if param.Status != "" {
		m = m.Where("r.status = ", param.Status)
	}

	if param.RoleKey != "" {
		m = m.Where("r.role_key like ?", "%"+param.RoleKey+"%")
	}

	if param.DataScope != "" {
		m = m.Where("r.data_scope = ", param.DataScope)
	}

	if param.BeginTime != "" {
		m = m.Where("date_format(r.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
	}

	if param.EndTime != "" {
		m = m.Where("date_format(r.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
	}

	total, err := m.Count()

	if err != nil {
		return nil
	}

	page := page.CreatePaging(param.PageNum, param.PageSize, total)

	m = m.Limit(page.StartNum, page.Pagesize)
	if param.OrderByColumn != "" {
		m = m.Order(param.OrderByColumn + " " + param.IsAsc)
	}
	result := &define.RoleServiceList{
		Page:  page.PageNum,
		Size:  page.Pagesize,
		Total: page.Total,
	}
	if err = m.Structs(&result.List); err != nil {
		return nil
	}
	return result
}

//添加数据
func (s *roleService)Create(ctx context.Context,req *define.RoleApiCreateReq) (int64, error) {
	user := shared.Context.Get(ctx).User

	if s.CheckRoleNameUniqueAll(req.RoleName){
		return 0,gerror.New("角色名称已存在")
	}

	if s.CheckRoleKeyUniqueAll(req.RoleKey) {
		return 0,gerror.New("角色权限已存在")
	}

	var role model.SysRole
	role.RoleName = req.RoleName
	role.RoleKey = req.RoleKey
	role.Status = req.Status
	role.Remark = req.Remark
	role.CreateTime = gtime.Now()
	role.CreateBy = ""
	role.DelFlag = "0"
	role.DataScope = "1"
	role.CreateBy = user.LoginName
	var editReq *define.RoleApiEditReq
	gconv.Struct(req,&editReq)
	return s.save(&role,editReq)
}

//修改数据
func (s *roleService)Update(ctx context.Context,req *define.RoleApiEditReq) (int64, error) {
	if s.CheckRoleNameUnique(req.RoleName, req.RoleId){
		return 0,gerror.New("角色名称已存在")
	}

	if s.CheckRoleKeyUnique(req.RoleKey, req.RoleId)  {
		return 0,gerror.New("角色权限已存在")
	}

	user := shared.Context.Get(ctx).User

	role, err := dao.SysRole.FindOne(dao.SysRole.Columns.RoleId,req.RoleId)

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
	role.UpdateBy = user.LoginName
	role.CreateBy = user.LoginName
	return s.save(role,req)
}

// 添加 修改统一处理
func (s *roleService) save(role *model.SysRole,req *define.RoleApiEditReq) (int64, error) {
	tx, err := g.DB().Begin()
	if err != nil {
		return 0, err
	}
	result, err := dao.SysRole.TX(tx).Data(role).Save()
	var rid int64
	if role.RoleId == 0 {
		// 新增
		rid, err = result.LastInsertId()
		if err != nil || rid <= 0 {
			tx.Rollback()
			return 0, err
		}
	}else{
		// 修改
		rid = role.RoleId
	}
	if req.MenuIds != ""{
		menus := convert.ToInt64Array(req.MenuIds, ",")
		if len(menus) > 0 {
			roleMenus := make([]model.SysRoleMenu, 0)
			for i := range menus {
				if menus[i] > 0 {
					var tmp model.SysRoleMenu
					tmp.RoleId = rid
					tmp.MenuId = menus[i]
					roleMenus = append(roleMenus, tmp)
				}
			}
			if len(roleMenus) > 0 {
				if role.RoleId != 0 {
					dao.SysRoleMenu.TX(tx).Delete(dao.SysRoleMenu.Columns.RoleId,role.RoleId)
				}
				_, err := dao.SysRoleMenu.TX(tx).Data(roleMenus).Insert()
				if err != nil {
					tx.Rollback()
					return 0, err
				}
				// 加载权限
				go s.ReloadPermissionsForUser(req.RoleKey)
				// 清空缓存
				go Menu.ClearCache()
			}
		}
	}
	return rid,tx.Commit()
}

func (s *roleService)Delete(ids string) bool {
	idarr := convert.ToInt64Array(ids, ",")

	field := fmt.Sprintf("%s in(?)",dao.SysRole.Columns.RoleId)
	roles,err := dao.SysRole.FindAll(field,idarr)

	result, err := dao.SysRole.Delete(field,idarr)
	if err != nil {
		return false
	}
	nums, _ := result.RowsAffected()
	for _, role := range roles {
		go casbin.DeletePermissionsForUser(role.RoleKey)
	}
	// 删除对应的权限
	dao.SysRoleMenu.Delete(field,idarr)
	dao.SysRoleDept.Delete(field,idarr)
	go Menu.ClearCache()
	return nums > 0
}

// 导出excel
func (s *roleService)Export(param *define.RoleApiSelectPageReq) (string, error) {
	m := dao.SysRole.As("r").Where("r.del_flag = '0'")
	if param != nil {
		if param.RoleName != "" {
			m = m.Where("r.role_name like ?", "%"+param.RoleName+"%")
		}
		if param.Status != "" {
			m = m.Where("r.status = ", param.Status)
		}
		if param.RoleKey != "" {
			m = m.Where("r.role_key like ?", "%"+param.RoleKey+"%")
		}

		if param.DataScope != "" {
			m = m.Where("r.data_scope = ", param.DataScope)
		}

		if param.BeginTime != "" {
			m = m.Where("date_format(r.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			m = m.Where("date_format(r.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}
	//角色序号	角色名称	角色权限	角色排序	数据范围	角色状态
	m = m.Fields("r.role_id,r.role_name,r.role_key,r.role_sort,r.data_scope,r.status")
	result, err := m.M.All()
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

// 查询角色详情
func (s *roleService) Info(id int64) (*model.SysRole,error) {
	return dao.SysRole.FindOne(dao.SysRole.Columns.RoleId,id)
}

//保存数据权限
func (s *roleService)AuthDataScope(ctx context.Context,req *define.RoleApiDataScopeReq) (int64, error) {

	user := shared.Context.Get(ctx).User
	if IsAdmin(req.RoleId) {
		return 0,gerror.New("不允许操作超级管理员角色")
	}
	role, err := s.Info(req.RoleId)
	if err != nil {
		return 0, err
	}
	if req.DataScope != "" {
		role.DataScope = req.DataScope
	}
	role.UpdateBy = user.LoginName
	role.UpdateTime = gtime.Now()
	role.UpdateBy = user.LoginName

	tx, err := g.DB().Begin()
	if err != nil {
		return 0, err
	}
	_, err = dao.SysRole.TX(tx).Data(role).Save()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if req.DeptIds != "" {
		deptids := convert.ToInt64Array(req.DeptIds, ",")
		if len(deptids) > 0 {
			roleDepts := make([]model.SysRoleDept, 0)
			for i := range deptids {
				if deptids[i] > 0 {
					var tmp model.SysRoleDept
					tmp.RoleId = role.RoleId
					tmp.DeptId = deptids[i]
					roleDepts = append(roleDepts, tmp)
				}
			}
			if len(roleDepts) > 0 {
				dao.SysRoleDept.TX(tx).Delete(dao.SysRoleDept.Columns.RoleId,role.RoleId)
				_, err := dao.SysRoleDept.TX(tx).Data(roleDepts).Insert()
				if err != nil {
					tx.Rollback()
					return 0, err
				}
			}
		}
	}
	return 1, tx.Commit()

}

func (s *roleService)ChangeStatus(roleId int64, status string) error {
	if IsAdmin(roleId) {
		return gerror.New("不能停用超级管理员")
	}
	_, err := dao.SysRole.Data(g.Map{
		"status": status,
	}).Where(dao.SysRole.Columns.RoleId,roleId).Update()
	if err != nil {
		return err
	}
	return nil
}


func (s *roleService) GetRolePermission(ctx context.Context) g.Array {
	customCtx := shared.Context.Get(ctx)
	var roles g.Array
	if IsAdmin(customCtx.User.UserId) {
		roles = append(roles,"admin" )
	}else{
		var roleEntitys []model.SysRoleFlag
		if err := dao.SysRole.As("r").LeftJoin("sys_user_role ur", "ur.role_id = r.role_id").LeftJoin("sys_user u", "u.user_id = ur.user_id").Where("r.del_flag = '0' and u.del_flag = '0'").Where("ur.user_id = ?", customCtx.User.UserId).Structs(&roleEntitys); err != nil {
			return roles
		}
		for _,roleEntity := range roleEntitys  {
			if roleEntity.RoleKey != "" {
				roles = append(roles, roleEntity.RoleKey)
			}
		}
	}
	return roles
}


//查询角色
func (s *roleService)SelectRoleContact() (g.Array, error) {
	var userRoleFlags []model.SysRoleFlag
	if err := dao.SysRole.As("r").Fields("distinct r.role_id, r.role_name, r.role_key, r.role_sort, r.data_scope,r.status, r.del_flag, r.create_time, r.remark").Where("r.del_flag = '0'").Structs(&userRoleFlags);err != nil{
		return nil,gerror.New("未查询到用户角色数据")
	}
	var userRoles g.Array
	for _, userRoleFlag := range userRoleFlags {
		userRoles = append(userRoles, userRoleFlag)
	}
	return userRoles, nil
}

func (s *roleService) GetRoleListByUid(uid int64) ([]int64,error) {
	m := dao.SysRole.As("r")
	m = m.LeftJoin("sys_user_role ur", "ur.role_id = r.role_id")
	m = m.LeftJoin("sys_user u", "u.user_id = ur.user_id")
	m = m.Where("r.del_flag = '0' and u.del_flag = '0'")
	m = m.Where("ur.user_id = ?", uid)
	//var result g.Array
	var result []model.SysRoleFlag
	err := m.Structs(&result)
	if err != nil {
		return nil, gerror.New("未查询到用户角色数据数据")
	}
	var roleIds []int64
	for _, roleEntityFlag := range result {
		roleIds = append(roleIds, roleEntityFlag.RoleId)
	}
	return roleIds, nil
}


// ============= 校验操作 ================
//检查角色名是否唯一
func (s *roleService)CheckRoleNameUniqueAll(roleName string) bool {
	role, err := dao.SysRole.FindOne(dao.SysRole.Columns.RoleName,roleName)
	if err == nil && role != nil && role.RoleId > 0 {
		return true
	}
	return false
}

//检查角色键是否唯一
func (s *roleService)CheckRoleKeyUniqueAll(roleKey string) bool {
	role, err := dao.SysRole.FindOne(dao.SysRole.Columns.RoleKey,roleKey)
	if err == nil && role != nil && role.RoleId > 0 {
		return true
	}
	return false
}

//检查角色名是否唯一
func (s *roleService)CheckRoleNameUnique(roleName string, roleId int64) bool {
	role, err := dao.SysRole.FindOne(g.Map{
		dao.SysRole.Columns.RoleName:roleName,
		fmt.Sprintf("%s <> ",dao.SysRole.Columns.RoleId): roleId,
	})
	if err == nil && role != nil && role.RoleId > 0 {
		return true
	}
	return false
}

//检查角色键是否唯一
func (s *roleService)CheckRoleKeyUnique(roleKey string, roleId int64) bool {
	role, err := dao.SysRole.FindOne(g.Map{
		dao.SysRole.Columns.RoleKey:roleKey,
		fmt.Sprintf("%s <> ",dao.SysRole.Columns.RoleId): roleId,
	})
	if err == nil && role != nil && role.RoleId > 0 {
		return true
	}
	return false
}


// 重置角色权限
func (s *roleService)ReloadPermissionsForUser(roleName string) {
	casbin.DeletePermissionsForUser(roleName)
	s.LoadRolePolicy(roleName)
}

// 加载所有角色权限
func (s *roleService)LoadRolePolicy(roleName string) {
	permissionForRoles := s.GetRoleMenuPolicy(roleName)
	for _, permissionForRole := range permissionForRoles {
		casbin.AddPermissionForUser(permissionForRole.RoleName, permissionForRole.Path, permissionForRole.Method)
	}
}

func (s *roleService)GetRoleMenuPolicy(roleName ...string) []define.RoleServicePermissionForRole {
	var result []define.RoleServicePermissionForRole
	m := dao.SysMenu.As("m")
	m = m.LeftJoin("sys_role_menu rm", "m.menu_id = rm.menu_id")
	m = m.LeftJoin("sys_role ro", "rm.role_id = ro.role_id")
	m = m.Fields("distinct ro.role_key as roleName,m.perms as path,m.method")
	m = m.Where("m.status = '0' and ro.status = '0' and m.menu_type != 'M'")
	if len(roleName) > 0 && roleName[0] != "" {
		m = m.Where("ro.role_key = ?", roleName[0])
	}
	_ = m.Structs(&result)
	return result
}