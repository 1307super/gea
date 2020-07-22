package logininfor

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"gea/app/utils/page"
)

// Fill with you ideas below.
//查询列表请求参数
type SelectPageReq struct {
	LoginName     string `p:"loginName"`     //登陆名
	Status        string `p:"status"`        //状态
	Ipaddr        string `p:"ipaddr"`        //登录地址
	BeginTime     string `p:"beginTime"`     //数据范围
	EndTime       string `p:"endTime"`       //开始时间
	PageNum       int    `p:"pageNum"`       //当前页码
	PageSize      int    `p:"pageSize"`      //每页数
	OrderByColumn string `p:"orderByColumn"` //排序字段
	IsAsc         string `p:"isAsc"`         //排序方式
}

// 根据条件分页查询用户列表
func SelectPageList(param *SelectPageReq) ([]Entity, *page.Paging, error) {
	db, err := gdb.Instance()
	if err != nil {
		return nil, nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_logininfor")

	if param != nil {
		if param.LoginName != "" {
			model.Where("login_name like ?", "%"+param.LoginName+"%")
		}

		if param.Ipaddr != "" {
			model.Where("ipaddr like ?", "%"+param.Ipaddr+"%")
		}

		if param.Status != "" {
			model.Where("status = ?", param.Status)
		}

		if param.BeginTime != "" {
			model.Where("date_format(login_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(login_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
		}
	}

	total, err := model.Count()

	if err != nil {
		return nil, nil, gerror.New("读取行数失败")
	}

	page := page.CreatePaging(param.PageNum, param.PageSize, total)

	model.Fields("info_id,login_name,ipaddr,login_location,browser,os,status,msg,login_time")

	model.Order("login_time desc")

	model.Limit(page.StartNum, page.Pagesize)

	var result []Entity

	err = model.Structs(&result)
	return result, page, nil
}

// 导出excel
func SelectExportList(param *SelectPageReq) (gdb.Result, error) {
	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_logininfor")

	if param != nil {
		if param.LoginName != "" {
			model.Where("login_name like ?", "%"+param.LoginName+"%")
		}

		if param.Ipaddr != "" {
			model.Where("ipaddr like ?", "%"+param.Ipaddr+"%")
		}

		if param.Status != "" {
			model.Where("status = ?", param.Status)
		}

		if param.BeginTime != "" {
			model.Where("date_format(login_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(login_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
		}
	}

	//"访问编号", "登录名称", "登录地址", "登录地点", "浏览器", "操作系统", "登录状态", "操作信息", "登录时间"
	model.Fields("info_id,login_name,ipaddr,login_location,browser,os,status,msg,login_time")

	result, err := model.All()
	if err != nil {
		return nil, err
	}
	return result, nil
}
