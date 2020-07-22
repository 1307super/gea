package dict_data

import (
	"gea/app/utils/page"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
)

// Fill with you ideas below.
//新增页面请求参数
type AddReq struct {
	DictLabel string `p:"dictLabel"  v:"required#字典标签不能为空"`
	DictValue string `p:"dictValue"  v:"required#字典键值不能为空"`
	DictType  string `p:"dictType"  v:"required#字典类型不能为空"`
	DictSort  int    `p:"dictSort"  v:"required#字典排序不能为空"`
	Status    string `p:"status"    v:"required#状态不能为空"`
	Remark    string `p:"remark"`
}

//修改页面请求参数
type EditReq struct {
	DictCode  int64  `p:"dictCode" v:"required#主键ID不能为空"`
	DictLabel string `p:"dictLabel"  v:"required#字典标签不能为空"`
	DictValue string `p:"dictValue"  v:"required#字典键值不能为空"`
	DictType  string `p:"dictType"`
	DictSort  int    `p:"dictSort"  v:"required#字典排序不能为空"`
	Status    string `p:"status"    v:"required#状态不能为空"`
	Remark    string `p:"remark"`
}

//分页请求参数
type SelectPageReq struct {
	DictType  string `p:"dictType"`  //字典名称
	DictLabel string `p:"dictLabel"` //字典标签
	Status    string `p:"status"`    //状态
	BeginTime string `p:"beginTime"` //开始时间
	EndTime   string `p:"endTime"`   //结束时间
	PageNum   int    `p:"pageNum"`   //当前页码
	PageSize  int    `p:"pageSize"`  //每页数
}

//根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) ([]Entity, *page.Paging, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil,nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_dict_data t")

	if param != nil {
		if param.DictLabel != "" {
			model.Where("t.dict_label like ?", "%"+param.DictLabel+"%")
		}

		if param.Status != "" {
			model.Where("t.status = ", param.Status)
		}

		if param.DictType != "" {
			model.Where("t.dict_type like ?", "%"+param.DictType+"%")
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

	model := db.Table("sys_dict_data t")

	if param != nil {
		if param.DictLabel != "" {
			model.Where("t.dict_label like ?", "%"+param.DictLabel+"%")
		}

		if param.Status != "" {
			model.Where("t.status = ", param.Status)
		}

		if param.DictType != "" {
			model.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	//"字典编码","字典排序","字典标签","字典键值","字典类型","样式属性","表格回显样式","是否默认","状态","创建者","创建时间","更新者","更新时间","备注"
	model.Fields("t.dict_code,t.dict_sort,t.dict_label,t.dict_value,t.dict_type,t.css_class,t.list_class,t.is_default,t.status,t.create_by,t.create_time,t.update_by,t.update_time,t.remark")

	result, _ := model.All()
	return result, nil
}

//获取所有数据
func SelectListAll(param *SelectPageReq) ([]Entity, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_dict_data t")

	if param != nil {
		if param.DictLabel != "" {
			model.Where("t.dict_label like ?", "%"+param.DictLabel+"%")
		}

		if param.Status != "" {
			model.Where("t.status = ", param.Status)
		}

		if param.DictType != "" {
			model.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []Entity

	err = model.Structs(&result)
	return result, err
}
