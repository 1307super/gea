package service

import (
	"encoding/json"
	"fmt"
	"gea/app/dao"
	"gea/app/model"
	"gea/app/shared"
	"gea/app/system/admin/internal/define"
	"gea/app/utils/convert"
	"gea/app/utils/excel"
	"gea/app/utils/ip"
	"gea/app/utils/page"
	"gea/app/utils/response"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var OperLog = &operLogService{}

type operLogService struct{}

// 分页列表
func (s *operLogService) GetList(param *define.OperlogApiSelectPageReq) *define.OperlogServiceList {
	m := dao.SysOperLog.As("t")
	if param != nil {
		if param.Title != "" {
			m = m.Where("t.title like ?", "%"+param.Title+"%")
		}
		if param.OperName != "" {
			m = m.Where("t.oper_name like ?", "%"+param.OperName+"%")
		}
		if param.Status != "" {
			m = m.Where("t.status = ?", param.Status)
		}
		if param.BusinessType != "" {
			m = m.Where("t.business_type = ?", param.BusinessType)
		}

		if param.BeginTime != "" {
			m = m.Where("date_format(t.oper_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}

		if param.EndTime != "" {
			m = m.Where("date_format(t.oper_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
		}
	}
	total, err := m.Count()
	m = m.Fields("oper_id, title, business_type, method, request_method, operator_type, oper_name, dept_name, oper_url, oper_ip, oper_location, oper_param, json_result, status, error_msg, oper_time")
	if err != nil {
		return nil
	}
	page := page.CreatePaging(param.PageNum, param.PageSize, total)
	m = m.Order("t.oper_time desc")
	m = m.Limit(page.StartNum, page.Pagesize)
	result := &define.OperlogServiceList{
		Page:  page.PageNum,
		Size:  page.Pagesize,
		Total: page.Total,
	}
	if err = m.Structs(&result.List); err != nil {
		return nil
	}
	return result
}

//新增记录
func (s *operLogService) Create(r *ghttp.Request, title, inContent string, outContent *response.CommonRes) error {
	user := shared.Context.Get(r.Context()).User
	if user == nil {
		return gerror.New("用户未登陆")
	}
	var operLog model.SysOperLog
	// 清除data
	outContent.Data = ""
	outJson, _ := json.Marshal(outContent)
	outJsonStr := string(outJson)
	operLog.Title = title
	operLog.OperParam = inContent
	operLog.JsonResult = outJsonStr
	operLog.BusinessType = gconv.Int(outContent.Btype)
	if operLog.BusinessType == 0 {
		operLog.BusinessType = s.GetBusinessTypeByMethod(r.Method)
	}
	//操作类别（0其它 1后台用户 2手机端用户）
	operLog.OperatorType = 1
	//操作状态（0正常 1异常）
	if outContent.Code == 0 {
		operLog.Status = 0
	} else {
		operLog.Status = 1
	}
	operLog.OperName = user.LoginName
	operLog.RequestMethod = r.Method
	//获取用户部门
	//dept := deptServic.SelectDeptById(user.DeptId)

	//if user.DeptName != "" {
	operLog.DeptName = user.Dept.DeptName
	//} else {
	//	operLog.DeptName = ""
	//}

	operLog.OperUrl = r.RequestURI
	operLog.Method = r.RequestURI
	operLog.OperIp = r.GetClientIp()

	operLog.OperLocation = ip.GetCityByIp(operLog.OperIp)
	operLog.OperTime = gtime.Now()

	_, err := dao.SysOperLog.Data(operLog).Insert()
	return err
}

//批量删除记录
func (s *operLogService) Delete(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := dao.SysOperLog.Delete(fmt.Sprintf("%s in(?)", dao.SysOperLog.Columns.OperId), idarr)
	if err != nil {
		return 0
	}
	nums, _ := result.RowsAffected()
	return nums
}

//清空记录
func (s *operLogService) Clean() int64 {
	result, err := dao.SysOperLog.Delete(fmt.Sprintf("%s > ?", dao.SysOperLog.Columns.OperId), "0")
	if err != nil {
		return 0
	}
	nums, _ := result.RowsAffected()
	return nums
}

// 导出excel
func (s *operLogService) Export(param *define.OperlogApiSelectPageReq) (string, error) {
	//"日志主键", "模块标题", "业务类型", "方法名称", "请求方式", "操作类别", "操作人员", "部门名称", "请求URL", "主机地址", "操作地点", "请求参数", "返回参数", "操作状态","操作时间"
	m := dao.SysOperLog.Fields("oper_id, title, business_type, method, request_method, operator_type, oper_name, dept_name, oper_url, oper_ip, oper_location, oper_param, json_result, status, error_msg, oper_time")
	if param != nil {
		if param.Title != "" {
			m = m.Where("title like ?", "%"+param.Title+"%")
		}

		if param.OperName != "" {
			m = m.Where("oper_name like ?", "%"+param.OperName+"%")
		}

		if param.Status != "" {
			m = m.Where("status = ?", param.Status)
		}

		if param.BusinessType != "" {
			m = m.Where("business_type = ?", param.BusinessType)
		}

		if param.BeginTime != "" {
			m = m.Where("date_format(oper_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}

		if param.EndTime != "" {
			m = m.Where("date_format(oper_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
		}
	}
	result, err := m.M.All()
	if err != nil {
		return "", err
	}
	head := []string{"日志主键", "模块标题", "业务类型", "方法名称", "请求方式", "操作类别", "操作人员", "部门名称", "请求URL", "主机地址", "操作地点", "请求参数", "返回参数", "操作状态", "操作时间"}
	key := []string{"oper_id", "title", "business_type", "method", "request_method", "operator_type", "oper_name", "dept_name", "oper_url", "oper_ip", "oper_location", "oper_param", "json_result", "status", "error_msg", "oper_time"}
	url, err := excel.DownlaodExcel(head, key, result)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (s *operLogService) GetBusinessTypeByMethod(method string) int {
	switch method {
	case "POST":
		return gconv.Int(response.Buniss_Add)
	case "PUT":
		return gconv.Int(response.Buniss_Edit)
	case "DELETE":
		return gconv.Int(response.Buniss_Del)
	default:
		return gconv.Int(response.Buniss_Other)
	}
}
