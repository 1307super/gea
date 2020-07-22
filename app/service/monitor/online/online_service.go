// ==========================================================================
// GEAGO自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2020-02-17 14:03:51
// 生成路径: app/service/module/online/online_service.go
// 生成人：yunjie
// ==========================================================================
package online

import (
	onlineModel "gea/app/model/monitor/online"
	"gea/app/utils/convert"
	"gea/app/utils/page"
	"gea/app/utils/token"
	"github.com/gogf/gf/frame/g"
)

//根据主键查询数据
func SelectRecordById(id int64) (*onlineModel.Entity, error) {
	return onlineModel.FindOne("sessionId", id)
}

//根据主键删除数据
func DeleteRecordById(id int64) bool {
	result, err := onlineModel.Delete("sessionId", id)
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
	result, err := onlineModel.Delete("sessionId in (?)", idarr)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

//批量删除数据
func DeleteRecordNotInIds(ids []string) int64 {
	result, err := onlineModel.Delete("sessionId not in (?)", ids)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

//添加数据
func AddSave(entity onlineModel.Entity) (int64, error) {
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

//根据条件查询数据
func SelectListAll(params *onlineModel.SelectPageReq) ([]onlineModel.Entity, error) {
	return onlineModel.SelectListAll(params)
}

//根据条件分页查询数据
func SelectListByPage(params *onlineModel.SelectPageReq) ([]onlineModel.Entity, *page.Paging, error) {
	return onlineModel.SelectListByPage(params)
}

// 强退
func ForceLogout(tokenStr string) {
	token.RemoveCache(tokenStr)
	// 修改状态
	onlineModel.Delete(g.Map{"token":tokenStr})
}
