package config

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"gea/app/utils/page"
)

// Fill with you ideas below.
//新增页面请求参数
type AddReq struct {
	ConfigName  string `p:"configName"  v:"required#参数名称不能为空"`
	ConfigKey   string `p:"configKey"  v:"required#参数键名不能为空"`
	ConfigValue string `p:"configValue"  v:"required#参数键值不能为空"`
	ConfigType  string `p:"configType"    v:"required#系统内置不能为空"`
	Remark      string `p:"remark"`
}

//修改页面请求参数
type EditReq struct {
	ConfigId    int64  `p:"configId" v:"required#主键ID不能为空"`
	ConfigName  string `p:"configName"  v:"required#参数名称不能为空"`
	ConfigKey   string `p:"configKey"  v:"required#参数键名不能为空"`
	ConfigValue string `p:"configValue"  v:"required#参数键值不能为空"`
	ConfigType  string `p:"configType"    v:"required#系统内置不能为空"`
	Remark      string `p:"remark"`
}

//分页请求参数
type SelectPageReq struct {
	ConfigName string `p:"configName"` //参数名称
	ConfigKey  string `p:"configKey"`  //参数键名
	ConfigType string `p:"configType"` //状态
	BeginTime  string `p:"beginTime"`  //开始时间
	EndTime    string `p:"endTime"`    //结束时间
	PageNum    int    `p:"pageNum"`    //当前页码
	PageSize   int    `p:"pageSize"`   //每页数
}

//检查参数键名请求参数
type CheckConfigKeyReq struct {
	ConfigId  int64  `p:"configId"  v:"required#ID不能为空"`
	ConfigKey string `p:"configKey"  v:"required#参数键名不能为空"`
}

//检查参数键名请求参数
type CheckPostCodeALLReq struct {
	ConfigKey string `p:"configKey"  v:"required#参数键名不能为空"`
}

//根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) ([]Entity, *page.Paging, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_config t")

	if param != nil {
		if param.ConfigName != "" {
			model.Where("t.config_name like ?", "%"+param.ConfigName+"%")
		}

		if param.ConfigType != "" {
			model.Where("t.status = ", param.ConfigType)
		}

		if param.ConfigKey != "" {
			model.Where("t.config_key like ?", "%"+param.ConfigKey+"%")
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

	model := db.Table("sys_config t")

	if param != nil {
		if param.ConfigName != "" {
			model.Where("t.config_name like ?", "%"+param.ConfigName+"%")
		}

		if param.ConfigType != "" {
			model.Where("t.status = ", param.ConfigType)
		}

		if param.ConfigKey != "" {
			model.Where("t.config_key like ?", "%"+param.ConfigKey+"%")
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	//"参数主键","参数名称","参数键名","参数键值","系统内置（Y是 N否）","状态"
	model.Fields("t.config_id,t.config_name,t.config_key,t.config_value,t.config_type")

	result, _ := model.All()
	return result, nil
}

//获取所有数据
func SelectListAll(param *SelectPageReq) ([]Entity, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_config t")

	if param != nil {
		if param.ConfigName != "" {
			model.Where("t.config_name like ?", "%"+param.ConfigName+"%")
		}

		if param.ConfigType != "" {
			model.Where("t.status = ", param.ConfigType)
		}

		if param.ConfigKey != "" {
			model.Where("t.config_key like ?", "%"+param.ConfigKey+"%")
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

//校验参数键名是否唯一
func CheckPostCodeUnique(configKey string, configId int64) (*Entity, error) {
	return FindOne("config_id !=? and config_key=?", configId, configKey)
}

//校验参数键名是否唯一
func CheckPostCodeUniqueAll(configKey string) (*Entity, error) {
	return FindOne("config_key=?", configKey)
}
