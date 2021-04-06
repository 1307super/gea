package service

import (
	"fmt"
	"gea/app/dao"
	"gea/app/model"
	"gea/app/system/admin/internal/define"
	"gea/app/utils/convert"
	"gea/app/utils/excel"
	"gea/app/utils/ip"
	"gea/app/utils/page"
	"github.com/gogf/gf/os/gtime"
	"github.com/mssola/user_agent"
)

var Logininfor = &logininforService{}

type logininforService struct{}

// 根据条件分页查询列表
func (s *logininforService) GetList(param *define.LogininforApiSelectPageReq) *define.LogininforServiceList {

	m := dao.SysLogininfor.As("t")

	if param != nil {
		if param.LoginName != "" {
			m = m.Where("login_name like ?", "%"+param.LoginName+"%")
		}

		if param.Ipaddr != "" {
			m = m.Where("ipaddr like ?", "%"+param.Ipaddr+"%")
		}

		if param.Status != "" {
			m = m.Where("status = ?", param.Status)
		}

		if param.BeginTime != "" {
			m = m.Where("date_format(login_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}
		if param.EndTime != "" {
			m = m.Where("date_format(login_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
		}
	}
	total, err := m.Count()
	if err != nil {
		return nil
	}
	page := page.CreatePaging(param.PageNum, param.PageSize, total)
	m = m.Fields("info_id,login_name,ipaddr,login_location,browser,os,status,msg,login_time")
	m = m.Order("login_time desc")
	m = m.Limit(page.StartNum, page.Pagesize)
	result := &define.LogininforServiceList{
		Page:  page.PageNum,
		Size:  page.Pagesize,
		Total: page.Total,
	}
	if err = m.Structs(&result.List); err != nil {
		return nil
	}
	return result
}

func (s *logininforService) Create(status, username, ipaddr, userAgent, msg string) {
	var logininfor model.SysLogininfor
	logininfor.Status = status
	logininfor.LoginName = username
	logininfor.Ipaddr = ipaddr
	ua := user_agent.New(userAgent)
	os := ua.OS()
	browser, _ := ua.Browser()
	loginLocation := ip.GetCityByIp(ipaddr)
	logininfor.Os = os
	logininfor.Browser = browser
	logininfor.LoginTime = gtime.Now()
	logininfor.LoginLocation = loginLocation
	logininfor.Msg = msg
	dao.SysLogininfor.Insert(logininfor)
}

//批量删除记录
func (s *logininforService) Delete(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := dao.SysLogininfor.Delete(fmt.Sprintf("%s in(?)", dao.SysLogininfor.Columns.InfoId), idarr)
	if err != nil {
		return 0
	}
	nums, _ := result.RowsAffected()
	return nums
}

//清空记录
func (s *logininforService) Clean() int64 {
	result, err := dao.SysLogininfor.Delete(fmt.Sprintf("%s > ?", dao.SysLogininfor.Columns.InfoId), "0")
	if err != nil {
		return 0
	}
	nums, _ := result.RowsAffected()

	return nums
}

// 导出excel
func (s *logininforService) Export(param *define.LogininforApiSelectPageReq) (string, error) {
	//"访问编号", "用户名称", "登录地址", "登录地点", "浏览器", "操作系统", "登录状态", "操作信息", "登录日期"
	m := dao.SysLogininfor.Fields("info_id,login_name,ipaddr,login_location,browser,os,status,msg,login_time")
	if param != nil {
		if param.LoginName != "" {
			m = m.Where("login_name like ?", "%"+param.LoginName+"%")
		}

		if param.Ipaddr != "" {
			m = m.Where("ipaddr like ?", "%"+param.Ipaddr+"%")
		}

		if param.Status != "" {
			m = m.Where("status = ?", param.Status)
		}

		if param.BeginTime != "" {
			m = m.Where("date_format(login_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}
		if param.EndTime != "" {
			m = m.Where("date_format(login_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
		}
	}
	result, err := m.M.All()
	if err != nil {
		return "", err
	}
	head := []string{"访问编号", "用户名称", "登录地址", "登录地点", "浏览器", "操作系统", "登录状态", "操作信息", "登录日期"}
	key := []string{"info_id", "login_name", "ipaddr", "login_location", "browser", "os", "status", "msg", "login_time"}
	url, err := excel.DownlaodExcel(head, key, result)
	if err != nil {
		return "", err
	}
	return url, nil
}
