package operlog

import (
	"encoding/json"
	"fmt"
	"gea/app/model"
	"gea/app/model/monitor/oper_log"
	deptServic "gea/app/service/system/dept"
	userService "gea/app/service/system/user"
	"gea/app/utils/convert"
	"gea/app/utils/excel"
	"gea/app/utils/ip"
	"gea/app/utils/page"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

//新增记录
func Add(r *ghttp.Request, title, inContent string, outContent *model.CommonRes) error {
	user,_ := userService.GetProfileApi(r.GetInt64("jwtUid"))
	if user == nil {
		fmt.Println("用户未登陆")
		return gerror.New("用户未登陆")
	}
	var operLog oper_log.Entity
	// 清除data
	outContent.Data = ""
	outJson, _ := json.Marshal(outContent)
	outJsonStr := string(outJson)
	operLog.Title = title
	operLog.OperParam = inContent
	operLog.JsonResult = outJsonStr
	operLog.BusinessType = gconv.Int(outContent.Btype)
	if operLog.BusinessType == 0  {
		operLog.BusinessType = GetBusinessTypeByMethod(r.Method)
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
	dept := deptServic.SelectDeptById(user.DeptId)

	if dept != nil {
		operLog.DeptName = dept.DeptName
	} else {
		operLog.DeptName = ""
	}

	operLog.OperUrl = r.RequestURI
	operLog.Method = r.RequestURI
	operLog.OperIp = r.GetClientIp()

	operLog.OperLocation = ip.GetCityByIp(operLog.OperIp)
	operLog.OperTime = gtime.Now()

	_, err := operLog.Insert()
	return err
}

func GetBusinessTypeByMethod(method string) int{
	switch method {
	case "POST":
		return gconv.Int(model.Buniss_Add)
	case "PUT":
		return gconv.Int(model.Buniss_Edit)
	case "DELETE":
		return gconv.Int(model.Buniss_Del)
	default:
		return gconv.Int(model.Buniss_Other)
	}
}
// 根据条件分页查询用户列表
func SelectPageList(param *oper_log.SelectPageReq) ([]oper_log.Entity, *page.Paging, error) {
	return oper_log.SelectPageList(param)
}

//根据主键查询用户信息
func SelectRecordById(id int64) (*oper_log.Entity, error) {
	return oper_log.FindOne("oper_id", id)
}

//根据主键删除用户信息
func DeleteRecordById(id int64) bool {
	result, err := oper_log.Delete("oper_id", id)
	if err == nil {
		affected, _ := result.RowsAffected()
		if affected > 0 {
			return true
		}
	}

	return false
}

//批量删除记录
func DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := oper_log.Delete("oper_id in (?)", idarr)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

//清空记录
func DeleteRecordAll() int64 {
	result, err := oper_log.Delete()
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

// 导出excel
func Export(param *oper_log.SelectPageReq) (string, error) {
	result, err := oper_log.SelectExportList(param)
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
