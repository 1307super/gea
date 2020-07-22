package dict_data

import (
	"gea/app/utils/page"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	dictDataModel "gea/app/model/system/dict_data"
	userService "gea/app/service/system/user"
	"gea/app/utils/convert"
	"gea/app/utils/excel"
)

//根据主键查询数据
func SelectRecordById(id int64) (*dictDataModel.Entity, error) {
	return dictDataModel.FindOne("dict_code", id)
}

//根据主键删除数据
func DeleteRecordById(id int64) bool {
	result, err := dictDataModel.Delete("dict_code", id)
	if err == nil {
		affected, _ := result.RowsAffected()
		if affected > 0 {
			return true
		}
	}

	return false
}

//批量删除数据记录
func DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := dictDataModel.Delete("dict_code in (?)", idarr)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

//添加数据
func AddSave(req *dictDataModel.AddReq, r *ghttp.Request) (int64, error) {
	var entity dictDataModel.Entity
	entity.DictType = req.DictType
	entity.Status = req.Status
	entity.DictLabel = req.DictLabel
	entity.DictSort = req.DictSort
	entity.DictValue = req.DictValue
	entity.Remark = req.Remark
	entity.CreateTime = gtime.Now()
	entity.CreateBy = ""

	user,_ := userService.GetProfileApi(r.GetInt64("jwtUid"))

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	result, err := entity.Insert()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil || id <= 0 {
		return 0, err
	}
	return id, nil
}

//修改数据
func EditSave(req *dictDataModel.EditReq, r *ghttp.Request) (int64, error) {

	entity, err := dictDataModel.FindOne("dict_code=?", req.DictCode)

	if err != nil {
		return 0, err
	}

	if entity == nil {
		return 0, gerror.New("数据不存在")
	}

	entity.DictType = req.DictType
	entity.Status = req.Status
	entity.DictLabel = req.DictLabel
	entity.DictSort = req.DictSort
	entity.DictValue = req.DictValue
	entity.Remark = req.Remark
	entity.UpdateTime = gtime.Now()
	entity.UpdateBy = ""

	user,_ := userService.GetProfileApi(r.GetInt64("jwtUid"))

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	result, err := entity.Update()

	if err != nil {
		return 0, err
	}

	rs, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rs, nil
}

//根据条件分页查询角色数据
func SelectListAll(params *dictDataModel.SelectPageReq) ([]dictDataModel.Entity, error) {
	return dictDataModel.SelectListAll(params)
}

//根据条件分页查询角色数据
func SelectListByPage(params *dictDataModel.SelectPageReq) ([]dictDataModel.Entity, *page.Paging,  error) {
	return dictDataModel.SelectListByPage(params)
}

// 导出excel
func Export(param *dictDataModel.SelectPageReq) (string, error) {
	result, err := dictDataModel.SelectListExport(param)
	if err != nil {
		return "", err
	}

	head := []string{"字典编码", "字典排序", "字典标签", "字典键值", "字典类型", "样式属性", "表格回显样式", "是否默认", "状态", "创建者", "创建时间", "更新者", "更新时间", "备注"}
	key := []string{"dict_code", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "status", "create_by", "create_time", "update_by", "update_time", "remark"}
	url, err := excel.DownlaodExcel(head, key, result)

	if err != nil {
		return "", err
	}

	return url, nil
}
