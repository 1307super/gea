package permission

import (
	menuModel "gea/app/model/system/menu"
	roleModel "gea/app/model/system/role"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/frame/g"
)

func IsAdmin(uid int64) bool {
	if uid == 1 {
		return true
	}
	return false
}
// 获取角色数据权限
func GetRolePermission(uid int64) g.Array {
	var roles g.Array
	if IsAdmin(uid) {
		roles = append(roles,"admin" )
	}else{
		roleEntitys,err := roleModel.SelectRoleListByUserId(uid)
		if err != nil {
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

// 获取菜单数据权限
func GetMenuPermission(uid int64) *garray.StrArray {
	perms := garray.NewStrArray()
	if IsAdmin(uid) {
		perms.Append("*:*:*")
	}else{
		perms = menuModel.SelectMenuPermsByUserId(uid)
	}
	return perms
}
