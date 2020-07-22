package dict_type

import (
	"gea/app/utils/page"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
)

// Fill with you ideas below.
//新增页面请求参数
type AddReq struct {
	DictName string `p:"dictName"  v:"required#字典名称不能为空"`
	DictType string `p:"dictType"  v:"required#字典类型不能为空"`
	Status   string `p:"status"  v:"required#状态不能为空"`
	Remark   string `p:"remark"`
}

//修改页面请求参数
type EditReq struct {
	DictId   int64  `p:"dictId" v:"required#主键ID不能为空"`
	DictName string `p:"dictName"  v:"required#字典名称不能为空"`
	DictType string `p:"dictType"  v:"required#字典类型不能为空"`
	Status   string `p:"status"  v:"required#状态不能为空"`
	Remark   string `p:"remark"`
}

//分页请求参数
type SelectPageReq struct {
	DictName      string `p:"dictName"`      //字典名称
	DictType      string `p:"dictType"`      //字典类型
	Status        string `p:"status"`        //字典状态
	BeginTime     string `p:"beginTime"`     //开始时间
	EndTime       string `p:"endTime"`       //结束时间
	OrderByColumn string `p:"orderByColumn"` //排序字段
	IsAsc         string `p:"isAsc"`         //排序方式
	PageNum       int    `p:"pageNum"`       //当前页码
	PageSize      int    `p:"pageSize"`      //每页数
}

//检查字典类型请求参数
type CheckDictTypeReq struct {
	DictId   int64  `p:"dictId"  v:"required#ID不能为空"`
	DictType string `p:"dictType"  v:"required#参数键名不能为空"`
}

//检查字典类型请求参数
type CheckDictTypeALLReq struct {
	DictType string `p:"dictType"  v:"required#参数键名不能为空"`
}

//根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) ([]Entity,  *page.Paging, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil,nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_dict_type t")

	if param != nil {
		if param.DictName != "" {
			model.Where("t.dict_name like ?", "%"+param.DictName+"%")
		}

		if param.DictType != "" {
			model.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.Status != "" {
			model.Where("t.status = ", param.Status)
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


	if param.OrderByColumn != "" {
		model.Order(param.OrderByColumn + " " + param.IsAsc)
	}

	var result []Entity
	model.Structs(&result)
	return result,page, nil
}

// 导出excel
func SelectListExport(param *SelectPageReq) (gdb.Result, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_dict_type t")

	if param != nil {
		if param.DictName != "" {
			model.Where("t.dict_name like ?", "%"+param.DictName+"%")
		}

		if param.DictType != "" {
			model.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.Status != "" {
			model.Where("t.status = ", param.Status)
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	//"字典主键","字典名称","字典类型","状态","创建者","创建时间","更新者","更新时间","备注"
	model.Fields("t.dict_id,t.dict_name,t.dict_type,t.status,t.create_by,t.create_time,t.update_by,t.update_time,t.remark")

	result, _ := model.All()
	return result, nil
}

//获取所有数据
func SelectListAll(param *SelectPageReq) ([]Entity, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_dict_type t")

	if param != nil {
		if param.DictName != "" {
			model.Where("t.dict_name like ?", "%"+param.DictName+"%")
		}

		if param.DictType != "" {
			model.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.Status != "" {
			model.Where("t.status = ", param.Status)
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

//校验字典类型是否唯一
func CheckDictTypeUnique(dictType string, dictId int64) (*Entity, error) {
	return FindOne("dict_id !=? and dict_type=?", dictId, dictType)
}

//校验字典类型是否唯一
func CheckDictTypeUniqueAll(dictType string) (*Entity, error) {
	return FindOne("dict_type=?", dictType)
}
