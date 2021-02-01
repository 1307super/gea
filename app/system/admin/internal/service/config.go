package service

import (
	"context"
	"fmt"
	"gea/app/dao"
	"gea/app/model"
	"gea/app/shared"
	"gea/app/system/admin/internal/define"
	"gea/app/utils/convert"
	"gea/app/utils/excel"
	"gea/app/utils/page"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var Config = &configService{}

type configService struct{}

//根据主键查询数据
func (s *configService) Info(id int64) (*model.SysConfig, error) {
	return dao.SysConfig.FindOne(dao.SysConfig.Columns.ConfigId, id)
}

// 获取分页数据
func (s *configService) GetList(param *define.ConfigApiSelectPageReq) *define.ConfigServiceList {
	m := dao.SysConfig.As("t")
	if param != nil {
		if param.ConfigName != "" {
			m = m.Where("t.config_name like ?", "%"+param.ConfigName+"%")
		}

		if param.ConfigType != "" {
			m = m.Where("t.status = ", param.ConfigType)
		}

		if param.ConfigKey != "" {
			m = m.Where("t.config_key like ?", "%"+param.ConfigKey+"%")
		}

		if param.BeginTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}
	total, err := m.Count()
	if err != nil {
		return nil
	}
	page := page.CreatePaging(param.PageNum, param.PageSize, total)
	m = m.Limit(page.StartNum, page.Pagesize)
	result := &define.ConfigServiceList{
		Page:  page.PageNum,
		Size:  page.Pagesize,
		Total: page.Total,
	}
	if err = m.Structs(&result.List); err != nil {
		return nil
	}
	return result
}

//添加数据
func (s *configService) Create(ctx context.Context, req *define.ConfigApiCreateReq) (int64, error) {
	if s.CheckConfigKeyUniqueAll(req.ConfigKey) {
		return 0, gerror.New("参数键名已存在")
	}
	user := shared.Context.Get(ctx).User
	var entity model.SysConfig
	entity.CreateTime = gtime.Now()
	entity.CreateBy = user.LoginName
	var editReq *define.ConfigApiUpdateReq
	gconv.Struct(req, &editReq)
	return s.save(&entity, editReq)
}

//修改数据
func (s *configService) Update(ctx context.Context, req *define.ConfigApiUpdateReq) (int64, error) {
	if s.CheckConfigKeyUnique(req.ConfigKey, req.ConfigId) {
		return 0, gerror.New("参数键名已存在")
	}
	user := shared.Context.Get(ctx).User
	entity, err := dao.SysConfig.FindOne(dao.SysConfig.Columns.ConfigId, req.ConfigId)
	if err != nil {
		return 0, err
	}
	if entity == nil {
		return 0, gerror.New("数据不存在")
	}
	entity.UpdateTime = gtime.Now()
	entity.UpdateBy = user.LoginName
	return s.save(entity, req)
}

func (s *configService) save(config *model.SysConfig, req *define.ConfigApiUpdateReq) (int64, error) {
	var (
		rs  int64
		err error
	)
	config.ConfigName = req.ConfigName
	config.ConfigKey = req.ConfigKey
	config.ConfigType = req.ConfigType
	config.ConfigValue = req.ConfigValue
	config.Remark = req.Remark
	result, err := dao.SysConfig.Data(config).Save()
	if err != nil {
		return 0, err
	}
	if config.ConfigId == 0 {
		// 新增
		rs, err = result.LastInsertId()
	} else {
		rs, err = result.RowsAffected()
	}
	if err != nil {
		return 0, err
	}
	if rs > 0 {
		gcache.Set(config.ConfigKey, config.ConfigValue, 0)
	}
	return rs, nil
}

//批量删除数据记录
func (s *configService) Delete(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	field := fmt.Sprintf("%s in(?)", dao.SysConfig.Columns.ConfigId)
	list, _ := dao.SysConfig.FindAll(field, idarr)
	result, err := dao.SysConfig.Delete(field, idarr)
	if err != nil {
		return 0
	}
	if len(list) > 0 {
		for _, item := range list {
			gcache.Remove(item.ConfigKey)
		}
	}
	nums, _ := result.RowsAffected()
	return nums
}

// 导出excel
func (s *configService) Export(param *define.ConfigApiSelectPageReq) (string, error) {
	m := dao.SysConfig.As("t")
	if param != nil {
		if param.ConfigName != "" {
			m = m.Where("t.config_name like ?", "%"+param.ConfigName+"%")
		}

		if param.ConfigType != "" {
			m = m.Where("t.status = ", param.ConfigType)
		}

		if param.ConfigKey != "" {
			m = m.Where("t.config_key like ?", "%"+param.ConfigKey+"%")
		}

		if param.BeginTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}
	//"参数主键","参数名称","参数键名","参数键值","系统内置（Y是 N否）","状态"
	m = m.Fields("t.config_id,t.config_name,t.config_key,t.config_value,t.config_type")
	result, err := m.M.All()
	if err != nil {
		return "", err
	}
	head := []string{"参数主键", "参数名称", "参数键名", "参数键值", "系统内置（Y是 N否）", "状态"}
	key := []string{"config_id", "config_name", "config_key", "config_value", "config_type"}
	url, err := excel.DownlaodExcel(head, key, result)
	if err != nil {
		return "", err
	}
	return url, nil
}

//根据键获取值
func (s *configService) GetValueByKey(key string) string {
	resultStr := ""
	result, _ := gcache.Get(key)
	if result == nil {
		configRecord, err := dao.SysConfig.FindOne(dao.SysConfig.Columns.ConfigKey, key)
		if err != nil {
			return ""
		}
		resultStr = configRecord.ConfigValue
		gcache.Set(key, resultStr, 0)
	} else {
		resultStr = result.(string)
	}

	return resultStr
}

//检查角色名是否唯一
func (s *configService) CheckConfigKeyUniqueAll(configKey string) bool {
	entity, err := dao.SysConfig.FindOne(dao.SysConfig.Columns.ConfigKey, configKey)
	if err != nil {
		return true
	}
	if entity != nil && entity.ConfigId > 0 {
		return true
	}
	return false
}

//检查岗位名称是否唯一
func (s *configService) CheckConfigKeyUnique(configKey string, configId int64) bool {
	entity, err := dao.SysConfig.FindOne(g.Map{
		fmt.Sprintf("%s != ?", dao.SysConfig.Columns.ConfigId): configId,
		dao.SysConfig.Columns.ConfigKey:                        configKey,
	})
	if err != nil {
		return true
	}
	if entity != nil && entity.ConfigId > 0 {
		return true
	}
	return false
}
