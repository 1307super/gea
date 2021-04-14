package casbin

import (
	gdbadapter "github.com/1307super/gdb-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/frame/g"
	"log"
)

var Enforcer *casbin.Enforcer
var AuthConf = "config/casbin_auth.conf"

func Register() {
	casbinConfig := g.Cfg().GetMapStrStr("casbin")
	a, err := gdbadapter.NewAdapter(casbinConfig["driverName"], casbinConfig["dataSourceName"])
	if err != nil {
		log.Fatalln("casbin 初始化失败")
	}
	Enforcer, err = casbin.NewEnforcer(AuthConf, a)
	if err != nil {
		log.Fatalln("casbin NewEnforcer 初始化失败")
	}
	err = Enforcer.LoadPolicy()
	if err != nil {
		log.Fatalln("casbin LoadPolicy 初始化失败")
	}
}

// 添加用户角色关联关系
func AddRoleForUser(userName, roleName string) (bool, error) {
	return Enforcer.AddRoleForUser(userName, roleName)
}

// 删除用户角色关联关系
func DeleteRoleForUser(userName, roleName string) (bool, error) {
	return Enforcer.DeleteRoleForUser(userName, roleName)
}

// 清空用户角色关联关系
func DeleteRolesForUser(userName string) (bool, error) {
	return Enforcer.DeleteRolesForUser(userName)
}

// 添加角色与资源关系
func AddPermissionForUser(roleName, path, method string) (bool, error) {
	return Enforcer.AddPermissionForUser(roleName, path, method)
}

// 删除角色与资源关系
func DeletePermissionForUser(roleName, path, method string) (bool, error) {
	return Enforcer.DeletePermissionForUser(roleName, path, method)
}

// 清空角色与资源关系
func DeletePermissionsForUser(roleName string) (bool, error) {
	return Enforcer.DeletePermissionsForUser(roleName)
}

// 校验
func Enforce(userName, path, method string) (bool, error) {
	return Enforcer.Enforce(userName, path, method)
}
