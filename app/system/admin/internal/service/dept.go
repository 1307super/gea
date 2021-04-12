package service

import (
	"context"
	"gea/app/dao"
	"gea/app/model"
	"gea/app/shared"
	"gea/app/system/admin/internal/define"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

var Dept = &deptService{}

type deptService struct{}

//查询部门管理数据
func (s *deptService)GetAll(ctx context.Context,param *define.DeptApiSelectPageReq) ([]*model.SysDeptExtend, error) {
	m := dao.SysDept.As("d").Where(dao.SysDept.Columns.DelFlag,"0")
	var result []*model.SysDeptExtend
	if param != nil {
		if param.ParentId > 0 {
			m = m.Where("d.parent_id = ?", param.ParentId)
		}
		if param.DeptName != "" {
			m = m.Where("d.dept_name like ?", "%"+param.DeptName+"%")
		}
		if param.Status != "" {
			m = m.Where("d.status = ?", param.Status)
		}
	}
	// 获取资源权限
	dataScope := DataScopeFilter(ctx,"d","")
	if dataScope != "" {
		m = m.Where(dataScope)
	}
	m = m.Order("d.parent_id, d.order_num")
	err := m.Structs(&result)
	return s.BuildDeptTree(result), err
}


//新增保存信息
func (s *deptService)Create(ctx context.Context, req *define.DeptApiCreateReq) (int64, error) {
	if s.CheckDeptNameUniqueAll(req.DeptName, req.ParentId) {
		return 0, gerror.New("部门名称已存在")
	}
	user := shared.Context.Get(ctx).User
	var dept model.SysDept
	if req.ParentId != 0 {
		pdept, err := dao.SysDept.FindOne(dao.SysDept.Columns.DeptId,req.ParentId)
		if err == nil && pdept != nil {
			if pdept.Status != "0" {
				return 0, gerror.New("部门停用，不允许新增")
			} else {
				dept.Ancestors = pdept.Ancestors + "," + gconv.String(pdept.DeptId)
			}
		}
	}
	dept.DeptName = req.DeptName
	dept.Status = req.Status
	dept.ParentId = req.ParentId
	dept.DelFlag = "0"
	dept.Email = req.Email
	dept.Leader = req.Leader
	dept.Phone = req.Phone
	dept.OrderNum = req.OrderNum
	dept.CreateBy = user.UserExtend.LoginName
	dept.CreateTime = gtime.Now()
	rs, err := dao.SysDept.Data(dept).Insert()
	if err != nil {
		return 0, err
	}
	did, _ := rs.LastInsertId()
	return did, nil
}

//修改保存信息
func (s *deptService)Update(ctx context.Context,req *define.DeptApiUpdateReq) (int64, error) {
	if s.CheckDeptNameUnique(req.DeptName, req.DeptId, req.ParentId) {
		return 0, gerror.New("部门名称已存在")
	}
	user := shared.Context.Get(ctx).User
	dept, err := dao.SysDept.FindOne(dao.SysDept.Columns.DeptId,req.DeptId)
	if err != nil || dept == nil || dept.DeptId <= 0 {
		return 0, gerror.New("数据不存在")
	}
	var pdept *model.SysDept
	if dept.ParentId == 0 {
		// 顶级部门
		pdept = dept
	}else{
		pdept, err = dao.SysDept.FindOne(dao.SysDept.Columns.DeptId,req.ParentId)
	}
	if pdept != nil {
		if pdept.Status != "0" {
			return 0, gerror.New("部门停用，不允许新增")
		} else {
			newAncestors := pdept.Ancestors + "," + gconv.String(pdept.DeptId)
			s.UpdateDeptChildren(dept.DeptId, newAncestors, dept.Ancestors)

			dept.DeptName = req.DeptName
			dept.Status = req.Status
			dept.ParentId = req.ParentId
			dept.DelFlag = "0"
			dept.Email = req.Email
			dept.Leader = req.Leader
			dept.Phone = req.Phone
			dept.OrderNum = req.OrderNum
			dept.UpdateBy = user.UserExtend.LoginName
			dept.UpdateTime = gtime.Now()
			dao.SysDept.Data(dept).Save()
			return 1, nil
		}

	} else {
		return 0, gerror.New("父部门不能为空")
	}
}

//修改子元素关系
func (s *deptService)UpdateDeptChildren(deptId int64, newAncestors, oldAncestors string) {
	deptList, err := dao.SysDept.FindAll("find_in_set(?, ancestors)",deptId)
	if err != nil {
		return
	}
	if deptList == nil || len(deptList) <= 0 {
		return
	}
	for _, tmp := range deptList {
		tmp.Ancestors = strings.ReplaceAll(tmp.Ancestors, oldAncestors, newAncestors)
	}

	ancestors := " case dept_id"
	idStr := ""

	for _, dept := range deptList {
		ancestors += " when " + gconv.String(dept.DeptId) + " then " + dept.Ancestors
		if idStr == "" {
			idStr = gconv.String(dept.DeptId)
		} else {
			idStr += "," + gconv.String(dept.DeptId)
		}
	}
	ancestors += " end "
	dao.SysDept.Update("ancestors = ?", "dept_id in(?)", ancestors, idStr)
}

//删除部门管理信息
func (s *deptService)Delete(deptId int64) int64 {
	if deptId == 0{
		return 0
	}
	rs, err := dao.SysDept.Where(dao.SysDept.Columns.DeptId,deptId).Data(g.Map{dao.SysDept.Columns.DelFlag:"2"}).Update()
	if err != nil {
		return 0
	}
	rows, _ := rs.RowsAffected()
	return rows
}

// 部门信息
func (s *deptService)Info(deptId int64) (*model.SysDept) {
	if deptId == 0{
		return nil
	}
	var result *model.SysDept
	m := dao.SysDept.As("d").Fields("d.dept_id, d.parent_id, d.ancestors, d.dept_name, d.order_num, d.leader, d.phone, d.email, d.status,(select dept_name from sys_dept where dept_id = d.parent_id) parent_name").Where("d.dept_id = ?", deptId)
	if err := m.Struct(&result);err != nil{
		return nil
	}
	return result
}

//加载部门列表树
func (s *deptService)DeptTree(ctx context.Context,parentId int64, deptName, status string) ([]*model.SysDeptExtend, error) {
	list, err := s.GetAll(ctx,&define.DeptApiSelectPageReq{
		ParentId:  parentId,
		DeptName:  "",
		Status:    "",
	})
	if err != nil {
		return nil, err
	}
	return list,nil
}

// 加载角色部门（数据权限）列表树
func (s *deptService)RoleDeptTreeData(ctx context.Context,roleId int64) (*define.DeptServiceRoleTreeData,error) {
	// 1 先查出所有部门
	result, err := s.DeptTree(ctx,0,"","")
	if err != nil {
		return nil,gerror.New("获取角色权限失败")
	}
	// 2 查出权限
	deptIds, err := s.GetRoleDeptIds(roleId)
	if err != nil {
		return nil,gerror.New("获取角色权限失败")
	}
	return &define.DeptServiceRoleTreeData{
		Depts:s.BuildDepts(result),
		CheckedKeys:deptIds,
	},nil
}

//根据角色ID查询部门
func (s *deptService)GetRoleDeptIds(roleId int64) (g.Array, error) {
	var deptIds g.Array
	roleDepts, err := dao.SysRoleDept.FindAll(dao.SysRoleDept.Columns.RoleId,roleId)
	if err != nil {
		return nil,err
	}
	for _,roleDept := range roleDepts {
		deptIds = append(deptIds, roleDept.DeptId)
	}

	return deptIds, nil
}

func (s *deptService)BuildDeptTree(depts []*model.SysDeptExtend)[]*model.SysDeptExtend {
	var returnList []*model.SysDeptExtend
	tempList := garray.NewIntArray(true)
	for _, dept := range depts  {
		tempList.Append(gconv.Int(dept.DeptId))
	}
	for _,dept := range depts{
		if !tempList.Contains(gconv.Int(dept.ParentId)) {
			s.recursionFn(depts,dept)
			returnList = append(returnList, dept)
		}
	}
	if len(returnList) == 0 {
		returnList = depts
	}
	return returnList
}

// 递归
func (s *deptService)recursionFn(depts []*model.SysDeptExtend, d *model.SysDeptExtend){
	childMenus := s.getChildList(depts,d)
	d.Children = childMenus
	for _,childMenu := range childMenus  {
		if s.hasChild(depts,childMenu){
			for _,menu := range childMenus {
				s.recursionFn(depts,menu)
			}
		}
	}
}

// 获取子级列表
func (s *deptService)getChildList(depts []*model.SysDeptExtend, d *model.SysDeptExtend)[]*model.SysDeptExtend {
	var tlist []*model.SysDeptExtend
	for _,dept := range depts  {
		if dept.ParentId == d.DeptId {
			tlist = append(tlist, dept)
		}
	}
	return tlist
}

// 判断是是否还有子级
func (s *deptService)hasChild(depts []*model.SysDeptExtend, d *model.SysDeptExtend) bool {
	return len(s.getChildList(depts,d)) > 0
}

// 构建前端路由所需要的菜单
func (s *deptService)BuildDepts(depts []*model.SysDeptExtend) []model.DeptExtend {
	var deptsExtends []model.DeptExtend
	for _,dept := range depts{
		var deptsExtend model.DeptExtend
		deptsExtend.Id = dept.DeptId
		deptsExtend.Label = dept.DeptName
		cDeptes := dept.Children
		if len(cDeptes) > 0 {
			deptsExtend.Children = s.BuildDepts(cDeptes)
		}
		deptsExtends = append(deptsExtends, deptsExtend)
	}
	return deptsExtends
}


//校验部门名称是否唯一
func (s *deptService)CheckDeptNameUniqueAll(deptName string, parentId int64) bool {
	dept, err := dao.SysDept.FindOne(g.Map{
		dao.SysDept.Columns.DeptName: deptName,
		dao.SysDept.Columns.ParentId: parentId,
	})
	if err != nil {
		return true
	}
	if dept != nil && dept.DeptId > 0 {
		return true
	} else {
		return false
	}
}

//校验部门名称是否唯一
func (s *deptService)CheckDeptNameUnique(deptName string, deptId, parentId int64) bool {
	dept, err := dao.SysDept.FindOne(g.Map{
		dao.SysDept.Columns.DeptName: deptName,
		dao.SysDept.Columns.ParentId: parentId,
	})
	if err != nil {
		return true
	}
	if dept != nil && dept.DeptId != deptId {
		return true
	} else {
		return false
	}
}