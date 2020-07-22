// ==========================================================================
// GEAGO自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2020-02-17 14:03:51
// 生成路径: app/model/module/online/online.go
// 生成人：yunjie
// ==========================================================================
package online

import (
	"gea/app/utils/page"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
)

//新增页面请求参数
type AddReq struct {
	LoginName      string      `p:"loginName" v:"required#登录账号不能为空"`
	DeptName       string      `p:"deptName" v:"required#部门名称不能为空"`
	Ipaddr         string      `p:"ipaddr" `
	LoginLocation  string      `p:"loginLocation" `
	Browser        string      `p:"browser" `
	Os             string      `p:"os" `
	Status         string      `p:"status" v:"required#在线状态on_line在线off_line离线不能为空"`
	StartTimestamp *gtime.Time `p:"startTimestamp" `
	LastAccessTime *gtime.Time `p:"lastAccessTime" `
	ExpireTime     int         `p:"expireTime" `
}

//修改页面请求参数
type EditReq struct {
	Token          string      `p:"token" v:"required#主键ID不能为空"`
	LoginName      string      `p:"loginName" v:"required#登录账号不能为空"`
	DeptName       string      `p:"deptName" v:"required#部门名称不能为空"`
	Ipaddr         string      `p:"ipaddr" `
	LoginLocation  string      `p:"loginLocation" `
	Browser        string      `p:"browser" `
	Os             string      `p:"os" `
	Status         string      `p:"status" v:"required#在线状态on_line在线off_line离线不能为空"`
	StartTimestamp *gtime.Time `p:"startTimestamp" `
	LastAccessTime *gtime.Time `p:"lastAccessTime" `
	ExpireTime     int         `p:"expireTime" `
}

//分页请求参数
type SelectPageReq struct {
	Token          string      `p:"token"`          //用户会话id
	LoginName      string      `p:"loginName"`      //登录账号
	DeptName       string      `p:"deptName"`       //部门名称
	Ipaddr         string      `p:"ipaddr"`         //登录IP地址
	LoginLocation  string      `p:"loginLocation"`  //登录地点
	Browser        string      `p:"browser"`        //浏览器类型
	Os             string      `p:"os"`             //操作系统
	Status         string      `p:"status"`         //在线状态on_line在线off_line离线
	StartTimestamp *gtime.Time `p:"startTimestamp"` //session创建时间
	LastAccessTime *gtime.Time `p:"lastAccessTime"` //session最后访问时间
	ExpireTime     int         `p:"expireTime"`     //超时时间，单位为分钟
	BeginTime      string      `p:"beginTime"`      //开始时间
	EndTime        string      `p:"endTime"`        //结束时间
	PageNum        int         `p:"pageNum"`        //当前页码
	PageSize       int         `p:"pageSize"`       //每页数
	OrderByColumn  string      `p:"orderByColumn"`  //排序字段
	IsAsc          string      `p:"isAsc"`          //排序方式
}

//根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) ([]Entity, *page.Paging, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_user_online t")

	if param != nil {

		if param.Token != "" {
			model.Where("t.token = ?", param.Token)
		}

		if param.LoginName != "" {
			model.Where("t.login_name like ?", "%"+param.LoginName+"%")
		}

		if param.DeptName != "" {
			model.Where("t.dept_name like ?", "%"+param.DeptName+"%")
		}

		if param.Ipaddr != "" {
			model.Where("t.ipaddr = ?", param.Ipaddr)
		}

		if param.LoginLocation != "" {
			model.Where("t.login_location = ?", param.LoginLocation)
		}

		if param.Browser != "" {
			model.Where("t.browser = ?", param.Browser)
		}

		if param.Os != "" {
			model.Where("t.os = ?", param.Os)
		}

		if param.Status != "" {
			model.Where("t.status = ?", param.Status)
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	total, err := model.Count()

	if err != nil {
		return nil, nil, gerror.New("读取行数失败")
	}

	page := page.CreatePaging(param.PageNum, param.PageSize, total)

	model.Limit(page.StartNum, page.Pagesize)

	if param.OrderByColumn != "" {
		model.Order(param.OrderByColumn + " " + param.IsAsc)
	}

	var result []Entity
	model.Structs(&result)
	return result, page, nil
}

// 导出excel
func SelectListExport(param *SelectPageReq) (gdb.Result, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_user_online t")

	if param != nil {

		if param.Token != "" {
			model.Where("t.token = ?", param.Token)
		}

		if param.LoginName != "" {
			model.Where("t.login_name like ?", "%"+param.LoginName+"%")
		}

		if param.DeptName != "" {
			model.Where("t.dept_name like ?", "%"+param.DeptName+"%")
		}

		if param.Ipaddr != "" {
			model.Where("t.ipaddr = ?", param.Ipaddr)
		}

		if param.LoginLocation != "" {
			model.Where("t.login_location = ?", param.LoginLocation)
		}

		if param.Browser != "" {
			model.Where("t.browser = ?", param.Browser)
		}

		if param.Os != "" {
			model.Where("t.os = ?", param.Os)
		}

		if param.Status != "" {
			model.Where("t.status = ?", param.Status)
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	//"用户会话id","登录账号","部门名称","登录IP地址","登录地点","浏览器类型","操作系统","在线状态on_line在线off_line离线","session创建时间","session最后访问时间","超时时间，单位为分钟",
	model.Fields(" t.token ,t.login_name ,t.dept_name ,t.ipaddr ,t.login_location ,t.browser ,t.os ,t.status ,t.start_timestamp ,t.last_access_time ,t.expire_time")

	result, _ := model.All()
	return result, nil
}

//获取所有数据
func SelectListAll(param *SelectPageReq) ([]Entity, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_user_online t")

	if param != nil {

		if param.Token != "" {
			model.Where("t.token = ?", param.Token)
		}

		if param.LoginName != "" {
			model.Where("t.login_name like ?", "%"+param.LoginName+"%")
		}

		if param.DeptName != "" {
			model.Where("t.dept_name like ?", "%"+param.DeptName+"%")
		}

		if param.Ipaddr != "" {
			model.Where("t.ipaddr = ?", param.Ipaddr)
		}

		if param.LoginLocation != "" {
			model.Where("t.login_location = ?", param.LoginLocation)
		}

		if param.Browser != "" {
			model.Where("t.browser = ?", param.Browser)
		}

		if param.Os != "" {
			model.Where("t.os = ?", param.Os)
		}

		if param.Status != "" {
			model.Where("t.status = ?", param.Status)
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []Entity
	model.Structs(&result)
	return result, nil
}
