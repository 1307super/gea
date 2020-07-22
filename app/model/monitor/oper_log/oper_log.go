package oper_log

import (
	"gea/app/utils/page"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
)

//
//查询列表请求参数
type SelectPageReq struct {
	Title         string `p:"title"`         //系统模块
	OperName      string `p:"operName"`      //操作人员
	BusinessType  string  `p:"businessType"`  //操作类型
	Status        string `p:"status"`        //操作类型
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

	model := db.Table("sys_oper_log")

	if param != nil {
		if param.Title != "" {
			model.Where("title like ?", "%"+param.Title+"%")
		}

		if param.OperName != "" {
			model.Where("oper_name like ?", "%"+param.OperName+"%")
		}

		if param.Status != "" {
			model.Where("status = ?", param.Status)
		}

		if param.BusinessType != "" {
			model.Where("business_type = ?", param.BusinessType)
		}

		if param.BeginTime != "" {
			model.Where("date_format(oper_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(oper_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
		}
	}

	total, err := model.Count()

	if err != nil {
		return nil, nil, gerror.New("读取行数失败")
	}

	page := page.CreatePaging(param.PageNum, param.PageSize, total)

	model.Fields("oper_id, title, business_type, method, request_method, operator_type, oper_name, dept_name, oper_url, oper_ip, oper_location, oper_param, json_result, status, error_msg, oper_time")

	model.Order("oper_time desc")

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

	model := db.Table("sys_oper_log")

	if param != nil {
		if param.Title != "" {
			model.Where("title like ?", "%"+param.Title+"%")
		}

		if param.OperName != "" {
			model.Where("oper_name like ?", "%"+param.OperName+"%")
		}

		if param.Status != "" {
			model.Where("status = ?", param.Status)
		}

		if param.BusinessType != "" {
			model.Where("business_type = ?", param.BusinessType)
		}

		if param.BeginTime != "" {
			model.Where("date_format(oper_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(oper_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
		}
	}

	//"日志主键", "模块标题", "业务类型", "方法名称", "请求方式", "操作类别", "操作人员", "部门名称", "请求URL", "主机地址", "操作地点", "请求参数", "返回参数", "操作状态","操作时间"
	model.Fields("oper_id, title, business_type, method, request_method, operator_type, oper_name, dept_name, oper_url, oper_ip, oper_location, oper_param, json_result, status, error_msg, oper_time")

	result, err := model.All()
	if err != nil {
		return nil, err
	}
	return result, nil
}
