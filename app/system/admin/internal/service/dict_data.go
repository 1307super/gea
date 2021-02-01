package service

import (
	"context"
	"fmt"
	"gea/app/dao"
	"gea/app/model"
	"gea/app/shared"
	"gea/app/system/admin/internal/define"
	"gea/app/utils/convert"
	"gea/app/utils/page"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var DictData = &dictDataService{}

type dictDataService struct{}

func (s *dictDataService) Info(id int64) (*model.SysDictData,error) {
	return dao.SysDictData.FindOne(dao.SysDictData.Columns.DictCode,id)
}
//根据条件分页查询数据
func (s *dictDataService)GetList(param *define.DictDataApiSelectPageReq) *define.DictDataServiceList {
	m := dao.SysDictData.As("t")
	if param != nil {
		if param.DictLabel != "" {
			m = m.Where("t.dict_label like ?", "%"+param.DictLabel+"%")
		}
		if param.Status != "" {
			m = m.Where("t.status = ", param.Status)
		}
		if param.DictType != "" {
			m = m.Where("t.dict_type like ?", "%"+param.DictType+"%")
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

	result := &define.DictDataServiceList{
		Page:  page.PageNum,
		Size:  page.Pagesize,
		Total: page.Total,
	}
	if err = m.Structs(&result.List); err != nil {
		return nil
	}
	return result
}

//获取所有数据
func (s *dictDataService)GetAll(param *define.DictDataApiSelectPageReq) ([]model.SysDictData, error) {
	m := dao.SysDictData.As("t")
	if param != nil {
		if param.DictLabel != "" {
			m = m.Where("t.dict_label like ?", "%"+param.DictLabel+"%")
		}
		if param.Status != "" {
			m = m.Where("t.status = ", param.Status)
		}
		if param.DictType != "" {
			m = m.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}
		if param.BeginTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}
		if param.EndTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []model.SysDictData
	if err := m.Structs(&result); err != nil {
		return nil,gerror.New("未获取到数据")
	}
	return result, nil
}

//添加数据
func (s *dictDataService)Create(ctx context.Context, req *define.DictDataApiCreateReq) (int64, error) {
	user := shared.Context.Get(ctx).User
	var entity model.SysDictData
	entity.CreateTime = gtime.Now()
	entity.CreateBy = user.LoginName
	var editReq *define.DictDataApiUpdateReq
	gconv.Struct(req,&editReq)
	return s.save(&entity,editReq)
}

//修改数据
func (s *dictDataService)Update(ctx context.Context, req *define.DictDataApiUpdateReq) (int64, error) {
	user := shared.Context.Get(ctx).User
	entity, err := dao.SysDictData.FindOne(dao.SysDictData.Columns.DictCode,req.DictCode)
	if err != nil {
		return 0, err
	}
	if entity == nil {
		return 0, gerror.New("数据不存在")
	}
	entity.UpdateTime = gtime.Now()
	entity.UpdateBy = user.LoginName
	return s.save(entity,req)
}

func (s *dictDataService) save(dictData *model.SysDictData, req *define.DictDataApiUpdateReq) (int64,error){
	var (
		rs int64
		err error
	)
	dictData.DictType = req.DictType
	dictData.Status = req.Status
	dictData.DictLabel = req.DictLabel
	dictData.DictSort = req.DictSort
	dictData.DictValue = req.DictValue
	dictData.Remark = req.Remark
	result, err := dao.SysDictData.Data(dictData).Save()
	if err != nil {
		return 0, err
	}
	if dictData.DictCode == 0 {
		// 新增
		rs, err = result.LastInsertId()
	}else{
		rs, err = result.RowsAffected()
	}
	if err != nil {
		return 0, err
	}
	return rs, nil
}

func (s *dictDataService) Delete(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := dao.SysDictData.Delete(fmt.Sprintf("%s in(?)",dao.SysDictData.Columns.DictCode),idarr)
	if err != nil {
		return 0
	}
	nums, _ := result.RowsAffected()
	return nums
}

// 导出excel
func (s *dictDataService)Export(param *define.DictDataApiSelectPageReq) (gdb.Result, error) {
	m := dao.SysDictData.As("t")
	if param != nil {
		if param.DictLabel != "" {
			m = m.Where("t.dict_label like ?", "%"+param.DictLabel+"%")
		}
		if param.Status != "" {
			m = m.Where("t.status = ", param.Status)
		}
		if param.DictType != "" {
			m = m.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}
		if param.BeginTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}
		if param.EndTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}
	//"字典编码","字典排序","字典标签","字典键值","字典类型","样式属性","表格回显样式","是否默认","状态","创建者","创建时间","更新者","更新时间","备注"
	m = m.Fields("t.dict_code,t.dict_sort,t.dict_label,t.dict_value,t.dict_type,t.css_class,t.list_class,t.is_default,t.status,t.create_by,t.create_time,t.update_by,t.update_time,t.remark")

	result, _ := m.M.All()
	return result, nil
}