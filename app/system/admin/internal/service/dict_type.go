package service

import (
	"context"
	"gea/app/dao"
	"gea/app/model"
	"gea/app/shared"
	"gea/app/system/admin/internal/define"
	"gea/app/utils/convert"
	"gea/app/utils/excel"
	"gea/app/utils/page"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var DictType = &dictTypeService{}

type dictTypeService struct{}

func (s *dictTypeService)Info(id int64) (*model.SysDictType,error) {
	return dao.SysDictType.FindOne(dao.SysDictType.Columns.DictId,id)
}
//根据条件分页查询数据
func (s *dictTypeService)GetList(param *define.DictTypeApiSelectPageReq) *define.DictTypeServiceList{

	m := dao.SysDictType.As("t")

	if param != nil {
		if param.DictName != "" {
			m = m.Where("t.dict_name like ?", "%"+param.DictName+"%")
		}

		if param.DictType != "" {
			m = m.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.Status != "" {
			m = m.Where("t.status = ", param.Status)
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
	if param.OrderByColumn != "" {
		m = m.Order(param.OrderByColumn + " " + param.IsAsc)
	}
	result := &define.DictTypeServiceList{
		Page:page.PageNum,
		Total:page.Total,
		Size:page.Pagesize,
	}
	if err = m.Structs(&result.List); err != nil {
		return nil
	}
	return result
}

//获取所有数据
func (s *dictTypeService)GetAll(param *define.DictTypeApiSelectPageReq) ([]model.SysDictType, error) {

	m := dao.SysDictType.As("t")
	if param != nil {
		if param.DictName != "" {
			m = m.Where("t.dict_name like ?", "%"+param.DictName+"%")
		}

		if param.DictType != "" {
			m = m.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.Status != "" {
			m = m.Where("t.status = ", param.Status)
		}

		if param.BeginTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []model.SysDictType
	m.Structs(&result)
	return result, nil
}

//添加数据
func (s *dictTypeService)Create(ctx context.Context,req *define.DictTypeApiCreateReq) (int64, error) {
	if s.CheckDictTypeUniqueAll(req.DictType){
		return 0,gerror.New("字典类型已存在")
	}
	user := shared.Context.Get(ctx).User
	var entity model.SysDictType
	entity.CreateTime = gtime.Now()
	entity.CreateBy = user.UserExtend.LoginName

	var editReq *define.DictTypeApiUpdateReq
	gconv.Struct(req,&editReq)
	return s.save(&entity,editReq)
}

//修改数据
func (s *dictTypeService)Update(ctx context.Context, req *define.DictTypeApiUpdateReq) (int64, error) {
	if s.CheckDictTypeUnique(req.DictType, req.DictId){
		return 0,gerror.New("字典类型已存在")
	}
	user := shared.Context.Get(ctx).User
	entity, err := dao.SysDictType.FindOne(dao.SysDictType.Columns.DictId,req.DictId)
	if err != nil {
		return 0, err
	}
	if entity == nil {
		return 0, gerror.New("数据不存在")
	}

	entity.UpdateTime = gtime.Now()
	entity.UpdateBy = user.UserExtend.LoginName
	return s.save(entity,req)
}

func (s *dictTypeService)save(dictType *model.SysDictType,req *define.DictTypeApiUpdateReq) (int64,error){
	dictType.Status = req.Status
	dictType.DictType = req.DictType
	dictType.DictName = req.DictName
	dictType.Remark = req.Remark
	result, err := dao.SysDictType.Data(dictType).Save()
	if err != nil {
		return 0, err
	}

	if dictType.DictId == 0 {
		// 新增
		id, err := result.LastInsertId()

		if err != nil || id <= 0 {
			return 0, err
		}
	}else{
		rs, err := result.RowsAffected()
		if err != nil || rs <= 0 {
			return 0, err
		}
	}
	return 1, nil
}

// 导出excel
func (s *dictTypeService)Export(param *define.DictTypeApiSelectPageReq) (string, error) {
	m := dao.SysDictType.As("t")
	if param != nil {
		if param.DictName != "" {
			m = m.Where("t.dict_name like ?", "%"+param.DictName+"%")
		}

		if param.DictType != "" {
			m = m.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.Status != "" {
			m = m.Where("t.status = ", param.Status)
		}

		if param.BeginTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	//"字典主键","字典名称","字典类型","状态","创建者","创建时间","更新者","更新时间","备注"
	m = m.Fields("t.dict_id,t.dict_name,t.dict_type,t.status,t.create_by,t.create_time,t.update_by,t.update_time,t.remark")

	result, err := m.M.All()
	if err != nil {
		return "", err
	}
	head := []string{"字典主键", "字典名称", "字典类型", "状态", "创建者", "创建时间", "更新者", "更新时间", "备注"}
	key := []string{"dict_id", "dict_name", "dict_type", "status", "create_by", "create_time", "update_by", "update_time", "remark"}
	url, err := excel.DownlaodExcel(head, key, result)
	if err != nil {
		return "", err
	}

	return url, nil
}

//批量删除数据记录
func (s *dictTypeService)Delete(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	dictTypes,_ := dao.SysDictType.Fields(dao.SysDictType.Columns.DictType).Where("dict_id in (?)",idarr).FindAll()
	result, err := dao.SysDictType.Delete("dict_id in (?)", idarr)
	if err != nil {
		return 0
	}
	if len(dictTypes) > 0 {
		for _, dictType := range dictTypes {
			g.DB().GetCache().Remove("sys_dict:"+dictType.DictType)
		}
	}
	nums, _ := result.RowsAffected()
	return nums
}

//检查字典类型是否唯一
func (s *dictTypeService)CheckDictTypeUniqueAll(configKey string) bool {
	entity, err := dao.SysDictType.FindOne(dao.SysDictType.Columns.DictType,configKey)
	if err != nil {
		return true
	}
	if entity != nil && entity.DictId > 0 {
		return true
	}
	return false
}

//检查字典类型是否唯一
func (s *dictTypeService)CheckDictTypeUnique(configKey string, configId int64) bool {
	entity, err := dao.SysDictType.FindOne(g.Map{
		dao.SysDictType.Columns.DictId:configId,
		dao.SysDictType.Columns.DictType:configKey,
	})
	if err != nil {
		return true
	}
	if entity != nil && entity.DictId > 0 {
		return true
	}
	return false
}