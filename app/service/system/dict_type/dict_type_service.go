package dict_type

import (
	"gea/app/model"
	dictTypeModel "gea/app/model/system/dict_type"
	userService "gea/app/service/system/user"
	"gea/app/utils/convert"
	"gea/app/utils/excel"
	"gea/app/utils/page"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

//根据主键查询数据
func SelectRecordById(id int64) (*dictTypeModel.Entity, error) {
	return dictTypeModel.FindOne("dict_id", id)
}

//根据主键删除数据
func DeleteRecordById(id int64) bool {
	result, err := dictTypeModel.Delete("dict_id", id)
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
	result, err := dictTypeModel.Delete("dict_id in (?)", idarr)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

//添加数据
func AddSave(req *dictTypeModel.AddReq, r *ghttp.Request) (int64, error) {
	var entity dictTypeModel.Entity
	entity.Status = req.Status
	entity.DictType = req.DictType
	entity.DictName = req.DictName
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
func EditSave(req *dictTypeModel.EditReq, r *ghttp.Request) (int64, error) {

	entity, err := dictTypeModel.FindOne("dict_id=?", req.DictId)

	if err != nil {
		return 0, err
	}

	if entity == nil {
		return 0, gerror.New("数据不存在")
	}

	entity.Status = req.Status
	entity.DictType = req.DictType
	entity.DictName = req.DictName
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
func SelectListAll(params *dictTypeModel.SelectPageReq) ([]dictTypeModel.Entity, error) {
	return dictTypeModel.SelectListAll(params)
}

//根据条件分页查询角色数据
func SelectListByPage(params *dictTypeModel.SelectPageReq) ([]dictTypeModel.Entity, *page.Paging,error) {
	return dictTypeModel.SelectListByPage(params)
}

//根据字典类型查询信息
func SelectDictTypeByType(dictType string) *dictTypeModel.Entity {
	rs, err := dictTypeModel.FindOne("dict_type=?", dictType)
	if err != nil {
		return nil
	}
	return rs
}

// 导出excel
func Export(param *dictTypeModel.SelectPageReq) (string, error) {
	result, err := dictTypeModel.SelectListExport(param)
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

//检查字典类型是否唯一
func CheckDictTypeUniqueAll(configKey string) string {
	entity, err := dictTypeModel.CheckDictTypeUniqueAll(configKey)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.DictId > 0 {
		return "1"
	}
	return "0"
}

//检查字典类型是否唯一
func CheckDictTypeUnique(configKey string, configId int64) string {
	entity, err := dictTypeModel.CheckDictTypeUnique(configKey, configId)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.DictId > 0 {
		return "1"
	}
	return "0"
}

//查询字典类型树
func SelectDictTree(params *dictTypeModel.SelectPageReq) *[]model.Ztree {
	var result []model.Ztree
	dictList, err := dictTypeModel.SelectListAll(params)
	if err == nil && dictList != nil {
		for _, item := range dictList {
			var tmp model.Ztree
			tmp.Id = item.DictId
			tmp.Name = transDictName(item)
			tmp.Title = item.DictType
			result = append(result, tmp)
		}
	}
	return &result
}

func transDictName(entity dictTypeModel.Entity) string {
	return `(` + entity.DictName + `)&nbsp;&nbsp;&nbsp;` + entity.DictType
}
