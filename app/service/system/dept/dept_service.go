package dept

import (
	"gea/app/model"
	deptModel "gea/app/model/system/dept"
	userModel "gea/app/model/system/user"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

//新增保存信息
func AddSave(req *deptModel.AddReq, r *ghttp.Request) (int64, error) {
	var dept deptModel.Entity

	pdept, err := deptModel.FindOne("dept_id=?", req.ParentId)
	if err == nil && pdept != nil {
		if pdept.Status != "0" {
			return 0, gerror.New("部门停用，不允许新增")
		} else {
			dept.Ancestors = pdept.Ancestors + "," + gconv.String(pdept.DeptId)
		}

	} else {
		return 0, gerror.New("父部门不能为空")
	}

	dept.DeptName = req.DeptName
	dept.Status = req.Status
	dept.ParentId = req.ParentId
	dept.DelFlag = "0"
	dept.Email = req.Email
	dept.Leader = req.Leader
	dept.Phone = req.Phone
	dept.OrderNum = req.OrderNum

	user,_ := userModel.SelectUserByUid(r.GetInt64("jwtUid"))

	if user != nil && user.UserId > 0 {
		dept.CreateBy = user.LoginName
	}

	dept.CreateTime = gtime.Now()

	rs, err := dept.Insert()
	if err != nil {
		return 0, err
	}
	did, _ := rs.LastInsertId()
	return did, nil
}

//修改保存信息
func EditSave(req *deptModel.EditReq, r *ghttp.Request) (int64, error) {
	dept, err := deptModel.FindOne("dept_id=?", req.DeptId)
	if err != nil || dept == nil || dept.DeptId <= 0 {
		return 0, gerror.New("数据不存在")
	}
	var pdept *deptModel.Entity
	if dept.DeptId == 100 {
		// 顶级部门
		pdept = dept
	}else{
		pdept, err = deptModel.FindOne("dept_id=?", req.ParentId)
	}
	if pdept != nil {
		if pdept.Status != "0" {
			return 0, gerror.New("部门停用，不允许新增")
		} else {
			newAncestors := pdept.Ancestors + "," + gconv.String(pdept.DeptId)
			deptModel.UpdateDeptChildren(dept.DeptId, newAncestors, dept.Ancestors)

			dept.DeptName = req.DeptName
			dept.Status = req.Status
			dept.ParentId = req.ParentId
			dept.DelFlag = "0"
			dept.Email = req.Email
			dept.Leader = req.Leader
			dept.Phone = req.Phone
			dept.OrderNum = req.OrderNum

			user,_ := userModel.SelectUserByUid(r.GetInt64("jwtUid"))

			if user != nil && user.UserId > 0 {
				dept.UpdateBy = user.LoginName
			}

			dept.UpdateTime = gtime.Now()

			dept.Update()
			return 1, nil
		}

	} else {
		return 0, gerror.New("父部门不能为空")
	}
}

//根据分页查询部门管理数据
func SelectListAll(param *deptModel.SelectPageReq) ([]*deptModel.EntityExtend, error) {
	return deptModel.SelectDeptList(param.ParentId, param.DeptName, param.Status)
}

//删除部门管理信息
func DeleteDeptById(deptId int64) int64 {
	return deptModel.DeleteDeptById(deptId)
}

//根据部门ID查询信息
func SelectDeptById(deptId int64) *deptModel.EntityExtend {
	deptEntity, err := deptModel.SelectDeptById(deptId)
	if err != nil {
		return nil
	}
	return deptEntity
}

//根据ID查询所有子部门
func SelectChildrenDeptById(deptId int64) []*deptModel.Entity {
	return deptModel.SelectChildrenDeptById(deptId)
}

//加载部门列表树
func SelectDeptTree(parentId int64, deptName, status string) ([]*deptModel.EntityExtend, error) {
	list, err := deptModel.SelectDeptList(parentId, deptName, status)
	if err != nil {
		return nil, err
	}

	list = BuildDeptTree(list)
	//return InitZtree(list, nil), nil
	return list,nil

}

////查询部门管理数据
//func SelectDeptList(parentId int64, deptName, status string) ([]deptModel.Entity, error) {
//	return deptModel.SelectDeptList(parentId, deptName, status)
//}
//
//根据角色ID查询部门（数据权限）
func SelectRoleDeptIds(roleId int64) (g.Array, error) {
	return deptModel.SelectRoleDeptIds(roleId)
}

//对象转部门树
func InitZtree(deptList []deptModel.Entity, roleDeptList *[]string) *[]model.Ztree {
	var result []model.Ztree
	isCheck := false
	if roleDeptList != nil && len(*roleDeptList) > 0 {
		isCheck = true
	}

	for i := range deptList {
		if deptList[i].Status == "0" {
			var ztree model.Ztree
			ztree.Id = deptList[i].DeptId
			ztree.Pid = deptList[i].ParentId
			ztree.Name = deptList[i].DeptName
			ztree.Title = deptList[i].DeptName
			if isCheck {
				tmp := gconv.String(deptList[i].DeptId) + deptList[i].DeptName
				tmpcheck := false
				for j := range *roleDeptList {
					if strings.EqualFold((*roleDeptList)[j], tmp) {
						tmpcheck = true
						break
					}
				}
				ztree.Checked = tmpcheck
			}
			result = append(result, ztree)
		}
	}
	return &result
}

//查询部门是否存在用户
func CheckDeptExistUser(deptId int64) bool {
	return deptModel.CheckDeptExistUser(deptId)
}

//查询部门人数
func SelectDeptCount(deptId, parentId int64) int {
	return deptModel.SelectDeptCount(deptId, parentId)
}

//校验部门名称是否唯一
func CheckDeptNameUniqueAll(deptName string, parentId int64) string {
	dept, err := deptModel.CheckDeptNameUniqueAll(deptName, parentId)
	if err != nil {
		return "1"
	}
	if dept != nil && dept.DeptId > 0 {
		return "1"
	} else {
		return "0"
	}
}

//校验部门名称是否唯一
func CheckDeptNameUnique(deptName string, deptId, parentId int64) string {
	dept, err := deptModel.CheckDeptNameUnique(deptName, deptId, parentId)
	if err != nil {
		return "1"
	}
	if dept != nil && dept.DeptId > 0 {
		return "1"
	} else {
		return "0"
	}
}

func BuildDeptTree(depts []*deptModel.EntityExtend)[]*deptModel.EntityExtend {
	var returnList []*deptModel.EntityExtend
	tempList := garray.NewIntArray(true)
	for _, dept := range depts  {
		tempList.Append(gconv.Int(dept.DeptId))
	}
	for _,dept := range depts{
		if !tempList.Contains(gconv.Int(dept.ParentId)) {
			recursionFn(depts,dept)
			returnList = append(returnList, dept)
		}
	}
	if len(returnList) == 0 {
		returnList = depts
	}
	return returnList
}


// 递归
func recursionFn(depts []*deptModel.EntityExtend, d *deptModel.EntityExtend){
	childMenus := getChildList(depts,d)
	d.Children = childMenus
	for _,childMenu := range childMenus  {
		if hasChild(depts,childMenu){
			for _,menu := range childMenus {
				recursionFn(depts,menu)
			}
		}
	}
}

// 获取子级列表
func getChildList(depts []*deptModel.EntityExtend, d *deptModel.EntityExtend)[]*deptModel.EntityExtend {
	var tlist []*deptModel.EntityExtend
	for _,dept := range depts  {
		if dept.ParentId == d.DeptId {
			tlist = append(tlist, dept)
		}
	}
	return tlist
}

// 判断是是否还有子级
func hasChild(depts []*deptModel.EntityExtend, d *deptModel.EntityExtend) bool {
	return len(getChildList(depts,d)) > 0
}

// 构建前端路由所需要的菜单
func BuildDepts(depts []*deptModel.EntityExtend) []deptModel.DeptExtend {
	var deptsExtends []deptModel.DeptExtend
	for _,dept := range depts{
		var deptsExtend deptModel.DeptExtend
		deptsExtend.Id = dept.DeptId
		deptsExtend.Label = dept.DeptName
		cDeptes := dept.Children
		if len(cDeptes) > 0 {
			deptsExtend.Children = BuildDepts(cDeptes)
		}
		deptsExtends = append(deptsExtends, deptsExtend)
	}
	return deptsExtends
}
